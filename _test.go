package main 

	//Creating a test array
	tester1 := ec2Deployer{
		Count:         1,
		Region:        "us-east-1",
		SecurityGroup: "tester1243",
		KeypairFile:   "/Users/mike.hodges/.ssh/do_rsa.pub",
		KeypairName:   "do_rsa",
		NewKeypair:    false,
	}
	tester2 := ec2Deployer{
		Count:         1,
		Region:        "us-west-1",
		SecurityGroup: "tester1243",
		KeypairFile:   "/Users/mike.hodges/.ssh/do_rsa.pub",
		KeypairName:   "do_rsa",
		NewKeypair:    false,
	}
	tester3 := ec2Deployer{
		Count:         1,
		Region:        "eu-west-1",
		SecurityGroup: "tester1243",
		KeypairFile:   "/Users/mike.hodges/.ssh/do_rsa.pub",
		KeypairName:   "do_rsa",
		NewKeypair:    false,
	}
	testers := [...]ec2Deployer{tester1, tester2, tester3}