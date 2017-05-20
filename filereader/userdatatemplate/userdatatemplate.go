package userdatatemplate

import (
	"bytes"
	"errors"
	"encoding/json"

	ct "github.com/coreos/container-linux-config-transpiler/config"
	ctp "github.com/coreos/container-linux-config-transpiler/config/templating"
	"github.com/kubernetes-incubator/kube-aws/filereader/texttemplate"
)

func GetString(filename string, data interface{}) (string, error) {
	buf, err := texttemplate.GetBytesBuffer(filename, data)

	if err != nil {
		return "", err
	}

	conf, report := ct.Parse(buf.Bytes())
	if len(report.Entries) > 0 {
		return "", errors.New(report.String())
	}

	ignConf, report := ct.ConvertAs2_0(conf, ctp.PlatformEC2)
	if len(report.Entries) > 0 {
		return "", errors.New(report.String())
	}

	b, err := json.Marshal(&ignConf)
	return bytes.NewBuffer(b).String(), err
}
