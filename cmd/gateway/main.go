package main

import (
	"jrpc/internal/gateway"
	"jrpc/internal/grpc"
	"log"
	"net/http"
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
	client := grpc.NewClient("grpc:50051")

	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		gateway.HandleJSONRPC(w, r, client)
	})

	log.Println("JSON-RPC gateway running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Server failed to start at 8080")
	}

}
