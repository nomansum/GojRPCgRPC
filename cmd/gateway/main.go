package main

import (
	"log"
	"net/http"

	"jrpc/internal/gateway"
	"jrpc/internal/grpc"
	//"google.golang.org/grpc"
)

func main() {
	// conn, err := grpc.Dial(
	// 	"localhost:50051",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//client := pb.NewOrderServiceClient(conn)
	client := grpc.NewClient("localhost:50051")

	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		gateway.HandleJSONRPC(w, r, client)
	})

	log.Println("JSON-RPC gateway running on :8080")
	http.ListenAndServe(":8080", nil)
}
