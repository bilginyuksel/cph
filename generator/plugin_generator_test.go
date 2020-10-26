package generator

import "testing"

func TestGeneratePluginWithGivenName_ExpectPlugin(t *testing.T) {
	GeneratePlugin(".", "example")
	// cordova-plugin-yusa-example
	// cordova-plugin-yusa-example/src
	// cordova-plugin-yusa-example/www
	// cordova-plugin-yusa-example/scripts
	// cordova-plugin-yusa-example/tests
	// cordova-plugin-yusa-example/hooks
	// cordova-plugin-yusa-example/resources
	// cordova-plugin-yusa-example/resources/build.gradle
	// cordova-plugin-yusa-example/gradle/gradleWrapper.jar
	// cordova-plugin-yusa-example/gradle/plugin.gradle
	// cordova-plugin-yusa-example/plugin.xml
	// cordova-plugin-yusa-example/package.json
	// cordova-plugin-yusa-example/app_define.json
	// cordova-plugin-yusa-example/tsconfig.json
}
