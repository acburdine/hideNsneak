Welcome to hideNsneak.
===============================
![Alt text](assets/logo.png "hideNsneak")
This application assists in managing attack infrastructure by providing an interface to rapidly deploy, manage, and take down various cloud services. These include VMs, domain fronting, Cobalt Strike servers, API gateways, and firewalls.


How to use the tool & use cases
-------------------------------
(You run commands to do this and this, for example if you wanna do this you do that)

For hosts, need to assume Ubuntu 16.04 Linux host


Running locally
---------------
1. download go
2. download terraform
3. download ansible
4. download docker
5. `git clone https://github.com/rmikehodges/hideNsneak.git`
6. `cd hideNsneak/main`
7. `go get github.com/rmikehodges/hideNsneak/cloud`
8. `go get github.com/rmikehodges/hideNsneak/misc`
9. `go get github.com/rmikehodges/hideNsneak/sshext`
10. `go run main.go`
11. fill in the values in config.yaml with API keys, file paths, etc
12. set up your ssh key in your config.yaml file with all cloud provider you'd like to use (AWS, Google, Digital Ocean)
6. Need to create a file under package main titled `secrets.go`
7. Need to create ~/.terraform.d/plugins and add https://github.com/nbering/terraform-provider-ansible/ to it

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

terraformer.go --> wrapper for terraform functionality. This is where the good stuff happens

Developing locally & contributions
----------------------------------

We would love to have you contribute to hideNsneak. Feel free to fork the repo and start contributing, we will review pull requests as we receive them. If you feel like some things need improvement or some features need adding, feel free to open up an issue and hopefully -- someone will pick it up. 

For those who decide to contribute regularly... We've got some real comfy t-shirts for you.

License 
-------

MIT