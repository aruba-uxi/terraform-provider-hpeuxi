---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "hpeuxi_service_test_group_assignment Data Source - hpeuxi"
subcategory: ""
description: |-
  Retrieves a specific service test group assignment.
---

# hpeuxi_service_test_group_assignment (Data Source)

Retrieves a specific service test group assignment.

## Example Usage

```terraform
# Retrieve data for a service test group assignment
data "hpeuxi_service_test_group_assignment" "my_service_test_group_assignment" {
  filter = {
    id = "<my_service_test_group_assignment_id>"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `filter` (Attributes) The filter used to filter the specific service test group assignment. (see [below for nested schema](#nestedatt--filter))

### Read-Only

- `group_id` (String) The identifier of the assigned group.
- `id` (String) The identifier of the service test group assignment.
- `service_test_id` (String) The identifier of the assigned service test.

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Required:

- `id` (String) The identifier of the service test group assignment.
