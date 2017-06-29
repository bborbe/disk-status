package status

import (
	"bytes"
	"fmt"
	"syscall"
)

type DiskStatus struct {
	BytesTotal  uint64 `json:"bytes_total"`
	BytesUsed   uint64 `json:"bytes_used"`
	BytesFree   uint64 `json:"bytes_free"`
	InodesTotal uint64 `json:"inodes_total"`
	InodesUsed  uint64 `json:"inodes_used"`
	InodesFree  uint64 `json:"inodes_free"`
}

func (d *DiskStatus) String() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "Bytes-Total: %d\n", d.BytesTotal)
	fmt.Fprintf(b, "Bytes-Used: %d\n", d.BytesUsed)
	fmt.Fprintf(b, "Bytes-Free: %d\n", d.BytesFree)
	fmt.Fprintf(b, "Inodes-Total: %d\n", d.InodesTotal)
	fmt.Fprintf(b, "Inodes-Used: %d\n", d.InodesUsed)
	fmt.Fprintf(b, "Inodes-Free: %d\n", d.InodesFree)
	return b.String()
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus, err error) {
	fs := syscall.Statfs_t{}
	err = syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.BytesTotal = fs.Blocks * uint64(fs.Bsize)
	disk.BytesFree = fs.Bfree * uint64(fs.Bsize)
	disk.BytesUsed = disk.BytesTotal - disk.BytesFree

	disk.InodesTotal = uint64(fs.Files)
	disk.InodesFree = uint64(fs.Ffree)
	disk.InodesUsed = disk.InodesTotal - disk.InodesFree
	return
}
