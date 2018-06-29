package deployer

type DOProvider struct {
	Instances []DOInstance `json:"instances"`
}

type DOInstance struct {
	Config  DORegionConfig    `json:"config"`
	IPIDMap map[string]string `json:"ip_id"`
}

type DORegionConfig struct {
}

//Deprecated
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
