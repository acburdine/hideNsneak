package deployer

//Output Parsing Structs

type TerraformOutput struct {
	Master OuterLevel `json:"providers"`
}

type OuterLevel struct {
	ProviderValues Providers `json:"value"`
}

type Providers struct {
	AWSProvider    AWSProvider    `json:"AWS"`
	DOProvider     DOProvider     `json:"DO"`
	GoogleProvider GoogleProvider `json:"GOOGLE"`
	AzureProvider  AzureProvider  `json:"AZURE"`
}

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
