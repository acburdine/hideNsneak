package deployer

type GoogleProvider struct {
	Instances []GoogleInstance `json:"instances"`
}

type GoogleInstance struct {
	Count int
}

//Deprecated
// type googleCloudDeployer struct {
// 	Region            string
// 	Project           string
// 	Count             int
// 	SSHUser           string
// 	SSHPubKeyFile     string
// 	SSHPrivateKeyFile string
// 	MachineType       string
// 	Image             string
// }
