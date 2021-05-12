---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fugue_azure_environment Resource - terraform-provider-fugue"
subcategory: ""
description: |-
  fugue_azure_environment manages an Environment in Fugue corresponding to one Azure subscription.
---

# fugue_azure_environment (Resource)

`fugue_azure_environment` manages an Environment in Fugue corresponding to one Azure subscription.

## Example Usage

```terraform
variable "tenant_id" {
  description = "Azure tenant ID"
  type        = string
}

variable "subscription_id" {
  description = "Azure subscription ID"
  type        = string
}

variable "application_id" {
  description = "Azure application ID"
  type        = string
}

variable "client_secret" {
  description = "Azure client secret"
  type        = string
  sensitive   = true
}

resource "fugue_azure_environment" "example" {
  name                   = "example"
  tenant_id              = var.tenant_id
  subscription_id        = var.subscription_id
  application_id         = var.application_id
  client_secret          = var.client_secret
  compliance_families    = ["CISAZURE"]
  survey_resource_groups = ["*"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **application_id** (String) The Azure Active Directory application ID used for Fugue.
- **client_secret** (String, Sensitive) The Azure secret generated for the Active Directory application.
- **name** (String) The name for the environment.
- **subscription_id** (String) The Azure subscription ID.
- **survey_resource_groups** (Set of String) Survey resource groups.
- **tenant_id** (String) The Azure Tenant ID.

### Optional

- **compliance_families** (Set of String) The set of compliance families to enable in this environment.
- **scan_interval** (Number) Controls the time in seconds between scheduled scans of this environment.
- **scan_schedule_enabled** (Boolean) Controls whether this environment is scanned on a schedule.

### Read-Only

- **id** (String) The unique ID for this environment as generated by Fugue.
- **scan_status** (String) Indicates whether a scan on this environment is currently running.

