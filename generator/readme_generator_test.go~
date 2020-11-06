package generator

import "testing"

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

	function fun4(callback: (data: string[], innerCallback: ()=>void)) {
		// function
	}
	`

	actualResult := MarshalTS(sampleTs)
	expectedResult := createSampleReadmeResult()

	if len(actualResult.Functions) != len(expectedResult.Functions) {
		t.Error()
	}

	for key, element := range actualResult.Functions {
		// iterate map
		_, hasElement := expectedResult[key]
		if !hasElement {
			t.Error()
		}
	}

}
