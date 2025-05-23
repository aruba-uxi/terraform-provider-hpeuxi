---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "hpeuxi_network_group_assignment Resource - hpeuxi"
subcategory: ""
description: |-
  Manages a network group assignment.
  Note: it is recommended to use a hpeuxi_group resource id as the group_id. This will help maintain dependencies between resources. This is useful when a destructive action is performed on an ancestor of the assigned group.
---

# hpeuxi_network_group_assignment (Resource)

Manages a network group assignment.

Note: it is recommended to use a `hpeuxi_group` **resource** `id` as the `group_id`. This will help maintain dependencies between resources. This is useful when a destructive action is performed on an ancestor of the assigned group.

## Example Usage

```terraform
resource "hpeuxi_network_group_assignment" "my_network_group_assignment" {
  network_id = hpeuxi_wired_network.my_network.id
  group_id   = hpeuxi_group.my_group.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_id` (String) The identifier of the group to be assigned to. Use `hpeuxi_group` resource id field (recommended); `data.hpeuxi_group` id field or group id here.
- `network_id` (String) The identifier of the network to be assigned. Use `hpeuxi_wired_network` resource id field; `data.hpeuxi_wired_network` id field; `hpeuxi_wireless_network` resource id field; `data.hpeuxi_wireless_network` id field wired network id; or wireless network id here.

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
