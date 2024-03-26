// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package disk_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bborbe/disk-status/disk"
)

var _ = Describe("Disk Usage", func() {
	It("returns InodesUsed > 0", func() {
		disk, err := disk.Usage("/")
		Expect(err).NotTo(HaveOccurred())
		Expect(disk.InodesUsed).To(BeNumerically(">", 0))
	})
	It("returns InodesTotal > 0", func() {
		disk, err := disk.Usage("/")
		Expect(err).NotTo(HaveOccurred())
		Expect(disk.InodesTotal).To(BeNumerically(">", 0))
	})
	It("returns InodesFree > 0", func() {
		disk, err := disk.Usage("/")
		Expect(err).NotTo(HaveOccurred())
		Expect(disk.InodesFree).To(BeNumerically(">", 0))
	})
	It("returns BytesFree > 0", func() {
		disk, err := disk.Usage("/")
		Expect(err).NotTo(HaveOccurred())
		Expect(disk.BytesFree).To(BeNumerically(">", 0))
	})
	It("returns BytesTotal > 0", func() {
		disk, err := disk.Usage("/")
		Expect(err).NotTo(HaveOccurred())
		Expect(disk.BytesTotal).To(BeNumerically(">", 0))
	})
	It("returns BytesUsed > 0", func() {
		disk, err := disk.Usage("/")
		Expect(err).NotTo(HaveOccurred())
		Expect(disk.BytesUsed).To(BeNumerically(">", 0))
	})
})
