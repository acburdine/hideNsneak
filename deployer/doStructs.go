package deployer

type Token struct {
	AccessToken string
}

type DOProvider struct {
	Instances []DOInstance `json:"instances"`
}

type DOInstance struct {
	ModuleName  string
	Image       string `json:"image"`
	PrivateKey  string `json:"private_key_file"`
	Fingerprint string `json:"fingerprint"`
	Size        string `json:"size"`
	Count       int    `json:"region_count,string"`
	Region      string `json:"region"`
	DefaultUser string `json:"default_user"`
}

type DOConfigWrapper struct {
	ModuleName  string
	Image       string
	PrivateKey  string
	Fingerprint string
	Size        string
	DefaultUser string
	RegionMap   map[string]int
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
