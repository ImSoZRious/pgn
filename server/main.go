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

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
	id int32
}

var logger *zap.Logger

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	logger.Info("Received", zap.String("Message", in.GetName()))
	msg := fmt.Sprintf("Hello %v from %v", in.GetName(), s.id)
	return &pb.HelloReply{Message: msg}, nil
}

func main() {
	flag.Parse()
	initLog()

	logger.Info("Starting in v1")

	port, err := strconv.ParseInt(os.Getenv("SERVER_PORT"), 10, 32)
	if err != nil {
		logger.Info("Failed to parse `SERVER_PORT` using default port (50051)")
		port = 50051
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error("failed to listen", zap.Error(err))
		os.Exit(1)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{
		id: rand.Int31(),
	})

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)

	logger.Info("server listening at", zap.String("Port", lis.Addr().String()))
	if err := s.Serve(lis); err != nil {
		logger.Error("failed to serve", zap.Error(err))
		os.Exit(1)
	}
}

func initLog() {
	logger, _ = zap.NewProduction()
}
