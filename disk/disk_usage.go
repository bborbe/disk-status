// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package disk

import (
	"syscall"
)

// Status contains disk informations.
type Status struct {
	BytesTotal  uint64 `json:"bytes_total"`
	BytesUsed   uint64 `json:"bytes_used"`
	BytesFree   uint64 `json:"bytes_free"`
	InodesTotal uint64 `json:"inodes_total"`
	InodesUsed  uint64 `json:"inodes_used"`
	InodesFree  uint64 `json:"inodes_free"`
}

// Usage returns the free, used and total bytes and inodes for the given path.
func Usage(path string) (*Status, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return nil, err
	}
	var disk Status
	disk.BytesTotal = fs.Blocks * uint64(fs.Bsize)
	disk.BytesFree = fs.Bfree * uint64(fs.Bsize)
	disk.BytesUsed = disk.BytesTotal - disk.BytesFree

	disk.InodesTotal = uint64(fs.Files)
	disk.InodesFree = uint64(fs.Ffree)
	disk.InodesUsed = disk.InodesTotal - disk.InodesFree
	return &disk, nil
}
