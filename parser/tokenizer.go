package parser

import (
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
	"@":  true,
	"(":  true,
	")":  true,
	"{":  true,
	"}":  true,
	"\"": true,
	"+":  true,
	"'":  true,
	",":  true,
	"`":  true,
	":":  true,
	";":  true,
	"=":  true,
}


var tokens []string

func Tokenize(content string) []string {
	// I couldn't pass slice by reference so I created a global slice and initalized it in
	// this method. Whenever this method called it will be initialized again.
	tokens = []string{}
	currentWord, currentElem := "", ""

	for idx := 0; idx < len(content); idx++ {
		currentElem = string(content[idx])

		if _, ok := symbols[currentElem]; ok {
			idx = tokenizeSymbolThenReturnEndIdx(content, currentWord, idx)
			currentWord = ""
			continue
		}

		// Advance improvement.
		// Create a special character hashMap for ",',`,/,@ and
		// map the functions to this characters
		// if hashMap returns ok result then call the function for the returned value.
		if currentElem == "/" {
			idx = tokenizeCommentThenReturnEndIdx(content, idx)
			continue
		}

		currentWord = strings.Trim(currentWord, " ")
		currentWord = strings.Trim(currentWord, "\n")
		currentWord = strings.Trim(currentWord, "\t")
		if (currentElem == " " || currentElem == "\n") && len(currentWord) > 0 {
			tokens = append(tokens, currentWord)
			currentWord = ""
			continue
		}
		if currentElem == "\n" || currentElem == " " || currentElem == "\t" {
			continue
		}
		currentWord += currentElem
	}

	// If currentWord is not empty add to tokens
	currentWord = strings.Trim(currentWord, " ")
	currentWord = strings.Trim(currentWord, "\n")
	currentWord = strings.Trim(currentWord, "\t")
	if len(currentWord) > 0 {
		tokens = append(tokens, currentWord)
	}

	return tokens
}


func tokenizeCommentThenReturnEndIdx(content string, startIdx int) int {
	endIdx := startIdx + 1
	comment := ""

	if string(content[endIdx]) == "/" {
		endIdx++
		// oneline comment
		// it should be \n or \r\n find the difference
		for ;endIdx<len(content) && string(content[endIdx]) != "\n"; endIdx++ {
			comment += string(content[endIdx])
		}

	} else if string(content[endIdx]) == "*" {
		endIdx++
		// multiline comment
		for ;endIdx<len(content)-1; endIdx++ {
			if string(content[endIdx]) == "*" && string(content[endIdx+1]) == "/" {
				endIdx += 2
				break
			}
			comment += string(content[endIdx])
		}
	}

	tokens = append(tokens, comment)

	return endIdx
}


func findEndIndexOfString(content string, startIdx int) int {
	endIdx := startIdx + 1
	for ;content[endIdx] != content[startIdx]; {
		if string(content[endIdx]) == "\\" {
			endIdx++
		}
		endIdx++
	}
	return endIdx
}

func tokenizeSymbolThenReturnEndIdx(content string, currentWord string, startIdx int) int {
	symbol := string(content[startIdx])
	isStringStart := (symbol == "\"" || symbol == "'" || symbol == "`")
	if isStringStart {
		endIdx := findEndIndexOfString(content, startIdx)
		tokens = append(tokens, strings.Replace(content[startIdx+1:endIdx], "\\", "", -1))
		return endIdx
	}

	currentWord = strings.Trim(currentWord, " ")
	if len(currentWord) > 0 {
		tokens = append(tokens, currentWord)
	}
	tokens = append(tokens, symbol)

	return startIdx
}


