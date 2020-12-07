package parser

import "testing"

func TestTokenize_Variables1(t *testing.T) {
	content := `const foo  ="Life \"is all about quotes\"";let l;     var baz:string; const koz:   number   =     5;var soz; var loz = 'hi'`

	expected := []string{"const", "foo", "=", "Life \"is all about quotes\"", ";", "let", "l", ";", "var", "baz", ":", "string", ";", "const", "koz", ":", "number", "=", "5", ";", "var", "soz", ";", "var", "loz", "=", "hi"}
	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}
}

func TestTokenize_Variables2(t *testing.T) {
	content := `const f='\'\'hello\'\'';let l="\"\"world\"\"";`

	expected := []string{"const", "f", "=", "''hello''", ";", "let", "l", "=", "\"\"world\"\"", ";"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}
}

func TestTokenize_Variables3(t *testing.T) {
	content := `const f: string="hello world"
                const l="wow"
let yey; let how: string
var v
var l = "lo\"l\"ol";let l;`

	expected := []string{"const", "f", ":", "string", "=", "hello world", "const", "l", "=", "wow", "let",
		"yey", ";", "let", "how", ":", "string", "var", "v", "var", "l", "=", "lo\"l\"ol", ";", "let", "l", ";"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}

}

func TestTokenize_VariablesAndComments1(t *testing.T) {
	content := `const f: string="hello world"
    /*This is my multiline comment
        This is my multiline comment
        This is my multiline comment
        */
                const l="wow"
                // This is my first comment
let yey; let how: string
var v
// This is my second comment
var l = "lo\"l\"ol";let l;`
	multilineComment := `This is my multiline comment
        This is my multiline comment
        This is my multiline comment
        `
	expected := []string{"const", "f", ":", "string", "=", "hello world", "/*", multilineComment, "const", "l", "=", "wow", "//", " This is my first comment", "let",
		"yey", ";", "let", "how", ":", "string", "var", "v", "//", " This is my second comment", "var", "l", "=", "lo\"l\"ol", ";", "let", "l", ";"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}

}

func TestTokenize_VariablesAndAnnotations1(t *testing.T) {
	content := `const f: string="hello world"
                const l="wow"
@NotNull
let yey; let how: string
@Nullable()
var v
var l = "lo\"l\"ol";let l;`

	expected := []string{"const", "f", ":", "string", "=", "hello world", "const", "l", "=", "wow", "@", "NotNull", "let",
		"yey", ";", "let", "how", ":", "string", "@", "Nullable", "(", ")", "var", "v", "var", "l", "=", "lo\"l\"ol", ";", "let", "l", ";"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}

}

func TestTokenize_VariablesAndAnnotationsAndComments1(t *testing.T) {
	content := `const f: string="hello world"
                const l="wow"
@NotNull
let yey; let how: string //comment1
@Nullable() //comment2
var v
/*Hello world_This is my multilecomment*/
var l = "lo\"l\"ol";let l;`

	expected := []string{"const", "f", ":", "string", "=", "hello world", "const", "l", "=", "wow", "@", "NotNull", "let",
		"yey", ";", "let", "how", ":", "string", "//", "comment1", "@", "Nullable", "(", ")", "//", "comment2", "var", "v", "/*", "Hello world_This is my multilecomment", "var", "l", "=", "lo\"l\"ol", ";", "let", "l", ";"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}

}

func TestTokenize_Functions1(t *testing.T) {
	content := `function foo(){
        const foo = "foo";
        let va:string = "lol";
        helloWorld();
    }
    export function baz(param1: number, param2: string) string{
        return param1.toString() + param2;	
}`

	expected := []string{"function", "foo", "(", ")", "{", "const", "foo", "=", "foo", ";", "let", "va", ":", "string", "=", "lol", ";", "helloWorld", "(", ")", ";", "}", "export", "function", "baz", "(", "param1", ":", "number", ",", "param2", ":", "string", ")", "string", "{", "return", "param1.toString", "(", ")", "+", "param2", ";", "}"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}
}

