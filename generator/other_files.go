package generator

const PACKAGE_JSON = `{
    "name": "@hmscore/%s",
    "description": "Cordova HMS $NAME$ Plugin",
    "version": "$VERSION$",
    "main": "./www/index.js",
    "types": "./types/index.d.ts",
    "repository": {
      "type": "git",
      "url": "https://github.com/HMS-Core/hms-cordova-plugin.git",
      "directory": "%s"
    },
    "bugs": "https://github.com/HMS-Core/hms-cordova-plugin/issues",
    "homepage": "https://developer.huawei.com/consumer/en/doc/overview/HMS-Core-Plugin",
    "license": "Apache-2.0",
    "licenseFilename": "LICENCE",
    "readmeFilename": "README.md",
    "cordova": {
      "id": "%s",
      "platforms": [
        "android"
      ]
    },
    "keywords": [
      "cordova",
      "ecosystem:cordova",
      "cordova-android",
      "%s",
      "hms-%s",
      "huawei-%s",
      "hms",
      "huawei"
    ],
    "files": [
      "hooks",
      "src/android/src/main/java",
      "src/android/plugin.gradle",
      "types",
      "www",
      "LICENCE",
      "package.json",
      "plugin.xml",
      "README.md"
    ],
    "dependencies": {
      "@types/cordova": "0.0.34"
    },
    "devDependencies": {
      "typescript": "^4.1.2"
    },
    "scripts": {
      "build:core": "tsc",
      "build": "npm run build:core",
      "watch": "tsc --watch"
    }
  }
  `

const APP_DEFINE = `{
    "fileVersion": "1",
    "name": "Cordova-Hms-$NAME$",
    "description": "Cordova HMS $NAME$ Plugin",
    "zipName": "cordova-%s-package",
    "packageName": "%s",
    "version": "$VERSION$",
    "type": "microService",
    "processes": {
        "Cordova-Hms-$NAME$": {
            "subscribes": []
        }
    },
    "files": [
        "hooks",
        "src/android/src/main/java",
        "src/android/plugin.gradle",
        "types",
        "www",
        "LICENCE",
        "package.json",
        "plugin.xml",
        "README.md"
    ]
}
`

const TS_CONFIG = `{
    "compilerOptions": {
        "strict": true,
        "noImplicitAny": true,
        "noEmitOnError": true,
        "removeComments": false,
        "esModuleInterop": true,
        "sourceMap": true,
        "inlineSources": true,
        "target": "ES2015",
        "module": "commonjs",
        "newLine": "LF",
        "outDir": "www",
        "declaration": true,
        "declarationDir": "types"
    },
    "compileOnSave": true,
    "exclude": ["node_modules", "**/*.spec.ts", "types", "ionic-native"]
}`

const PLUGIN_XML = `<?xml version='1.0' encoding='utf-8'?>
<plugin id="%s"
        version="$VERSION$"
        xmlns="http://apache.org/cordova/ns/plugins/1.0"
        xmlns:android="http://schemas.android.com/apk/res/android">
    <name>HMS $NAME$</name>
    <description>Cordova HMS $NAME$ Plugin</description>
    <license>Apache 2.0</license>
    <keywords>cordova,%s,hms-%s,huawei-%s,hms,huawei</keywords>

    <platform name="android">
    </platform>
    </plugin>`

const PLUGIN_GRADLE = `cdvPluginPostBuildExtras.add({
    apply plugin: 'com.huawei.agconnect'
})`

