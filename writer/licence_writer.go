package writer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const LICENCE = `/*
    Copyright 2020. Huawei Technologies Co., Ltd. All rights reserved.

    Licensed under the Apache License, Version 2.0 (the "License")
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        https://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/
`

func CheckIfLicenceExists(fileName string) (bool, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return false, err
	}
	scanner := bufio.NewScanner(file)
	fileContent := ""
	for scanner.Scan() {
		fileContent += scanner.Text() + "\n"
	}
	return strings.Contains(fileContent, LICENCE), err
}

func CheckIfLicenceFormatIsValid(file *os.File) bool {
	firstLine := "Copyright 2020. Huawei Technologies Co., Ltd. All rights reserved."
	scanner := bufio.NewScanner(file)
	fileContent := ""
	for scanner.Scan() {
		fileContent += scanner.Text() + "\n"
	}
	if strings.Contains(fileContent, firstLine) {
		index := strings.Index(fileContent, firstLine)
		tempLicence := ""
		for i := index; fileContent[i-1] != '*' || fileContent[i] != '/'; i++ {
			tempLicence += string(fileContent[i])
		}
		fmt.Println(tempLicence)
	}
	return true
}
