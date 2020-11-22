package parser

import "testing"

func TestParseLoop_ParseSample1(t *testing.T) {
	content := `
	class HuaweiMapImpl implements HuaweiMap,HuaweiMap2,HuaweiMap3 {

		public readonly components: Map<string, any> = new Map<string, any>();
		private readonly id: number;
		private readonly uiSettings: UiSettings;

		scroll(): void {
			const mapRect = document.getElementById(this.divId).getBoundingClientRect();
			this.forceUpdateXAndY(mapRect.x, mapRect.y);
		}

		async hideMap(): Promise<void> {
			return asyncExec("HMSMap", "hideMap", [this.divId]);
		}

		async on(event: MapEvent, callback: (val: any) => void): Promise<void> {
			const fixedFunctionNameForJavaScript: string = '${event}_${this.id}';
			const fixedFunctionNameForJava: string = 'set${event[0].toUpperCase()}${event.substr(1)}Listener';

			return asyncExec('HMSMap', 'mapOptions', [this.divId, 'setListener', fixedFunctionNameForJava, {'content': callback.toString()}])
				.then(value => {
					//(<any>window)[fixedFunctionNameForJavaScript] = callback;
					window.subscribeHMSEvent(fixedFunctionNameForJavaScript, callback);
				}).catch(err => console.log(err));
		}

		async addCircle(circleOptions: CircleOptions): Promise<Circle> {
			if (!circleOptions["center"]) return Promise.reject(ErrorCodes.toString(ErrorCodes.CENTER_PROPERTY_MUST_DEFINED));
			const componentId = await asyncExec('HMSMap', 'addComponent', [this.divId, "CIRCLE", circleOptions]);
			const circle: Circle = new CircleImpl(this.divId, this.id, componentId);
			this.components.set(circle.getId(), circle);
			return circle;
		}
	}
	function demo(data: string, data1: () => void) {
		console.log("hello world");
		return asyncExec('HMSMap', 'ssda', {'dothis':'nottodothis'});
	}
	const foo: string
	const baz: number = 5
	function fooCreator(safa: number): Promise<string> {
		return "promise<string>";
	}`
	Tokenize(content)
	given := ParseLoop()
	expected := &TSFile{
		classes: []class{
			class{
				name:               "HuaweiMapImpl",
				implements:         true,
				implementedClasses: []string{"HuaweiMap", "HuaweiMap2", "HuaweiMap3"},
				variables: []variable{{accessModifier: "public", readonly: true, name: "components", dType: "Map<string,any>", dValue: "newMap<string,any>()"},
					{accessModifier: "private", readonly: true, name: "id", dType: "number"},
					{accessModifier: "private", readonly: true, name: "uiSettings", dType: "UiSettings"}},
				functions: []function{{name: "scroll", rtype: "void", sbody: "constmapRect=document.getElementById(this.divId).getBoundingClientRect();this.forceUpdateXAndY(mapRect.x,mapRect.y);"},
					{async: true, name: "hideMap", rtype: "Promise<void>", sbody: "returnasyncExec(HMSMap,hideMap,[this.divId]);"},
					{async: true, name: "on", params: []param{{name: "callback", dtype: "(val:any)=>void"}, {name: "event", dtype: "MapEvent"}}, rtype: "Promise<void>", sbody: "constfixedFunctionNameForJavaScript:string=${event}_${this.id};constfixedFunctionNameForJava:string=set${event[0].toUpperCase()}${event.substr(1)}Listener;returnasyncExec(HMSMap,mapOptions,[this.divId,setListener,fixedFunctionNameForJava,{content:callback.toString()}]).then(value=>{(<any>window)[fixedFunctionNameForJavaScript] = callback;window.subscribeHMSEvent(fixedFunctionNameForJavaScript,callback);}).catch(err=>console.log(err));"},
					{async: true, name: "addCircle", params: []param{{name: "circleOptions", dtype: "CircleOptions"}}, rtype: "Promise<Circle>", sbody: "if(!circleOptions[center])returnPromise.reject(ErrorCodes.toString(ErrorCodes.CENTER_PROPERTY_MUST_DEFINED));constcomponentId=awaitasyncExec(HMSMap,addComponent,[this.divId,CIRCLE,circleOptions]);constcircle:Circle=newCircleImpl(this.divId,this.id,componentId);this.components.set(circle.getId(),circle);returncircle;"}},
			}},
		functions: []function{
			function{
				name: "demo", rtype: "any", sbody: "console.log(hello world);returnasyncExec(HMSMap,ssda,{dothis:nottodothis});",
				params: []param{{name: "data1", dtype: "()=>void"}, {name: "data", dtype: "string"}},
			},
			function{
				name: "fooCreator", rtype: "Promise<string>", sbody: "returnpromise<string>;",
				params: []param{{name: "safa", dtype: "number"}},
			}},
		variables: []variable{
			{name: "foo", dType: "string"},
			{name: "baz", dType: "number", dValue: "5"},
		},
	}

	if len(given.classes) != len(expected.classes) {
		t.Errorf("given:%d but expected:%d", len(given.classes), len(expected.classes))
	}
	if len(given.functions) != len(expected.functions) {
		t.Errorf("given:%d but expected:%d", len(given.functions), len(expected.functions))
	}
	if len(given.variables) != len(expected.variables) {
		t.Errorf("given:%d but expected:%d", len(given.variables), len(expected.variables))
	}
	for i := 0; i < len(expected.classes); i++ {
		compareClass(t, &given.classes[i], &expected.classes[i])
	}
	for i := 0; i < len(expected.functions); i++ {
		compareFunctions(t, &given.functions[i], &expected.functions[i])
	}
	for i := 0; i < len(expected.variables); i++ {
		compareVariables(t, &given.variables[i], &expected.variables[i])
	}
}

