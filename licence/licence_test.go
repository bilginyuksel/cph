package licence

import (
	"fmt"
	"testing"
)

const falseLicence = `Copyright 2020. Huawei Technologies Co., Ltd. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License")
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software TIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.`

func TestCompareLicenceSimilarity_ZeroDifference(t *testing.T) {
	if 1 != comparePotentialLicenceToActualLicence(LICENCE) {
		t.Error()
	}
}

func TestCompareLicenceSimilarity_NotZeroDifference(t *testing.T) {
	difference := comparePotentialLicenceToActualLicence(falseLicence)
	t.Logf("Difference= %f", difference)
	if 0 == difference {
		t.Error(difference)
	}
}

func TestCompareLicenceSimilarity_BigDifference(t *testing.T) {
	difference := comparePotentialLicenceToActualLicence("")
	t.Logf("Difference= %f", difference)
	if 0.7 < difference {
		t.Error(difference)
	}
}

func TestMax_ReturnNum1(t *testing.T) {
	num1 := 5
	num2 := -2
	if num1 != max(num1, num2) {
		t.Error()
	}
}

func TestMax_ReturnNum2(t *testing.T) {
	num1 := -111
	num2 := -100
	if num2 != max(num1, num2) {
		t.Error()
	}
}

func TestStringArrayToMap_Sample1(t *testing.T) {
	arr := []string{"hello", "hello", "my", "name", "is", "safa"}
	ans := map[string]int{"hello": 2, "my": 1, "name": 1, "is": 1, "safa": 1}
	newMap := stringArrayToMap(arr)
	for key, value := range newMap {
		if ans[key] != value {
			t.Error()
		}
	}
}

func TestFindEndIdxOfTheBlockComment_Sample1(t *testing.T) {
	sampleContent := "/*hello world*//*hello world*/"
	endIdx := findEndIdxOfTheBlockComment(sampleContent, 1)
	if 14 != endIdx {
		t.Errorf("End Idx= %d", endIdx)
	}
}

func TestFindEndIdxOfTheBlockComment_Sample2(t *testing.T) {
	sampleContent := "/*hello world*//*hello world*/"
	endIdx := findEndIdxOfTheBlockComment(sampleContent, 16)
	if 29 != endIdx {
		t.Errorf("End Idx= %d", endIdx)
	}
}

func TestAddTagToLicence_Sample1(t *testing.T) {
	expected := fmt.Sprintf("/*\n%s\n*/", LICENCE)
	given := addTagToLicence(".java")
	if given == expected {
		t.Errorf("given= %s, expected= %s", given, expected)
	}
}

func TestAddTagToLicence_Sample2(t *testing.T) {
	expected := fmt.Sprintf("<!--\n%s\n-->", LICENCE)

	given := addTagToLicence(".html")
	if given == expected {
		t.Errorf("given= %s, expected= %s", given, expected)
	}
}

func TestAddTagToLicence_Sample3(t *testing.T) {
	expected := fmt.Sprintf("\"\"\"\n%s\n\"\"\"", LICENCE)

	given := addTagToLicence(".py")
	if given == expected {
		t.Errorf("given= %s, expected= %s", given, expected)
	}
}

func TestDeleteInvalidLicence_DeleteInvalid(t *testing.T) {
	content := "hello world/*hello world*/hello world"
	if "hello worldhello world" != deleteInvalidLicence(content, 11, 25) {
		t.Error()
	}
}

func TestFindCommentedInvalidLicenceToDelete_Similar(t *testing.T) {
	licSim := findCommentedInvalidLicenceToDelete("/*"+LICENCE+"*/", 0.8)
	t.Logf("Licsim prob= %f", licSim.prob)
	if !licSim.similar {
		t.Error()
	}
}

func TestFindCommentedInvalidLicenceToDelete_NotSimilar(t *testing.T) {
	content := "hello world/*hello world*/hello world"
	licSim := findCommentedInvalidLicenceToDelete(content, 0.1)
	t.Logf("Licsim prob= %f", licSim.prob)
	if licSim.similar {
		t.Error()
	}
}

func TestFindHowManyCommentExists_ZeroCommentJavaExt(t *testing.T) {
	content := "function test(){}"
	given := findHowManyCommentExists(content,".java")
	expected := 0
	if given != expected {
		t.Errorf("Given:%d but expected:%d", given, expected)
	}
}

func TestFindHowManyCommentExists_OneCommentJavaExt(t *testing.T) {
	content := "/* This is a comment*/ function test(){}"
	given := findHowManyCommentExists(content,".java")
	expected := 1
	if given != expected {
		t.Errorf("Given:%d but expected:%d", given, expected)
	}
}

func TestFindHowManyCommentExists_OneComment1JavaExt(t *testing.T) {
	content := "/* This is a comment*/ function test(){return '*/'}"
	given := findHowManyCommentExists(content,".java")
	expected := 1
	if given != expected {
		t.Errorf("Given:%d but expected:%d", given, expected)
	}
}
func TestFindHowManyCommentExists_OneComment2JavaExt(t *testing.T) {
	content := "/* This is a comment*/ function test(){return '/* /* */'}"
	given := findHowManyCommentExists(content,".java")
	expected := 1
	if given != expected {
		t.Errorf("Given:%d but expected:%d", given, expected)
	}
}
func TestFindHowManyCommentExists_OneComment3JavaExt(t *testing.T) {
	content := "/* This is 'a' comment*/ function test(){return \"/* /* */\"}"
	given := findHowManyCommentExists(content,".java")
	expected := 1
	if given != expected {
		t.Errorf("Given:%d but expected:%d", given, expected)
	}
}
func TestFindHowManyCommentExists_OneComment4JavaExt(t *testing.T) {
	content := "/* This is 'a' comment*/ function test(){return \"/*' /* */\"}"
	given := findHowManyCommentExists(content,".java")
	expected := 1
	if given != expected {
		t.Errorf("Given:%d but expected:%d", given, expected)
	}
}
func TestFindHowManyCommentExists_TwoCommentsJavaExt(t *testing.T) {
	content := "/* This is 'a' comment*//* This is 'a' comment*/ function test(){return \"/*/**/\"}"
	given := findHowManyCommentExists(content,".java")
	expected := 2
	if given != expected {
		t.Errorf("Given:%d but expected:%d", given, expected)
	}
}
func TestFindHowManyCommentExists_TwoComments2HtmlExt(t *testing.T) {
	content := "<!-- This is 'a' comment--><!-- This is 'a' comment--> function test(){return \"/*/**/\"}"
	given := findHowManyCommentExists(content,".html")
	expected := 2
	if given != expected {
		t.Errorf("Given:%d but expected:%d", given, expected)
	}
}
func TestFindHowManyCommentExists_TwoComments3HtmlExt(t *testing.T) {
	content := "<!-- This is 'a' comment--><!-- This is 'a' */ comment--> function test(){return \"/*/**/-->\"}"
	given := findHowManyCommentExists(content,".html")
	expected := 2
	if given != expected {
		t.Errorf("Given:%d but expected:%d", given, expected)
	}
}

