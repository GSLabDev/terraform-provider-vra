---
layout: "vra"
page_title: "Provider: vRealize Automation"
sidebar_current: "docs-vra-index"
description: |-
  The vRealize Automation provider is used to interact with the resources supported by
  vRealize Automation. The provider needs to be configured with the proper credentials
  before it can be used.
---

# vRealize Automation Provider

The vRealize Automation provider is used to interact with the resources supported by
vRealize Automation.
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

~> **NOTE:** The vRealize Automation Provider currently represents _initial support_
and therefore may undergo significant changes as the community improves it. This
provider at this time only supports Executing blueprint Resource

## Example Usage

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

## Argument Reference

The following arguments are used to configure the Active Directory Provider:

* `host_url` - (Required) This is the host url for vRealize Api. Can also
  be specified with the `VRA_URL` environment variable.
* `tenant` - (Required) This is the tenant of the user for vRealize API operations. Can
  also be specified with the `VRA_TENANT` environment variable.
* `user_name` - (Required) This is the username for vRealize Api
  operations. Can also be specified with the `VRA_USER` environment
  variable.
* `user_password` - (Required) This is the password for vRealize User. Can also be specified with `VRA_PASSWORD`

## Acceptance Tests

The vRealize Automation provider's acceptance tests require the above provider
configuration fields to be set using the documented environment variables.

Once all these variables are in place, the tests can be run like this:

```
make testacc TEST=./builtin/providers/vra
```