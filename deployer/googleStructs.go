package deployer

type GooglefrontConfigWrapper struct {
	ModuleName          string
	FrontedDomain       string
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
