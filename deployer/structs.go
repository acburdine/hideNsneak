package deployer

type ec2Deployer struct {
	SecurityGroup   string
	SecurityGroupID string
	Count           int
	CustomAmi       string
	InstanceType    string
	DefaultUser     string
	Region          string
	NewKeypair      bool
	KeypairName     string
	PrivateKeyFile  string
	PublicKeyFile   string
}

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

//ReadList contains a list of all of the resources
//across different providers per region
type ReadList struct {
	ec2DeployerList          []ec2Deployer
	azureCdnDeployerList     []azureCdnDeployer
	azureDeployerList        []azureDeployer
	cloudFrontDeployerList   []cloudFrontDeployer
	digitalOceanDeployerList []digitalOceanDeployer
	googleCloudDeployerList  []googleCloudDeployer
	apiGatewayDeployerList   []apiGatewayDeployer
}
