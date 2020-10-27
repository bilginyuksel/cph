package writer

import (
	"bufio"
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

	// open the file to be appended to for read
	f, err := os.Open(fileName)

	if err != nil {
		return false, err
	}

	defer f.Close()

	// make a temporary outfile
	outfile, err := os.Create("temp.java")

	if err != nil {
		return false, err
	}

	defer outfile.Close()

	// append at the start
	LICENCE := ReadFile(licenceFile)
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
