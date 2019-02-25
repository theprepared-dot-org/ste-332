// Copyright 2019 Mark Spicer
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
// WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package application provides a reusable definition for a gRPC server that comes out of the box with metrics, logging,
// and configuration.
package application

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Application defines a gRPC based server application.
type Application struct {
	name            string
	banner          string
	version         string
	Server          *grpc.Server
	log             *logrus.Logger
	serverInterface string
	adminInterface  string
}

// NewApplication provides an instantiated Application instance with the provided settings. You will still need to
// register your gRPC server instance with the gRPC server in the application returned.
func NewApplication(name string, banner string, version string) (*Application, error) {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logEntry := logrus.NewEntry(logger)

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_logrus.StreamServerInterceptor(logEntry),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(logEntry),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	app := &Application{
		name:            name,
		banner:          banner,
		version:         version,
		Server:          s,
		log:             logger,
		serverInterface: ":50051",
		adminInterface:  ":8081",
	}

	return app, nil
}

// Run starts the gRPC server and the admin interface in a blocking fashion.
func (app *Application) Run() error {
	app.enablePrometheusMetrics()

	lis, err := net.Listen("tcp", app.serverInterface)
	if err != nil {
		return err
	}

	app.printBanner()
	app.log.Infof("listening on Server interface %s and admin interface %s", app.serverInterface, app.adminInterface)

	return app.Server.Serve(lis)
}

func (app *Application) enablePrometheusMetrics() {
	grpc_prometheus.EnableHandlingTimeHistogram()
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(app.adminInterface, nil)
}

func (app *Application) printBanner() {
	fmt.Fprint(os.Stderr, fmt.Sprintf(app.banner, app.version))
}
