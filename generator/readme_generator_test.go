package generator

import "testing"


func createSampleReadmeResult() *TSFile{
	return *TSFile{Name: "Sample",
		Functions: map[string]Function{
			"fun": Function{Export:true, Name:"fun"},
			"fun1": Function{Export:true, Name:"fun1", Return: "Promise<void>",
			Parameters: []Parameter{Parameter{Name:"data", Type:"string"}, Parameter{Name:"data1", Type: "any[]"}}},
			"fun2": Function{Name:"fun2", Return: "void", Parameters: []Parameter{
			 Parameter{Name:"callback", Type:"()=>void"}},
			Annotations: []Annotation{Annotation{Name:"Annotation", Param:"{annotation:paprameter}"}}},
			"fun3": Function{Name:"fun3", Parameters: []Parameter{
				Parameter{Name:"callback1", Type:"()=>string"},
				Parameter{Name:"callback2", Type:"(data:string)=>void"},
			}},
			"fun4": Function{Name:"fun4", Parameters: []Parameter{
				Parameter{Name:"callback", Type:"(data:string[], innerCallback:()=>void)=>string)"},
			}},
		},
	}
}


func parse_expectFunction(t *testing.T){
	sampleTs := `export function fun(){
	}
	export function fun1(data: string, data1: any[]): Promise<void>{
		return asyncExec('str', 'str1', data, data1);
	}
	@Annotation({ annotation: parameter })
	function fun2(callback: () => void): void{
		
	}
	const var1 = () => {
		// function
	};

	const var2 = function(){
		// function
	}
	
	function fun3(callback1: () => string, callback2: (data: string) => void){
		// function
	}

	function fun4(callback: (data: string[], innerCallback: ()=>void)=>string) {
		// function
	}
	`

	actualResult := Parse(sampleTs)
	expectedResult := createSampleReadmeResult()

	if len(actualResult.Functions) != len(expectedResult.Functions) {
		t.Error()
	}

	for key, element := range expectedResult.Functions {
		// iterate map
		expectedElement, hasElement := actualResult[key]
		if !hasElement {
			t.Error()
		}

		// Control contents.
		if expectedElement != element {
			t.Error()
		}
	}

}
