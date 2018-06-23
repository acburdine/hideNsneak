Hey you. Welcome to hideNsneak.
===============================
(What is hideNsneak?)


How to use the tool & use cases
-------------------------------
(You run commands to do this and this, for example if you wanna do this you do that)


Running locally
---------------
1. download go
2. download terraform
3. download ansible
4. download docker
5. clone git repo
6. Need to create a file under package main titled `secrets.go`

	```const tfvars = 
		aws_access_key = "YOUR_SECRET_KEY"
		aws_secret_key = "YOUR_SECRET_KEY"
		do_token = "YOUR_SECRET_KEY"
		azure_tenant_id = "YOUR_SECRET_KEY"
		azure_client_id = "YOUR_SECRET_KEY"
		azure_client_secret = "YOUR_SECRET_KEY"
		azure_subscription_id = "YOUR_SECRET_KEY"```
	

Commands & functionality
------------------------
(put commands, expected outputs and funcs here)


Organization
------------
_terraform --> has all the terraform related stuff

_test.go --> test file that is ignored by go code (underscores do that)
secrets.go --> ignored in gitignore. where you keep your secret keys
constants.go --> where you can find all the constants for deploying modules
lib.go --> utility functions
structs.go --> basic structs for all modules
terraformer.go --> wrapper for terraform functionality. This is where the good stuff happens.

Developing locally & contributions
----------------------------------

