package main

type ec2Deployer struct {
	Count         int
	Region        string
	SecurityGroup string
	CustomAmi     string
	KeypairFile   string
	KeypairName   string
	NewKeypair    bool
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
}

type cloudFrontDeployer struct {
	Origin string
}

type digitalOceanDeployer struct {
	Image          string
	Fingerprint    string
	PrivateKey     string
	SSHFingerprint string
	Size           string
	Count          int
}

type googleCloudDeployer struct {
	Region        string
	Project       string
	InstanceCount int
	SSHUser       string
	SSHPubKeyFile string
	MachineType   string
}

type apiGatewayDeployer struct {
	TargetURI string
}
