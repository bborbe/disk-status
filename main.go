// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/bborbe/argument/v2"
	"github.com/golang/glog"

	"github.com/bborbe/disk-status/disk"
)

type application struct {
	Port int    `arg:"port" env:"PORT" default:"8080" usage:"Port"`
	Path string `arg:"path" env:"PATH" default:"/" usage:"Path"`
}

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	runtime.GOMAXPROCS(runtime.NumCPU())
	_ = flag.Set("logtostderr", "true")

	ctx := contextWithSig(context.Background())
	app := &application{}
	if err := argument.Parse(ctx, app); err != nil {
		glog.Exitf("parse args failed: %v", err)
	}

	glog.V(0).Infof("create http server on %d", app.Port)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Port),
		Handler: disk.NewMetricsHandler(app.Path),
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
