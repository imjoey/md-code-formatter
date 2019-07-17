package formatter

import (
	"bytes"
	"testing"
)

func TestHCLFormatter(t *testing.T) {
	cases := []struct {
		Code     *CodeDiff
		ErrCount int
	}{
		{
			Code: &CodeDiff{
				Before: []byte(`
resource "ovirt_vm" "vm" {
			name = "my_vm"
			cluster_id = "3e7e71ed-24ea-4812-8ef9-a09a858d31e4"
			memory = 4096
			template_id = "5ba458ca-00a4-0358-00cb-000000000223"
}`),
				After: []byte(`
resource "ovirt_vm" "vm" {
  name        = "my_vm"
  cluster_id  = "3e7e71ed-24ea-4812-8ef9-a09a858d31e4"
  memory      = 4096
  template_id = "5ba458ca-00a4-0358-00cb-000000000223"
}`),
				Lang: HCL,
			},
			ErrCount: 0,
		},
		{
			Code: &CodeDiff{
				Before: []byte(`
resource "ovirt_vm" "vm" {
  name       = "my_vm"
  cluster_id = "3e7e71ed-24ea-4812-8ef9-a09a858d31e4"
  memory     = 4096 # in megabytes

  block_device {
    disk_id   = "${ovirt_disk.boot_disk_1.id}"
    interface = "virtio"
  }
}`),
				After: []byte(`
resource "ovirt_vm" "vm" {
  name       = "my_vm"
  cluster_id = "3e7e71ed-24ea-4812-8ef9-a09a858d31e4"
  memory     = 4096 # in megabytes

  block_device {
    disk_id   = "${ovirt_disk.boot_disk_1.id}"
    interface = "virtio"
  }
}`),
				Lang: HCL,
			},
			ErrCount: 0,
		},
		{
			Code: &CodeDiff{
				Before: []byte(`
resource "ovirt_disk" "boot_disk_1" {
  name              = "boot_disk_1 # syntax error here
  alias             = "boot_disk_1"
  size              = 60 # in gigabytes
  format            = "cow"
  storage_domain_id = "5ba458ca-00a4-0358-00cb-000000000223"
  sparse            = true
}`),
				After: nil,
				Lang:  HCL,
			},
			ErrCount: 7,
		},
	}

	f := NewHCLFormatter()
	for _, c := range cases {
		after, errs := f.Format(c.Code.Before)
		if len(errs) != c.ErrCount {
			var errStr string
			if len(errs) > 0 {
				errStr = errs[0].Error()
			}
			t.Fatalf("Expected %d errors, got %d: %s (may more)", c.ErrCount, len(errs), errStr)
		}
		if !bytes.Equal(c.Code.After, after) {
			t.Fatalf("Expected \n%s\n after formatting, got \n%s\n", after, c.Code.After)
		}
	}

}
