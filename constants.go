package main

const state = `
	terraform {
		backend "s3" {
		  bucket         = "hidensneak-terraform"
		  key            = "filename.tfstate"
		  dynamodb_table = "terraform-state-lock-dynamo"
		  region         = "us-east-1"
		  encrypt        = true
		}
	  }
`

const variables = `variable "do_token" {}

variable "aws_access_key" {}

variable "aws_secret_key" {}

variable "azure_tenant_id" {}

variable "azure_client_id" {}

variable "azure_client_secret" {}

variable "azure_subscription_id" {}
`

///////////////////// MODULES /////////////////////
const ec2Module = `
	module "aws-{{.Region}}" {
		source         		 = "modules/ec2-deployment"
		default_sg_name 	 = "{{.SecurityGroup}}"
		region_count   		 = {{.Count}}
		custom_ami 			 = "{{.CustomAmi}}"
		aws_instance_type	 = "{{.InstanceType}}"
		ec2_default_user	 = "{{.DefaultUser}}"
		aws_access_key 		 = "${var.aws_access_key}"
		aws_secret_key 		 = "${var.aws_secret_key}"
		aws_region    		 = "{{.Region}}"
		aws_new_keypair      = "{{.NewKeypair}}"
		aws_keypair_name     = "{{.KeypairName}}"
		aws_private_key_file = "{{.PrivateKeyFile}}"
		aws_public_key_file  = "{{.PublicKeyFile}}"
	}
`

const azureCdnModule = `
	module "azure-cdn-{{.Endpoint}}" {
		source                  = "modules/azure-cdn-deployment"
		azure_subscription_id   = "${var.azure_subscription_id}"
		azure_tenant_id         = "${var.azure_tenant_id}"
		azure_client_id         = "${var.azure_client_id}"
		azure_client_secret     = "${var.azure_client_secret}"
		azure_cdn_hostname      = "{{.HostName}}"
		azure_cdn_profile_name  = "{{.ProfileName}}"
		azure_cdn_endpoint_name = "{{.EndpointName}}"
		azure_location          = "{{.Location}}"
		
	}
`

//TODO: need to run removeSpaces() on region when this is set
const azureModule = `
	module "azure-{{.Location}}" {
		source                = "modules/azure-deployment"
		azure_subscription_id = "${var.azure_subscription_id}"
		azure_tenant_id       = "${var.azure_tenant_id}"
		azure_client_id       = "${var.azure_client_id}"
		azure_client_secret   = "${var.azure_client_secret}"
		azure_location        = "{{.Location}}"
		azure_instance_count  = {{.InstanceCount}}
		azure_vm_size 		  = "{{.VMSize}}"
		azure_environment 	  = "{{.Environment}}"

	}
`

const cloudfrontModule = `
	module "cloudfront-{{.Region}}" {
		source            = "modules/cloudfront-deployment"
		cloudfront_origin = "{{.Origin}}"
		aws_access_key    = "${var.aws_access_key}"
		aws_secret_key    = "${var.aws_secret_key}"
		aws_region 		  = "{{.Region}}"
	}
`

const digitalOceanModule = `
	module "digital-ocean-{{.Region}}" {
		source           = "modules/droplet-deployment"
		do_token         = "${var.do_token}"
		do_image         = "{{.Image}}"
		pvt_key          = "{{.PrivateKey}}"
		ssh_fingerprint  = "{{.SSHFingerprint}}"
		do_region        = "{{.Region}}"
		do_size          = "{{.Size}}"
		do_count         = {{.Count}}
		do_default_user  = "{{.DefaultUser}}"
		do_name 		 = "{{.Name}}"
		do_firewall_name = "{{.FirewallName}}"
		do_ssh_source_ip = "{{.SSHSourceIP}}"
	}
`

const googleCloudModule = `
	module "google-cloud-{{.Region}}" {
		source               	 = "modules/gcp-deployment"
		gcp_region          	 = "{{.Region}}"
		gcp_project          	 = "{{.Project}}"
		gcp_instance_count   	 = {{.InstanceCount}}
		gcp_ssh_user         	 = "{{.SSHUser}}"
		gcp_ssh_pub_key_file 	 = "{{.SSHPubKeyFile}}"
		gcp_ssh_private_key_file = "{{.SSHPrivateKeyFile}}"
		gcp_machine_type	 	 = "{{.MachineType}}"
		gcp_image			 	 = "{{.Image}}"
	}
`

const apiGatewayModule = `
	module "apigateway-{{.TargetUri}}" {
		source 				 = "modules/api-gateway"
		aws_access_key    	 = "${var.aws_access_key}"
		aws_secret_key    	 = "${var.aws_secret_key}"
		aws_api_target_uri 	 = "{{.TargetURI}"
		aws_api_stage_name	 = "{{.StageName}}"
  	}
`

//////////////////////////////DIALOGUES///////////////////////////////////
const welcomeMessage = `
	Welcome to hideNsneak. Today's menu of cloud infrastructure:
	- EC2
	- API Gateway
	- Digital Ocean (DO)
	- Google Cloud Provider (GCP)
	- Azure CDN
	- Azure

	To start, run one of these commands: 
	- help : get list of commands to run
	- deploy : deploy new servers
	- destroy : destroy servers
	- start : start stopped servers
	- stop : stop running servers
	- list : list servers
	- shell : start and interact with a command shell on a server
	- socks-add : create a SOCKS proxy with a live server
	- socks-kill : kill an existing SOCKS proxy
	- domainfront : create a new domain front
	- domainfront-list : list existing domain fronts
	- nmap : initiate an nmapn scan and distriute it among hosts
	- proxyconf : print proxychains and SOCKSd configurations for SOCKS proxies
	- send : send a file or directory
	- get : retrieve a file or directory
	- firewall : create a firewall
	- firewall-list : list existing firewalls
	- quit : exit program
	- exit : exit program
`
const help = `
	- help : get list of commands to run
	- deploy : deploy new servers
	- destroy : destroy servers
	- start : start stopped servers
	- stop : stop running servers
	- list : list servers
	- shell : start and interact with a command shell on a server
	- socks-add : create a SOCKS proxy with a live server
	- socks-kill : kill an existing SOCKS proxy
	- domainfront : create a new domain front
	- domainfront-list : list existing domain fronts
	- nmap : initiate an nmapn scan and distriute it among hosts
	- proxyconf : print proxychains and SOCKSd configurations for SOCKS proxies
	- send : send a file or directory
	- get : retrieve a file or directory
	- firewall : create a firewall
	- firewall-list : list existing firewalls
	- quit : exit program
	- exit : exit program
`
const ascii = ` __     __     __         _______                              __    
|  |--.|__|.--|  |.-----.|    |  |.-----..-----..-----..---.-.|  |--.
|     ||  ||  _  ||  -__||       ||__ --||     ||  -__||  _  ||    < 
|__|__||__||_____||_____||__|____||_____||__|__||_____||___._||__|__|
                                                                     `

const prompt = "<hideNsneak> "

const shutdown = "<hideNsneak> Goodbye"

const doesntExist = "<hideNsneak> Looks like that command doesn't exist. Try running `help`."
