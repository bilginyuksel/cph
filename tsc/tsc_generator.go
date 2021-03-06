package tsc

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// CormetFun ...
type CormetFun struct {
	Name   string
	Params []Parameter
}

// CormetRef ...
type CormetRef struct {
	CormetList []CormetFun
	Reference  string
}

func cleanSplittedStringFromNonFunctions(strList []string) []string {
	newString := []string{}
	for _, val := range strList {
		if strings.Contains(val, "public void") {
			newString = append(newString, val)
		}
	}
	return newString
}

// GetCormetRef ...
func GetCormetRef(content string, reference string) *CormetRef {
	cormetList := []CormetFun{}
	splittedStrings := strings.Split(content, "@CordovaMethod")
	cleanStr := cleanSplittedStringFromNonFunctions(splittedStrings)
	// fmt.Println(splittedStrings)

	for _, val := range cleanStr {
		value := getCormet(val)
		// fmt.Println(value)
		cormetList = append(cormetList, *value)
	}
	return &CormetRef{CormetList: cormetList, Reference: reference}
}

// HasCormet ...
func HasCormet(content string) bool {
	return strings.Contains(content, "@CordovaMethod")
}

func countCormet(content string) int {
	return strings.Count(content, "@CordovaMethod")
}

func getCormet(content string) *CormetFun {
	methodIdx := strings.Index(content, "public void")
	name := ""
	idx := methodIdx + 12
	for ; idx < len(content); idx++ {
		if content[idx] == '(' {
			break
		}
		name += string(content[idx])
	}
	// find function body and check JSONArray var name
	paramInfoSTR := ""
	for ; idx < len(content); idx++ {
		if content[idx] == ')' || content[idx] == '(' {
			continue
		}
		if content[idx] == '{' {
			break
		}
		paramInfoSTR += string(content[idx])
	}
	jsonArrIdx := strings.Index(paramInfoSTR, "JSONArray")
	jsonArrIdx += 10
	argName := ""
	for i := jsonArrIdx; i < len(paramInfoSTR); i++ {
		if paramInfoSTR[i] == ',' || paramInfoSTR[i] == ')' {
			break
		}
		argName += string(paramInfoSTR[i])
	}
	// traverse the function body
	funBody := ""
	openBracket := 1
	for ; idx < len(content); idx++ {
		if openBracket == 0 {
			break
		}
		if content[idx] == '{' {
			openBracket++
			continue
		}
		if content[idx] == '}' {
			openBracket--
			continue
		}
		funBody += string(content[idx])
	}

	parameters := findUsagesAndReturnVarTypePairs(funBody, argName+".")
	convertParameterTypesToTS(parameters)

	return &CormetFun{Name: name, Params: parameters}
}

func createFile(filename string, content string) {
	d := []byte(content)
	err := ioutil.WriteFile(filename, d, 0644)
	if err != nil {
		panic(err)
	}
}

func createTSContentOf(ref string, cormetList []CormetFun) string {
	content := fmt.Sprintf(`
export class %s {
`, ref)
	for _, val := range cormetList {
		content += "\t" + val.Name + "("
		paramName := ""
		for i := 0; i < len(val.Params); i++ {
			content += val.Params[i].name + ":" + val.Params[i].typo
			if i != len(val.Params)-1 {
				content += ", "
			}
			paramName += ", " + val.Params[i].name
		}

		content += fmt.Sprintf("): Promise<any> {\n\t\treturn asyncExec('<class-name>', '%s', ['%s'%s]);\n\t}\n", ref, val.Name, paramName)
	}
	return content + "}\n"
}

// WriteCormetRefListToFiles ...
func WriteCormetRefListToFiles(cormetRef []CormetRef) {
	os.Mkdir("scripts", 0755)

	for _, val := range cormetRef {
		fname := fmt.Sprintf("scripts/%s.ts", val.Reference)
		fcontent := createTSContentOf(val.Reference, val.CormetList)
		createFile(fname, fcontent)
	}
}

