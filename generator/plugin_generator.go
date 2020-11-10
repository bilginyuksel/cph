package generator

import (
	"fmt"
	"io/ioutil"
	"os"
)

// CreateBasePlugin ...
func CreateBasePlugin(path string, group string, projectName string) {
	createFile := func(filename string, content string) {
		d := []byte(content)
		err := ioutil.WriteFile(filename, d, 0644)
		if err != nil {
			panic(err)
		}
	}
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s", group, projectName), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src", group, projectName), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main", group, projectName), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java", group, projectName), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java/com", group, projectName), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java/com/", group, projectName), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java/com/%s", group, projectName, group), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java/com/%s/cordova", group, projectName, group), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java/com/%s/cordova/%s", group, projectName, group, projectName), 0755)

	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/www", group, projectName), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/scripts", group, projectName), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/tests", group, projectName), 0755)
	os.Mkdir(fmt.Sprintf("cordova-plugin-%s-%s/types", group, projectName), 0755)

	createFile(fmt.Sprintf("cordova-plugin-%s-%s/README.md", group, projectName), fmt.Sprintf("## cordova-plugin-%s-%s", group, projectName))
	createFile(fmt.Sprintf("cordova-plugin-%s-%s/tsconfig.json", group, projectName), `{
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

	createFile(fmt.Sprintf("cordova-plugin-%s-%s/scripts/util.ts", group, projectName),
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
	capitalizedProjectName := projectName
	if projectName[0] >= 'A' || projectName[0] <= 'Z' {
		capitalizedProjectName = string(capitalizedProjectName[0]-32) + projectName[1:]
	}
	createFile(fmt.Sprintf("cordova-plugin-%s-%s/scripts/%s.ts", group, projectName, capitalizedProjectName),
		fmt.Sprintf(`import { asyncExec } from './util';

/*
Example function to show toast in the screen. Function calls asyncExec function from util.ts file. And async function 
sends information the java to process.

@param message :: The parameter sent here will be the message of the toast showed. 

*/
export function showToast(message: string): Promise<void> {
	return asyncExec('%s', 'showToast', [message]);
}`, capitalizedProjectName))

	createFile(fmt.Sprintf("cordova-plugin-%s-%s/src/main/java/com/%s/cordova/%s/%s.java", group, projectName, group, projectName, capitalizedProjectName),
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
}`, group, projectName, capitalizedProjectName, capitalizedProjectName))

	createFile(fmt.Sprintf("cordova-plugin-%s-%s/package.json", group, projectName), fmt.Sprintf(`{
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
}`, group, projectName, group, projectName, group, projectName))
	createFile(fmt.Sprintf("cordova-plugin-%s-%s/plugin.xml", group, projectName), fmt.Sprintf(`<?xml version='1.0' encoding='utf-8'?>
	<plugin id="cordova-plugin-%s-%s"
			version="1.0.0"
			xmlns="http://apache.org/cordova/ns/plugins/1.0"
			xmlns:android="http://schemas.android.com/apk/res/android">
		<name>Cordova Plugin %s %s</name>
		<description>Cordova Plugin %s %s</description>
		<license>Apache 2.0</license>
		<keywords>android, %s, %s</keywords>

		<js-module name="%s" src="www/%s.js">
        	<clobbers target="%s"/>
    	</js-module>

		<js-module name="util" src="www/util.js"/>
	
		<engines>
			<engine name="cordova" version=">=3.0.0"/>
		</engines>

		<platform name="android">
			 <source-file src="src/main/java/com/%s/cordova/%s/%s.java" 
				target-dir="src/com/%s/cordova/%s"/>
		</platform>
		</plugin>`, group, projectName, group, projectName, group,
		projectName, group, projectName, capitalizedProjectName,
		capitalizedProjectName, capitalizedProjectName, group, projectName,
		capitalizedProjectName, group, projectName))

}
