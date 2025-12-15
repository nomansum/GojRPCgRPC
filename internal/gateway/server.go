package gateway

// package gateway

// import (
// 	"log"
// 	"net/http"

// 	pb "jrpc/internal/grpc/pb"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// // Server represents the JSON-RPC gateway server
// type Server struct {
// 	addr   string
// 	client pb.OrderServiceClient
// }

// // NewServer initializes the gateway server and gRPC client
// func NewServer(addr string, grpcAddr string) (*Server, error) {
// 	conn, err := grpc.Dial(
// 		grpcAddr,
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	client := pb.NewOrderServiceClient(conn)

// 	return &Server{
// 		addr:   addr,
// 		client: client,
// 	}, nil
// }

// // Start runs the HTTP JSON-RPC server
// func (s *Server) Start() error {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
// 		HandleJSONRPC(w, r, s.client)
// 	})

// 	log.Println("JSON-RPC gateway listening on", s.addr)
// 	return http.ListenAndServe(s.addr, mux)
// }
