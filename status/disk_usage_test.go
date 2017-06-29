package status

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestDiskUsageInodesUsed(t *testing.T) {
	disk, err := DiskUsage("/")
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(int(disk.InodesUsed), Gt(0)); err != nil {
		t.Fatal(err)
	}
}

func TestDiskUsageInodesTotal(t *testing.T) {
	disk, err := DiskUsage("/")
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(int(disk.InodesTotal), Gt(0)); err != nil {
		t.Fatal(err)
	}
}

func TestDiskUsageInodesFree(t *testing.T) {
	disk, err := DiskUsage("/")
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(int(disk.InodesFree), Gt(0)); err != nil {
		t.Fatal(err)
	}
}

func TestDiskUsageBytesFree(t *testing.T) {
	disk, err := DiskUsage("/")
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(int(disk.BytesFree), Gt(0)); err != nil {
		t.Fatal(err)
	}
}

func TestDiskUsageBytesTotal(t *testing.T) {
	disk, err := DiskUsage("/")
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(int(disk.BytesTotal), Gt(0)); err != nil {
		t.Fatal(err)
	}
}

func TestDiskUsageBytesUsed(t *testing.T) {
	disk, err := DiskUsage("/")
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(int(disk.BytesUsed), Gt(0)); err != nil {
		t.Fatal(err)
	}
}
