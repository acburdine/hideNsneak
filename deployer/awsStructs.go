package deployer

type AWSApiConfigWrapper struct {
	ModuleName string
}

type AWSDomainFront struct{}

type AWSSecurityGroup struct{}

type cloudFrontDeployer struct {
	Origin string
	Region string
}

type EC2ConfigWrapper struct {
	ModuleName   string
	InstanceType string
	DefaultUser  string
	DefaultSG    string
	SgID         string
	PrivateKey   string
	PublicKey    string
	KeyPairName  string
	RegionMap    map[string]int
}
