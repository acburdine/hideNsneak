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

* hideNsneak can 

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

1. install [go](https://golang.org/dl/)
2. install [terraform](https://www.terraform.io/intro/getting-started/install.html)
3. install [ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)
4. download zip file [custom providers](https://github.com/nbering/terraform-provider-ansible/) --> then, `cd $HOME/.terraform.d/` and `mkdir plugins`--> then, download the executable for your platform here `https://github.com/nbering/terraform-provider-ansible/releases` and unzip it in `$HOME/.terraform.d/plugins` (make sure you move the binary into the /plugins/ directory)
5. `git clone https://github.com/rmikehodges/hideNsneak.git`
6. `cd hideNsneak`
7. `go get -u github.com/spf13/cobra/cobra`
8. `go get -u github.com/aws/aws-sdk-go/aws`
9. Fill in values in `config.yaml` with your keys and filepaths for the cloud providers you'd like to use:
		```
		aws_access_key = "YOUR_SECRET_KEY"
		aws_secret_key = "YOUR_SECRET_KEY"
		do_token = "YOUR_SECRET_KEY"
		azure_tenant_id = "YOUR_SECRET_KEY"
		azure_client_id = "YOUR_SECRET_KEY"
		azure_client_secret = "YOUR_SECRET_KEY"
		azure_subscription_id = "YOUR_SECRET_KEY"
		```
10. run `go build -o hidensneak main.go` to build the hidensneak executable
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
* `secrets.go` --> a file that you write yourself, with all your secret stuff


Contributions
-------------
We would love to have you contribute to hideNsneak. Feel free to fork the repo and start contributing, we will review pull requests as we receive them. If you feel like some things need improvement or some features need adding, feel free to open up an issue and hopefully -- someone will pick it up. 


License 
-------
MIT/BSD