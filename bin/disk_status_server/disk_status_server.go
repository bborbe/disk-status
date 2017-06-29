package main

import (
	"fmt"
	"github.com/bborbe/disk_utils/status"
	flag "github.com/bborbe/flagenv"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
	"net/http"
	"runtime"
)

const (
	PARAMETER_PORT     = "port"
	DEFAULT_PORT   int = 8080
	PARAMETER_PATH     = "path"
)

var (
	portPtr = flag.Int(PARAMETER_PORT, DEFAULT_PORT, "Port")
	pathPtr = flag.String(PARAMETER_PATH, "/", "Path")
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
	server := createServer()
	glog.V(2).Infof("start server")
	return gracehttp.Serve(server)
}

func createServer() *http.Server {
	glog.V(2).Infof("create http server on %s", *portPtr)
	handler := status.NewHandler(*pathPtr)
	return &http.Server{Addr: fmt.Sprintf(":%d", *portPtr), Handler: handler}
}