const BUILD_GRADLE = `apply plugin: 'com.android.library'

buildscript {
    repositories {
        google()
        jcenter()
        maven { url 'https://developer.huawei.com/repo/' }
    }

    dependencies {
        classpath 'com.android.tools.build:gradle:3.6.0'
    }
}

repositories {
    google()
    jcenter()
    maven { url 'https://developer.huawei.com/repo/' }
}

android {
    compileSdkVersion 29
    buildToolsVersion '29.0.3'

    defaultConfig {
        minSdkVersion 19
        targetSdkVersion 29
    }

    compileOptions {
        sourceCompatibility JavaVersion.VERSION_1_8
        targetCompatibility JavaVersion.VERSION_1_8
    }
}

dependencies {
    implementation 'org.apache.cordova:framework:8.1.0'
    implementation $KIT_DEPENDECIES$

    // Add only if you get error like 'can not find class HiAnalyticsUtils'
    // implementation 'com.huawei.hms:stats:5.0.3.301'
    // Add only if you get error like 'can not find class AGConnectServicesConfig'
    // implementation 'com.huawei.agconnect:agconnect-core:1.4.2.300'
}
`

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

function showToast(msg) {
    return asyncExec('%s', 'Test', ['showToast', msg]);
}
exports.showToast = showToast;
`

const ANDROID_MANIFEST = `<manifest xmlns:android="http://schemas.android.com/apk/res/android" package="com.huawei.hms.cordova.%s">
<uses-permission android:name="android.permission.ACCESS_NETWORK_STATE" />
</manifest>
`

const AFTER_PLUGIN_INSTALL = `"use strict";

var FSUtils = require("./FSUtils");

var ROOT_GRADLE_FILE = "platforms/android/build.gradle";
var COMMENT = "//This line is added by %s plugin";
var NEW_LINE = "\n";

module.exports = function (context) {
    if (!FSUtils.exists(ROOT_GRADLE_FILE)) {
        console.log(
            "root gradle file does not exist. after_plugin_install script wont be executed."
        );
    }

    var rootGradleContent = FSUtils.readFile(ROOT_GRADLE_FILE, "UTF-8");
    var lines = rootGradleContent.split(NEW_LINE);

    var depAddedLines = addAGConnectDependency(lines);
    var repoAddedLines = addHuaweiRepo(depAddedLines);

    FSUtils.writeFile(ROOT_GRADLE_FILE, repoAddedLines.join(NEW_LINE));
};

function addAGConnectDependency(lines) {
    var AG_CONNECT_DEPENDENCY =
        "classpath 'com.huawei.agconnect:agcp:1.4.2.300' " + COMMENT;
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
    var HUAWEI_REPO =
        "maven { url 'https://developer.huawei.com/repo/' } " + COMMENT;
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
}
`

const BEFORE_PLUGIN_UNINSTALL = `
"use strict";

var FSUtils = require("./FSUtils");

var ROOT_GRADLE_FILE = "platforms/android/build.gradle";
var COMMENT = "//This line is added by %s plugin";
var NEW_LINE = "\n";

module.exports = function (context) {
    if (!FSUtils.exists(ROOT_GRADLE_FILE)) {
        console.log(
            "root gradle file does not exist. before_plugin_uninstall script wont be executed."
        );
    }

    var rootGradleContent = FSUtils.readFile(ROOT_GRADLE_FILE, "UTF-8");
    var lines = rootGradleContent.split(NEW_LINE);

    var linesAfterRemove = removeLinesAddedByPlugin(lines);

    FSUtils.writeFile(ROOT_GRADLE_FILE, linesAfterRemove.join(NEW_LINE));
};

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
"use strict";

var fs = require("fs");

var FSUtils = (function () {
    var api = {};

    api.exists = function (path) {
        try {
            return fs.existsSync(path);
        } catch (err) {
            /*NOPE*/
        }
        return false;
    };

    api.readFile = function (path, encoding) {
        return fs.readFileSync(path, encoding);
    };

    api.writeFile = function (path, content) {
        fs.writeFileSync(path, content);
    };

    return api;
})();

module.exports = FSUtils;
`

const LICENCE_FILE = `Apache License

Version 2.0, January 2004

http://www.apache.org/licenses/

TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

1. Definitions.

"License" shall mean the terms and conditions for use, reproduction, and distribution as defined by Sections 1 through 9 of this document.

