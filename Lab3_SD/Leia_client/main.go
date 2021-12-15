package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

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

func gnr() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:9003", grpc.WithInsecure())
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
	cantRebelds := int(r.NR)
	if cantRebelds >= 0 {
		log.Printf("La cantidad de rebeldes es: %d", cantRebelds)
		fmt.Printf("El reloj encontrado es: {%d %d %d}", r.X, r.Y, r.Z)
		println("En la IP: %s", r.Ip)
	} else if cantRebelds == -1 {
		fmt.Println("El planeta no existe.")
	} else if cantRebelds == -2 {
		fmt.Println("La ciudad no existe.")
	} else {
		fmt.Println("uh, esto no deberia pasar nunca.")
	}
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
