---
layout: "ibm"
page_title: "IBM : multivlan_firewall"
sidebar_current: "docs-ibm-resource-multivlan-firewall"
description: |-
  Manages IBM Multi Vlan Firewall.
---

# ibm\_multivlan_firewall

Provides an SSL certificate resource. This allows SSL certificates to be requested, and delete request for ssl certificates.

For additional details, see the [IBM Cloud (SoftLayer) multivlan firewall Request docs](https://softlayer.github.io/reference/datatypes/SoftLayer_Container_Product_Order_Network_Protection_Firewall_Dedicated/)

## Example Usage

In the following example, you can create a multi-vlan firewall:

```hcl
resource "ibm_multi_vlan_firewall" "firewall_first" {
	datacenter = "dal13"
	pod = "pod01"
	name = "Checkdelete1"
	public_vlan_id = 2213543
	firewall_type = "FortiGate Firewall Appliance HA Option"
	addon_configuration = ["FortiGate Security Appliance - Web Filtering Add-on (High Availability)","FortiGate Security Appliance - NGFW Add-on (High Availability)","FortiGate Security Appliance - AV Add-on (High Availability)"]
	}
```


## Argument Reference

The following arguments are supported:

* `datacenter` - (Required, string) The data center in which the firewall appliance resides.
* `pod` - (Required, string) The pod in which the firewall resides
* `name` - (Required, string) The name of the firewall device
* `firewall_type` - (Required, string) The type of the firwall device. Allowed values are:- FortiGate Security Appliance,FortiGate Firewall Appliance HA Option
* `addon_configuration` - (Required, list) The list of addons that are allowed. Allowed values:- ["FortiGate Security Appliance - Web Filtering Add-on (High Availability)","FortiGate Security Appliance - NGFW Add-on (High Availability)","FortiGate Security Appliance - AV Add-on (High Availability)"] or ["FortiGate Security Appliance - Web Filtering Add-on","FortiGate Security Appliance - NGFW Add-on","FortiGate Security Appliance - AV Add-on"]