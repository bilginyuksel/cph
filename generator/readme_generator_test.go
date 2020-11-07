package generator

func createSampleReadmeResult() *TSFile{
	return &TSFile{Name: "Sample",
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

