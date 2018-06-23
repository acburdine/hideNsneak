package main

type awsDeployer struct {
	Count         int
	Region        string
	SecurityGroup string
	CustomAmi     string
	KeypairFile   string
	KeypairName   string
	NewKeypair    bool
}
