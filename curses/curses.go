package curses

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lcurses
// #include <curses.h>
// #include <locale.h>
import "C"

func InitScr() {
	C.setlocale(C.LC_ALL, C.CString(""))
	C.initscr()
}

func EndWin() {
	C.endwin()
}

func NoEcho() {
	C.noecho()
}

func Echo() {
	C.echo()
}

func Cbreak() {
	C.cbreak()
}

func NoCbreak() {
	C.nocbreak()
}

func CursSet(visibility int) {
	C.curs_set(C.int(visibility))
}

func GetMaxX() int {
	return int(C.getmaxx(C.stdscr))
}

func GetMaxY() int {
	return int(C.getmaxy(C.stdscr))
}

func Clear() {
	C.clear()
}

func Refresh() {
	C.refresh()
}

func NoDelay(enabled bool) {
	C.nodelay(C.stdscr, C.bool(enabled))
}

func GetCh() rune {
	return rune(C.getch())
}

func MvAddStr(y int, x int, str string) {
	C.mvaddstr(C.int(y), C.int(x), C.CString(str))
}
