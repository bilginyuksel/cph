package parser

import "testing"

func TestParse_ParseSample1(t *testing.T) {
	content := `

	export enum Color {
		RED = -65536,
		DARK_GRAY = -12303292,
		TRANSPARENT = 0
	}

	export interface Projection {
		fromScreenLocation(point: Point): Promise<LatLng>;
		getVisibleRegion(): Promise<VisibleRegion>;
		toScreenLocation(latLng: LatLng): Promise<Point>;
	}
	//@class This is a base Class represent map object.
	class HuaweiMapImpl implements HuaweiMap,HuaweiMap2,HuaweiMap3 {

		public readonly components: Map<string, any> = new Map<string, any>();
		private readonly id: number;
		private readonly uiSettings: UiSettings;

		/**
		* This is an interface.
		* @param callback callback Function to pass bilmem ne
		* @return any
		*/
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
				.then(Value => {
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

	function fooCreator(safa: number = 5): Promise<string> {
		return "promise<string>";
	}
	
	
	/**
	* This is an interface.
	* @param callback callback Function to pass bilmem ne
	* @return any
	*/
	function considerCase(callback: ()=>void = () => {console.log("hello world")}) {
	
	}

	function considerCase2(callback: (data1: string, data2: any) => void = (data1, data2) => {
		console.log("Hello world");
		const test = () => {
			console.log(data1);
			console.log(data2);
		}}) {
	
	}

	`
	Tokenize(content)
	given := Parse()
	expected := &TSFile{
		Classes: []Class{
			{
				Name:               "HuaweiMapImpl",
				Implements:         true,
				ImplementedClasses: []string{"HuaweiMap", "HuaweiMap2", "HuaweiMap3"},
				Variables: []Variable{
					{AccessModifier: "public", Readonly: true, Name: "components", DType: "Map<string,any>", DValue: "newMap<string,any>()"},
					{AccessModifier: "private", Readonly: true, Name: "id", DType: "number"},
					{AccessModifier: "private", Readonly: true, Name: "uiSettings", DType: "UiSettings"},
				},
				Functions: []Function{
					{Name: "scroll", Rtype: "void", Sbody: "constmapRect=document.getElementById(this.divId).getBoundingClientRect();this.forceUpdateXAndY(mapRect.x,mapRect.y);"},
					{Async: true, Name: "hideMap", Rtype: "Promise<void>", Sbody: "returnasyncExec(HMSMap,hideMap,[this.divId]);"},
					{Async: true, Name: "on", Params: []FParam{{Name: "callback", Dtype: "(val:any)=>void"}, {Name: "event", Dtype: "MapEvent"}}, Rtype: "Promise<void>", Sbody: "constfixedFunctionNameForJavaScript:string=${event}_${this.id};constfixedFunctionNameForJava:string=set${event[0].toUpperCase()}${event.substr(1)}Listener;returnasyncExec(HMSMap,mapOptions,[this.divId,setListener,fixedFunctionNameForJava,{content:callback.toString()}]).then(Value=>{//(<any>window)[fixedFunctionNameForJavaScript] = callback;window.subscribeHMSEvent(fixedFunctionNameForJavaScript,callback);}).catch(err=>console.log(err));"},
					{Async: true, Name: "addCircle", Params: []FParam{{Name: "circleOptions", Dtype: "CircleOptions"}}, Rtype: "Promise<Circle>", Sbody: "if(!circleOptions[center])returnPromise.reject(ErrorCodes.toString(ErrorCodes.CENTER_PROPERTY_MUST_DEFINED));constcomponentId=awaitasyncExec(HMSMap,addComponent,[this.divId,CIRCLE,circleOptions]);constcircle:Circle=newCircleImpl(this.divId,this.id,componentId);this.components.set(circle.getId(),circle);returncircle;"},
				},
			}},
		Functions: []Function{
			{
				Name: "demo", Rtype: "any", Sbody: "console.log(hello world);returnasyncExec(HMSMap,ssda,{dothis:nottodothis});",
				Params: []FParam{{Name: "data1", Dtype: "()=>void"}, {Name: "data", Dtype: "string"}},
			},
			{
				Name: "fooCreator", Rtype: "Promise<string>", Sbody: "returnpromise<string>;",
				Params: []FParam{{Name: "safa", Dtype: "number", Value: "5"}},
			},
			{
				Name:   "considerCase",
				Rtype:  "any",
				Params: []FParam{{Name: "callback", Dtype: "()=>void", Value: "()=>{console.log(hello world)}"}},
			},
			{
				Name:  "considerCase2",
				Rtype: "any",
				Params: []FParam{
					{
						Name:  "callback",
						Dtype: "(data1:string,data2:any)=>void",
						Value: "(data1,data2)=>{console.log(Hello world);consttest=()=>{console.log(data1);console.log(data2);}}",
					},
				},
			},
		},
		Variables: []Variable{
			{
				Name:  "foo",
				DType: "string",
			},
			{
				Name: "baz", DType: "number", DValue: "5",
			},
		},
		Interfaces: []Tinterface{{
			Name: "Projection",
			Functions: []Function{
				{Name: "fromScreenLocation", Params: []FParam{{Name: "point", Dtype: "Point"}}, Rtype: "Promise<LatLng>"},
				{Name: "getVisibleRegion", Rtype: "Promise<VisibleRegion>"},
				{Name: "toScreenLocation", Params: []FParam{{Name: "latLng", Dtype: "LatLng"}}, Rtype: "Promise<Point>"}},
		},
		},
		Enums: []Enum{
			{
				Export: true, Name: "Color",
				Items: []EnumItem{{Name: "RED", Value: "-65536"}, {Name: "DARK_GRAY", Value: "-12303292"}, {Name: "TRANSPARENT", Value: "0"}},
			},
		},
	}

	if len(given.Classes) != len(expected.Classes) {
		t.Errorf("given:%d but expected:%d", len(given.Classes), len(expected.Classes))
	}
	if len(given.Functions) != len(expected.Functions) {
		t.Errorf("given:%d but expected:%d", len(given.Functions), len(expected.Functions))
	}
	if len(given.Variables) != len(expected.Variables) {
		t.Errorf("given:%d but expected:%d", len(given.Variables), len(expected.Variables))
	}
	if len(given.Interfaces) != len(expected.Interfaces) {
		t.Errorf("given:%d but expected:%d", len(given.Interfaces), len(expected.Interfaces))
	}
	if len(given.Enums) != len(expected.Enums) {
		t.Errorf("given:%d but expected:%d", len(given.Enums), len(expected.Enums))
	}
	for i := 0; i < len(given.Interfaces); i++ {
		compareInterface(t, &given.Interfaces[i], &expected.Interfaces[i])
	}
	for i := 0; i < len(given.Enums); i++ {
		compareEnums(t, &given.Enums[i], &expected.Enums[i])
	}
	for i := 0; i < len(expected.Classes); i++ {
		compareClass(t, &given.Classes[i], &expected.Classes[i])
	}
	for i := 0; i < len(expected.Functions); i++ {
		compareFunctions(t, &given.Functions[i], &expected.Functions[i])
	}
	for i := 0; i < len(expected.Variables); i++ {
		compareVariables(t, &given.Variables[i], &expected.Variables[i])
	}
}

