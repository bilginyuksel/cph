package writer

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func CheckIfFileContainsLicenceAlready(fileName string, licenceFile string) (bool, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return false, err
	}
	fileContent := ReadFile(file.Name())
	err = file.Close()
	if err != nil {
		return false, err
	}
	LICENCE := ReadFile(licenceFile)
	return strings.Contains(fileContent, LICENCE), err
}

func CheckIfLicenceFormatIsValid(file *os.File) bool {
	content := ReadFile(file.Name())
	lines := strings.Split(content, "\n")
	firstLine := lines[1]
	fileContent := ReadFile(file.Name())
	if strings.Contains(fileContent, firstLine) {
		index := strings.Index(fileContent, firstLine)
		tempLicence := ""
		for i := index; fileContent[i-1] != '*' || fileContent[i] != '/'; i++ {
			tempLicence += string(fileContent[i])
		}
	}
	return true
}

func WriteLicenceToFile(fileName string, licenceFile string) (bool, error) {
	ok, err := CheckIfFileContainsLicenceAlready(fileName, licenceFile)
	if ok {
		return false, err
	}
	bytes, _ := ioutil.ReadFile(fileName)
	var content string
	licence := ReadFile(licenceFile)
	content = licence + string(bytes)
	_ = ioutil.WriteFile(fileName, []byte(content), 0644)
	return true, nil
}

// ReadFile ...
func ReadFile(fileName string) string {
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
