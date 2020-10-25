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

func CheckIfFileContainsLicenceAlready(fileName string) (bool, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return false, err
	}
	scanner := bufio.NewScanner(file)
	fileContent := ""
	for scanner.Scan() {
		fileContent += scanner.Text() + "\n"
	}
	err2 := file.Close()
	if err2 != nil {
		return false, err2
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

func WriteToFileLicence(fileName string) (bool, error) {
	ok, err := CheckIfFileContainsLicenceAlready(fileName)

	if ok {
		return false, err
	}

	// make a temporary outfile
	outfile, err := os.Create("temp.java")

	if err != nil {
		return false, err
	}

	defer outfile.Close()

	// open the file to be appended to for read
	f, err := os.Open(fileName)

	if err != nil {
		return false, err
	}

	defer f.Close()

	// append at the start
	_, err = outfile.WriteString(LICENCE)
	if err != nil {
		return false, err
	}

	scanner := bufio.NewScanner(f)

	// read the file to be appended to and output all of it
	for scanner.Scan() {
		_, err = outfile.WriteString(scanner.Text())
		_, err = outfile.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}
	// ensure all lines are written
	outfile.Sync()
	//close the file to rename it, otherwise error will be thrown
	err = outfile.Close()
	if err != nil {
		return false, err
	}
	err = f.Close()
	if err != nil {
		return false, err
	}
	// overwrite the old file with the new one
	err = os.Remove(fileName)
	if err != nil {
		return false, err
	}
	err = os.Rename(outfile.Name(), fileName)
	if err != nil {
		return false, err
	}
	return true, nil
}