func compareEnums(t *testing.T, given *Enum, expected *Enum) {
	if given.Export != expected.Export {
		t.Errorf("given:%t but expected:%t", given.Export, expected.Export)
	}
	if given.Name != expected.Name {
		t.Errorf("given:%s but expected:%s", given.Name, expected.Name)
	}
	for i := 0; i < len(given.Items); i++ {
		compareEnumItem(t, &given.Items[i], &expected.Items[i])
	}
}

func compareEnumItem(t *testing.T, given *EnumItem, expected *EnumItem) {
	if given.Name != expected.Name {
		t.Errorf("given:%s but expected:%s", given.Name, expected.Name)
	}
	if given.Value != expected.Value {
		t.Errorf("given:%s but expected:%s", given.Value, expected.Value)
	}
}

func compareInterface(t *testing.T, given *Tinterface, expected *Tinterface) {
	if given.Name != expected.Name {
		t.Errorf("given:%s but expected:%s", given.Name, expected.Name)
	}
	for i := 0; i < len(given.Functions); i++ {
		compareFunctions(t, &given.Functions[i], &expected.Functions[i])
	}

	for i := 0; i < len(given.Variables); i++ {
		compareVariables(t, &given.Variables[i], &expected.Variables[i])
	}

}

func compareVariables(t *testing.T, given *Variable, expected *Variable) {
	if given.AccessModifier != expected.AccessModifier {
		t.Errorf("given:%s but expected:%s", given.AccessModifier, expected.AccessModifier)
	}
	if given.Name != expected.Name {
		t.Errorf("given:%s but expected:%s", given.Name, expected.Name)
	}
	if given.Readonly != expected.Readonly {
		t.Errorf("given:%t but expected:%t", given.Readonly, expected.Readonly)
	}
	if given.Static != expected.Static {
		t.Errorf("given:%t but expected:%t", given.Static, expected.Static)
	}
	if given.Declare != expected.Declare {
		t.Errorf("given:%t but expected:%t", given.Declare, expected.Declare)
	}
	if given.DType != expected.DType {
		t.Errorf("given:%s but expected:%s", given.DType, expected.DType)
	}
	if given.DValue != expected.DValue {
		t.Errorf("given:%s but expected:%s", given.DValue, expected.DValue)
	}
}

