---
layout: "vra"
page_title: "vRealize Automation: vra_execute_blueprint"
sidebar_current: "docs-vra-resource-inventory-folder"
description: |-
  Provides a vRealize Blueprint Executing resource. This can be used to execute blueprint in vRealize Automation.
---

# vra\_execute\_blueprint

Provides a vRealize Automation executing resource. This can be used to execute blueprints in vRealize.

## Example Usage

```hcl
# Execute a blueprint
resource "vra_execute_blueprint" "ExecuteBlueprint" {
       blueprint_name = "Create simple virtual machine"
       input_file_name = "data.json"
       time_out = 20
}
```

## Argument Reference

The following arguments are supported:

* `blueprint_name` - (Required) The Blueprint Name which we want to execute
* `input_file_name` - (Required) The configuration file which is required to execute blueprint,to create that json file or the configuaration file refer CREATE_JSON.MD file
* `timeout` - (Optional) Timeout should be in seconds and its optional i.e user can gives in seconds 
 otherwise its default value is 50 sec.