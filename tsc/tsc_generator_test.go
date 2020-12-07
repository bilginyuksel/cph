package tsc

import "testing"

func TestHasCordovaMethod_NoCordovaMethod(t *testing.T) {
	given := HasCormet("content")
	expected := false
	if given != expected {
		t.Error()
	}
}

func TestHasCordovaMethod_CordovaMethodExists(t *testing.T) {
	content := `
    @CordovaMethod
    public void something(final CorPack corPack, JSONArray args, final Promise promise) {
        promise.success();
    }`
	given := HasCormet(content)
	expected := true
	if given != expected {
		t.Error()
	}
}

func TestCountCordovaMethod_FindTwoCordovaMethods(t *testing.T) {
	content := `
    @CordovaMethod
    public void something(final CorPack corPack, JSONArray args, final Promise promise){
        promise.success();
    }

    @CordovaMethod
    public void something(final CorPack corPack, JSONArray args, final Promise promise){
        promise.success();
    }
    `
	given := countCormet(content)
	expected := 2
	if expected != given {
		t.Error()
	}
}

func TestGetCormet_GetMethodInfoSample1(t *testing.T) {
	content := `
    @CordovaMethod
    public void something(final CorPack corPack, JSONArray args, final Promise promise){
        promise.success();
    }
    `
	given := getCormet(content)
	expected := &CormetFun{Name: "something"}
	if !isCormetFunctionsSame(given, expected) {
		t.Errorf("given: %v, expected: %v", given, expected)
	}
}

func TestGetCormet_GetMethodInfoSample2(t *testing.T) {
	content := `
    @CordovaMethod
    public void myCormet(final CorPack corPack, JSONArray args, final Promise promise){
        promise.success();
    }
    `
	given := getCormet(content)
	expected := &CormetFun{Name: "myCormet"}
	if !isCormetFunctionsSame(given, expected) {
		t.Errorf("given: %v, expected: %v", given, expected)
	}
}

func TestGetCormet_GetMethodInfoWithParameters(t *testing.T) {
	content := `
    @CordovaMethod
    public void myCormet(final CorPack corPack, JSONArray args, final Promise promise) {
        String id = args.getString(0);
        int no = args.getInt(1);
        promise.success();
    }
    `
	given := getCormet(content)
	expected := &CormetFun{Name: "myCormet", Params: []Parameter{{"id", 0, "string"},
		{"no", 1, "number"}}}
	if !isCormetFunctionsSame(given, expected) {
		t.Error()
	}
}

func TestGetCormet_GetMethodInfoWithParameters2(t *testing.T) {
	content := `
    @CordovaMethod
    public void myCormet(final CorPack corPack, JSONArray args, final Promise promise) {
        String id = args.getString(0);
        int no = args.getInt(1);
        float obj = (float) args.getDouble(3);
        JSONObject myObj = args.getJSONObject(2);
        promise.success();
    }
    `
	given := getCormet(content)
	expected := &CormetFun{Name: "myCormet", Params: []Parameter{{"id", 0, "string"},
		{"no", 1, "number"}, {"obj", 3, "number"}, {"myObj", 2, "object"}}}
	if !isCormetFunctionsSame(given, expected) {
		t.Error()
	}
}

func TestGetCormetRef_GetMultipleMethods(t *testing.T) {
	content := `
    @CordovaMethod
    public void met1(final CorPack corPack, JSONArray args, final Promise promise){
        String id = args.getString(0);
        int no = args.getInt(1);
        promise.success();
    }

    @CordovaMethod
    public void myCormet(final CorPack corPack, JSONArray args, final Promise promise) {
        String id = args.getString(0);
        int no = args.getInt(1);
        float obj = (float) args.getDouble(3);
        JSONObject myObj = args.getJSONObject(2);
        promise.success();
    }

    @CordovaMethod
    public void emptyBody(final CorPack corPack, JSONArray args, final Promise promise) {
        promise.success();
    }
    `
	given := GetCormetRef(content, "MyReference")
	cormetList := []CormetFun{
		CormetFun{Name: "met1", Params: []Parameter{{"id", 0, "string"}, {"no", 1, "number"}}},
		CormetFun{Name: "myCormet", Params: []Parameter{{"id", 0, "string"}, {"no", 1, "number"}, {"obj", 3, "number"}, {"myObj", 2, "object"}}},
		CormetFun{Name: "emptyBody"},
	}
	expected := &CormetRef{CormetList: cormetList, Reference: "MyReference"}
	if given.Reference != expected.Reference {
		t.Error()
	}
	for i := 0; i < len(expected.CormetList); i++ {
		if !isCormetFunctionsSame(&given.CormetList[i], &expected.CormetList[i]) {
			t.Errorf("given: %v, expected: %v", given.CormetList[i], expected.CormetList[i])
		}
	}

}

func isCormetFunctionsSame(fun1 *CormetFun, fun2 *CormetFun) bool {
	if fun1.Name != fun2.Name {
		return false
	}

	if len(fun1.Params) != len(fun2.Params) {
		return false
	}

	for key, value := range fun1.Params {
		if fun2.Params[key] != value {
			return false
		}
	}

	return true
}
