package deployer

type AWSProvider struct {
	Instances      []AWSInstance      `json:"instances"`
	API            []AWSApi           `json:"api"`
	DomainFront    []AWSDomainFront   `json:"domain_front"`
	SecurityGroups []AWSSecurityGroup `json:"security_group"`
}

type AWSInstance struct {
	ModuleName      string
	SecurityGroup   string
	SecurityGroupID string
	Count           int
	CustomAmi       string
	InstanceType    string
	DefaultUser     string
	Region          string
	PrivateKeyFile  string
	PublicKeyFile   string
}

type AWSApi struct {
}

type AWSDomainFront struct{}

type AWSSecurityGroup struct{}

type cloudFrontDeployer struct {
	Origin string
	Region string
}

type EC2ConfigWrapper struct {
	ModuleName      string
	SecurityGroup   string
	SecurityGroupID string
	Count           int
	Ami             string
	InstanceType    string
	DefaultUser     string
	Region          string
	PrivateKeyFile  string
	PublicKeyFile   string
	RegionMap       map[string]int
}

//Deprecated
type apiGatewayDeployer struct {
	TargetURI string
	StageName string
}
