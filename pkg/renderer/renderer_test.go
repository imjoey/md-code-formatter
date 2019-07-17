package renderer

import (
	"bytes"
	"testing"
)

func TestRenderHCL(t *testing.T) {
	cases := []struct {
		Before []byte
		After  []byte
	}{
		{
			Before: []byte(`
---
layout: "ovirt"
page_title: "oVirt: ovirt_vm"
sidebar_current: "docs-ovirt-resource-vm"
description: |-
  Manages a VM resource within oVirt.
---

# ovirt\_vm

Manages a VM resource within oVirt.

## Example Usage

### Boot VM From an Existing Template (Disk)` +

				"\n```hcl\n" +
				`resource "ovirt_vm" "vm"{
  name = "my_vm"
  cluster_id = "3e7e71ed-24ea-4812-8ef9-a09a858d31e4"
  memory = 4096
  # there has one or more disks in the specified template
  template_id= "5ba458ca-00a4-0358-00cb-000000000223"
}` +
				"\n```"),
			After: []byte(`
---
layout: "ovirt"
page_title: "oVirt: ovirt_vm"
sidebar_current: "docs-ovirt-resource-vm"
description: |-
  Manages a VM resource within oVirt.
---

# ovirt\_vm

Manages a VM resource within oVirt.

## Example Usage

### Boot VM From an Existing Template (Disk)` +
				"\n```hcl\n" +
				`resource "ovirt_vm" "vm" {
  name       = "my_vm"
  cluster_id = "3e7e71ed-24ea-4812-8ef9-a09a858d31e4"
  memory     = 4096
  # there has one or more disks in the specified template
  template_id = "5ba458ca-00a4-0358-00cb-000000000223"
}` +
				"\n```"),
		},
	}

	for _, c := range cases {
		r := NewMdCodeRenderer(c.Before, WithHCLEnabled())
		var buf bytes.Buffer

		r.Render(&buf)

		if !bytes.Equal(buf.Bytes(), c.After) {
			t.Fatalf("Expected \n%s\n after formatting, got \n%s\n", c.After, buf.Bytes())
		}

	}

}
