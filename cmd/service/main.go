package main

import (
	"log"
	"net"
	"net/http"

	grpcsvc "jrpc/internal/grpc"
	pb "jrpc/internal/grpc/pb"
	"jrpc/internal/observability"

	grpcprom "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	observability.Init()

	grpcMetrics := grpcprom.NewServerMetrics()
	grpcMetrics.EnableHandlingTimeHistogram()

	if err := prometheus.Register(grpcMetrics); err != nil {
		log.Printf("Note: gRPC Prometheus metrics registration skipped (likely already registered): %v", err)
		// Do NOT fatal — the metrics will still work if already registered
	} else {
		log.Println("gRPC Prometheus metrics registered successfully")
	}

	go func() {
		log.Println("Metrics server started on :2112")
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatalf("metrics server failed: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
	)

	orderSvc := grpcsvc.NewOrderServer()
	pb.RegisterOrderServiceServer(server, orderSvc)
	pb.RegisterCancelServiceServer(server, orderSvc)

	// This is crucial — initializes per-method counters even if collector was pre-registered
	grpcMetrics.InitializeMetrics(server)

	log.Println("gRPC service running on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
