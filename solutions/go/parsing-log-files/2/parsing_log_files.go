package parsinglogfiles

import (
    "fmt"
    "regexp"
)

func IsValidLine(text string) bool {
    re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*-=]*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    sum := 0
    re := regexp.MustCompile(`(?i)".*password.*"`)

    for _, line := range lines {
        if re.MatchString(line) {
            sum++
        }
    }
    
	return sum
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+([\w]+)`)
    
    for index, line := range lines {
        result := re.FindStringSubmatch(line)
        if result != nil && result[1] != ""{
            lines[index] = fmt.Sprintf("[USR] %s %s", result[1], line)
        }
    }
    
    return lines
}
