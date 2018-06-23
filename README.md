1. How to run locally 
2. Quirks 
3. What is this tool for 
4. Handy use cases
5. Commands 
6. Organization of code 
7. Contributing and developing locally 
8. 


Need to create a file under package main titled secrets.go

const tfvars = `
	aws_access_key = "YOUR_SECRET_KEY"
	aws_secret_key = "YOUR_SECRET_KEY"
	do_token = "YOUR_SECRET_KEY"
	azure_tenant_id = "YOUR_SECRET_KEY"
	azure_client_id = "YOUR_SECRET_KEY"
	azure_client_secret = "YOUR_SECRET_KEY"
	azure_subscription_id = "YOUR_SECRET_KEY"
`
