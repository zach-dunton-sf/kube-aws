package config

import (
	"github.com/kubernetes-incubator/kube-aws-ng/model"
	"github.com/kubernetes-incubator/kube-aws-ng/types"
	"github.com/kubernetes-incubator/kube-aws-ng/types/coreos"
	"github.com/kubernetes-incubator/kube-aws-ng/types/dex"
	"github.com/kubernetes-incubator/kube-aws-ng/types/ec2"
	"github.com/kubernetes-incubator/kube-aws-ng/types/kubernetes"
	"net"
	"net/url"
)

type SubnetConfRef struct {
	Name ec2.SubnetName
}

type SubnetIdConf struct {
	Id                ec2.SubnetId  //conflicts: IdFromStackOutput
	IdFromStackOutput ec2.StackName //conflicts: Id;
}

type NGWIdConf struct {
	Id                ec2.NGWId      //conflicts: IdFromStackOutput
	IdFromStackOutput ec2.StackName  //conflicts: Id;
	eipAllocationId   ec2.EIPAllocId //conflicts: Id, IdFromStackOutput
}

type RouteTableIdConf struct {
	Id                ec2.RouteTableId //conflicts: IdFromStackOutput
	IdFromStackOutput ec2.StackName    //conflicts: Id;
}

type SubnetConf struct {
	Name             ec2.SubnetName
	AvailabilityZone ec2.AvailabilityZone
	InstanceCIDR     net.IPNet
	Private          bool
	SubnetIdConf
	NatGateway NGWIdConf
	RouteTable RouteTableIdConf
}

type APIEndpointLoadBalancer struct { // need invariants generator
	Id              ec2.ELBName
	CreateRecordSet bool
	RecordSetTTL    uint

	Subnets                     []SubnetConfRef
	Private                     bool
	HostedZone                  ec2.HostedZoneId
	ApiAccessAllowedSourceCIDRs []net.IPNet
	SecurityGroupIds            []ec2.SecurityGroupId
}

type APIEndpointName string

type APIEndpointConf struct {
	Name    APIEndpointName
	DnsName types.DNSName

	LoadBalancer APIEndpointLoadBalancer
}

type ASGConf struct {
	MinSize                            uint
	MaxSize                            uint
	RollingUpdateMinInstancesInService uint
}

type IAMConf struct {
	Role struct {
		ManagedPolicies []ec2.IAMPolicyARN
		InstanceProfile ec2.InstanceProfileARN
	}
}

type VolumeConf struct {
	Size uint
	Type ec2.VolumeType
	Iops uint
}

type VolumeMountConf struct {
	VolumeConf
	Device types.BlockDeviceName
	Path   types.FilesystemPath
}

type MaybeEncryptedOrEphemeralVolume struct {
	VolumeConf
	Encrypted bool // conflicts: Ephemeral
	Ephemeral bool // conflicts: Encrypted, broken
}

type ControllerConf struct {
	Count              uint
	CreateTimeout      ec2.Timeout
	InstanceType       ec2.InstanceType
	RootVolume         VolumeConf //validate: must be empty/default device and path
	SecurityGroupIds   []ec2.SecurityGroupId
	AutoScalingGroup   ASGConf
	IAM                IAMConf
	Subnets            []SubnetConfRef
	NodeLabels         map[kubernetes.LabelName]kubernetes.LabelValue
	CustomFiles        []map[string]string
	CustomSystemdUnits []map[string]string
}

// validate: g2 or p2 instance, docker runtime
type GpuConf struct {
	Nvidia struct {
		Enabled bool
		Version types.NvidiaDriverVersion
	}
}

type NodePoolName string

type SpotFleetConf struct{} //TODO

type ContainerImages struct {
	HyperkubeImage                     model.Image `yaml:"omitempty",default:"{'quay.io/coreos/hyperkube', 'v1.7.3_coreos.0', false}"`
	AwsCliImage                        model.Image `yaml:"omitempty"`
	CalicoNodeImage                    model.Image `yaml:"omitempty"`
	CalicoCniImage                     model.Image `yaml:"omitempty"`
	CalicoCtlImage                     model.Image `yaml:"omitempty"`
	CalicoPolicyControllerImage        model.Image `yaml:"omitempty"`
	ClusterAutoscalerImage             model.Image `yaml:"omitempty"`
	ClusterProportionalAutoscalerImage model.Image `yaml:"omitempty"`
	KubeDnsImage                       model.Image `yaml:"omitempty"`
	KubeDnsMasqImage                   model.Image `yaml:"omitempty"`
	KubeReschedulerImage               model.Image `yaml:"omitempty"`
	DnsMasqMetricsImage                model.Image `yaml:"omitempty"`
	ExecHealthzImage                   model.Image `yaml:"omitempty"`
	HeapsterImage                      model.Image `yaml:"omitempty"`
	AddonResizerImage                  model.Image `yaml:"omitempty"`
	KubeDashboardImage                 model.Image `yaml:"omitempty"`
	PauseImage                         model.Image `yaml:"omitempty"`
	FlannelImage                       model.Image `yaml:"omitempty"`
	DexImage                           model.Image `yaml:"omitempty"`
	JournaldCloudWatchLogsImage        model.Image `yaml:"omitempty"`
}

