package deployer

type GooglefrontConfigWrapper struct {
	ModuleName          string
	InvokeURI           string
	Host                string
	HostURL             string
	FunctionName        string
	SourceFile          string
	PackageFile         string
	RestrictUA          string
	RestrictSubnet      string
	RestrictHeader      string
	RestrictHeaderValue string
	Enabled             bool
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
