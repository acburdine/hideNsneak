package deployer

// //State Parsing
// type TerraformState struct {
// 	Version int
// 	Serial  int
// 	Backend *TerraformBackend
// 	Modules []TerraformStateModule
// }

// // The structure of the "backend" section of the Terraform .tfstate file
// type TerraformBackend struct {
// 	Type   string
// 	Config map[string]interface{}
// }

// // The structure of a "module" section of the Terraform .tfstate file
// type TerraformStateModule struct {
// 	Path      []string
// 	Outputs   map[string]interface{}
// 	Resources map[string]interface{}
// }

//Deprecated
//ReadList contains a list of all of the resources
//across different providers per region
// type ReadList struct {
// 	azureCdnDeployerList     []azureCdnDeployer
// 	azureDeployerList        []azureDeployer
// 	cloudFrontDeployerList   []cloudFrontDeployer
// 	digitalOceanDeployerList []digitalOceanDeployer
// 	googleCloudDeployerList  []googleCloudDeployer
// 	apiGatewayDeployerList   []apiGatewayDeployer
// }

// type TerraformOutput struct {
// 	Master OuterLevel `json:"providers"`
// }

// type OuterLevel struct {
// 	ProviderValues Providers `json:"value"`
// }

// type Providers struct {
// 	AWSProvider    AWSProvider    `json:"AWS"`
// 	DOProvider     DOProvider     `json:"DO"`
// 	GoogleProvider GoogleProvider `json:"GOOGLE"`
// 	AzureProvider  AzureProvider  `json:"AZURE"`
// }

// type IPID struct {
// 	IPList []string `json:"ip"`
// 	IDList []string `json:"id"`
// }
