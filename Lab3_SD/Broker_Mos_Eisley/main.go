package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"

	pb "github.com/CodeZeo/T3-SD/Lab3_SD/comms"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedBrokerServer
}

func getLatestServerPlanet(planet string) (string, int, int, int) {
	ips := []string{"localhost:9005", "localhost:9006", "localhost:9007"}
	latest := 0 // el clock mas tardio
	posLatest := rand.Intn(len(ips))
	clockChosen := []int{0, 0, 0}
	for i := 0; i < len(ips); i++ {
		conn, err := grpc.Dial(ips[i], grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewFulcrumClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		fmt.Println("Calling RNR")
		r, err := c.GetClock(ctx, &pb.Planet{Planet: planet})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		clock := []int{int(r.X), int(r.Y), int(r.Z)}
		if clock[i] > latest {
			latest = clock[i]
			posLatest = i
			clockChosen = clock
		}
	}
	// Set up a connection to the server.

	return ips[posLatest], clockChosen[0], clockChosen[1], clockChosen[2]
}

//func randomIP() string {
//	ips := []string{"localhost:9005", "localhost:9006", "localhost:9007"}
//	return ips[0] //mientras tanto para hacer pruebas
//	return ips[rand.Intn(len(ips))]
//}

func commandValid(c []string) bool {
	for _, cosa := range c {
		if len(cosa) > 0 {
			return true
		}
	}
	return false
}

func (s *Server) GetIP(ctx context.Context, command *pb.Command) (*pb.Conn, error) {
	strings := strings.Fields(command.C)
	var retorno string
	if commandValid(strings) {
		retorno, _, _, _ = getLatestServerPlanet(strings[1])
	} else {
		retorno = ""
	}
	fmt.Println("GetIP invoked")

	/* for Vivos < 15 {
		time.Sleep(1 * time.Second) // Ojala funcione [si no chao]
	} */
	return &pb.Conn{Ip: retorno}, nil
}

func (s *Server) GetNumberRebelds(ctx context.Context, locateCity *pb.LocateCity) (*pb.NumberRebeldsClock, error) {
	fmt.Println("GNR invoked")
	/* for Vivos < 15 {
		time.Sleep(1 * time.Second) // Ojala funcione [si no chao]
	} */
	nr, x, y, z, ip := gnr(locateCity.NombrePlaneta, locateCity.NombreCiudad)
	return &pb.NumberRebeldsClock{NR: int32(nr), X: int32(x), Y: int32(y), Z: int32(z), Ip: ip}, nil
}

func gnr(planet string, city string) (int, int, int, int, string) {
	flag.Parse()
	// Set up a connection to the server.
	ip, x, y, z := getLatestServerPlanet(planet)
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFulcrumClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("Calling RNR")
	r, err := c.ReturnNumberRebelds(ctx, &pb.LocateCity{NombrePlaneta: planet, NombreCiudad: city})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return int(r.NR), x, y, z, ip
}

// Main, basicamente corre todo
func main() {

	fmt.Println("Soy el Broker!")
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
