Welcome to hideNsneak. *Note: The official release of the tool is during Black Hat Arsenal this year, so this is WIP*
===============================
![Alt text](assets/logo.png "hideNsneak")
This application assists in managing attack infrastructure for penetration testers by providing an interface to rapidly deploy, manage, and take down various cloud services. These include VMs, domain fronting, Cobalt Strike servers, API gateways, and firewalls.


Table of contents 
------------------
  * [Overview](#overview)
  * [Running locally](#running-locally)
  * [Commands](#commands)
  * [Organization](#organization)
  * [Contributions](#contributions)
  * [License](#license)


Overview
---------
hideNsneak provides a simple interface that allows penetration testers to build ephemeral infrastructure -- one that requires minimal overhead. 
hideNsneak can: 

* *`deploy`, `destroy`, and `list`*
	1. Cloud instances via EC2, Google Cloud, Digital Ocean, Azure, and Alibaba Cloud
	2. API Gateway (AWS)
	3. Domain fronts via CloudFront and Azure Cloudfront

* *Proxy into said infrastructure*
* *Send and receive files*
* *Port scanning via NMAP*
* *Remote installations of Burp Collab, Cobalt Strike, Socat, LetsEncrypt, GoPhish, and SQLMAP*


Running locally
---------------
At this time, all hosts are assumed `Ubuntu 16.04 Linux`. In the future, we're hoping to add on a docker container to decrease initial setup time. 

1. Create a new AWS S3 bucket for your state
	- Ensure this is not public as it will hold your terraform state
2. run `./setup.sh`
3. `cp config/example-config.json to config/config.json` 
	- fill in the values
	- aws_access_id, aws_secret_key, aws_bucket_name are required at minimum
4. `go build -o hidensneak main.go`
11. now you can use with `hidensneak [command]`


Commands
---------
* `hidensneak help` --> run this anytime to get available commands 
* `hidensneak instance deploy`
* `hidensneak instance destroy`
* `hidensneak instance list`
* `hidensneak api deploy`
* `hidensneak api destroy`
* `hidensneak api list`
* `hidensneak domainfront enable`
* `hidensneak domainfront disable`
* `hidensneak domainfront deploy`
* `hidensneak domainfront destroy`
* `hidensneak domainfront list`
* `hidensneak socks deploy`
* `hidensneak socks list`
* `hidensneak socks destroy`
* `hidensneak socks proxychains`
* `hidensneak socks socksd`
* `hidensneak install burp`
* `hidensneak install cobaltstrike`
* `hidensneak install socat`
* `hidensneak install letsencrypt`
* `hidensneak install gophish`
* `hidensneak install nmap`
* `hidensneak install sqlmap`
* `hidensneak file push`
* `hidensneak file pull`

For all commands, you can run `--help` after any of them to get guidance on what flags to use.


Organization
------------
* `_terraform` --> stuff related to deploying, destroying, and listing infrastucture
* `_ansible` --> stuff related to ssh
* `_assets` --> random assets for the beauty of this project
* `_cmd` --> frontend interface 
* `_deployer` --> backend commands and structs
* `main.go` --> where the magic happens 


Contributions
-------------
We would love to have you contribute to hideNsneak. Feel free to pull the repo and start contributing, we will review pull requests as we receive them. If you feel like some things need improvement or some features need adding, feel free to open up an issue and hopefully -- someone will pick it up. 

Miscellaneous
-------------
All firewall rules are performed at the host level. All instances are restricted to port 22 upon deployment. In order to achieve this, a default security group is made in every AWS region named hideNsneak which is by default to the world. This security group should not be used for any other infrastructure.

If there are issues dealing with existing resources causing you to not be able to deploy or destroy
then cd into the terraform directory and run the following command for the resource that is giving the error

terraform state rm (resource name)

This removes the resource from your state allowing you to proceed but cleanup of the resource must be done manually


License 
-------
MIT
