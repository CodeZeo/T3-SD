package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/CodeZeo/T3-SD/Lab3_SD/comms"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func agCity(Data *pb.DataCity) pb.Clock {
	//Consultar por la ip de un Server Fulcrum
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9003", grpc.WithInsecure())
	if err != nil {
		conn, err = grpc.Dial("localhost:9004", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
	}
	defer conn.Close()
	cc := pb.NewBrokerClient(conn)
	comando := "AddCity " + Data.NombrePlaneta + " " + Data.NombreCiudad
	response, err := cc.GetIP(context.Background(), &pb.Command{C: comando})
	//conectar al Fulcrum
	conn, err = grpc.Dial(response.Ip+":9005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	ccc := pb.NewFulcrumClient(conn)
	//realizar Create
	reloj, err := ccc.AddCity(context.Background(), &pb.DataCity{})

	return pb.Clock{X: int32(reloj.X), Y: int32(reloj.Y), Z: int32(reloj.Z)}
}

func gnr() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//c := pb.NewBrokerClient(conn)

	var planeta string
	var ciudad string

	fmt.Println("Ingrese el planeta a buscar: ")
	fmt.Scanln(&planeta)

	fmt.Println("Ingrese la ciudad a buscar: ")
	fmt.Scanln(&ciudad)
	// Contact the server and print out its response.
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	r := agCity(&pb.DataCity{NombrePlaneta: planeta, NombreCiudad: ciudad, NuevoValor: 0})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("El reloj es: %s", strconv.Itoa(int(r.NR)))
}

func main() {
	//conectar al Broker

	//agCity()

	flag := true
	var opcion int
	for flag {
		fmt.Println("Ingrese una opcion: ")
		fmt.Println("1) Ver rebeldes en ciudad.")
		fmt.Println("2) Salir.")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			gnr()
		} else {
			flag = false

		}
	}
}
