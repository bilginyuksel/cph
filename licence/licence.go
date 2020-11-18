package licence

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"strings"
)

// LICENCE ... -- apache licence
// Copyright 2020. Huawei Technologies Co., Ltd. All rights reserved.
var LICENCE = `    Copyright 2020. Huawei Technologies Co., Ltd. All rights reserved.

	Licensed under the Apache License, Version 2.0 (the "License")
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.`

// IsExists ...
func IsExists(content string) bool {
	return strings.Contains(content, LICENCE)
}

// var knownTypes = string[]{}
var extensions = map[string][]string{
	".html": {"<!--", "-->"},
	".py":   {"\"\"\"", "\"\"\""},
	".java": {"/*", "*/"},
	".css":  {"/*", "*/"},
	".scss": {"/*", "*/"},
	".js":   {"/*", "*/"},
	".ts":   {"/*", "*/"},
}

func findHowManyCommentExists(content string, extension string) int {
	endTag := extensions[extension][1]
	count := 0
	singleQuoteFlag := false
	doubleQuoteFlag := false
	isNotQuoteStarted := false
	for i := 2; i < len(content); i++ {
		if content[i] == '\'' {
			singleQuoteFlag = !singleQuoteFlag
		}
		if content[i] == '"' {
			doubleQuoteFlag = !doubleQuoteFlag
		}
		isNotQuoteStarted = !doubleQuoteFlag && !singleQuoteFlag
		if isNotQuoteStarted && isCommentEndingFound(content, i, endTag) {
			count++
		}
	}
	return count
}

func isCommentEndingFound(content string, startIdx int, endTag string) bool {
	return content[startIdx-len(endTag)+1:startIdx+1] == endTag
}

// Write ...
func Write(filePath string) {

	content := readFile(filePath)
	if IsExists(content) {
		return
	}
	similarity := findCommentedInvalidLicenceToDelete(content, 0.8)
	if similarity.similar {
		content = deleteInvalidLicence(content, similarity.startIdx, similarity.endIdx)
	}
	extension := filepath.Ext(filePath)
	if _, ok := extensions[extension]; !ok {
		fmt.Printf("Unknown file extension= %s! Can't licence it.", extension)
		return
	}
	content = addTagToLicence(extension) + content
	ioutil.WriteFile(filePath, []byte(content), 0644)
}

type licenceSimilarity struct {
	similar  bool
	startIdx int
	endIdx   int
	prob     float64
}

func deleteInvalidLicence(content string, startIdx int, endIdx int) string {
	return strings.Replace(content, content[startIdx:endIdx+1], "", 1)
}

func findCommentedInvalidLicenceToDelete(content string, bound float64) licenceSimilarity {
	var prob float64 = 0
	var startIdx = 0
	var endIdx = 0
	for i := 0; i < len(content); i++ {
		if content[i] == '*' && content[i-1] == '/' {
			end := findEndIdxOfTheBlockComment(content, i)
			potentialLicence := content[i+1 : end+1]
			tempProb := comparePotentialLicenceToActualLicence(potentialLicence)
			if tempProb > prob {
				prob = tempProb
				startIdx = i - 1
				endIdx = end
			}
			i = end + 1
		}
	}

	if bound <= prob {
		return licenceSimilarity{similar: true, startIdx: startIdx, endIdx: endIdx, prob: prob}
	}
	return licenceSimilarity{similar: false, prob: prob}
}

func comparePotentialLicenceToActualLicence(potentialLicence string) float64 {
	potentialLicence = strings.ReplaceAll(potentialLicence, "\n", " ")
	licence := strings.ReplaceAll(LICENCE, "\n", " ")

	potentialLicenceMap := stringArrayToMap(strings.Split(potentialLicence, " "))
	originalLicenceMap := stringArrayToMap(strings.Split(licence, " "))

	actualTotal := 0
	for _, value := range originalLicenceMap {
		actualTotal += value
	}
	match := 0
	for key, value := range potentialLicenceMap {
		match += int(math.Min(float64(value), float64(originalLicenceMap[key])))
	}

	return float64(match) / float64(actualTotal)
}

func max(num1 int, num2 int) int {
	if num1 < num2 {
		return num2
	}
	return num1
}

func stringArrayToMap(arr []string) map[string]int {
	tempMap := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		tempMap[arr[i]]++
	}
	return tempMap
}

func findEndIdxOfTheBlockComment(content string, startIdx int) int {
	for i := startIdx + 1; i < len(content); i++ {
		if content[i] == '/' && content[i-1] == '*' {
			return i
		}
	}
	return startIdx
}

func addTagToLicence(extension string) string {
	tag := extensions[extension]
	licence := tag[0] + "\n" + LICENCE + "\n" + tag[1] + "\n\n"
	return licence
}

// !NOT TESTED!
func readFile(filename string) string {
	bytes, _ := ioutil.ReadFile(filename)
	return string(bytes)
}
