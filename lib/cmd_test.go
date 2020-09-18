package lib

import (
	"testing"
)

func Test_createNewDomainConfig(t *testing.T) {
	config, err := loadConfig()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	domain := "sam-test.demo2.mixmedia.com"
	ip := "192.168.33.127"

	err = createNewDomainConfig(domain, ip, config.DomainCfgSaveDir,
		config.EntryTemplate, config.DomainCfgFileNameFormat)

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	//err, filename := parseName(config.DomainCfgFileNameFormat, domain, ip)
	//if err != nil {
	//	t.Error(err)
	//	t.Fail()
	//	return
	//}
	//
	//tmpPath := filepath.Join(config.DomainCfgSaveDir, filename)
	//defer os.Remove(tmpPath)
}