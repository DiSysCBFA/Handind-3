package server

const (
	port = ":5000" // default port
	name = "Chitty-Chat-Server"
)

/*
func setupServer() {
	log.Println("Setting up server on port:", port, "with the name", name)
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer, err := CreateGrpcServer(name)
	if err != nil {
		log.Fatalf("Failed to create gRPC server: %v", err)
	}
}
*/
