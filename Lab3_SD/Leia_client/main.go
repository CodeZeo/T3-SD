package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "../comms"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func gnr() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBrokerClient(conn)

	var planeta string
	var ciudad string

	fmt.Println("Ingrese el planeta a buscar: ")
	fmt.Scanln(&planeta)

	fmt.Println("Ingrese la ciudad a buscar: ")
	fmt.Scanln(&ciudad)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetNumberRebelds(ctx, &pb.LocateCity{NombrePlaneta: planeta, NombreCiudad: ciudad})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("La cantidad de rebeldes es: %s", strconv.Itoa(int(r.NR)))
}

func main() {
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