"Licensor" shall mean the copyright owner or entity authorized by the copyright owner that is granting the License.

"Legal Entity" shall mean the union of the acting entity and all other entities that control, are controlled by, or are under common control with that entity. For the purposes of this definition, "control" means (i) the power, direct or indirect, to cause the direction or management of such entity, whether by contract or otherwise, or (ii) ownership of fifty percent (50%) or more of the outstanding shares, or (iii) beneficial ownership of such entity.

"You" (or "Your") shall mean an individual or Legal Entity exercising permissions granted by this License.

"Source" form shall mean the preferred form for making modifications, including but not limited to software source code, documentation source, and configuration files.

"Object" form shall mean any form resulting from mechanical transformation or translation of a Source form, including but not limited to compiled object code, generated documentation, and conversions to other media types.

"Work" shall mean the work of authorship, whether in Source or Object form, made available under the License, as indicated by a copyright notice that is included in or attached to the work (an example is provided in the Appendix below).

"Derivative Works" shall mean any work, whether in Source or Object form, that is based on (or derived from) the Work and for which the editorial revisions, annotations, elaborations, or other modifications represent, as a whole, an original work of authorship. For the purposes of this License, Derivative Works shall not include works that remain separable from, or merely link (or bind by name) to the interfaces of, the Work and Derivative Works thereof.

"Contribution" shall mean any work of authorship, including the original version of the Work and any modifications or additions to that Work or Derivative Works thereof, that is intentionally submitted to Licensor for inclusion in the Work by the copyright owner or by an individual or Legal Entity authorized to submit on behalf of the copyright owner. For the purposes of this definition, "submitted" means any form of electronic, verbal, or written communication sent to the Licensor or its representatives, including but not limited to communication on electronic mailing lists, source code control systems, and issue tracking systems that are managed by, or on behalf of, the Licensor for the purpose of discussing and improving the Work, but excluding communication that is conspicuously marked or otherwise designated in writing by the copyright owner as "Not a Contribution."

"Contributor" shall mean Licensor and any individual or Legal Entity on behalf of whom a Contribution has been received by Licensor and subsequently incorporated within the Work.

2. Grant of Copyright License. Subject to the terms and conditions of this License, each Contributor hereby grants to You a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable copyright license to reproduce, prepare Derivative Works of, publicly display, publicly perform, sublicense, and distribute the Work and such Derivative Works in Source or Object form.

3. Grant of Patent License. Subject to the terms and conditions of this License, each Contributor hereby grants to You a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable (except as stated in this section) patent license to make, have made, use, offer to sell, sell, import, and otherwise transfer the Work, where such license applies only to those patent claims licensable by such Contributor that are necessarily infringed by their Contribution(s) alone or by combination of their Contribution(s) with the Work to which such Contribution(s) was submitted. If You institute patent litigation against any entity (including a cross-claim or counterclaim in a lawsuit) alleging that the Work or a Contribution incorporated within the Work constitutes direct or contributory patent infringement, then any patent licenses granted to You under this License for that Work shall terminate as of the date such litigation is filed.

4. Redistribution. You may reproduce and distribute copies of the Work or Derivative Works thereof in any medium, with or without modifications, and in Source or Object form, provided that You meet the following conditions:

You must give any other recipients of the Work or Derivative Works a copy of this License; and
You must cause any modified files to carry prominent notices stating that You changed the files; and
You must retain, in the Source form of any Derivative Works that You distribute, all copyright, patent, trademark, and attribution notices from the Source form of the Work, excluding those notices that do not pertain to any part of the Derivative Works; and
If the Work includes a "NOTICE" text file as part of its distribution, then any Derivative Works that You distribute must include a readable copy of the attribution notices contained within such NOTICE file, excluding those notices that do not pertain to any part of the Derivative Works, in at least one of the following places: within a NOTICE text file distributed as part of the Derivative Works; within the Source form or documentation, if provided along with the Derivative Works; or, within a display generated by the Derivative Works, if and wherever such third-party notices normally appear. The contents of the NOTICE file are for informational purposes only and do not modify the License. You may add Your own attribution notices within Derivative Works that You distribute, alongside or as an addendum to the NOTICE text from the Work, provided that such additional attribution notices cannot be construed as modifying the License. 

