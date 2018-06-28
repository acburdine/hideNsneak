package deployer

type azureCdnDeployer struct {
	HostName     string
	ProfileName  string
	EndpointName string
	Location     string
}

type azureDeployer struct {
	Location    string
	Count       int
	VMSize      string
	Environment string
	PrivateKey  string
	PublicKey   string
}

type cloudFrontDeployer struct {
	Origin string
	Region string
}

type digitalOceanDeployer struct {
	Image        string
	Fingerprint  string
	PrivateKey   string
	PublicKey    string
	Size         string
	Count        int
	Region       string
	DefaultUser  string
	Name         string
	FirewallName string
}

type googleCloudDeployer struct {
	Region            string
	Project           string
	Count             int
	SSHUser           string
	SSHPubKeyFile     string
	SSHPrivateKeyFile string
	MachineType       string
	Image             string
}

type apiGatewayDeployer struct {
	TargetURI string
	StageName string
}

//Output Parsing Structs
type TerraformOutput struct {
	Master OuterLevel `json:"providers"`
}

type OuterLevel struct {
	ProviderValues Providers `json:"value"`
}

type Providers struct {
	AWSProvider    AWSProvider    `json:"AWS"`
	DoProvider     DOProvider     `json:"DO"`
	GoogleProvider GoogleProvider `json:"GOOGLE"`
	AzureProvider  AzureProvider  `json:"AZURE"`
}

type AWSProvider struct {
	Instances      []AWSInstance      `json:"instances"`
	API            []AWSApi           `json:"api"`
	DomainFront    []AWSDomainFront   `json:"domain_front"`
	SecurityGroups []AWSSecurityGroup `json:"security_group"`
}

type AWSInstance struct {
	UsEast1      AWSRegion `json:"us-east-1"`
	UsEast2      AWSRegion `json:"us-east-2"`
	UsWest1      AWSRegion `json:"us-west-1"`
	UsWest2      AWSRegion `json:"us-west-2"`
	CaCentral1   AWSRegion `json:"ca-central-1"`
	EuCentral1   AWSRegion `json:"eu-central-1"`
	EuWest1      AWSRegion `json:"eu-west-1"`
	EuWest2      AWSRegion `json:"eu-west-2"`
	EuWest3      AWSRegion `json:"eu-west-3"`
	ApNorthEast1 AWSRegion `json:"ap-northeast-1"`
	ApNorthEast2 AWSRegion `json:"ap-northeast-2"`
	ApNorthEast3 AWSRegion `json:"ap-northeast-3"`
	ApSouthEast1 AWSRegion `json:"ap-southeast-1"`
	ApSouthEast2 AWSRegion `json:"ap-southeast-2"`
	ApSouth1     AWSRegion `json:"ap-south-1"`
	Saeast1      AWSRegion `json:"sa-east-1"`
}

type AWSRegion struct {
	Config  AWSRegionConfig `json:"config"`
	IpIdMap IPIDMap         `json:"ip_id"`
}

type AWSRegionConfig struct {
	SecurityGroup   string `json:"hidensneak"`
	SecurityGroupID string `json:"aws_sg_id"`
	Count           string `json:"region_count"`
	CustomAmi       string `json:"custom_ami"`
	InstanceType    string `json:"aws_instance_type"`
	DefaultUser     string `json:"ec2_default_user"`
	KeypairName     string
	Region          string `json:"region"`
	PrivateKeyFile  string `json:"private_key_file"`
	PublicKeyFile   string `json:"public_key_file"`
}

type IPIDMap struct {
	IpMap map[string]string `json:"ip_id"`
}

type AWSApi struct {
}

type AWSDomainFront struct{}

type AWSSecurityGroup struct{}

type DOProvider struct{}

type GoogleProvider struct{}

type AzureProvider struct{}

//ReadList contains a list of all of the resources
//across different providers per region
type ReadList struct {
	ec2DeployerList          []AWSRegionConfig
	azureCdnDeployerList     []azureCdnDeployer
	azureDeployerList        []azureDeployer
	cloudFrontDeployerList   []cloudFrontDeployer
	digitalOceanDeployerList []digitalOceanDeployer
	googleCloudDeployerList  []googleCloudDeployer
	apiGatewayDeployerList   []apiGatewayDeployer
}
