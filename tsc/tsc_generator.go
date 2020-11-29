package generator

import (
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

func hasCormet(content string) bool {
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
	"getJSONArray":  "Array",
	"getString":     "string",
	"getInt":        "number",
	"getDouble":     "number",
	"getBoolean":    "boolean",
	"optJSONObject": "object",
	"optJSONArray":  "Array",
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
