package writer

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestWriteLicence_LicenseGivenFile(t *testing.T)  {
	WriteLicence("test/test2/HmsPushMessaging.java")
	f, err := os.OpenFile("HmsPushMessaging.java", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	rawBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawBytes), "\n")
	actualLicence := ""
	for i:=0; i<16; i++ {
		actualLicence += lines[i]
	}
	if actualLicence != LICENCE {
		t.Errorf("Licences should be same.Actual licence: %s, Expected licence: %s ", actualLicence, LICENCE)
	}
}