func createGlobalTSFunctionsToSingleFile(ref string, cormetList []CormetFun) string {
	content := ""
	for _, val := range cormetList {
		content += "export function " + val.Name + "("
		paramName := ""
		for i := 0; i < len(val.Params); i++ {
			content += val.Params[i].name + ":" + val.Params[i].typo
			if i != len(val.Params)-1 {
				content += ", "
			}
			paramName += ", " + val.Params[i].name
		}

		content += fmt.Sprintf("): Promise<any> {\n\treturn asyncExec('<class-name>', '%s', ['%s'%s]);\n}\n", ref, val.Name, paramName)
	}
	return content + "\n"
}

// WriteCormetRefListToFile ...
func WriteCormetRefListToFile(filename string, cormetRef []CormetRef) {
	// os.Mkdir("scripts", 0755)

	generalContent := "import { asyncExec } from './utils';\n"
	for _, val := range cormetRef {
		fcontent := createGlobalTSFunctionsToSingleFile(val.Reference, val.CormetList)
		generalContent += fcontent
	}
	generalContent += "\n\n// EVENT REGISTERATION FUNCTION FOR SINGLE FILE\n"
	generalContent += "export function on(event: string, callback: ()=>void){\n\twindow.subscribeHMSEvent(event, callback);\n}\n"
	createFile("src/www/"+filename, generalContent)
}

func findUsagesAndReturnVarTypePairs(content string, key string) []Parameter {
	// O(len(content)*len(m)) it could be better
	params := []Parameter{}
	keyLength := len(key)
	for idx := 0; idx < len(content)-keyLength-1; idx++ {
		if content[idx] == key[0] && content[idx:idx+keyLength] == key {
			argName := getArgName(content, idx+keyLength)
			argIdx := getArgIdx(content, idx+keyLength)
			argType := getArgType(content, idx+keyLength)
			params = append(params, Parameter{argName, argIdx, argType})
			idx += keyLength
		}
	}
	return params
}

// Parameter ...
type Parameter struct {
	name string
	idx  int
	typo string
}

func getArgName(content string, idx int) string {
	// go back until you see '=' sign
	for ; idx >= 0; idx-- {
		if content[idx] == '=' {
			break
		}
	}
	idx-- // skip the equal sign
	// collect the parameter data
	name := ""
	for ; idx >= 0; idx-- {
		if content[idx] == ' ' && len(strings.Trim(name, " ")) > 0 {
			return reverse(strings.Trim(name, " "))
		}
		name += string(content[idx])
	}
	return ""
}

func reverse(content string) string {
	newContent := ""
	for i := len(content) - 1; i >= 0; i-- {
		newContent += string(content[i])
	}
	return newContent
}

func getArgIdx(content string, idx int) int {
	// find idx opening bracket
	for ; idx < len(content); idx++ {
		if content[idx] == '(' {
			break
		}
	}
	idx++ // skip opening bracket
	// go for the idx closing bracket
	argIdx := ""
	for ; idx < len(content); idx++ {
		if content[idx] == ')' {
			break
		}
		argIdx += string(content[idx])
	}
	num, err := strconv.Atoi(argIdx)
	if err != nil {
		return 99999 // set as last index
	}
	return num
}

func getArgType(content string, idx int) string {
	argType := ""
	for ; idx < len(content); idx++ {
		if content[idx] == '(' {
			break
		}
		argType += string(content[idx])
	}
	return argType
}

var javaToTSTypes = map[string]string{
	"getJSONObject": "object",
	"getJSONArray":  "any",
	"getString":     "string",
	"getInt":        "number",
	"getDouble":     "number",
	"getBoolean":    "boolean",
	"optJSONObject": "object",
	"optJSONArray":  "any",
	"optString":     "string",
	"optInt":        "string",
	"optDouble":     "number",
	"optBoolean":    "boolean",
	"get":           "any",
	"opt":           "any",
}

func convertParameterTypesToTS(params []Parameter) {
	for i := 0; i < len(params); i++ {
		params[i].typo = javaToTSTypes[params[i].typo]
	}
}
