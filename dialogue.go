package main

////////////////////////////// dialogues for CLI ///////////////////////////////////
const ascii = ` __     __     __         _______                              __    
|  |--.|__|.--|  |.-----.|    |  |.-----..-----..-----..---.-.|  |--.
|     ||  ||  _  ||  -__||       ||__ --||     ||  -__||  _  ||    < 
|__|__||__||_____||_____||__|____||_____||__|__||_____||___._||__|__|
                                                                     `

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

const prompt = "<hideNsneak> "
const shutdown = "<hideNsneak> Goodbye"
const doesntExist = "<hideNsneak> Looks like that command doesn't exist. Try running `help`."

////////////////// deploy /////////////////////
const chooseProviders = "<hideNsneak/deploy> Enter the cloud providers you would like to use, separated by commas. [Default: EC2,DO, Google]: "
const unknownProvider = "<hideNsneak/deploy> Unknown Cloud Provider, please check your input. Choices: EC2, DO, Google, Azure, AzureCDN, ApiGateway"

// EC2
const setupEC2 = "<hideNSneak/deploy/setup/EC2> Please enter your EC2 deploy setup preferences..."
const secretKeyEC2 = "<hideNSneak/deploy/setup/EC2> Secret key is not set up. Enter your AWS secret key: "
const accessKeyEC2 = "<hideNSneak/deploy/setup/EC2> Access key is not set up. Enter your AWS access key: "
const regionsEC2 = "<hideNSneak/deploy/setup/EC2> Enter regions to deploy in, separated by commas: "
const numServersToDeployEC2 = "<hideNSneak/deploy/setup/EC2> Enter the number of servers to deploy [this will be distributed across desired regions]: "
const privateKeyFileEC2 = "<hideNSneak/deploy/setup/EC2> Enter filepath to private key file: "

const keyPairNameEC2 = "<hideNSneak/deploy/setup/EC2> Enter keypair name: "

// if keypair name not found, then ask for publickeyfile
const publicKeyFileEC2 = "<hideNSneak/deploy/setup/EC2> Enter filepath to public key file: "
const defaultSecurityGroupNameEC2 = "<hideNSneak/deploy/setup/EC2> Enter name for default security group: "

// optional
const wantInstanceTypeEC2 = "<hideNSneak/deploy/setup/EC2> Would you like to specify instance types? [yes/no]: "
const instanceTypeEC2 = "<hideNSneak/deploy/setup/EC2> Enter instance type: "
const wantDefaultUserEC2 = "<hideNSneak/deploy/setup/EC2> Would you like to specify instance types? [yes/no]: "
const defaultUserEC2 = "<hideNSneak/deploy/setup/EC2> Enter default user: "
const wantCustomAmiEC2 = "<hideNSneak/deploy/setup/EC2> Would you like to specify instance types? [yes/no]: "
const customAmiEC2 = "<hideNSneak/deploy/setup/EC2> Enter custom AMI: "

const invalidResponseEC2 = "<hideNSneak/deploy/setup/EC2> Invalid response. Enter appropriate response or refer to EC2 deployment specifications for guidance on naming conventions."
const confirmPreferencesEC2 = "<hideNSneak/deploy/setup/EC2> Please confirm your setup below. Choose a number to edit. [yes/int]"
const confirmSaveEC2 = "<hideNSneak/deploy/setup/EC2> Preferences saved. Moving on..."

// Digital Ocean
const setupDO = "<hideNSneak/deploy/setup/DO> Please enter your Digital Ocean deploy setup preferences..."
const numServersToDeployDO = "<hideNSneak/deploy/setup/DO> Enter the number of servers to deploy: "
const confirmSetupDO = "<hideNSneak/deploy/setup/DO> Preferences saved."

const setupGoogle = "<hideNSneak/deploy/setup/GCP> Please enter your Google deploy setup preferences..."
const numServersToDeployGoogle = "<hideNSneak/deploy/setup/GCP> Enter the number of servers to deploy: "
const confirmSetupGoogle = "<hideNSneak/deploy/setup/GCP> Preferences saved."

const setupAzure = "<hideNSneak/deploy/setup/azure> Please enter your Azure deploy setup preferences..."
const numServersToDeployAzure = "<hideNSneak/deploy/setup/azure> Enter the number of servers to deploy: "
const confirmSetupAzure = "<hideNSneak/deploy/setup/azure> Preferences saved."

const setupAzureCDN = "<hideNSneak/deploy/setup/azure-cdn> Please enter your Azure-CDN deploy setup preferences..."
const numServersToDeployAzureCDN = "<hideNSneak/deploy/setup/azure-cdn> Enter the number of servers to deploy: "
const confirmSetupAzureCDN = "<hideNSneak/deploy/setup/azure-cdn> Preferences saved."

const setupAPIGateway = "<hideNSneak/deploy/setup/api-gateway> Please enter your API Gateway deploy setup preferences..."
const numServersToDeployAPIGateway = "<hideNSneak/deploy/setup/api-gateway> Enter the number of servers to deploy: "
const confirmSetupAPIGateway = "<hideNSneak/deploy/setup/api-gateway> Preferences saved."

////////////////// destroy /////////////////////
////////////////// start /////////////////////
////////////////// stop /////////////////////
////////////////// list /////////////////////
////////////////// shell /////////////////////
////////////////// socks-add /////////////////////
////////////////// socks-kill /////////////////////
////////////////// domainfront /////////////////////
////////////////// domainfront-list /////////////////////
////////////////// nmap /////////////////////
////////////////// proxyconf /////////////////////
////////////////// send /////////////////////
////////////////// get /////////////////////
////////////////// firewall /////////////////////
////////////////// firewall-list /////////////////////
