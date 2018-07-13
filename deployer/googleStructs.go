package deployer

type googleDomainFront struct {
	Host                string
	HostURL             string
	RestrictUA          string
	RestrictSubnet      string
	RestrictHeader      string
	RestrictHeaderValue string
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
