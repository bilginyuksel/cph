package writer

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func IsLicenceExist(filePath string, licenceFilePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	fileContent := readFile(file.Name())
	err = file.Close()
	if err != nil {
		return false, err
	}
	LICENCE := readFile(licenceFilePath)
	return strings.Contains(fileContent, LICENCE), err
}

func IsLicenceFormatValid(filePath string, licenceFilePath string) bool {
	licence := readFile(licenceFilePath)
	licenceArray := strings.Split(licence, "\n")
	fileContent := readFile(filePath)
	for i := 0; i < len(licenceArray); i++ {
		line := licenceArray[i]
		if strings.Contains(fileContent, line) {
			startIndex := strings.Index(fileContent, "/*")
			endIndex := strings.Index(fileContent, "*/")
			if startIndex >= 0 && endIndex > 0 {
				tempLicence := ""
				for i := startIndex; i < endIndex+2; i++ {
					tempLicence += string(fileContent[i])
				}
				fmt.Println(tempLicence)
				return licence == tempLicence
			}
		}
	}
	return true
}

func WriteLicenceToFile(filePath string, licenceFilePath string, startTag string, endTag string) (bool, error) {
	licenceExist, err := IsLicenceExist(filePath, licenceFilePath)
	licenceFormatValid := IsLicenceFormatValid(filePath, licenceFilePath)
	if licenceExist || !licenceFormatValid {
		return false, err
	}
	if startTag == "" || endTag == "" {
		startTag = "/*"
		endTag = "*/"
	}
	bytes, _ := ioutil.ReadFile(filePath)
	var content string
	licence := readFile(licenceFilePath)
	if !isLicenceAlreadyContainsTags(licence, startTag) {
		licence = addTagToLicence(startTag, endTag, licence)
	}
	content = licence + string(bytes)
	_ = ioutil.WriteFile(filePath, []byte(content), 0644)
	return true, nil
}
func isLicenceAlreadyContainsTags(licence string, startTag string) bool {
	lines := strings.Split(licence, "\n")
	return strings.Contains(lines[0], startTag)
}

func addTagToLicence(startTag string, endTag string, licence string) string {
	licence = startTag + "\n" + licence + endTag + "\n"
	return licence
}

func readFile(filePath string) string {
	file, err := os.Open(filePath)
	checkErr(err)
	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	file.Close()
	return content
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
