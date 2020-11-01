package writer

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func isLicenceExist(fileName string, licenceFile string) (bool, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return false, err
	}
	fileContent := readFile(file.Name())
	err = file.Close()
	if err != nil {
		return false, err
	}
	LICENCE := readFile(licenceFile)
	return strings.Contains(fileContent, LICENCE), err
}

func isLicenceFormatValid(file *os.File) bool {
	content := readFile(file.Name())
	lines := strings.Split(content, "\n")
	firstLine := lines[1]
	fileContent := readFile(file.Name())
	if strings.Contains(fileContent, firstLine) {
		index := strings.Index(fileContent, firstLine)
		tempLicence := ""
		for i := index; fileContent[i-1] != '*' || fileContent[i] != '/'; i++ {
			tempLicence += string(fileContent[i])
		}
	}
	return true
}

func WriteLicenceToFile(fileName string, licenceFile string, startTag string, endTag string) (bool, error) {
	ok, err := isLicenceExist(fileName, licenceFile)
	if ok {
		return false, err
	}
	if startTag == "" || endTag == "" {
		startTag = "/*"
		endTag = "*/"
	}
	bytes, _ := ioutil.ReadFile(fileName)
	var content string
	licence := readFile(licenceFile)
	if !isLicenceAlreadyContainsTags(licence, startTag) {
		licence = addTagToLicence(startTag, endTag, licence)
	}
	content = licence + string(bytes)
	_ = ioutil.WriteFile(fileName, []byte(content), 0644)
	return true, nil
}
func isLicenceAlreadyContainsTags(licence string, startTag string) bool {
	lines := strings.Split(licence, "\n")
	return strings.Contains(lines[0], startTag)
}

func addTagToLicence(startTag string, endTag string, licence string) string {
	licence = startTag + "\n" + licence + "\n" + endTag
	return licence
}

// readFile ...
func readFile(fileName string) string {
	file, err := os.Open(fileName)
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
