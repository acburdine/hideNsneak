package main

type ec2Deployer struct {
	SecurityGroup  string
	Count          int
	CustomAmi      string
	InstanceType   string
	DefaultUser    string
	Region         string
	NewKeypair     bool
	KeypairName    string
	PrivateKeyFile string
	PublicKeyFile  string
}

type azureCdnDeployer struct {
	HostName     string
	ProfileName  string
	EndpointName string
	Location     string
}

type azureDeployer struct {
	Location      string
	InstanceCount string
	VMSize        string
	Environment   string
}

type cloudFrontDeployer struct {
	Origin string
	Region string
}

type digitalOceanDeployer struct {
	Image          string
	Fingerprint    string
	PrivateKey     string
	SSHFingerprint string
	Size           string
	Count          int
	Region         string
	DefaultUser    string
	Name           string
	FirewallName   string
	SSHSourceIP    string
}

type googleCloudDeployer struct {
	Region            string
	Project           string
	InstanceCount     int
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

type readList struct {
	ec2DeployerList          []ec2Deployer
	azureCdnDeployerList     []azureCdnDeployer
	azureDeployerList        []azureDeployer
	cloudFrontDeployerList   []cloudFrontDeployer
	digitalOceanDeployerList []digitalOceanDeployer
	googleCloudDeployerList  []googleCloudDeployer
	apiGatewayDeployerList   []apiGatewayDeployer
}
