package main

import (
	"encoding/json"
	"fmt"
)

var stateFile = []byte(`{
    "version": 3,
    "terraform_version": "0.11.0",
    "serial": 2,
    "lineage": "493f46cf-d49e-4e5e-a6eb-9d21744a27a6",
    "modules": [
        {
            "path": [
                "root"
            ],
            "outputs": {},
            "resources": {
                "ansible_host.db": {
                    "type": "ansible_host",
                    "depends_on": [],
                    "primary": {
                        "id": "db.example.com",
                        "attributes": {
                            "groups.#": "2",
                            "groups.0": "example",
                            "groups.1": "db",
                            "id": "db.example.com",
                            "inventory_hostname": "db.example.com",
                            "vars.%": "2",
                            "vars.bar": "ddd",
                            "vars.fooo": "ccc"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.ansible"
                },
                "ansible_host.www": {
                    "type": "ansible_host",
                    "depends_on": [],
                    "primary": {
                        "id": "www.example.com",
                        "attributes": {
                            "groups.#": "2",
                            "groups.0": "example",
                            "groups.1": "web",
                            "id": "www.example.com",
                            "inventory_hostname": "www.example.com",
                            "vars.%": "2",
                            "vars.bar": "bbb",
                            "vars.fooo": "aaa"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.ansible"
                }
            },
            "depends_on": []
        }
    ]
}
`)

//////////////////////Terraform Inventory//////////////////////
////////Based on: https://github.com/nbering/terraform-inventory/blob/master/terraform.py///

// #! /usr/bin/env python2

// import json
// import os
// import re
// import subprocess
// import sys

// TERRAFORM_PATH = os.environ.get('ANSIBLE_TF_BIN', 'terraform')
// TERRAFORM_DIR = os.environ.get('ANSIBLE_TF_DIR', os.getcwd())

const TERRAFORM_PATH = "/usr/local/bin/terraform"

type State struct {
}

// {
//     "version": 3,
//     "terraform_version": "0.11.0",
//     "serial": 2,
//     "lineage": "493f46cf-d49e-4e5e-a6eb-9d21744a27a6",
//     "modules": [
//         {
//             "path": [
//                 "root"
//             ],
//             "outputs": {},
//             "resources": {
//                 "ansible_host.db": {
//                     "type": "ansible_host",
//                     "depends_on": [],
//                     "primary": {
//                         "id": "db.example.com",
//                         "attributes": {
//                             "groups.#": "2",
//                             "groups.0": "example",
//                             "groups.1": "db",
//                             "id": "db.example.com",
//                             "inventory_hostname": "db.example.com",
//                             "vars.%": "2",
//                             "vars.bar": "ddd",
//                             "vars.fooo": "ccc"
//                         },
//                         "meta": {},
//                         "tainted": false
//                     },
//                     "deposed": [],
//                     "provider": "provider.ansible"
//                 },
//                 "ansible_host.www": {
//                     "type": "ansible_host",
//                     "depends_on": [],
//                     "primary": {
//                         "id": "www.example.com",
//                         "attributes": {
//                             "groups.#": "2",
//                             "groups.0": "example",
//                             "groups.1": "web",
//                             "id": "www.example.com",
//                             "inventory_hostname": "www.example.com",
//                             "vars.%": "2",
//                             "vars.bar": "bbb",
//                             "vars.fooo": "aaa"
//                         },
//                         "meta": {},
//                         "tainted": false
//                     },
//                     "deposed": [],
//                     "provider": "provider.ansible"
//                 }
//             },
//             "depends_on": []
//         }
//     ]
// }

// def _extract_dict(attrs, key):
//     out = {}
//     for k in attrs.keys():
//         match = re.match(r"^" + key + r"\.(.*)", k)
//         if not match or match.group(1) == "%":
//             continue

//         out[match.group(1)] = attrs[k]
//     return out

// def _extract_list(attrs, key):
//     out = []

//     length_key = key + ".#"
//     if length_key not in attrs.keys():
//         return []

//     length = int(attrs[length_key])
//     if length < 1:
//         return []

//     for i in range(0, length):
//         out.append(attrs["{}.{}".format(key, i)])

//     return out

// def _init_group(children=None, hosts=None, vars=None):
//     return {
//         "hosts": [] if hosts is None else hosts,
//         "vars": {} if vars is None else vars,
//         "children": [] if children is None else children
//     }

// def _add_host(inventory, hostname, groups, host_vars):
//     inventory["_meta"]["hostvars"][hostname] = host_vars
//     for group in groups:
//         if group not in inventory.keys():
//             inventory[group] = _init_group(hosts=[hostname])
//         elif hostname not in inventory[group]:
//             inventory[group]["hosts"].append(hostname)

// def _add_group(inventory, group_name, children, group_vars):
//     if group_name not in inventory.keys():
//         inventory[group_name] = _init_group(children=children, vars=group_vars)
//     else:
//         # Start out with support for only one "group" with a given name
//         # If there's a second group by the name, last in wins
//         inventory[group_name]["children"] = children
//         inventory[group_name]["vars"] = group_vars

// def _init_inventory():
//     return {
//         "all": _init_group(),
//         "_meta": {
//             "hostvars": {}
//         }
//     }

// def _handle_host(attrs, inventory):
//     host_vars = _extract_dict(attrs, "vars")
//     groups = _extract_list(attrs, "groups")
//     hostname = attrs["inventory_hostname"]

//     if "all" not in groups:
//         groups.append("all")

//     _add_host(inventory, hostname, groups, host_vars)

// def _handle_group(attrs, inventory):
//     group_vars = _extract_dict(attrs, "vars")
//     children = _extract_list(attrs, "children")
//     group_name = attrs["inventory_group_name"]

//     _add_group(inventory, group_name, children, group_vars)

// def _walk_state(tfstate, inventory):
//     for module in tfstate["modules"]:
//         for resource in module["resources"].values():
//             if not resource["type"].startswith("ansible_"):
//                 continue

//             attrs = resource["primary"]["attributes"]

//             if resource["type"] == "ansible_host":
//                 _handle_host(attrs, inventory)
//             if resource["type"] == "ansible_group":
//                 _handle_group(attrs, inventory)

//     return inventory

// def _main():
//     try:

//         tf_command = [TERRAFORM_PATH, 'state', 'pull', '-input=false']

//         proc = subprocess.Popen(tf_command, cwd=TERRAFORM_DIR, stdout=subprocess.PIPE)
//         tfstate = json.load(proc.stdout)
//         inventory = _walk_state(tfstate, _init_inventory())
//         sys.stdout.write(json.dumps(inventory, indent=2))
//     except:
//         sys.exit(1)

// if __name__ == '__main__':
//     _main()

type stateStruct struct {
	version           int           `json:"version"`
	terraform_version string        `json:"terraform_version"`
	serial            int           `json:"serial"`
	lineage           string        `json:"lineage"`
	modules           []interface{} `json:"modules"`
}

func mainFunc() {
	// args := []string{"state", "pull", "-input=false"}
	// execCmd(TERRAFORM_PATH, args)
	var f stateStruct
	err := json.Unmarshal(stateFile, &f)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}
	fmt.Println(f)

}
