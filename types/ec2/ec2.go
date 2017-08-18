package ec2

import (
)

type Region string
type AvailabilityZone string

// +gen * regex
type IAMRoleName string
const IAMRoleName_regex = `^[\w+=,.@-]+$`

// +gen * regex
type ELBName string
const ELBName_regex = `^[^-][\w-]{1,30}[^-]$`

type SubnetName string

// +gen * regex
type SecurityGroupId string
const SecurityGroupId_regex = `^sg-[0-9a-z]{8}$`

type IAMPolicyARN string
type InstanceProfileARN string
type ALBTargetGroupARN string
type KMSKeyARN string

// +gen * regex
type Timeout string      //PT15M
const Timeout_regex = `^PT(?:\d+H)?(?:\d+M)?(?:\d+S)?$`

type InstanceType string //enum?
type VolumeType string   //enum?
type InstanceTenancy string

type SSHKeyPairName string
type HostedZoneId string

// +gen * regex
type AmiId string
const AmiId_regex = `^ami-[0-9a-z]{8}$`

type EFSId string
type VPCId string
type IGWId string

// +gen * regex
type SubnetId string
const SubnetId_regex = `^subnet-[0-9a-z]{8}$`

// +gen * regex
type RouteTableId string
const RouteTableId_regex = `^rtb-[0-9a-z]{8}$`

type NGWId string
type EIPAllocId string

type StackName string

type TagName string
type TagValue string
