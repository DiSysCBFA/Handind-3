package server

import (
	"context"
	"handin-3/service"
	"log"

	"google.golang.org/grpc"

	tasks "handin-3/api"
)

type server struct {
	tasks.UnimplementedTaskServiceServer
	clock service.LamportClock

	name string
}

// CreateTask implements tasks.TaskServiceServer.
func (s server) CreateTask(context.Context, *tasks.CreateTaskRequest) (*tasks.Task, error) {
	panic("unimplemented")
}

// GetTask implements tasks.TaskServiceServer.
func (s server) GetTask(context.Context, *tasks.GetTaskRequest) (*tasks.Task, error) {
	panic("unimplemented")
}

// ListTasks implements tasks.TaskServiceServer.
func (s server) ListTasks(*tasks.ListTasksRequest, grpc.ServerStreamingServer[tasks.Task]) error {
	panic("unimplemented")
}

// RecordTasks implements tasks.TaskServiceServer.
func (s server) RecordTasks(grpc.ClientStreamingServer[tasks.CreateTaskRequest, tasks.TaskSummary]) error {
	panic("unimplemented")
}

// TaskChat implements tasks.TaskServiceServer.
func (s server) TaskChat(grpc.BidiStreamingServer[tasks.TaskComment, tasks.TaskComment]) error {
	panic("unimplemented")
}

// mustEmbedUnimplementedTaskServiceServer implements tasks.TaskServiceServer.
func (s server) mustEmbedUnimplementedTaskServiceServer() {
	panic("unimplemented")
}

func (s *server) init() {
	// Init clock on server
	s.clock.AddClock(s.name)
}

func (s *server) incrementClock() {
	s.clock.Tick(s.name)
}

func (s *server) determineNewClock(sender string) {
	s.clock.DetermineNewClock(sender, s.name)
}

func (s *server) getName() string {
	return s.name
}

func CreateServer(name string) (*server, error) {

	chittyChatServer := server{
		clock: service.LamportClock{},
		name:  name,
	}

	return &chittyChatServer, nil
}

func CreateGrpcServer(name string) (*grpc.Server, error) {
	grpcServer := grpc.NewServer()
	chatServer, err := CreateServer(name)
	chatServer.init()

	if err != nil {
		return nil, err
	}

	tasks.RegisterTaskServiceServer(grpcServer, *chatServer)

	log.Printf("Starting gRPC server with name: %s", name)

	chatServer.incrementClock()

	return grpcServer, nil

}
