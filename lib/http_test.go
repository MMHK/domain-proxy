package lib

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_AddEntry(t *testing.T) {
	cfg, err := loadConfig()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	httpServer := NewService(cfg.Listen, cfg)

	formData := url.Values{
		"domain": []string{"sam-test.demo2.mixmedia.com"},
		"ip": []string{"192.168.33.127"},
	}

	t.Log(formData.Encode())

	req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(formData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	writer := httptest.NewRecorder()

	httpServer.AddEntry(writer, req)

	resp := writer.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is %v", resp.StatusCode)
		t.Fail()
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)

	t.Log(string(data))
}

func Test_RemoveEntry(t *testing.T) {
	cfg, err := loadConfig()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	httpServer := NewService(cfg.Listen, cfg)

	formData := url.Values{
		"domain": []string{"sam-test.demo2.mixmedia.com"},
		"ip": []string{"192.168.33.127"},
	}

	t.Log(formData.Encode())

	req := httptest.NewRequest(http.MethodPost, "/remove", strings.NewReader(formData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	writer := httptest.NewRecorder()

	httpServer.RemoveEntry(writer, req)

	resp := writer.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is %v", resp.StatusCode)
		t.Fail()
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)

	t.Log(string(data))
}