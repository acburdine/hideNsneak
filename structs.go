package structs

type Aws_Deployer struct {
	Count          int
	Region         string
	Security_Group string
	Custom_Ami     string
	Keypair_File   string
	Keypair_Name   string
	New_Keypair    bool
}
