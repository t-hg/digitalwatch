package style

import "strings"

func Apply(text string, charset []string) string {
	styledText := ""
	for _, char := range text {
		styledChar := ""
		switch char {
		case '0':
			styledChar = charset[0]
		case '1':
			styledChar = charset[1]
		case '2':
			styledChar = charset[2]
		case '3':
			styledChar = charset[3]
		case '4':
			styledChar = charset[4]
		case '5':
			styledChar = charset[5]
		case '6':
			styledChar = charset[6]
		case '7':
			styledChar = charset[7]
		case '8':
			styledChar = charset[8]
		case '9':
			styledChar = charset[9]
		case ':':
			styledChar = charset[10]
		}
		if styledChar == "" {
			continue
		}
		styledChar = strings.TrimPrefix(styledChar, "\n")
		styledChar = strings.TrimSuffix(styledChar, "\n")
		if styledText == "" {
			styledText = styledChar
			continue
		}
		styledTextLines := strings.Split(styledText, "\n")
		styledCharLines := strings.Split(styledChar, "\n")
		if len(styledTextLines) != len(styledCharLines) {
			continue
		}
		for idx := range styledTextLines {
			styledTextLines[idx] += styledCharLines[idx]
		}
		styledText = strings.Join(styledTextLines, "\n")
	}
	return styledText
}
