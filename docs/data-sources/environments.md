---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fugue_environments Data Source - terraform-provider-fugue"
subcategory: ""
description: |-
  Multi-result version of fugue_environment where the results are placed in the ids element.
---

# fugue_environments (Data Source)

Multi-result version of `fugue_environment` where the results are placed in the `ids` element.

## Example Usage

```terraform
data "fugue_environments" "all_aws_environments" {
  filter {
    name   = "cloud_provider"
    values = ["aws"]
  }
}

output "environment_ids" {
  value = data.fugue_environment.all_aws_environments.ids
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **filter** (Block Set) Filters. The result is restricted to the intersection of the result of all filters. (see [below for nested schema](#nestedblock--filter))
- **id** (String) The ID of this resource.

### Read-Only

- **ids** (List of String) List of the ID of each result.

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Required:

- **name** (String) Name of the field you want to apply this filter to.
- **values** (List of String) Value to match


