---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "hpeuxi_network_group_assignment Resource - hpeuxi"
subcategory: ""
description: |-
  Manages a network group assignment.
---

# hpeuxi_network_group_assignment (Resource)

Manages a network group assignment.

## Example Usage

```terraform
resource "hpeuxi_network_group_assignment" "my_network_group_assignment" {
    network_id = hpeuxi_wired_network.my_network.id
    group_id = hpeuxi_group.my_group.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_id` (String) The identifier of the group to be assigned to. Use group id; `uxi_group` resource id field or `uxi_group` datasource id field here.
- `network_id` (String) The identifier of the network to be assigned. Use wired network id; `uxi_wired_network` resource id field; `uxi_wired_network` datasource id field; wireless network id; `uxi_wireless_network` resource id field or `uxi_wireless_network` datasource id field here.

### Read-Only

- `id` (String) The identifier of the network group assignment

## Import

Import is supported using the following syntax:

```shell
# Import network group assignment using its ID
terraform import hpeuxi_network_group_assignment.my_network_group_assignment <my_network_group_assignment_id>

# Import network group assignment using its ID with an import block
import {
    to = hpeuxi_network_group_assignment.my_network_group_assignment
    id = "<my_network_group_assignment_id>"
}
```
