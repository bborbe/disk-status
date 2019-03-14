// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package disk

import (
	"fmt"
	"io"
	"net/http"

	"github.com/golang/glog"
)

type metricsHandler struct {
	path string
}

// NewMetricsHandler returns a http.Handler for the given path.
func NewMetricsHandler(path string) http.Handler {
	return &metricsHandler{
		path: path,
	}
}

func (s *metricsHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	glog.V(4).Infof("get metrics for path %s", s.path)
	usage, err := Usage(s.path)
	if err != nil {
		glog.V(1).Infof("get metrics for path %s failed: %v", s.path, err)
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	write(resp, "available bytes", "diskstatus_bytestotal", usage.BytesTotal)
	write(resp, "used bytes", "diskstatus_bytesused", usage.BytesUsed)
	write(resp, "free bytes", "diskstatus_bytesfree", usage.BytesFree)
	write(resp, "available inodes", "diskstatus_inodestotal", usage.InodesTotal)
	write(resp, "used inodes", "diskstatus_inodesused", usage.InodesUsed)
	write(resp, "free inodes", "diskstatus_inodesfree", usage.InodesFree)
}

func write(w io.Writer, des string, name string, value uint64) {
	fmt.Fprintf(w, "# HELP %s %s\n", name, des)
	fmt.Fprintf(w, "# TYPE %s gauge\n", name)
	fmt.Fprintf(w, "%s %d\n", name, value)
}