func compareVariables(t *testing.T, given *variable, expected *variable) {
	if given.accessModifier != expected.accessModifier {
		t.Errorf("given:%s but expected:%s", given.accessModifier, expected.accessModifier)
	}
	if given.name != expected.name {
		t.Errorf("given:%s but expected:%s", given.name, expected.name)
	}
	if given.readonly != expected.readonly {
		t.Errorf("given:%t but expected:%t", given.readonly, expected.readonly)
	}
	if given.static != expected.static {
		t.Errorf("given:%t but expected:%t", given.static, expected.static)
	}
	if given.declare != expected.declare {
		t.Errorf("given:%t but expected:%t", given.declare, expected.declare)
	}
	if given.dType != expected.dType {
		t.Errorf("given:%s but expected:%s", given.dType, expected.dType)
	}
	if given.dValue != expected.dValue {
		t.Errorf("given:%s but expected:%s", given.dValue, expected.dValue)
	}
}

func compareFunctions(t *testing.T, given *function, expected *function) {

	if given.export != expected.export {
		t.Errorf("given:%t but expected:%t", given.export, expected.export)
	}
	if given.accessModifier != expected.accessModifier {
		t.Errorf("given:%s but expected:%s", given.accessModifier, expected.accessModifier)
	}
	if given.name != expected.name {
		t.Errorf("given:%s but expected:%s", given.name, expected.name)
	}
	if given.async != expected.async {
		t.Errorf("given:%t but expected:%t", given.async, expected.async)
	}
	if given.static != expected.static {
		t.Errorf("given:%t but expected:%t", given.static, expected.static)
	}
	if given.declare != expected.declare {
		t.Errorf("given:%t but expected:%t", given.declare, expected.declare)
	}
	if given.rtype != expected.rtype {
		t.Errorf("given:%s but expected:%s", given.rtype, expected.rtype)
	}
	if given.sbody != expected.sbody {
		t.Errorf("given:%s but expected:%s", given.sbody, expected.sbody)
	}
	for i := 0; i < len(given.params); i++ {
		compareFunctionParameters(t, &given.params[i], &expected.params[i])
	}
}

func compareFunctionParameters(t *testing.T, given *param, expected *param) {
	if given.dtype != expected.dtype {
		t.Errorf("parameter dtype --> given:%s but expected:%s", given.dtype, expected.dtype)
	}

	if given.name != expected.name {
		t.Errorf("parameter name -->  given:%s but expected:%s", given.name, expected.name)
	}
}

func compareClass(t *testing.T, given *class, expected *class) {
	if given.name != expected.name {
		t.Errorf("given:%s but expected:%s", given.name, expected.name)
	}
	compareImplementedInterfacesAndExtendedClassOfClass(t, given, expected)
	compareVariablesOfClass(t, given, expected)
	compareFunctionsOfClass(t, given, expected)
}

func compareImplementedInterfacesAndExtendedClassOfClass(t *testing.T, given *class, expected *class) {
	if given.extends != expected.extends {
		t.Errorf("given:%t but expected:%t", given.extends, expected.extends)
	}
	if given.implements != expected.implements {
		t.Errorf("given:%t but expected:%t", given.implements, expected.implements)
	}
	if len(given.implementedClasses) != len(expected.implementedClasses) {
		t.Errorf("given:%d but expected:%d", len(given.implementedClasses), len(expected.implementedClasses))
	}
	if given.extendedClass != expected.extendedClass {
		t.Errorf("given:%s but expected:%s", given.extendedClass, expected.extendedClass)
	}
	for i := 0; i < len(given.implementedClasses); i++ {
		if given.implementedClasses[i] != expected.implementedClasses[i] {
			t.Errorf("given:%s but expected:%s", given.implementedClasses[i], expected.implementedClasses[i])
		}
	}
}

func compareVariablesOfClass(t *testing.T, given *class, expected *class) {
	if len(given.variables) != len(expected.variables) {
		t.Errorf("given:%d but expected:%d", len(given.variables), len(expected.variables))
	}
	for i := 0; i < len(given.variables); i++ {
		compareVariables(t, &given.variables[i], &expected.variables[i])
	}
}

func compareFunctionsOfClass(t *testing.T, given *class, expected *class) {
	if len(given.functions) != len(expected.functions) {
		t.Errorf("given:%d but expected:%d", len(given.functions), len(expected.functions))
	}
	for i := 0; i < len(given.functions); i++ {
		compareFunctions(t, &given.functions[i], &expected.functions[i])
	}
}
