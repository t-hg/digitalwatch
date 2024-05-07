package main

import (
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/t-hg/digitalwatch/curses"
	"github.com/t-hg/digitalwatch/style"
)

func render(text string) {
	// find y, x so that given
	// text is centered
	maxLineLen := 0
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		runes := []rune(line)
		if len(runes) > maxLineLen {
			maxLineLen = len(runes)
		}
	}
	maxY := curses.GetMaxY()
	maxX := curses.GetMaxX()
	y := maxY/2 - len(lines)/2
	x := maxX/2 - maxLineLen/2

	// print lines respectively
	curses.Clear()
	for idx, line := range lines {
		curses.MvAddStr(y+idx, x, line)
	}
	curses.Refresh()
}

func main() {
	// flags
	flagStyle := flag.Int("style", 1, "different styles (1, 2 or 3)")
	flag.Parse()

	// setup
	curses.InitScr()
	curses.Cbreak()
	curses.NoEcho()
	curses.CursSet(0)
	curses.NoDelay(true)

	// signal handlers
	sigint := make(chan os.Signal, 1)
	sigwinch := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)
	signal.Notify(sigwinch, syscall.SIGWINCH)

	// text to be displayed
	memoizedText := ""
	text := ""

	// styling to be used
	var charset []string
	switch *flagStyle {
	case 1:
		charset = style.Charset1
	case 2:
		charset = style.Charset2
	case 3:
		charset = style.Charset3
	}

loop:
	for {
		// handle signals
		select {
		case <-sigint:
			break loop
		case <-sigwinch:
			curses.EndWin()
			curses.Refresh()
		default:
		}

		// handle character input
		switch curses.GetCh() {
		case 'q':
			break loop
		}

		// update watch
		text = time.Now().Format("15:04:05")
		text = style.Apply(text, charset)

		// display text
		if text != memoizedText {
			// call render only if text has not changed.
			// this reduces flickering
			render(text)
			memoizedText = text
		}

		// little time interval
		// to avoid busy wait
		time.Sleep(500 * time.Millisecond)
	}

	// cleanup
	curses.CursSet(2)
	curses.Echo()
	curses.NoCbreak()
	curses.EndWin()
}
