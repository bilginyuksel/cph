package parser

import "testing"

func TestParseLoop_ParseClass(t *testing.T) {
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
	}`
	Tokenize(content)
	given := ParseLoop()
	expected := &class{
		name:               "HuaweiMapImpl",
		implements:         true,
		implementedClasses: []string{"HuaweiMap", "HuaweiMap2", "HuaweiMap3"},
		variables: []variable{{accessModifier: "public", readonly: true, name: "components", dType: "Map<string,any>", dValue: "newMap<string,any>()"},
			{accessModifier: "private", readonly: true, name: "id", dType: "number"},
			{accessModifier: "private", readonly: true, name: "uiSettings", dType: "UiSettings"}},
		functions: []function{{name: "scroll", rtype: "void", sbody: "constmapRect=document.getElementById(this.divId).getBoundingClientRect();this.forceUpdateXAndY(mapRect.x,mapRect.y);"},
			{async: true, name: "hideMap", rtype: "Promise<void>", sbody: "returnasyncExec(HMSMap,hideMap,[this.divId]);"},
			{async: true, name: "on", params: []param{{name: "event", dtype: "MapEvent"}, {name: "callback", dtype: "(val: any)=>void"}}, rtype: "Promise<void>", sbody: "constfixedFunctionNameForJavaScript:string=${event}_${this.id};constfixedFunctionNameForJava:string=set${event[0].toUpperCase()}${event.substr(1)}Listener;returnasyncExec(HMSMap,mapOptions,[this.divId,setListener,fixedFunctionNameForJava,{content:callback.toString()}]).then(value=>{(<any>window)[fixedFunctionNameForJavaScript] = callback;window.subscribeHMSEvent(fixedFunctionNameForJavaScript,callback);}).catch(err=>console.log(err));"},
			{async: true, name: "addCircle", params: []param{{name: "circleOptions", dtype: "CircleOptions"}}, rtype: "Promise<Circle>", sbody: "if(!circleOptions[center])returnPromise.reject(ErrorCodes.toString(ErrorCodes.CENTER_PROPERTY_MUST_DEFINED));constcomponentId=awaitasyncExec(HMSMap,addComponent,[this.divId,CIRCLE,circleOptions]);constcircle:Circle=newCircleImpl(this.divId,this.id,componentId);this.components.set(circle.getId(),circle);returncircle;"}},
	}
	if given.name != expected.name {
		t.Errorf("given:%s but expected:%s", given.name, expected.name)
	}
	if given.extends != expected.extends {
		t.Errorf("given:%t but expected:%t", given.extends, expected.extends)
	}
	if given.implements != expected.implements {
		t.Errorf("given:%t but expected:%t", given.implements, expected.implements)
	}
	//Compare implemented classes
	if len(given.implementedClasses) != len(expected.implementedClasses) {
		t.Errorf("given:%d but expected:%d", len(given.implementedClasses), len(expected.implementedClasses))
	}
	//Compare implemented classes one by one
	for i := 0; i < len(given.implementedClasses); i++ {
		givenCurrent := given.implementedClasses[i]
		expectedCurrent := expected.implementedClasses[i]
		if givenCurrent != expectedCurrent {
			t.Errorf("given:%s but expected:%s", givenCurrent, expectedCurrent)
		}
	}
	//Compare variables length
	if len(given.variables) != len(expected.variables) {
		t.Errorf("given:%d but expected:%d", len(given.variables), len(expected.variables))
	}
	//Compare variables one by one
	for i := 0; i < len(given.variables); i++ {
		givenCurrent := given.variables[i]
		expectedCurrent := expected.variables[i]
		if givenCurrent.accessModifier != expectedCurrent.accessModifier {
			t.Errorf("given:%s but expected:%s", givenCurrent.accessModifier, expectedCurrent.accessModifier)
		}
		if givenCurrent.name != expectedCurrent.name {
			t.Errorf("given:%s but expected:%s", givenCurrent.name, expectedCurrent.name)
		}
		if givenCurrent.readonly != expectedCurrent.readonly {
			t.Errorf("given:%t but expected:%t", givenCurrent.readonly, expectedCurrent.readonly)
		}
		if givenCurrent.static != expectedCurrent.static {
			t.Errorf("given:%t but expected:%t", givenCurrent.static, expectedCurrent.static)
		}
		if givenCurrent.declare != expectedCurrent.declare {
			t.Errorf("given:%t but expected:%t", givenCurrent.declare, expectedCurrent.declare)
		}
		if givenCurrent.dType != expectedCurrent.dType {
			t.Errorf("given:%s but expected:%s", givenCurrent.dType, expectedCurrent.dType)
		}
		if givenCurrent.dValue != expectedCurrent.dValue {
			t.Errorf("given:%s but expected:%s", givenCurrent.dValue, expectedCurrent.dValue)
		}
	}

	//Compare functions length
	if len(given.functions) != len(expected.functions) {
		t.Errorf("given:%d but expected:%d", len(given.functions), len(expected.functions))
	}

	//Compare functions one by one
	for i := 0; i < len(given.functions); i++ {
		givenCurrent := given.functions[i]
		expectedCurrent := expected.functions[i]
		if givenCurrent.export != expectedCurrent.export {
			t.Errorf("given:%t but expected:%t", givenCurrent.export, expectedCurrent.export)
		}
		if givenCurrent.accessModifier != expectedCurrent.accessModifier {
			t.Errorf("given:%s but expected:%s", givenCurrent.accessModifier, expectedCurrent.accessModifier)
		}
		if givenCurrent.name != expectedCurrent.name {
			t.Errorf("given:%s but expected:%s", givenCurrent.name, expectedCurrent.name)
		}
		if givenCurrent.async != expectedCurrent.async {
			t.Errorf("given:%t but expected:%t", givenCurrent.async, expectedCurrent.async)
		}
		if givenCurrent.static != expectedCurrent.static {
			t.Errorf("given:%t but expected:%t", givenCurrent.static, expectedCurrent.static)
		}
		if givenCurrent.declare != expectedCurrent.declare {
			t.Errorf("given:%t but expected:%t", givenCurrent.declare, expectedCurrent.declare)
		}
		if givenCurrent.rtype != expectedCurrent.rtype {
			t.Errorf("given:%s but expected:%s", givenCurrent.rtype, expectedCurrent.rtype)
		}
		if givenCurrent.sbody != expectedCurrent.sbody {
			t.Errorf("given:%s but expected:%s", givenCurrent.sbody, expectedCurrent.sbody)
		}
	}
}
