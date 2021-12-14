package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/CodeZeo/T3-SD/Lab3_SD/comms"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedBrokerServer
}

// anexo para tratar errores
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (s *Server) getIP(ctx context.Context, command *pb.Command) (*pb.Conn, error) {

	/* for Vivos < 15 {
		time.Sleep(1 * time.Second) // Ojala funcione [si no chao]
	} */
	return &pb.Conn{Ip: randomIP()}, nil
}

func (s *Server) getNumberRebelds(ctx context.Context, locateCity *pb.LocateCity) (*pb.NumberRebelds, error) {

	/* for Vivos < 15 {
		time.Sleep(1 * time.Second) // Ojala funcione [si no chao]
	} */
	return &pb.NumberRebelds{NR: int32(gnr(locateCity.NombrePlaneta, locateCity.NombreCiudad))}, nil
}

func randomIP() string {
	ips := []string{"localhost:9005", "localhost:9006", "localhost:9007"}
	return ips[rand.Intn(len(ips))]
}

func gnr(planet string, city string) int {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(randomIP(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFulcrumClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ReturnNumberRebelds(ctx, &pb.LocateCity{NombrePlaneta: planet, NombreCiudad: city})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return int(r.NR)
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
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9003))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterBrokerServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
