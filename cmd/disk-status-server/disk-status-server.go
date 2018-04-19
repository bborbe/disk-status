package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/bborbe/disk_utils/status"
	flag "github.com/bborbe/flagenv"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
)

var (
	portPtr = flag.Int("port", 8080, "Port")
	pathPtr = flag.String("path", "/", "Path")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := do(); err != nil {
		glog.Exit(err)
	}
}

func do() error {
	glog.V(0).Infof("create http server on %d", *portPtr)
	return gracehttp.Serve(&http.Server{Addr: fmt.Sprintf(":%d", *portPtr), Handler: status.NewHandler(*pathPtr)})
}
