# Terraform vRealize Automation Provider

This is the repository for the Terraform [vRealize Automation][1] Provider, which one can use
with Terraform to work with vRealize Automation.

[1]: https://www.vmware.com/products/vrealize-automation.html

Support Execution of Blueprint, work in progress to support additional usecase of vRealize Automation.
Watch this space!

For general information about Terraform, visit the [official website][3] and the
[GitHub project page][4].

[3]: https://terraform.io/
[4]: https://github.com/hashicorp/terraform

# Using the Provider

The current version of this provider requires Terraform v0.10.2 or higher to
run.

Note that you need to run `terraform init` to fetch the provider before
deploying. Read about the provider split and other changes to TF v0.10.0 in the
official release announcement found [here][4].

[4]: https://www.hashicorp.com/blog/hashicorp-terraform-0-10/

## Full Provider Documentation

The provider is usefull in executing any blueprint of vRA from Terraform.

### Example
```hcl

# Configure the vRealize Automation Provider
provider "vra" {
        host_url = "${var.vra_server_host_url}"
        tenant = "${var.var_server_tenant_name}"
        user_name = "${var.vra_server_username}"
        user_password = "${var.vra_server_password}"
}

# Execute a blueprint
resource "vra_execute_blueprint" "ExecuteBlueprint" {
       blueprint_name = "Create simple virtual machine"
       input_file_name = "data.json"
       time_out = 20
}
```

# Building The Provider

**NOTE:** Unless you are [developing][7] or require a pre-release bugfix or feature,
you will want to use the officially released version of the provider (see [the
section above][8]).

[7]: #developing-the-provider
[8]: #using-the-provider


## Cloning the Project

First, you will want to clone the repository to
`$GOPATH/src/github.com/terraform-providers/terraform-provider-vra`:

```sh
mkdir -p $GOPATH/src/github.com/terraform-providers
cd $GOPATH/src/github.com/terraform-providers
git clone git@github.com:terraform-providers/terraform-provider-vra
```

## Running the Build

After the clone has been completed, you can enter the provider directory and
build the provider.

```sh
cd $GOPATH/src/github.com/terraform-providers/terraform-provider-vra
make build
```

## Installing the Local Plugin

After the build is complete, copy the `terraform-provider-vra` binary into
the same path as your `terraform` binary, and re-run `terraform init`.

After this, your project-local `.terraform/plugins/ARCH/lock.json` (where `ARCH`
matches the architecture of your machine) file should contain a SHA256 sum that
matches the local plugin. Run `shasum -a 256` on the binary to verify the values
match.

# Developing the Provider

If you wish to work on the provider, you'll first need [Go][9] installed on your
machine (version 1.9+ is **required**). You'll also need to correctly setup a
[GOPATH][10], as well as adding `$GOPATH/bin` to your `$PATH`.

[9]: https://golang.org/
[10]: http://golang.org/doc/code.html#GOPATH

See [Building the Provider][11] for details on building the provider.

[11]: #building-the-provider

# Testing the Provider

**PreRequisite:** Testing the vRA provider need vRA setup up and running. Also, need input json for blueprint planning to execute.

## Configuring Environment Variables

Most of the tests in this provider require a comprehensive list of environment
variables to run. See the individual `*_test.go` files in the
[`vra/`](vra/) directory for more details. The next section also
describes how you can manage a configuration file of the test environment
variables.

### Using the `.tf-vra-devrc.mk` file

The [`tf-vra-devrc.mk.example`](tf-odl-devrc.mk.example) file contains
an up-to-date list of environment variables required to run the acceptance
tests. Copy this to `$HOME/.tf-vra-devrc.mk` and change the permissions to
something more secure (ie: `chmod 600 $HOME/.tf-vra-devrc.mk`), and
configure the variables accordingly.

## Running the Acceptance Tests

After this is done, you can run the acceptance tests by running:

```sh
$ make testacc
```

If you want to run against a specific set of tests, run `make testacc` with the
`TESTARGS` parameter containing the run mask as per below:

```sh
make testacc TESTARGS="-run=TestAccVRA"
```

This following example would run all of the acceptance tests matching
`TestAccVRA`. Change this for the specific tests you want to
run.



vRealize Automation:-
      
        vRealize Automation provides UI where developers,users can request IT 
        services and manage IT resources. vRA provider is created newly. It 
        utilizes go library. It uses rest apis of vRA and send http request.   


Resources:-

Execute Blueprint:

     -It executes any blueprint present in VRA.
     -Blueprint name is given in configuaration file.
     -Timeout should be in seconds and its optional i.e user can gives in seconds 
     otherwise its default value is 50 sec.
     -data.json and main.tf should be placed in same folder
     -To execute blueprint json data should pass from user in .json file, to generate 
     that json data we have to set value in vra blueprint and send http request
     to get that template using postman or restclient ,so we get the json format and 
     we can use it anytime only by changing the values in that json data.

The main.tf file contains the microservices of how to call the providers and resources. We need to specify required details for resource creation in this file.