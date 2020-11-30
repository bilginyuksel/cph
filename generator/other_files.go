package generator

const PACKAGE_JSON = `{
	"name": "%s",
	"title": "Cordova %s Plugin",
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
		"id": "%s",
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
}`

const TS_CONFIG = `{
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
}`

const PLUGIN_XML = `<?xml version='1.0' encoding='utf-8'?>
<plugin id="%s"
		version="1.0.0"
		xmlns="http://apache.org/cordova/ns/plugins/1.0"
		xmlns:android="http://schemas.android.com/apk/res/android">
	<name>Cordova Plugin %s</name>
	<description>Cordova Plugin %s</description>
	<license>Apache 2.0</license>
	<keywords>android, hms</keywords>

	<platform name="android">
	</platform>
    </plugin>`

const TS_UTILS = `import {exec} from 'cordova';

export function asyncExec(clazz: string, reference: string, args:any[] = []): Promise<any> {
    return new Promise((resolve, reject) => {
        exec(resolve, reject, clazz, reference, args);
    });
}

declare global {
    interface Window {
        hmsEvents: {
            [key: string]: (data: any) => void
        },
        runHMSEvent: (eventName: string, data: any) => void,
        subscribeHMSEvent: (eventName: string, callback: (data: any) => void) => void
        [key: string]: any
    }
}

function initEventHandler() {
    if (window.hmsEvents != null) return;
    window.hmsEvents = {};
    window.runHMSEvent = (eventName, data) => {
        if (window.hmsEvents.hasOwnProperty(eventName))
            window.hmsEvents[eventName](data);
    };
    window.subscribeHMSEvent = (eventName, handler) => {
        window.hmsEvents[eventName] = handler;
    };
}

initEventHandler()
`

const JS_MAIN = `const cordova = require("cordova");
function asyncExec(clazz, ref, args = []) {
    return new Promise((resolve, reject) => {
        cordova.exec(resolve, reject, clazz, ref, args);
    });
}
exports.asyncExec = asyncExec;
function initEventHandler() {
    if (window.hmsEvents != null)
        return;
    window.hmsEvents = {};
    window.runHMSEvent = (eventName, data) => {
        if (window.hmsEvents.hasOwnProperty(eventName))
            window.hmsEvents[eventName](data);
    };
    window.subscribeHMSEvent = (eventName, handler) => {
        window.hmsEvents[eventName] = handler;
    };
}
initEventHandler();

function showToast(msg) {
	return asyncExec('%s', 'Test', ['showToast', msg]);
}
exports.showToast = showToast;
`

const ANDROID_MANIFEST = `<?xml version='1.0' encoding='utf-8'?>
<manifest android:hardwareAccelerated="true"  package="com.huawei.hms.cordova.%s" xmlns:android="http://schemas.android.com/apk/res/android">
    <uses-permission android:name="android.permission.INTERNET" />
</manifest>`

