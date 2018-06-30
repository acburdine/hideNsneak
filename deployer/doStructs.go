package deployer

type DOProvider struct {
	Instances []DOInstance `json:"instances"`
}

type DOInstance struct {
	Config  DORegionConfig    `json:"config"`
	IPIDMap map[string]string `json:"ip_id"`
}

type DORegionConfig struct {
	ModuleName  string
	Image       string `json:"image"`
	PrivateKey  string `json:"private_key_file"`
	Fingerprint string `json:"fingerprint"`
	Size        string `json:"size"`
	Count       int    `json:"region_count"`
	Region      string `json:"region"`
	DefaultUser string `json:"default_user"`
}

type DOConfigWrapper struct {
	Config    DORegionConfig
	RegionMap map[string]int
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
