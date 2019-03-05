package status

import (
	"fmt"
	"io"
	"net/http"

	"github.com/golang/glog"
)

type statusHandler struct {
	path string
}

func NewHandler(path string) *statusHandler {
	s := new(statusHandler)
	s.path = path
	return s
}

func (s *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	glog.V(4).Infof("get metrics for path %s", s.path)
	usage, err := DiskUsage(s.path)
	if err != nil {
		glog.V(4).Infof("get metrics for path %s failed: %v", s.path, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	write(w, "available bytes", "diskstatus_bytestotal", usage.BytesTotal)
	write(w, "used bytes", "diskstatus_bytesused", usage.BytesUsed)
	write(w, "free bytes", "diskstatus_bytesfree", usage.BytesFree)
	write(w, "available inodes", "diskstatus_inodestotal", usage.InodesTotal)
	write(w, "used inodes", "diskstatus_inodesused", usage.InodesUsed)
	write(w, "free inodes", "diskstatus_inodesfree", usage.InodesFree)
}

func write(w io.Writer, des string, name string, value uint64) {
	fmt.Fprintf(w, "# HELP %s %s\n", name, des)
	fmt.Fprintf(w, "# TYPE %s gauge\n", name)
	fmt.Fprintf(w, "%s %d\n", name, value)
}