func compareFunctions(t *testing.T, given *Function, expected *Function) {

	if given.Export != expected.Export {
		t.Errorf("given:%t but expected:%t", given.Export, expected.Export)
	}
	if given.AccessModifier != expected.AccessModifier {
		t.Errorf("given:%s but expected:%s", given.AccessModifier, expected.AccessModifier)
	}
	if given.Name != expected.Name {
		t.Errorf("given:%s but expected:%s", given.Name, expected.Name)
	}
	if given.Async != expected.Async {
		t.Errorf("given:%t but expected:%t", given.Async, expected.Async)
	}
	if given.Static != expected.Static {
		t.Errorf("given:%t but expected:%t", given.Static, expected.Static)
	}
	if given.Declare != expected.Declare {
		t.Errorf("given:%t but expected:%t", given.Declare, expected.Declare)
	}
	if given.Rtype != expected.Rtype {
		t.Errorf("given:%s but expected:%s", given.Rtype, expected.Rtype)
	}
	if given.Sbody != expected.Sbody {
		t.Errorf("given:%s but expected:%s", given.Sbody, expected.Sbody)
	}
	for i := 0; i < len(given.Params); i++ {
		compareFunctionParameters(t, &given.Params[i], &expected.Params[i])
	}
}

func compareFunctionParameters(t *testing.T, given *FParam, expected *FParam) {
	if given.Dtype != expected.Dtype {
		t.Errorf("parameter Dtype --> given:%s but expected:%s", given.Dtype, expected.Dtype)
	}
	if given.Name != expected.Name {
		t.Errorf("parameter Name -->  given:%s but expected:%s", given.Name, expected.Name)
	}
	if given.Value != expected.Value {
		t.Errorf("parameter Value -->  given:%s but expected:%s", given.Value, expected.Value)
	}
}

func compareClass(t *testing.T, given *Class, expected *Class) {
	if given.Name != expected.Name {
		t.Errorf("given:%s but expected:%s", given.Name, expected.Name)
	}
	compareImplementedInterfacesAndExtendedClassOfClass(t, given, expected)
	compareVariablesOfClass(t, given, expected)
	compareFunctionsOfClass(t, given, expected)
}

func compareImplementedInterfacesAndExtendedClassOfClass(t *testing.T, given *Class, expected *Class) {
	if given.Extends != expected.Extends {
		t.Errorf("given:%t but expected:%t", given.Extends, expected.Extends)
	}
	if given.Implements != expected.Implements {
		t.Errorf("given:%t but expected:%t", given.Implements, expected.Implements)
	}
	if len(given.ImplementedClasses) != len(expected.ImplementedClasses) {
		t.Errorf("given:%d but expected:%d", len(given.ImplementedClasses), len(expected.ImplementedClasses))
	}
	if given.ExtendedClass != expected.ExtendedClass {
		t.Errorf("given:%s but expected:%s", given.ExtendedClass, expected.ExtendedClass)
	}
	for i := 0; i < len(given.ImplementedClasses); i++ {
		if given.ImplementedClasses[i] != expected.ImplementedClasses[i] {
			t.Errorf("given:%s but expected:%s", given.ImplementedClasses[i], expected.ImplementedClasses[i])
		}
	}
}

func compareVariablesOfClass(t *testing.T, given *Class, expected *Class) {
	if len(given.Variables) != len(expected.Variables) {
		t.Errorf("given:%d but expected:%d", len(given.Variables), len(expected.Variables))
	}
	for i := 0; i < len(given.Variables); i++ {
		compareVariables(t, &given.Variables[i], &expected.Variables[i])
	}
}

func compareFunctionsOfClass(t *testing.T, given *Class, expected *Class) {
	if len(given.Functions) != len(expected.Functions) {
		t.Errorf("given:%d but expected:%d", len(given.Functions), len(expected.Functions))
	}
	for i := 0; i < len(given.Functions); i++ {
		compareFunctions(t, &given.Functions[i], &expected.Functions[i])
	}
}
