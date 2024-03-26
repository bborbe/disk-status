// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	flag "github.com/bborbe/flagenv"
	"github.com/golang/glog"

	"github.com/bborbe/disk-status/disk"
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	runtime.GOMAXPROCS(runtime.NumCPU())
	_ = flag.Set("logtostderr", "true")
	portPtr := flag.Int("port", 8080, "Port")
	pathPtr := flag.String("path", "/", "Path")
	flag.Parse()

	ctx := contextWithSig(context.Background())

	glog.V(0).Infof("create http server on %d", *portPtr)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *portPtr),
		Handler: disk.NewMetricsHandler(*pathPtr),
	}
	go func() {
		select {
		case <-ctx.Done():
			if err := server.Shutdown(ctx); err != nil {
				glog.Warningf("shutdown failed: %v", err)
			}
		}
	}()
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		glog.V(0).Info(err)
		glog.Exitf("httpServer failed: %v", err)
	}
	glog.V(0).Infof("server finished")
}

func contextWithSig(ctx context.Context) context.Context {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	go func() {
		defer cancel()

		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-signalCh:
		case <-ctx.Done():
		}
	}()

	return ctxWithCancel
}
