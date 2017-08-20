package types

import(
	"net/url"
	"net"
	"encoding/base64"
	"gopkg.in/yaml.v2"
	"fmt"
)

type ClusterName string
type DNSName string
type NvidiaDriverVersion string
type SSHAuthorizedKey string

type BlockDeviceName string
type FilesystemPath string

type EtcdVersion string
type EtcdMemberIdentityProvider string //enum: eni, eip
type EtcdMemberIdentifier string       //validate: can't have '='

type ContainerRuntime string //enum: docker, rkt

// +gen * regex
type SHA1SUM string
const SHA1SUM_regex = `^[a-f0-9]{40}$`

// +gen parse
type Base64Yaml string
var Base64Yaml_parse = func(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	var i interface{}
	err = yaml.UnmarshalStrict(decoded, &i)
	return s, err
}

// +gen parse
type URL url.URL
var URL_parse = func(s string) (url.URL, error) {
	res, err := url.ParseRequestURI(s)
	if err != nil {
		return url.URL{}, err
	}
	return *res, nil
}

// +gen parse
type IPNet net.IPNet
var IPNet_parse = func(s string) (net.IPNet, error) {
	_, ipnet, err := net.ParseCIDR(s)
	if ipnet == nil {
		return net.IPNet{}, fmt.Errorf("'%s' is not IPNet", s)
	}
	return *ipnet, err
}

// +gen parse
type IP net.IP
var IP_parse = func(s string) (net.IP, error) {
	res := net.ParseIP(s)
	if res == nil {
		return net.IP{}, fmt.Errorf("'%s' is not IP address", s)
	}
	return res, nil
}
