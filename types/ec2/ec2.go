package ec2

import (
	"encoding/json"
	re "regexp"
	"fmt"
)


type Region string
type AvailabilityZone string


type IAMRoleName string
func (v *IAMRoleName) UnmarshalJSON(data []byte) error {
	if !re.MustCompile(`[\w+=,.@-]+`).Match(data) {
		return fmt.Errorf("Impossible IAMRoleName '%s'", string(data))
	}
	return json.Unmarshall(data, v)
}


type ELBName string
func (v *ELBName) UnmarshalJSON(data []byte) error {
	// This name must be unique within your set of load balancers for the region, must have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and cannot begin or end with a hyphen.
	if !re.MustCompile(`[\w-]{1,32}`).Match(data) {
		return fmt.Errorf("Impossible ELBName '%s'", string(data))
	}
	return json.Unmarshall(data, v)
}

type SubnetName string
type SecurityGroupId string
type IAMPolicyARN string
type InstanceProfileARN string

type Timeout string //PT15M
type InstanceType string //enum?
type VolumeType string //enum?
type InstanceTenancy string

type SSHKeyPairName string
type HostedZoneId string

type EFSId string
type VPCId string
type IGWId string
type SubnetId string
type RouteTableId string
type NGWId string
type StackName string


type TagName string
type TagValue string
