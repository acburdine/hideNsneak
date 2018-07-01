package deployer

type GoogleProvider struct {
	Instances []GoogleInstance `json:"instances"`
}

type GoogleInstance struct {
	Config  GoogleRegionConfig `json:"config"`
	IPIDMap map[string]string  `json:"ip_id"`
}

type GoogleRegionConfig struct {
	Count int
}

//Deprecated
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