type KubernetesContainerImages struct {
	KubernetesVersion string
	ContainerImages
}

type WaitSignalConf struct {
	Enabled      bool
	MaxBatchSize uint //validate: >0
}

type NodepoolConf struct {
	Name             NodePoolName
	Subnets          []SubnetConfRef
	SecurityGroupIds []ec2.SecurityGroupId
	LoadBalancer     struct {
		Enabled bool
		Names   []ec2.ELBName
		//SecurityGroupIds []ec2.SecurityGroup -- removed, duplicate of SGs above
	}
	ApiEndpointName APIEndpointName
	IAM             IAMConf
	TargetGroup     struct {
		Enabled bool
		Arns    []ec2.ALBTargetGroupARN
		//SecurityGroupIds []ec2.SecurityGroup -- removed, duplicate of SGs above
	}
	ManagedIamRoleSuffix      ec2.IAMRoleName `yaml:mangedIamRoleName`
	VolumeMounts              []VolumeMountConf
	NodeStatusUpdateFrequency kubernetes.TimePeriod // 10s, 5h
	Count                     uint
	InstanceType              ec2.InstanceType
	RootVolume                VolumeConf
	CreateTimeout             ec2.Timeout
	Tenancy                   ec2.InstanceTenancy
	Gpu                       GpuConf

	//SpotPrice uint -- removed, is listed in SpotFleetConf
	WaitSignal       WaitSignalConf
	AutoScalingGroup ASGConf
	SpotFleet        SpotFleetConf

	Autoscaling              AutoScalingConf
	AwsEnvironment           EtcAwsEnvironmentConf
	AwsNodeLabels            struct{ Enabled bool }
	ClusterAutoscalerSupport struct{ Enabled bool } //SuPER confusing with `autoscaling:`
	ElasticFileSystemId      ec2.EFSId
	EphemeralImageStorage    struct{ Enabled bool }
	Kube2IamSupport          struct{ Enabled bool }
	KubeletOpts              kubernetes.KubeletOptionsString

	NodeLabels map[kubernetes.LabelName]kubernetes.LabelValue
	Taints     []struct {
		Key    kubernetes.TaintKey
		Value  kubernetes.TaintValue
		Effect kubernetes.TaintEffect
	}
	KeyName        ec2.SSHKeyPairName
	ReleaseChannel coreos.ReleaseChannel
	AmiId          ec2.AmiId
	KubernetesContainerImages
	SSHAuthorizedKeys  []types.SSHAuthorizedKey
	CustomSettings     map[string]string
	CustomFiles        []map[string]string
	CustomSystemdUnits []map[string]string
}

type EtcdConf struct {
	Count            uint
	InstanceType     ec2.InstanceType
	RootVolume       VolumeConf
	DataVolume       MaybeEncryptedOrEphemeralVolume
	Tenancy          ec2.InstanceTenancy
	Subnets          []SubnetConfRef
	SecurityGroupIds []ec2.SecurityGroupId
	IAM              IAMConf

	Version                types.EtcdVersion
	Snapshot               struct{ Automated bool }
	DisasterRecovery       struct{ Automated bool }
	MemberIdentityProvider types.EtcdMemberIdentityProvider
	InternalDomainName     types.DNSName
	ManageRecordSets       bool
	HostedZone             struct{ Id ec2.HostedZoneId }
	Nodes                  []struct {
		Name types.EtcdMemberIdentifier
		Fqdn types.DNSName
	}

	KMSKeyArn          ec2.KMSKeyARN
	CustomFiles        []map[string]string
	CustomSystemdUnits []map[string]string
}

type AutoScalingConf struct {
	ClusterAutoScaler struct {
		Enabled bool
	}
}

type EtcAwsEnvironmentConf struct {
	Enabled     bool
	Environment map[string]string
}

