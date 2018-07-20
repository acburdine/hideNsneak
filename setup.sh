#!/bin/bash


terraformPathDir="/usr/local/bin/"

macTerraformLink="https://releases.hashicorp.com/terraform/0.11.7/terraform_0.11.7_darwin_amd64.zip"
linuxTerraformLink="https://releases.hashicorp.com/terraform/0.11.7/terraform_0.11.7_linux_amd64.zip"

ansibleProviderLinuxURL="https://github.com/nbering/terraform-provider-ansible/releases/download/v0.0.4/terraform-provider-ansible-linux_amd64.zip"
ansibleProviderMacURL="https://github.com/nbering/terraform-provider-ansible/releases/download/v0.0.4/terraform-provider-ansible-darwin_amd64.zip"
ansibleProviderFile="terraform-provider-ansible_v0.0.4"

echo "Checking initial requirements"
if ! [ -x $(command -v unzip) ]
then
    echo "error: unzip is required"
fi
if ! [ -x $(command -v python) ]
then
    echo "error: python is required"
fi
if ! [ -x $(command -v pip) ]
then
    echo "error: pip is required"
fi
if ! [ -x $(command -v go) ]
then
    echo "error: golang is required"
fi

echo "Checking Terraform Installation...."
if ! [ -x $(command -v terraform) ]
then
    if ! [ -x $(command -v unzip) ]
    then
        echo "unzip command is required. Exiting...."
        exit 
    fi

    if [ `uname` = "Linux" ]
    then
        test -x $(command -v terraform) || (wget -O /tmp/terraform_linux.zip $linuxTerraformLink; unzip /tmp/terraform_linux.zip -d /tmp/)
    elif [ `uname` = "Darwin" ]
    then
        test -x $(command -v terraform) || (wget -O /tmp/terraform_linux.zip $macTerraformLink; unzip /tmp/terraform_linux.zip -d /tmp/)
    else
        echo "System must be either Linux or OSx. Exiting...."
        exit 1
    fi

    if [ -f /tmp/terraform ]
    then
        sudo mv /tmp/terraform $terraformPathDir
    else
        echo "Terraform file not found in /tmp"
        exit 1
    fi
fi

echo "Installing Ansible Provider...."
if ! [ -f $HOME/.terraform.d/plugins/$ansibleProviderFile ]
then

    if ! [ -d $HOME/.terraform.d ]
    then
        mkdir $HOME/.terraform.d 
    fi

    if ! [ -d $HOME/.terraform.d/plugins ]
    then
        mkdir $HOME/.terraform.d/plugins
    fi

    if [ `uname` = "Linux" ]
    then
        wget -q -O "$HOME/.terraform.d/plugins/$ansibleProviderFile.zip" $ansibleProviderLinuxURL
        unzip "$HOME/.terraform.d/plugins/$ansibleProviderFile.zip" -d $HOME/.terraform.d/plugins/
        rm "$HOME/.terraform.d/plugins/$ansibleProviderFile.zip"
        mv $HOME/.terraform.d/plugins/*/$ansibleProviderFile $HOME/.terraform.d/plugins/
        rm -r "$HOME/.terraform.d/plugins/linux_amd64"
    elif [ `uname` = "Darwin" ]
    then
        wget -q -O "$HOME/.terraform.d/plugins/$ansibleProviderFile.zip" $ansibleProviderMacURL 
        unzip "$HOME/.terraform.d/plugins/$ansibleProviderFile.zip" -d $HOME/.terraform.d/plugins/
        rm "$HOME/.terraform.d/plugins/$ansibleProviderFile.zip"
        mv $HOME/.terraform.d/plugins/*/$ansibleProviderFile $HOME/.terraform.d/plugins/
        rm -r "$HOME/.terraform.d/plugins/darwin_amd64"
    fi
fi

echo "Cleaning up provider files..."

echo "Installing Ansible...."

if [ ! -x $(command -v ansible) ]
then
    if [ ! -x $(command -v pip) ]
    then
        echo "error: pip not installed. exiting...."
        exit 1
    fi
    # #If on Mac and experiencing errors, use the following command
    # sudo CFLAGS=-Qunused-arguments CPPFLAGS=-Qunused-arguments pip install ansible
    sudo pip install ansible
fi

echo "All requirements met!"