func TestTokenize_Functions2(t *testing.T) {
	content := `function foo(){
        const foo = "foo";
        let va:string = "lol";
        helloWorld();
    }
    export function baz(param1: number, param2: string) string{
        return param1.toString() + param2;	
    }
    function test(param: string, callback: ()=>void){}
`

	expected := []string{"function", "foo", "(", ")", "{", "const", "foo", "=", "foo", ";", "let", "va", ":", "string", "=", "lol", ";", "helloWorld", "(", ")", ";", "}", "export", "function", "baz", "(", "param1", ":", "number", ",", "param2", ":", "string", ")", "string", "{", "return", "param1.toString", "(", ")", "+", "param2", ";", "}", "function", "test", "(", "param", ":", "string", ",", "callback", ":", "(", ")", "=", ">", "void", ")", "{", "}"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}
}
func TestTokenize_Functions3(t *testing.T) {
	content := `
                function test(param: string, callback: (param: string) => void) {
                  callback("Test")
                }

                test("Test", function (param: string) {
                  console.log(param)
                })
`

	expected := []string{"function", "test", "(", "param", ":", "string", ",", "callback",
		":", "(", "param", ":", "string", ")", "=", ">", "void", ")", "{", "callback", "(",
		"Test", ")", "}", "test", "(", "Test", ",", "function", "(", "param", ":", "string", ")",
		"{", "console.log", "(", "param", ")", "}", ")"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}
}
func TestTokenize_Class1(t *testing.T) {
	content := `
                export class TestClass{
                  f: string="hello world"
                  l="wow"
                  @NotNull
                  yey:number;how: string
                  @Nullable()
                  v
                  async test():Promise<string>{
                    return "Test"
                  }
                }
`

	expected := []string{"export", "class", "TestClass", "{", "f", ":", "string",
		"=", "hello world", "l", "=", "wow", "@", "NotNull", "yey", ":", "number", ";",
		"how", ":", "string", "@", "Nullable", "(", ")", "v", "async", "test", "(", ")", ":",
		"Promise", "<", "string", ">", "{", "return", "Test", "}", "}"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}
}
func TestTokenize_Class2(t *testing.T) {
	content := `
                export abstract class TestClass{
                  f: string="hello world"
                  l="wow"
                  @NotNull
                  yey:number;how: string
                  @Nullable()
                  v
                  async test():Promise<string>{
                    await return "Test"
                  }
                
                  async abstract test2();
                
                }
`

	expected := []string{"export", "abstract", "class", "TestClass", "{", "f", ":", "string",
		"=", "hello world", "l", "=", "wow", "@", "NotNull", "yey", ":", "number", ";",
		"how", ":", "string", "@", "Nullable", "(", ")", "v", "async", "test", "(", ")", ":",
		"Promise", "<", "string", ">", "{", "await", "return", "Test", "}", "async",
		"abstract", "test2", "(", ")", ";", "}"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}
}
func TestTokenize_Interface1(t *testing.T) {
	content := `
                export interface Test {
                  param: string,
                  param2: number,
                  method1: () => string,
                  method2: (param: string) => void
                }
`

	expected := []string{"export", "interface", "Test", "{", "param", ":", "string",
		",", "param2", ":", "number", ",", "method1", ":", "(", ")", "=", ">", "string",
		",", "method2", ":", "(", "param", ":", "string", ")", "=", ">", "void", "}"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}
}
func TestTokenize_Interface2(t *testing.T) {
	content := `
                export interface Test {
                  test();
                  test2(param: string);
                  test3(): string;
                  test4(): () => void;
                  test5(param: string): () => string;
                  test6(callback: () => void): () => string;
                }
`

	expected := []string{"export", "interface", "Test", "{", "test", "(", ")", ";",
		"test2", "(", "param", ":", "string", ")", ";", "test3", "(", ")", ":", "string",
		";", "test4", "(", ")", ":", "(", ")", "=", ">", "void", ";", "test5", "(", "param",
		":", "string", ")", ":", "(", ")", "=", ">", "string", ";", "test6", "(", "callback",
		":", "(", ")", "=", ">", "void", ")", ":", "(", ")", "=", ">", "string", ";", "}"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}
}

func TestTokenize_Mix1(t *testing.T) {

	content := `import mm from './file';
    import {mm1, mm2, mm3} from './file2'
    export {mm1, mm2} from './file2';
    abstract class Foo {
        @Test("hello", "myHello")
        /*
        Docstring line-1
        Docstring line-2
        */
        abstract method1(param: string, param2: number): Promise<void>;
        method2(param): Promise<string> {
            doSomething();
            return new Promise((resolve, reject) => {
                console.log("hello world");
            });
        }
    }
    const globalParam: string = "print me";
    export function doSomething() {
        console.log("I did something");
        console.log(globalParam);
    }

    interface Data {
        foo: number;
        baz: string
    }

    function intertest(data: Data) : Data{
        let newData: data = {} as Data;
        if(data.foo) newData.foo = data.foo;
        if(data.baz) newData.baz = data.baz;
        return newData;
    }

    @ConstantAnnotation
    enum Constant {
        CONSTANT1="hello",
        CONSTANT2="world"
    }
    `

	expected := []string{"import", "mm", "from", "./file", ";", "import", "{", "mm1", ",", "mm2", ",", "mm3", "}", "from", "./file2",
		"export", "{", "mm1", ",", "mm2", "}", "from", "./file2", ";", "abstract", "class", "Foo", "{", "@", "Test", "(", "hello", ",", "myHello", ")", "/*", `
        Docstring line-1
        Docstring line-2
        `, "abstract", "method1", "(", "param", ":", "string", ",", "param2", ":", "number", ")", ":", "Promise", "<", "void", ">", ";",
		"method2", "(", "param", ")", ":", "Promise", "<", "string", ">", "{", "doSomething", "(", ")", ";", "return", "new", "Promise", "(",
		"(", "resolve", ",", "reject", ")", "=", ">", "{", "console.log", "(", "hello world", ")", ";", "}", ")", ";", "}", "}",
		"const", "globalParam", ":", "string", "=", "print me", ";", "export", "function", "doSomething", "(", ")", "{", "console.log",
		"(", "I did something", ")", ";", "console.log", "(", "globalParam", ")", ";", "}", "interface", "Data", "{", "foo", ":",
		"number", ";", "baz", ":", "string", "}", "function", "intertest", "(", "data", ":", "Data", ")", ":", "Data", "{", "let", "newData", ":", "data", "=",
		"{", "}", "as", "Data", ";", "if", "(", "data.foo", ")", "newData.foo", "=", "data.foo", ";", "if", "(", "data.baz", ")",
		"newData.baz", "=", "data.baz", ";", "return", "newData", ";", "}", "@", "ConstantAnnotation", "enum", "Constant", "{",
		"CONSTANT1", "=", "hello", ",", "CONSTANT2", "=", "world", "}"}

	given := Tokenize(content)

	if len(given) != len(expected) {
		t.Errorf("Length of the expected array is not satisfied. expected: %d, given: %d", len(expected), len(given))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != given[i] {
			t.Errorf(`tokens should be %s at index=%d but given is %s`, expected[i], i, given[i])
		}
	}
}
