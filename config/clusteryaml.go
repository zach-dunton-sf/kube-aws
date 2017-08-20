package config

import (
	"github.com/kubernetes-incubator/kube-aws-ng/model"
	"github.com/kubernetes-incubator/kube-aws-ng/types"
	"github.com/kubernetes-incubator/kube-aws-ng/types/coreos"
	"github.com/kubernetes-incubator/kube-aws-ng/types/dex"
	"github.com/kubernetes-incubator/kube-aws-ng/types/ec2"
	"github.com/kubernetes-incubator/kube-aws-ng/types/kubernetes"
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
	InstanceCIDR     types.IPNet
	Private          bool
	SubnetIdConf
	NatGateway NGWIdConf
	RouteTable RouteTableIdConf
}

type APIEndpointLoadBalancer struct { // need invariants generator
	Id              ec2.ELBName
	CreateRecordSet bool  `yaml:"createRecordSet"`
	RecordSetTTL    uint `yaml:"recordSetTTL"`

	Subnets                     []SubnetConfRef
	Private                     bool
	HostedZone                  ec2.HostedZoneId `yaml:hostedZone`
	ApiAccessAllowedSourceCIDRs []types.IPNet `yaml:ApiAccessAllowedSourceCIDRs`
	SecurityGroupIds            []ec2.SecurityGroupId
}

type APIEndpointName string

type APIEndpointConf struct {
	Name    APIEndpointName 
	DnsName types.DNSName `yaml:"dnsName"`

	LoadBalancer APIEndpointLoadBalancer `yaml:"loadBalancer"`
}

type ASGConf struct {
	MinSize                            uint `yaml:"minSize"`
	MaxSize                            uint `yaml:"maxSize"`
	RollingUpdateMinInstancesInService uint `yaml:"rollingUpdateMinInstancesInService"`
}

type IAMConf struct {
	Role struct {
		ManagedPolicies []ec2.IAMPolicyARN  `yaml:"managedPolicies"`
		InstanceProfile ec2.InstanceProfileARN `yaml:"instanceProfile"`
	}
}

type VolumeConf struct {
	Size uint `yaml:"size"`
	Type ec2.VolumeType `yaml:"type"`
	Iops uint `yaml:"iops"`
}

type VolumeMountConf struct {
	VolumeConf
	Device types.BlockDeviceName `yaml:"device"`
	Path   types.FilesystemPath `yaml:"path"`
}

type MaybeEncryptedOrEphemeralVolume struct {
	VolumeConf
	Encrypted bool `yaml:encrypted` // conflicts: Ephemeral
	Ephemeral bool `yaml:ephemeral` // conflicts: Encrypted, broken
}

type InstanceCommonDescrEmbed struct {
	Count              uint `yaml:"count"`
	CreateTimeout      ec2.Timeout `yaml:"createTimeout"`
	InstanceType       ec2.InstanceType `yaml:"instanceType"`
	Tenancy                   ec2.InstanceTenancy `yaml:"tenancy"` //new for Controller
	RootVolume         VolumeConf `yaml:rootVolume` //validate: must be empty/default device and path
	SecurityGroupIds   []ec2.SecurityGroupId `yaml:SecurityGroupIds`
	IAM                IAMConf `yaml:"iam"`
	Subnets            []SubnetConfRef
	CustomFiles        []map[string]string
	CustomSystemdUnits []map[string]string
	KeyName        ec2.SSHKeyPairName  // new for Etcd,Controller
	ReleaseChannel coreos.ReleaseChannel //  new for Etcd,Controller
	AmiId          ec2.AmiId  // new for Etcd, Controller
	ManagedIamRoleSuffix      ec2.IAMRoleName `yaml:mangedIamRoleName`  //new for Etcd,Ctrl
}

type ControllerConf struct {
	InstanceCommonDescrEmbed
	NodeLabels         map[kubernetes.LabelName]kubernetes.LabelValue
	AutoScalingGroup   ASGConf
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
	JournaldCloudWatchLogsImage        model.Image `yaml:"journaldCloudWatchLogsImage,omitempty"`
}

type KubernetesContainerImages struct {
	KubernetesVersion string
	ContainerImages
}

type WaitSignalConf struct {
	Enabled      bool 
	MaxBatchSize uint `yaml:"maxBatchSize"`//validate: >0
}

type NodepoolConf struct {
	InstanceCommonDescrEmbed `yaml:,inline`
	Name             NodePoolName
	LoadBalancer     struct {
		Enabled bool
		Names   []ec2.ELBName
		//SecurityGroupIds []ec2.SecurityGroup -- removed, duplicate of SGs above
	}
	ApiEndpointName APIEndpointName
	TargetGroup     struct {
		Enabled bool
		Arns    []ec2.ALBTargetGroupARN
		//SecurityGroupIds []ec2.SecurityGroup -- removed, duplicate of SGs above
	}
	VolumeMounts              []VolumeMountConf
	NodeStatusUpdateFrequency kubernetes.TimePeriod // 10s, 5h
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
	KubernetesContainerImages
	SSHAuthorizedKeys  []types.SSHAuthorizedKey
	CustomSettings     map[string]string
}

type EtcdConf struct {
	InstanceCommonDescrEmbed `yaml:inline`
	DataVolume       MaybeEncryptedOrEphemeralVolume

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
}

