package parser

import (
	"fmt"
	"strings"
)

var keywords = map[string]bool{
	"export":      true,
	"const":       true,
	"let":         true,
	"var":         true,
	"default":     true,
	"class":       true,
	"function":    true,
	"async":       true,
	"extends":     true,
	"implements":  true,
	"abstract":    true,
	"import":      true,
	"constructor": true,
	"return":      true,
	"//":          true,
	"/*":          true,
}

var symbols = map[string]bool{
	"(":  true,
	")":  true,
	"{":  true,
	"}":  true,
	"\"": true,
	"'":  true,
	"`":  true,
	":":  true,
	";":  true,
	"=":  true,
}

func ParseLine(line string) []string {
	var tokens []string
	match := ""
	var current string
	for i := 0; i < len(line); i++ {
		current = string(line[i])
		if _, ok := symbols[current]; ok {
			if current == "\"" || current == "'" || current == "`" {
				endIdx := findEndIndexOfString(line, i)
				tokens = append(tokens, strings.Replace(line[i+1:endIdx], "\\", "", -1))
				i = endIdx
				continue
			}
			match = strings.Trim(match, " ")
			if len(match) != 0 {
				tokens = append(tokens, match)
			}
			tokens = append(tokens, string(line[i]))
			match = ""
			continue
		}
		match = strings.Trim(match, " ")
		if current == " " && len(match) != 0 {
			tokens = append(tokens, match)
			match = ""
			continue
		}
		match += current
	}
	match = strings.Trim(match, " ")
	if match != "" {
		tokens = append(tokens, match)
	}
	fmt.Println(tokens)
	return tokens
}

func findEndIndexOfString(line string, startIdx int) int {
	endIdx := startIdx + 1
	for ; line[endIdx] != line[startIdx]; {
		if string(line[endIdx]) == "\\" {
			endIdx++
		}
		endIdx++
	}
	return endIdx
}
