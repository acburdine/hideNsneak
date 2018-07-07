package deployer

type AWSDomainFront struct{}

type AWSSecurityGroup struct{}

type AWSApiConfigWrapper struct {
	ModuleName string
	Name       string
	TargetURI  string
	InvokeURI  string
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

type CloudfrontConfigWrapper struct {
	ModuleName string
	ID         string
	Provider   string
	URL        string
	Origin     string
	Status     string
	Enabled    string
	Etag       string
}