type AutoScalingConf struct {
	ClusterAutoScaler struct {
		Enabled bool
	} `yaml:clusterAutoScaler`
}

type EtcAwsEnvironmentConf struct {
	Enabled     bool
	Environment map[string]string
}

type VPCConf struct {
	//VpcId ec2.VPCId -- removed as deprecated
	Vpc struct {
		Id                ec2.VPCId     //conflicts: IdFromStackOutput
		IdFromStackOutput ec2.StackName `yaml:idFromStackOutput` //conflicts: Id;
		//InternetGatewayId ec2.IGWId -- removed as deprecated
	} `yaml:vpc`

	InternetGateway struct {
		Id                ec2.IGWId     //conflicts: IdFromStackOutput
		IdFromStackOutput ec2.StackName `yaml:idFromStackOutput` //conflicts: Id
	}

	// RouteTableId ec2.RouteTableId -- removed in favour of subnets[].routeTable.id

	VpcCIDR types.IPNet `yaml:vpcCIDR` //future: should be conflicting with vpcID
	//InstanceCIDR net.IPNet // -- reomved in favour of nodepools[] and subnets[]

	Subnets []SubnetConf
}

type TLSConf struct {
	tlsCADurationDays   uint `default:"3650"`
	tlsCertDurationDays uint `default:"365"`
}

type CloudWatchLoggingConf struct {
	Enabled         bool
	RetentionInDays uint `yaml:retentionInDays`
	LocalStreaming  struct {
		Enabled  bool
		Filter   string
		Interval uint
	} `yaml:localStreaming`
}

type AmazonSSMAgentConf struct {
	Enabled     bool
	DownloadURL types.URL
	Sha1Sum     types.SHA1SUM
}

type AuditLogConf struct {
	Enabled bool
	MaxAge  uint `yaml:maxAge`
	LogPath string `yaml:logPath`
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
	URL             types.URL
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
	ClusterName    types.ClusterName `yaml:"clusterName"`
	ReleaseChannel coreos.ReleaseChannel `yaml:"releaseChannel"`

	AmiId                       ec2.AmiId `yaml:"amiId"`
	HostedZoneId                ec2.HostedZoneId `yaml:"hostedZoneId"`
	SshAccessAllowedSourceCIDRs []types.IPNet `yaml:"sshAccessAllowedSourceCIDRs"`

	AdminAPIEndpointName APIEndpointName `yaml:"adminAPIEndpointName"`
	ApiEndpoints         []APIEndpointConf `yaml:"apiEndpoints"`

	KeyName           ec2.SSHKeyPairName `yaml:"keyName"`
	SSHAuthorizedKeys []types.SSHAuthorizedKey `yaml:"sshAuthorizedKeys"`

	Region ec2.Region `yaml:"region"`
	//AvailabilityZone ec2.AvailabilityZone -- removed in favour of nodepools
	KMSKeyArn ec2.KMSKeyARN `yaml:"kmsKeyArn"`

	Controller ControllerConf `yaml:"controller"`

	//WorkerCount uint -- removing
	Worker struct {
		//apiEndpointName  -- removing
		NodePools []NodepoolConf `yaml:"nodePools"`
	} `yaml:"worker"`

	//WorkerCreationTimeout ec2.Timeout  -- removed in favour of nodepools
	//WorkerInstanceType ec2.InstanceType
	//WorkerRootVolumeSize
	//WorkerRootVolumeType
	//WorkerRootVolumeIOPS
	//WorkerTenancy
	//WorkerSpotPrice

	Etcd EtcdConf `yaml:"etcd"`

	VPCConf `yaml:",inline"`

	ServiceCIDR  types.IPNet `yaml:"serviceCIDR"`
	PodCIDR      types.IPNet `yaml:"podCIDR"`
	DnsServiceIP types.IP `yaml:"dnsServiceIP"`
	MapPublicIPs bool  `yaml:"mapPublicIPs"` //future:shouldn't it be per nodepool?

	TLSConf
	KubernetesContainerImages

	UseCalico              bool  `yaml:"useCalico"`
	ElasticFileSystemId    ec2.EFSId  `yaml:"elasticFileSystemId"`
	SharedPersistentVolume bool `yaml:"sharedPersistentVolume"`
	ContainerRuntime       types.ContainerRuntime `yaml:"containerRuntime"`

	ManageCertificates      bool  `yaml:"manageCertificates"`
	WaitSignal              WaitSignalConf  `yaml:"waitSignal"` // seems to be related to Controllers only, can be deprecate?
	KubeResourcesAutosave   struct{ Enabled bool } `yaml:"kubeResourcesAutosave"`
	CloudWatchLogging       CloudWatchLoggingConf `yaml:"cloudWatchLogging"`
	AmazonSsmAgent          AmazonSSMAgentConf `yaml:"amazonSsmAgent"`
	KubeDNS                 struct{ NodeLocalResolver bool } `yaml:"kubeDNS"`
	CloudFormationStreaming bool `yaml:"cloudFormationStreaming"`

	Addons AddonsConf  `yaml:"addons"`

	StackTags map[ec2.TagName]ec2.TagValue  `yaml:"StackTags"`

	CustomSettings map[string]interface{}  `yaml:"customSettings"`
}
