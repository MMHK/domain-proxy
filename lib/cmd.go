package lib

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type Entry struct {
	Domain string
	IP     string
}

func Run(cmd *exec.Cmd) (*bytes.Buffer, error) {
	var outPipe bytes.Buffer
	var errorPipe bytes.Buffer

	cmd.Stderr = &errorPipe
	cmd.Stdout = &outPipe
	log.Infof("start runing task:%s\n", cmd)
	err := cmd.Run()
	if err != nil {
		log.Error(err)
		log.Errorf("error:%s", outPipe.String())
		return nil, err
	}
	return &outPipe, nil
}

func ReloadService(cmd string) (err error) {
	buf, err := Run(exec.Command(cmd))

	log.Infof("run cmd:%s, result:%s", cmd, buf.String())

	return err
}

func parseTemplate(tpl_path string, domain string, ip string) (error, string) {
	var data []byte
	data, err := ioutil.ReadFile(tpl_path)
	if err != nil {
		log.Error(err)
		return err, ""
	}
	tpl := string(data)
	tplEntry, err := template.New("new_entry").Parse(tpl)
	if err != nil {
		log.Error(err)
		return err, ""
	}
	buf := new(bytes.Buffer)
	entry := Entry{Domain: domain, IP: ip}
	err = tplEntry.Execute(buf, entry)
	if err != nil {
		log.Error(err)
		return err, ""
	}
	return nil, buf.String()
}

func parseName(tpl string, domain string, ip string) (error, string) {
	entry := Entry{Domain: domain, IP: ip}
	tpl_entry, err := template.New("new_name").Parse(tpl)
	if err != nil {
		log.Error(err)
		return err, ""
	}
	buf := new(bytes.Buffer)
	err = tpl_entry.Execute(buf, entry)
	if err != nil {
		log.Error(err)
		return err, ""
	}
	return nil, buf.String()
}

func createNewDomainConfig(domain string, ip string, save_path string, tpl_path string, name string) (error) {
	var entry_content string
	err, entry_content := parseTemplate(tpl_path, domain, ip)
	if err != nil {
		log.Error(err)
		return err
	}
	var filename string
	err, filename = parseName(name, domain, ip)
	if err != nil {
		log.Error(err)
		return err
	}
	file_path := filepath.Join(save_path, filename)
	log.Info(file_path)
	data := []byte(entry_content)
	err = ioutil.WriteFile(file_path, data, os.ModePerm)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func RemoveDomainConfig(domain string, ip string, save_path string, name string) (error) {
	var filename string
	err, filename := parseName(name, domain, ip)
	if err != nil {
		log.Error(err)
		return err
	}
	filePath := filepath.Join(save_path, filename)
	err = os.Remove(filePath)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
