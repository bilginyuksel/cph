package writer

import (
	"fmt"
	"github.com/bilginyuksel/cordova-plugin-helper/parser"
	"github.com/bilginyuksel/cordova-plugin-helper/reader"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func beforeTestParseFileCreateSampleCordovaPluginXMLFile(filename string, t *testing.T) {

	content := []byte(`
	<?xml version='1.0' encoding='utf-8'?>
	<plugin id="cordova-plugin-hms-push"
			version="5.0.2"
			xmlns="http://apache.org/cordova/ns/plugins/1.0"
			xmlns:android="http://schemas.android.com/apk/res/android">
		<name>Cordova Plugin HMS Push</name>
		<description>Cordova Plugin HMS Push</description>
		<license>Apache 2.0</license>
		<keywords>android, huawei, hms, push</keywords>
	
		<engines>
			<engine name="cordova" version=">=3.0.0"/>
		</engines>
	
		<js-module name="HmsPush" src="www/HmsPush.js">
			<clobbers target="HmsPush"/>
		</js-module>
	
		<js-module name="HmsPushResultCode" src="www/HmsPushResultCode.js">
			<clobbers target="HmsPushResultCode"/>
		</js-module>
	
		<js-module name="HmsPushEvent" src="www/HmsPushEvent.js">
			<clobbers target="HmsPushEvent"/>
		</js-module>
	
		<js-module name="HmsLocalNotification" src="www/HmsLocalNotification.js">
			<clobbers target="HmsLocalNotification"/>
		</js-module>
	
	
		<js-module name="Interfaces" src="www/Interfaces.js"/>
		<js-module name="CordovaRemoteMessage" src="www/CordovaRemoteMessage.js"/>
		<js-module name="utils" src="www/utils.js"/>
	
		<platform name="android">
	
			<hook type="after_plugin_install" src="hooks/after_plugin_install.js"/>
			<hook type="before_plugin_uninstall" src="hooks/before_plugin_uninstall.js"/>
			<hook type="after_prepare" src="hooks/after_prepare.js"/>
	
			<config-file target="AndroidManifest.xml" parent="/*">
				<uses-permission android:name="android.permission.INTERNET"/>
				<uses-permission android:name="android.permission.ACCESS_NETWORK_STATE"/>
				<!-- Below permissions are to support vibration and send scheduled local notifications -->
				<uses-permission android:name="android.permission.VIBRATE"/>
				<uses-permission android:name="android.permission.RECEIVE_BOOT_COMPLETED"/>
				<uses-permission android:name="android.permission.WAKE_LOCK"/>
				<uses-permission android:name="android.permission.SYSTEM_ALERT_WINDOW"/>
			</config-file>
	
			<config-file target="AndroidManifest.xml" parent="application">
				<receiver android:name="com.huawei.hms.cordova.push.receiver.HmsLocalNotificationActionsReceiver"/>
				<!-- This receivers are for sending scheduled local notifications -->
				<receiver android:name="com.huawei.hms.cordova.push.receiver.HmsLocalNotificationBootEventReceiver">
					<intent-filter>
						<action android:name="android.intent.action.BOOT_COMPLETED"/>
					</intent-filter>
				</receiver>
				<receiver android:name="com.huawei.hms.cordova.push.receiver.HmsLocalNotificationScheduledPublisher"
						  android:enabled="true"
						  android:exported="true">
				</receiver>
				<meta-data
						android:name="push_kit_auto_init_enabled"
						android:value="true"/>
			</config-file>
			<config-file target="AndroidManifest.xml" parent="application/activity">
				<intent-filter>
					<action android:name="android.intent.action.VIEW"/>
					<category android:name="android.intent.category.DEFAULT"/>
					<category android:name="android.intent.category.BROWSABLE"/>
					<data android:scheme="app"/>
				</intent-filter>
			</config-file>
	
			<config-file target="AndroidManifest.xml" parent="application">
				<service android:name="com.huawei.hms.cordova.push.remote.HmsPushMessageService" android:exported="true">
					<intent-filter>
						<action android:name="com.huawei.push.action.MESSAGING_EVENT"/>
					</intent-filter>
				</service>
			</config-file>
	
			<config-file target="config.xml" parent="/*">
				<feature name="HMSPush">
					<param name="android-package" value="com.huawei.hms.cordova.push.HMSPush"/>
				</feature>
			</config-file>
	
			<framework src="androidx.core:core:1.3.1"/>
			<framework src="com.facebook.fresco:fresco:2.2.0"/>
			<framework src="com.huawei.hms:push:5.0.2.300"/>
	
			<framework src="resources/plugin.gradle" custom="true" type="gradleReference"/>
	
			<source-file src="src/main/java/com/huawei/hms/cordova/push/HMSPush.java"
						 target-dir="src/com/huawei/hms/cordova/push"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/hmslogger/HMSLogger.java"
						 target-dir="src/com/huawei/hms/cordova/push/hmslogger"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/config/NotificationAttributes.java"
						 target-dir="src/com/huawei/hms/cordova/push/config"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/constants/Core.java"
						 target-dir="src/com/huawei/hms/cordova/push/constants"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/constants/LocalNotification.java"
						 target-dir="src/com/huawei/hms/cordova/push/constants"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/constants/NotificationConstants.java"
						 target-dir="src/com/huawei/hms/cordova/push/constants"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/constants/RemoteMessageAttributes.java"
						 target-dir="src/com/huawei/hms/cordova/push/constants"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/constants/ResultCode.java"
						 target-dir="src/com/huawei/hms/cordova/push/constants"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/listeners/HmsLocalNotificationActionPublisher.java"
						 target-dir="src/com/huawei/hms/cordova/push/listeners"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/listeners/HmsMessagePublisher.java"
						 target-dir="src/com/huawei/hms/cordova/push/listeners"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/local/BitmapDataSubscriber.java"
						 target-dir="src/com/huawei/hms/cordova/push/local"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/local/HmsLocalNotification.java"
						 target-dir="src/com/huawei/hms/cordova/push/local"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/local/HmsLocalNotificationController.java"
						 target-dir="src/com/huawei/hms/cordova/push/local"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/local/HmsLocalNotificationPicturesLoader.java"
						 target-dir="src/com/huawei/hms/cordova/push/local"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/receiver/HmsLocalNotificationActionsReceiver.java"
						 target-dir="src/com/huawei/hms/cordova/push/receiver"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/receiver/HmsLocalNotificationBootEventReceiver.java"
						 target-dir="src/com/huawei/hms/cordova/push/receiver"/>
			<source-file
					src="src/main/java/com/huawei/hms/cordova/push/receiver/HmsLocalNotificationScheduledPublisher.java"
					target-dir="src/com/huawei/hms/cordova/push/receiver"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/receiver/NotificationActionHandler.java"
						 target-dir="src/com/huawei/hms/cordova/push/receiver"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/remote/HmsPushInstanceId.java"
						 target-dir="src/com/huawei/hms/cordova/push/remote"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/remote/HmsPushMessageService.java"
						 target-dir="src/com/huawei/hms/cordova/push/remote"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/remote/HmsPushMessaging.java"
						 target-dir="src/com/huawei/hms/cordova/push/remote"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/utils/ApplicationUtils.java"
						 target-dir="src/com/huawei/hms/cordova/push/utils"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/utils/ArrayUtil.java"
						 target-dir="src/com/huawei/hms/cordova/push/utils"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/utils/BundleUtils.java"
						 target-dir="src/com/huawei/hms/cordova/push/utils"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/utils/CordovaUtils.java"
						 target-dir="src/com/huawei/hms/cordova/push/utils"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/utils/MapUtils.java"
						 target-dir="src/com/huawei/hms/cordova/push/utils"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/utils/NotificationConfigUtils.java"
						 target-dir="src/com/huawei/hms/cordova/push/utils"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/utils/RemoteMessageUtils.java"
						 target-dir="src/com/huawei/hms/cordova/push/utils"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/utils/Action.java"
						 target-dir="src/com/huawei/hms/cordova/push/utils"/>
			<source-file src="src/main/java/com/huawei/hms/cordova/push/utils/ActionManager.java"
						 target-dir="src/com/huawei/hms/cordova/push/utils"/>
	
	
		</platform>
	</plugin>	
	`)
	err := ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckIfFileContainsLicenceAlready_LicenceExists(t *testing.T) {
	file := createFile("test.java")
	file, _ = os.OpenFile(file.Name(), os.O_RDWR, 0644)
	_, err := file.WriteString(ReadFile("licence.java"))

	_ = file.Close()
	ok, err := CheckIfFileContainsLicenceAlready("test.java", "licence.java")
	if err != nil {
		t.Error()
	}
	if !ok {
		t.Error()
	}
	os.Remove(file.Name())
}

func TestCheckIfFileContainsLicenceAlready_LicenceNotFound(t *testing.T) {
	file := createFile("test1.java")
	ok, err := CheckIfFileContainsLicenceAlready("test1.java", "licence.java")
	if err != nil {
		t.Error()
	}
	if ok {
		t.Error()
	}
	os.Remove(file.Name())
}

func TestCheckIfFileContainsLicenceAlready_FileNotFound(t *testing.T) {
	_, err := CheckIfFileContainsLicenceAlready("notFound.java", "licence.java")
	if err == nil {
		t.Error()
	}
}

func TestCheckIfFileContainsLicenceAlready_LicenceExistWithWrongFormat(t *testing.T) {
	file := createFile("test.java")
	_, _ = file.WriteString(ReadFile("licence.java"))
	_, _ = file.WriteString("test string")
	file, err := os.OpenFile("licence.java", os.O_RDONLY, 0644)
	if err != nil {
		t.Error()
	}
	_ = file.Close()
	CheckIfLicenceFormatIsValid(file)
	os.Remove("test.java")
}

func TestWriteToFileLicence_LicenceIsWrittenProperly(t *testing.T) {
	file := createFile("test.java")
	for i := 1; i < 100; i++ {
		_, _ = file.WriteString("This is a test file.\n")
	}
	_, _ = WriteLicenceToFile(file.Name(), "licence.java")
	s := ReadFile(file.Name())
	if !strings.Contains(s, ReadFile("licence.java")) {
		t.Error()
	}
	os.Remove(file.Name())
}

func TestWriteToFileLicence_LicenceIsNotWrittenProperly(t *testing.T) {
	ok, _ := WriteLicenceToFile("test.java", "licence.java")
	if ok {
		t.Error()
	}
}

func TestReadSourceFiles_AllJavaFilesAdded(t *testing.T) {
	createFolder("src")
	_ = createFile("src/test1.java")
	beforeTestParseFileCreateSampleCordovaPluginXMLFile("plugin.xml", t)
	plg, _ := parser.ParseXML("plugin.xml")
	javaFiles, _ := reader.FilePathWalkDir("src")
	plg.Platform.NewSource(javaFiles)
	_ = parser.CreateXML(plg, "plg.xml")
	newPlugin, _ := parser.ParseXML("plg.xml")
	newSourceFile := newPlugin.Platform.SourceFiles
	if newSourceFile[0].Src != "src\\test1.java" {
		t.Error()
	}
	os.Remove("plugin.xml")
	os.Remove("plg.xml")
	os.RemoveAll("src")
}

func TestReadJsModules_AllJsFilesAdded(t *testing.T) {
	createFolder("src")
	createFolder("www")
	_ = createFile("www/test1.js")
	_ = createFile("www/test2.js")
	_ = createFile("www/test3.js")
	beforeTestParseFileCreateSampleCordovaPluginXMLFile("plugin.xml", t)
	plg, _ := parser.ParseXML("plugin.xml")
	jsFiles, _ := reader.FilePathWalkDir("www")
	plg.NewJsModules(jsFiles)
	_ = parser.CreateXML(plg, "plg.xml")
	newPlugin, _ := parser.ParseXML("plg.xml")
	newJsModules := newPlugin.JsModule
	for i := 1; i <= 3; i++ {
		file := fmt.Sprintf("www\\test%d.js", i)
		if newJsModules[i-1].Src != file {
			t.Error()
		}
	}
	os.Remove("plugin.xml")
	os.Remove("plg.xml")
	os.RemoveAll("src")
	os.RemoveAll("www")
}

func TestReadJsModules_PluginXmlDoesntContainNonExistFile(t *testing.T) {
	createFolder("www")
	file := createFile("www/test1.js")
	beforeTestParseFileCreateSampleCordovaPluginXMLFile("plugin.xml", t)
	plg, _ := parser.ParseXML("plugin.xml")
	jsFiles, _ := reader.FilePathWalkDir("www")
	plg.NewJsModules(jsFiles)
	_ = parser.CreateXML(plg, "plg.xml")
	os.Remove(file.Name())
	resultPlugin, _ := parser.ParseXML("plg.xml")
	jsFiles, _ = reader.FilePathWalkDir("www")
	resultPlugin.NewJsModules(jsFiles)
	_ = parser.CreateXML(resultPlugin, "plg.xml")
	b, _ := ioutil.ReadFile("plg.xml")
	content := string(b)
	if strings.Contains(content, "www/test1.js") {
		t.Error()
	}
	os.Remove("plugin.xml")
	os.Remove("plg.xml")
	os.RemoveAll("www")
}

func createFolder(fileName string) {
	err := os.Mkdir(fileName, 0755)
	if err != nil {
		panic(err)
	}
}
func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
	return file
}
