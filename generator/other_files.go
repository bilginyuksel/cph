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

const TS_UTILS = `import { exec } from 'cordova';

export function asyncExec(clazz: string, reference: string, args: any = []) {
	return new Promise((resolve, reject) => {
		exec(resolve, reject, clazz, reference, args);
	});
}
`

const TS_MAIN = `import { asyncExec } from './utils';
`
