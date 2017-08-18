package types

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

type SHA1SUM string //validate: regex=[a-f0-9]{40}

type Base64Yaml string //validate: base64 decode, then yaml