type VPCConf struct {
	//VpcId ec2.VPCId -- removed as deprecated
	Vpc struct {
		Id                ec2.VPCId     //conflicts: IdFromStackOutput
		IdFromStackOutput ec2.StackName //conflicts: Id;
		//InternetGatewayId ec2.IGWId -- removed as deprecated
	}

	InternetGateway struct {
		Id                ec2.IGWId     //conflicts: IdFromStackOutput
		IdFromStackOutput ec2.StackName //conflicts: Id
	}

	// RouteTableId ec2.RouteTableId -- removed in favour of subnets[].routeTable.id

	VpcCIDR net.IPNet //future: should be conflicting with vpcID
	//InstanceCIDR net.IPNet // -- reomved in favour of nodepools[] and subnets[]

	Subnets []SubnetConf
}

type TLSConf struct {
	tlsCADurationDays   uint `default:"3650"`
	tlsCertDurationDays uint `default:"365"`
}

type CloudWatchLoggingConf struct {
	Enabled         bool
	RetentionInDays uint
	LocalStreaming  struct {
		Enabled  bool
		Filter   string
		Interval uint
	}
}

type AmazonSSMAgentConf struct {
	Enabled     bool
	DownloadURL url.URL
	Sha1Sum     types.SHA1SUM
}

type AuditLogConf struct {
	Enabled bool
	MaxAge  uint
	LogPath string
}

type ExperimentalAddonsConf struct {
	Admission struct {
		podSecurityPolicy  struct{ Enabled bool }
		denyEscalatingExec struct{ Enabled bool }
	}

	AwsEnvironment EtcAwsEnvironmentConf // q:found bare in nodepool, does it mean rest of experimental can be found there too?
	AuditLog       AuditLogConf
	Authentication struct {
		Webhook struct {
			Enabled      bool
			cacheTTL     kubernetes.TimePeriod
			configBase64 types.Base64Yaml
		}
	}
}

type DexConf struct {
	Enabled         bool
	URL             url.URL
	ClientID        dex.ClientID
	Username        dex.Username
	SelfSignedCa    bool
	Connectors      []map[string]interface{} //TODO better structs
	StaticClients   []map[string]string      //TODO better structs
	StaticPasswords []map[string]string      //TODO better structs
}

type AddonsConf struct {
	ClusterAutoscaler struct{ Enabled bool }
	Rescheduler       struct{ Enabled bool }

	ExperimentalAddonsConf

	AwsNodeLabels       struct{ Enabled bool }
	TlsBootstrap        struct{ Enabled bool }
	ephemeralImageStore struct{ Enabled bool }
	Kube2IamSupport     struct{ Enabled bool }
	NodeDrainer         struct{ Enabled bool }

	Dex     DexConf
	Plugins struct {
		Rbac struct{ Enabled bool }
	}
	DisableSecurityGroupIngress bool
	NodeMonitorGracePeriod      kubernetes.TimePeriod
}

type ClusterYAML struct {
	ClusterName    types.ClusterName
	ReleaseChannel coreos.ReleaseChannel

	AmiId                       ec2.AmiId
	HostedZoneId                ec2.HostedZoneId
	SshAccessAllowedSourceCIDRs []net.IPNet

	AdminAPIEndpointName APIEndpointName
	ApiEndpoints         []APIEndpointConf

	KeyName           ec2.SSHKeyPairName
	SSHAuthorizedKeys []types.SSHAuthorizedKey

	Region ec2.Region
	//AvailabilityZone ec2.AvailabilityZone -- removed in favour of nodepools
	KMSKeyArn ec2.KMSKeyARN

	Controller ControllerConf

	//WorkerCount uint -- removing
	Worker struct {
		//apiEndpointName  -- removing
		NodePools []NodepoolConf
	}

	//WorkerCreationTimeout ec2.Timeout  -- removed in favour of nodepools
	//WorkerInstanceType ec2.InstanceType
	//WorkerRootVolumeSize
	//WorkerRootVolumeType
	//WorkerRootVolumeIOPS
	//WorkerTenancy
	//WorkerSpotPrice

	Etcd EtcdConf

	VPCConf

	ServiceCIDR  net.IPNet
	PodCIDR      net.IPNet
	DnsServiceIP net.IP
	MapPublicIPs bool //future:shouldn't it be per nodepool?

	TLSConf
	KubernetesContainerImages

	UseCalico              bool
	ElasticFileSystemId    ec2.EFSId
	SharedPersistentVolume bool
	ContainerRuntime       types.ContainerRuntime

	ManageCertificates      bool
	WaitSignal              WaitSignalConf // seems to be related to Controllers only, can be deprecate?
	KubeResourcesAutosave   struct{ Enabled bool }
	CloudWatchLogging       CloudWatchLoggingConf
	AmazonSsmAgent          AmazonSSMAgentConf
	KubeDNS                 struct{ NodeLocalResolver bool }
	CloudFormationStreaming bool

	Addons AddonsConf

	StackTags map[ec2.TagName]ec2.TagValue

	CustomSettings map[string]interface{}
}
