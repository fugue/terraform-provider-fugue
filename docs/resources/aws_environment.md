---
page_title: "fugue_aws_environment Resource - terraform-provider-fugue"
subcategory: ""
description: |-
  
---

# Resource `fugue_aws_environment`





## Schema

### Required

- **name** (String, Required)
- **regions** (Set of String, Required)
- **resource_types** (Set of String, Required)
- **role_arn** (String, Required)

### Optional

- **compliance_families** (Set of String, Optional)
- **govcloud** (Boolean, Optional)
- **scan_interval** (Number, Optional)
- **scan_schedule_enabled** (Boolean, Optional)

### Read-only

- **id** (String, Read-only) The ID of this resource.
- **scan_status** (String, Read-only)


