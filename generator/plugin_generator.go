package generator

import (
	"fmt"
	"io/ioutil"
	"os"
)

// CreateBasePlugin ... Ready to develop example. Ready to run example.
func CreateBasePlugin(path string, group string, project string) {
	createFile := func(filename string, content string) {
		d := []byte(content)
		err := ioutil.WriteFile(filename, d, 0644)
		if err != nil {
			panic(err)
		}
	}

	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s", group, project), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src", group, project), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main", group, project), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java", group, project), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java/com", group, project), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java/com/", group, project), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java/com/%s", group, project, group), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java/com/%s/cordova", group, project, group), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java/com/%s/cordova/%s", group, project, group, project), 0755)
	// File...
	// os.Create("cordova-plugin-%s(group)-%s(project)/src/main/java/com/huawei/%s(group)/cordova/%s(project)/Example.java", 0755)

	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/www", group, project), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/scripts", group, project), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/tests", group, project), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/types", group, project), 0755)

	// cordova-plugin-%s(group)-%s(project)/resources/plugin.gradle
	createFile(fmt.Sprintf("cordova-plugin-%s-%s/README.md", group, project), fmt.Sprintf("## cordova-plugin-%s-%s", group, project))
	createFile(fmt.Sprintf("cordova-plugin-%s-%s/tsconfig.json", group, project), `{
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

	createFile(fmt.Sprintf("cordova-plugin-%s-%s/scripts/util.ts", group, project),
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

	createFile(fmt.Sprintf("cordova-plugin-%s-%s/scripts/%s.ts", group, project, project),
		fmt.Sprintf(`import { asyncExec } from './util';

/*
Example function to show toast in the screen. Function calls asyncExec function from util.ts file. And async function 
sends information the java to process.

@param message :: The parameter sent here will be the message of the toast showed. 

*/
export function showToast(message: string): Promise<void> {
	return asyncExec('%s', 'showToast', [message]);
}`, project))

	createFile(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java/com/%s/cordova/%s/%s.java", group, project, group, project, project),
		fmt.Sprintf(`package com.%s.cordova.%s;

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

public class %s extends CordovaPlugin {

	private final static String TAG = %s.class.getSimpleName();

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
}`, group, project, project, project))

	createFile(fmt.Sprintf("cordova-plugin-%s-%s/types/%s.d.ts", group, project, project), "export declare function showToast(message: string): Promise<void>;")
	createFile(fmt.Sprintf("cordova-plugin-%s-%s/types/util.d.ts", group, project), "export declare function asyncExec(clazz: string, func: string, args?: any[]): Promise<any>;")
	createFile(fmt.Sprintf("cordova-plugin-%s-%s/www/%s.js", group, project, project), fmt.Sprintf(`"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.showToast = void 0;
const util_1 = require("./util");
function showToast(message) {
	return util_1.asyncExec('%s', 'showToast', [message]);
}
exports.showToast = showToast;
//# sourceMappingURL=%s.js.map`, project, project))
	createFile(fmt.Sprintf("cordova-plugin-%s-%s/www/util.js", group, project), `"use strict";
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
	createFile(fmt.Sprintf("cordova-plugin-%s-%s/www/%s.js.map", group, project, project),
		fmt.Sprintf(`{"version":3,"file":"%s.js","sourceRoot":"","sources":["../scripts/%s.ts"],"names":[],"mappings":";;;AAAA,iCAAmC;AASnC,SAAgB,SAAS,CAAC,OAAe;IACxC,OAAO,gBAAS,CAAC,SAAS,EAAE,WAAW,EAAE,CAAC,OAAO,CAAC,CAAC,CAAC;AACrD,CAAC;AAFD,8BAEC","sourcesContent":["import { asyncExec } from './util';\n\n/*\nExample function to show toast in the screen. Function calls asyncExec function from util.ts file. And async function \nsends information the java to process.\n\n@param message :: The parameter sent here will be the message of the toast showed. \n\n*/\nexport function showToast(message: string): Promise<void> {\n\treturn asyncExec('%s', 'showToast', [message]);\n}"]}`, project, project, project))
	createFile(fmt.Sprintf("cordova-plugin-%s-%s/www/util.js.map", group, project), fmt.Sprintf(`{"version":3,"file":"util.js","sourceRoot":"","sources":["../scripts/util.ts"],"names":[],"mappings":";;;AAAA,qCAA+B;AAc/B,SAAgB,SAAS,CAAC,KAAa,EAAE,IAAY,EAAE,OAAc,EAAE;IACtE,OAAO,IAAI,OAAO,CAAC,CAAC,OAAO,EAAE,MAAM,EAAE,EAAE;QACtC,cAAI,CAAC,OAAO,EAAE,MAAM,EAAE,KAAK,EAAE,IAAI,EAAE,IAAI,CAAC,CAAC;IAC1C,CAAC,CAAC,CAAC;AACJ,CAAC;AAJD,8BAIC","sourcesContent":["import { exec } from 'cordova';\n\n/*\nasyncExec(clazz: string, func: string, args: any[] = []): Promise<any> \nIs a helper function to use when sending information to java with your cordova application. \n\n@param clazz :: you should write the class name you choose when creating the plugin the default is (%s(project).java). Basically you need to write the\nname of the java class which extends the CordovaPlugin. \n\n@param func :: This parameter matches with the action parameter in java files execute function. The value sent here will be captured\nas action from java.\n\n@paramm args :: You can send array of elements here to capture from java's execute method again.\n*/\nexport function asyncExec(clazz: string, func: string, args: any[] = []): Promise<any> {\n\treturn new Promise((resolve, reject) => {\n\t\texec(resolve, reject, clazz, func, args);\n\t});\n}"]}`, project))
	createFile(fmt.Sprintf("cordova-plugin-%s-%s/package.json", group, project), fmt.Sprintf(`{
	"name": "cordova-plugin-%s-%s",
	"title": "Cordova %s %s Plugin",
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
		"id": "cordova-plugin-%s-%s",
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
}`, group, project, group, project, group, project))
	createFile(fmt.Sprintf("cordova-plugin-%s-%s/plugin.xml", group, project), fmt.Sprintf(`<?xml version='1.0' encoding='utf-8'?>
	<plugin id="cordova-plugin-%s-%s"
			version="1.0.0"
			xmlns="http://apache.org/cordova/ns/plugins/1.0"
			xmlns:android="http://schemas.android.com/apk/res/android">
		<name>Cordova Plugin %s %s</name>
		<description>Cordova Plugin %s %s</description>
		<license>Apache 2.0</license>
		<keywords>android, %s, %s</keywords>
	
		<engines>
			<engine name="cordova" version=">=3.0.0"/>
		</engines>

		<platform name="android">
		</platform>
		</plugin>`, group, project, group, project, group, project, group, project))

}