const AFTER_PLUGIN_INSTALL = `'use strict';

var FSUtils = require('./FSUtils');

var ROOT_GRADLE_FILE = 'platforms/android/build.gradle';
var COMMENT = '//This line is added by cordova-plugin-hms-map plugin'
var NEW_LINE = '\n';

module.exports = function (context) {

    if (!FSUtils.exists(ROOT_GRADLE_FILE)) {
        console.log('root gradle file does not exist. after_plugin_install script wont be executed.');
    }

    var rootGradleContent = FSUtils.readFile(ROOT_GRADLE_FILE, 'UTF-8');
    var lines = rootGradleContent.split(NEW_LINE);

    var depAddedLines = addAGConnectDependency(lines);
    var repoAddedLines = addHuaweiRepo(depAddedLines);

    FSUtils.writeFile(ROOT_GRADLE_FILE, repoAddedLines.join(NEW_LINE));

}

function addAGConnectDependency(lines) {
    var AG_CONNECT_DEPENDENCY = 'classpath \'com.huawei.agconnect:agcp:1.4.1.300\' ' + COMMENT;

    var pattern = /(\s*)classpath(\s+)\'com.android.tools.build:gradle:([0-9-\.\:]+)/m;

    var index;
    for (var i = 0; i < lines.length; i++) {
        var line = lines[i];
        if (pattern.test(line)) {
            index = i;
            break;
        }
    }

    lines.splice(index + 1, 0, AG_CONNECT_DEPENDENCY);
    return lines;
}


function addHuaweiRepo(lines) {
    var HUAWEI_REPO = 'maven { url \'https://developer.huawei.com/repo/\' } ' + COMMENT
    var pattern = /(\s*)jcenter\(\)/m;

    var indexList = [];
    for (var i = 0; i < lines.length; i++) {
        var line = lines[i];
        if (pattern.test(line)) {
            indexList.push(i);
        }
    }

    for (var i = 0; i < indexList.length; i++) {
        lines.splice(indexList[i] + 1, 0, HUAWEI_REPO);
        if (i < indexList.length - 1) {
            indexList[i + 1] = indexList[i + 1] + 1;
        }
    }
    return lines;
}`

const AFTER_PREPARE = `'use strict';

var FSUtils = require('./FSUtils');

var DEST_DIR = 'platforms/android/app/';
var FILE_NAME = 'agconnect-services.json';


module.exports = function (context) {
    var platforms = context.opts.platforms;

    if (platforms.includes('android')) {
        copyAGConnect();
    }
}

function copyAGConnect() {

    if (!FSUtils.exists(FILE_NAME)) {
        console.log('agconnect-services.json does not exists!');
        return;
    }

    if (!FSUtils.exists(DEST_DIR)) {
        console.log('destination does not exist. dest : ' + DEST_DIR);
        return;
    }

    FSUtils.copyFile(FILE_NAME, DEST_DIR + FILE_NAME);
}
`

const BEFORE_PLUGIN_UNINSTALL = `
'use strict';

var FSUtils = require('./FSUtils');

var ROOT_GRADLE_FILE = 'platforms/android/build.gradle';
var COMMENT = '//This line is added by cordova-plugin-hms-map plugin'
var NEW_LINE = '\n';

module.exports = function (context) {

    if (!FSUtils.exists(ROOT_GRADLE_FILE)) {
        console.log('root gradle file does not exist. before_plugin_uninstall script wont be executed.');
    }

    var rootGradleContent = FSUtils.readFile(ROOT_GRADLE_FILE, 'UTF-8');
    var lines = rootGradleContent.split(NEW_LINE);

    var linesAfterRemove = removeLinesAddedByPlugin(lines);

    FSUtils.writeFile(ROOT_GRADLE_FILE, linesAfterRemove.join(NEW_LINE));

}


function removeLinesAddedByPlugin(lines) {
    var indexList = [];
    for (var i = 0; i < lines.length; i++) {
        var line = lines[i];
        if (line.includes(COMMENT)) {
            indexList.push(i);
        }
    }

    for (var i = 0; i < indexList.length; i++) {
        lines.splice(indexList[i], 1);

        //if a line is removed, indexes are changed
        if (i !== indexList.length - 1) {
            for (var j = i + 1; j < indexList.length; j++) {
                indexList[j] = indexList[j] - 1;
            }
        }
    }

    return lines;

}
`

const FS_UTILS = `
'use strict';

var fs = require('fs');


var FSUtils = (function () {
    var api = {};

    api.exists = function (path) {
        try {
            return fs.existsSync(path)
        } catch (err) {/*NOPE*/ }
        return false;
    }

    api.copyFile = function (src, dest) {
        fs.copyFileSync(src, dest);
    }

    api.readFile = function (path, encoding) {
        return fs.readFileSync(path, encoding);
    }

    api.writeFile = function (path, content) {
        fs.writeFileSync(path, content);
    }

    return api;
})();

module.exports = FSUtils;`
