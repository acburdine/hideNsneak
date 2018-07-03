package deployer

type AzureProvider struct {
	Instances   []AzureInstance    `json:"instances"`
	DomainFront []AzureDomainFront `json:"domain_front"`
}

type AzureDomainFront struct{}

type AzureInstance struct {
	Count int
}

//Deprecated
// type azureCdnDeployer struct {
// 	HostName     string
// 	ProfileName  string
// 	EndpointName string
// 	Location     string
// }

// type azureDeployer struct {
// 	Location    string
// 	Count       int
// 	VMSize      string
// 	Environment string
// 	PrivateKey  string
// 	PublicKey   string
// }