You may add Your own copyright statement to Your modifications and may provide additional or different license terms and conditions for use, reproduction, or distribution of Your modifications, or for any such Derivative Works as a whole, provided Your use, reproduction, and distribution of the Work otherwise complies with the conditions stated in this License.
5. Submission of Contributions. Unless You explicitly state otherwise, any Contribution intentionally submitted for inclusion in the Work by You to the Licensor shall be under the terms and conditions of this License, without any additional terms or conditions. Notwithstanding the above, nothing herein shall supersede or modify the terms of any separate license agreement you may have executed with Licensor regarding such Contributions.

6. Trademarks. This License does not grant permission to use the trade names, trademarks, service marks, or product names of the Licensor, except as required for reasonable and customary use in describing the origin of the Work and reproducing the content of the NOTICE file.

7. Disclaimer of Warranty. Unless required by applicable law or agreed to in writing, Licensor provides the Work (and each Contributor provides its Contributions) on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied, including, without limitation, any warranties or conditions of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE. You are solely responsible for determining the appropriateness of using or redistributing the Work and assume any risks associated with Your exercise of permissions under this License.

8. Limitation of Liability. In no event and under no legal theory, whether in tort (including negligence), contract, or otherwise, unless required by applicable law (such as deliberate and grossly negligent acts) or agreed to in writing, shall any Contributor be liable to You for damages, including any direct, indirect, special, incidental, or consequential damages of any character arising as a result of this License or out of the use or inability to use the Work (including but not limited to damages for loss of goodwill, work stoppage, computer failure or malfunction, or any and all other commercial damages or losses), even if such Contributor has been advised of the possibility of such damages.

9. Accepting Warranty or Additional Liability. While redistributing the Work or Derivative Works thereof, You may choose to offer, and charge a fee for, acceptance of support, warranty, indemnity, or other liability obligations and/or rights consistent with this License. However, in accepting such obligations, You may act only on Your own behalf and on Your sole responsibility, not on behalf of any other Contributor, and only if You agree to indemnify, defend, and hold each Contributor harmless for any liability incurred by, or claims asserted against, such Contributor by reason of your accepting any such warranty or additional liability.

END OF TERMS AND CONDITIONS`

const GITIGNORE = `# Generic
.idea/
*.iml
*.class
.classpath
.project
.vscode
.settings
*.bak
*.swp
*.user

# Android
.gradle
local.properties
logcat.log
build/
captures/

# Node JS + TS
logs
*.log
npm-debug.log*
yarn-debug.log*
yarn-error.log*
pids
*.pid
*.seed
*.pid.lock
lib-cov
coverage
.nyc_output
.grunt
bower_components
.lock-wscript
node_modules/
jspm_packages/
typings/
.npm
.eslintcache
.node_repl_history
*.tgz
.yarn-integrity
.env
.env.test
.cache
*.cache
.next
.nuxt
.vuepress/dist
.serverless/
.fusebox/
.dynamodb/
.tmp

# Mac
.DS_Store
*.DS_Store
.AppleDouble
.LSOverride
.DocumentRevisions-V100
.fseventsd
.Spotlight-V100
.TemporaryItems
.Trashes
.VolumeIcon.icns
.com.apple.timemachine.donotpresent
.AppleDB
.AppleDesktop
.apdisk
*.xcuserdatad

# Windows
Thumbs.db
ehthumbs.db
ehthumbs_vista.db
*.stackdump
[Dd]esktop.ini
$RECYCLE.BIN/
*.lnk

# Linux
*~
.fuse_hidden*
.directory
.Trash-*
.nfs*
`
