package ec2

import (
	re "regexp"
	"fmt"
)


type Region string
type AvailabilityZone string


type IAMRoleName string
func (v *IAMRoleName) UnmarshalYAML(f func(interface{}) error) error {
	var data string
	if err := f(&data); err != nil {
		return err
	}

	if !re.MustCompile(`^[\w+=,.@-]+$`).MatchString(data) {
		return fmt.Errorf("Impossible IAMRoleName '%s'", data)
	}

	*v = IAMRoleName(data)
	return nil
}

type ELBName string
func (v *ELBName) UnmarshalYAML(f func(interface{}) error) error {
	var data string
	if err := f(&data); err != nil {
		return err
	}

	if !re.MustCompile(`^[^-][\w-]{1,30}[^-]$`).MatchString(data) {
		return fmt.Errorf("Impossible ELBName '%s'", data)
	}

	*v = ELBName(data)
	return nil
}

type SubnetName string
type SecurityGroupId string
type IAMPolicyARN string
type InstanceProfileARN string
type ALBTargetGroupARN string
type KMSKeyARN string

type Timeout string //PT15M
type InstanceType string //enum?
type VolumeType string //enum?
type InstanceTenancy string

type SSHKeyPairName string
type HostedZoneId string

type AmiId string
type EFSId string
type VPCId string
type IGWId string
type SubnetId string
type RouteTableId string
type NGWId string
type EIPAllocId string

type StackName string


type TagName string
type TagValue string
