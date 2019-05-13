---
layout: "ibm"
page_title: "IBM : Bandwidth Pool"
sidebar_current: "docs-ibm-resource-bandwidth-pool"
description: |-
  Manages IBM Bandwidth Pool.
---

# ibm\_bandwidth_pool

This resource is used to order a Bandwidth Pool.

## Example Usage

```hcl
resource "ibm_bandwidth_pool" "bw_pool" {
	name="a-resource-to-mod"
	locationGroupId=1
	}
```

## Argument Reference

* `name` - (Required,  string) Name of Bandwidth Pool.
* `locationGroupId` - (Required,  Integer) Location Group ID is required.


## Attribute Reference

The following attributes are exported:

* `id` - The unique internal identifier of the Bandwidth Pool.
