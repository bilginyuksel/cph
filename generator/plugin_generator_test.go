package generator

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreatePluginWithGivenName_ExpectPlugin(t *testing.T) {
	CreateBasePlugin(".", "group", "project")

	createFile := func(filename string, content string) {
		d := []byte(content)
		err := ioutil.WriteFile(filename, d, 0644)
		if err != nil {
			t.Error(err)
		}
	}

	os.Mkdir("cordova-plugin-group-project", 0755)
	os.Mkdir("cordova-plugin-group-project/src", 0755)
	os.Mkdir("cordova-plugin-group-project/src/main", 0755)
	os.Mkdir("cordova-plugin-group-project/src/main/java", 0755)
	os.Mkdir("cordova-plugin-group-project/src/main/java/com", 0755)
	os.Mkdir("cordova-plugin-group-project/src/main/java/com/", 0755)
	os.Mkdir("cordova-plugin-group-project/src/main/java/com/group", 0755)
	os.Mkdir("cordova-plugin-group-project/src/main/java/com/group/cordova", 0755)
	os.Mkdir("cordova-plugin-group-project/src/main/java/com/group/cordova/project", 0755)
	// File...
	// os.Create("cordova-plugin-group-project/src/main/java/com/huawei/group/cordova/project/Example.java", 0755)

	os.Mkdir("cordova-plugin-group-project/www", 0755)
	os.Mkdir("cordova-plugin-group-project/scripts", 0755)
	os.Mkdir("cordova-plugin-group-project/tests", 0755)
	os.Mkdir("cordova-plugin-group-project/types", 0755)

	// cordova-plugin-group-project/resources/plugin.gradle
	createFile("cordova-plugin-group-project/README.md", "## cordova-plugin-group-project")
	createFile("cordova-plugin-group-project/tsconfig.json", `{
	"compileOnSave": true,
	"compilerOptions": {
		"noImplicitAny": true,
		"noEmitOnError": true,
		"removeComments": false,
		"sourceMap": true,
		"inlineSources": true,
		"outDir": "www",
		"module": "commonjs",
		"target": "es2015",
		"declaration": true,
		"declarationDir": "types"
	},
	"exclude": ["node_modules", "src", "www", "types"]
}`)

	createFile("cordova-plugin-group-project/scripts/util.ts",
		`import { exec } from 'cordova';

/*
asyncExec(clazz: string, func: string, args: any[] = []): Promise<any> 
Is a helper function to use when sending information to java with your cordova application. 

@param clazz :: you should write the class name you choose when creating the plugin the default is (Project.java). Basically you need to write the
name of the java class which extends the CordovaPlugin. 

@param func :: This parameter matches with the action parameter in java files execute function. The value sent here will be captured
as action from java.

@paramm args :: You can send array of elements here to capture from java's execute method again.
*/
export function asyncExec(clazz: string, func: string, args: any[] = []): Promise<any> {
	return new Promise((resolve, reject) => {
		exec(resolve, reject, clazz, func, args);
	});
}`)

	createFile("cordova-plugin-group-project/scripts/Project.ts",
		`import { asyncExec } from './util';

/*
Example function to show toast in the screen. Function calls asyncExec function from util.ts file. And async function 
sends information the java to process.

@param message :: The parameter sent here will be the message of the toast showed. 

*/
export function showToast(message: string): Promise<void> {
	return asyncExec('Project', 'showToast', [message]);
}`)

	createFile("cordova-plugin-group-project/src/main/java/com/group/cordova/project/Project.java",
		`package com.group.cordova.project;

import android.util.Log;
import android.widget.Toast;

import org.json.JSONArray;
import org.json.JSONException;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

import org.apache.cordova.CallbackContext;
import org.apache.cordova.CordovaInterface;
import org.apache.cordova.CordovaPlugin;
import org.apache.cordova.CordovaWebView;

public class Project extends CordovaPlugin {

	private final static String TAG = Project.class.getSimpleName();

	@Override
	public void initialize(CordovaInterface cordova, CordovaWebView webView) {
		super.initialize(cordova, webView);
	}

	@Override
	public boolean execute(String action, JSONArray args, CallbackContext callbackContext) throws JSONException {
		try {
			Method method = this.getClass().getDeclaredMethod(action, JSONArray.class, CallbackContext.class);
			method.invoke(this, args, callbackContext);
		} catch(NoSuchMethodException | InvocationTargetException | IllegalAccessException e) {
			Log.e(TAG, e.getMessage());
		}
		return true;
	}

	private void showToast(JSONArray args, final CallbackContext callbackContext) {
		String message = args.optString(0);
		Toast.makeText(webView.getContext(), message, Toast.LENGTH_SHORT).show();
		callbackContext.success();
	}
}`)

	createFile("cordova-plugin-group-project/types/Project.d.ts", "export declare function showToast(message: string): Promise<void>;")
	createFile("cordova-plugin-group-project/types/util.d.ts", "export declare function asyncExec(clazz: string, func: string, args?: any[]): Promise<any>;")
	createFile("cordova-plugin-group-project/www/Project.js", `"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.showToast = void 0;
const util_1 = require("./util");
function showToast(message) {
	return util_1.asyncExec('Project', 'showToast', [message]);
}
exports.showToast = showToast;
//# sourceMappingURL=Project.js.map`)
	createFile("cordova-plugin-group-project/www/util.js", `"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.asyncExec = void 0;
const cordova_1 = require("cordova");
function asyncExec(clazz, func, args = []) {
	return new Promise((resolve, reject) => {
		cordova_1.exec(resolve, reject, clazz, func, args);
	});
}
exports.asyncExec = asyncExec;
//# sourceMappingURL=util.js.map`)
	createFile("cordova-plugin-group-project/www/Project.js.map", `{"version":3,"file":"Project.js","sourceRoot":"","sources":["../scripts/Project.ts"],"names":[],"mappings":";;;AAAA,iCAAmC;AASnC,SAAgB,SAAS,CAAC,OAAe;IACxC,OAAO,gBAAS,CAAC,SAAS,EAAE,WAAW,EAAE,CAAC,OAAO,CAAC,CAAC,CAAC;AACrD,CAAC;AAFD,8BAEC","sourcesContent":["import { asyncExec } from './util';\n\n/*\nExample function to show toast in the screen. Function calls asyncExec function from util.ts file. And async function \nsends information the java to process.\n\n@param message :: The parameter sent here will be the message of the toast showed. \n\n*/\nexport function showToast(message: string): Promise<void> {\n\treturn asyncExec('Project', 'showToast', [message]);\n}"]}`)
	createFile("cordova-plugin-group-project/www/util.js.map", `{"version":3,"file":"util.js","sourceRoot":"","sources":["../scripts/util.ts"],"names":[],"mappings":";;;AAAA,qCAA+B;AAc/B,SAAgB,SAAS,CAAC,KAAa,EAAE,IAAY,EAAE,OAAc,EAAE;IACtE,OAAO,IAAI,OAAO,CAAC,CAAC,OAAO,EAAE,MAAM,EAAE,EAAE;QACtC,cAAI,CAAC,OAAO,EAAE,MAAM,EAAE,KAAK,EAAE,IAAI,EAAE,IAAI,CAAC,CAAC;IAC1C,CAAC,CAAC,CAAC;AACJ,CAAC;AAJD,8BAIC","sourcesContent":["import { exec } from 'cordova';\n\n/*\nasyncExec(clazz: string, func: string, args: any[] = []): Promise<any> \nIs a helper function to use when sending information to java with your cordova application. \n\n@param clazz :: you should write the class name you choose when creating the plugin the default is (Project.java). Basically you need to write the\nname of the java class which extends the CordovaPlugin. \n\n@param func :: This parameter matches with the action parameter in java files execute function. The value sent here will be captured\nas action from java.\n\n@paramm args :: You can send array of elements here to capture from java's execute method again.\n*/\nexport function asyncExec(clazz: string, func: string, args: any[] = []): Promise<any> {\n\treturn new Promise((resolve, reject) => {\n\t\texec(resolve, reject, clazz, func, args);\n\t});\n}"]}`)
	createFile("cordova-plugin-group-project/package.json", `{
	"name": "cordova-plugin-group-project",
	"title": "Cordova Group Project Plugin",
	"version":"1.0.0",
	"repository": {
		"type": "git",
		"url": ""
	},
	"licence": "Apache-2.0",
	"keywords": [
		"cordova"
	],
	"cordova": {
		"id": "cordova-plugin-group-project",
		"platforms": [
			"android"
		]
	},
	"engines": {
		"name": "cordova",
		"version": ">=3.0.0"
	},
	"devDependencies": {
		"eslint": "^3.19.0",
		"eslint-config-semistandard": "^11.0.0",
		"eslint-config-standard": "^10.2.1",
		"eslint-plugin-import": "^2.3.0",
		"eslint-plugin-node": "^5.0.0",
		"eslint-plugin-promise": "^3.5.0",
		"eslint-plugin-standard": "^3.0.1"
	},
	"dependencies": {},
	"files": [
		"src/**",
		"www/**",
		"LICENCE",
		"package.json",
		"plugin.xml",
		"README.md"
	]
}`)
	createFile("cordova-plugin-group-project/plugin.xml", "")

}
