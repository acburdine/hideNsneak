package deployer

const tfMainFile = "terraform/main.tf"
const tfVariablesFile = "terraform/variables.tf"
const tfVarsFile = "terraform/terraform.tfvars"
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
		aws_sg_id			 = "{{.SecurityGroupID}}"
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
		ansible_groups       = "[]"
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
		azure_public_key_file = "{{.PublicKey}}
		azure_private_key_file = "{{.PrivateKey}}
		ansible_groups       = "[]"
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
		ansible_groups       = "[]"
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
		ansible_groups       = "[]"
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
