package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "../comms"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedFulcrumServer
}

// anexo para tratar errores
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (s *Server) returnNumberRebelds(ctx context.Context, numberRebelds *pb.Empty) (*pb.NumberRebelds, error) {

	/* for Vivos < 15 {
		time.Sleep(1 * time.Second) // Ojala funcione [si no chao]
	} */
	return &pb.NumberRebelds{NR: int32(2)}, nil
}

func (s *Server) addCity(ctx context.Context, dataCity *pb.DataCity) (*pb.Clock, error) {

	/* for Vivos < 15 {
		time.Sleep(1 * time.Second) // Ojala funcione [si no chao]
	} */
	return &pb.Clock{X: 1, Y: 2, Z: 3}, nil
}

func (s *Server) updateName(ctx context.Context, dataCity *pb.DataCity) (*pb.Clock, error) {

	/* for Vivos < 15 {
		time.Sleep(1 * time.Second) // Ojala funcione [si no chao]
	} */
	return &pb.Clock{X: 2, Y: 2, Z: 3}, nil
}

func (s *Server) updateNumber(ctx context.Context, dataCity *pb.DataCity) (*pb.Clock, error) {

	/* for Vivos < 15 {
		time.Sleep(1 * time.Second) // Ojala funcione [si no chao]
	} */
	return &pb.Clock{X: 3, Y: 2, Z: 3}, nil
}

func (s *Server) deleteCity(ctx context.Context, locateCity *pb.LocateCity) (*pb.Clock, error) {

	/* for Vivos < 15 {
		time.Sleep(1 * time.Second) // Ojala funcione [si no chao]
	} */
	return &pb.Clock{X: 3, Y: 2, Z: 3}, nil
}

// Main, basicamente corre todo
func main() {

	fmt.Println("Soy el Fulcrum!")
	//q, errr, ch := openRMQ()
	// Estas variables se usan cada vez que se elimina alguien
	// Se debe llamar a sendJugadorEliminadoPozo()

	//parte cliente Lider-nameNode
	//parte Servidor Lider-Jugadores
	//cantRondasJuego1 := 1
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9005)) //9005 o 9006 o 9007
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterFulcrumServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
