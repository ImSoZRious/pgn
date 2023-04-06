/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	health "google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	_ "google.golang.org/grpc/xds"
)

const (
	defaultName = "world"
)

var (
	name = flag.String("name", defaultName, "Name to greet")

	logger *zap.Logger
)

func main() {
	flag.Parse()
	initLog()
	// Set up a connection to the server.

	logger.Info("Start runing v1")

	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		logger.Error("`SERVER_ADDR` is not provided")
		os.Exit(1)
	}

	fetchInterval, err := strconv.ParseInt(os.Getenv("FETCH_INTERVAL"), 10, 32)
	if err != nil {
		logger.Info("Error parsing `FETCH_INTERVAL` using default (15)")
	}

	healthPort, err := strconv.ParseInt(os.Getenv("HEALTH_PORT"), 10, 32)
	if err != nil {
		logger.Error("`HEALTH_PORT` is not provided / invalid")
		os.Exit(1)
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	logger.Info("Connecting", zap.String("Address", addr))
	if err != nil {
		logger.Error("did not connect", zap.Error(err))
		os.Exit(1)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	go healthServer(healthPort)

	// Contact the server and print out its response.
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			logger.Info("could not greet", zap.Error(err))
		} else {
			logger.Info("Received", zap.String("message", r.GetMessage()))
		}

		time.Sleep(time.Duration(fetchInterval) * time.Second)
	}
}

func healthServer(port int64) {
	server := health.NewServer()

	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error("failed to listen", zap.Error(err))
		os.Exit(1)
	}

	grpc_health_v1.RegisterHealthServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("failed to serve", zap.Error(err))
		os.Exit(1)
	}
}

func initLog() {
	logger, _ = zap.NewProduction()
}
