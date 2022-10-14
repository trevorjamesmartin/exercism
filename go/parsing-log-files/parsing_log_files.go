package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]?`)
	if err != nil {
		return false
	}
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<(.[^a-z|A-Z|0-9]*?)\>|<>`)
	if err != nil {
		return nil
	}
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`(?i)\"(.*?)password(.*?)\"`)
	t := 0
	for _, text := range lines {
		result := re.FindAllString(text, -1)
		if result != nil {
			t += len(result)
		}
	}
	return t
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line[0-9]*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, _ := regexp.Compile(`User\s+[a-zA-z]*[0-9]*`)
	r2, _ := regexp.Compile(`\s+`)
	for i, text := range lines {
		u := re.FindString(text)
		if u != "" {
			username := r2.Split(u, -1)
			lines[i] = fmt.Sprintf(`[USR] %s %s`, username[len(username)-1], text)
		}
	}
	return lines
}
