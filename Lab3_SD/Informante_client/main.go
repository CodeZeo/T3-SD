package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

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

//Función que realiza consultas al Broker para obtener la dirección de algún Fulcrum al que conectarse a realiza el add
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
	comando := "AddCity " + Data.NombrePlaneta + " " + Data.NombreCiudad + " " + strconv.Itoa(int(Data.NuevoValor))
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

//Función que realiza consultas al Broker para obtener la dirección de algún Fulcrum al que conectarse a realiza el update
func upNCity(Data *pb.ChangeNameCity) pb.Clock {
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
	comando := "UpdateName " + Data.NombrePlaneta + " " + Data.NombreCiudad + " " + Data.NuevoNombre
	response, err := cc.GetIP(context.Background(), &pb.Command{C: comando})
	//conectar al Fulcrum
	conn, err = grpc.Dial(response.Ip+":9005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	ccc := pb.NewFulcrumClient(conn)
	//realizar Create
	reloj, err := ccc.UpdateName(context.Background(), &pb.ChangeNameCity{})
	if err != nil {
		log.Fatalf("did not update: %s", err)
	}
	return pb.Clock{X: int32(reloj.X), Y: int32(reloj.Y), Z: int32(reloj.Z)}
}

//Función que realiza consultas al Broker para obtener la dirección de algún Fulcrum al que conectarse a realiza el update
func upVCity(Data *pb.DataCity) pb.Clock {
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
	comando := "UpdateNumber" + Data.NombrePlaneta + " " + Data.NombreCiudad + " " + strconv.Itoa(int(Data.NuevoValor))
	response, err := cc.GetIP(context.Background(), &pb.Command{C: comando})
	//conectar al Fulcrum
	conn, err = grpc.Dial(response.Ip+":9005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	ccc := pb.NewFulcrumClient(conn)
	//realizar Create
	reloj, err := ccc.UpdateNumber(context.Background(), &pb.DataCity{})
	if err != nil {
		log.Fatalf("did not update: %s", err)
	}
	return pb.Clock{X: int32(reloj.X), Y: int32(reloj.Y), Z: int32(reloj.Z)}
}

//Función que realiza consultas al Broker para obtener la dirección de algún Fulcrum al que conectarse a realiza el delete
func DeleteC(Data *pb.LocateCity) pb.Clock {
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
	comando := "DeleteCity" + Data.NombrePlaneta + " " + Data.NombreCiudad
	response, err := cc.GetIP(context.Background(), &pb.Command{C: comando})
	//conectar al Fulcrum
	conn, err = grpc.Dial(response.Ip+":9005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	ccc := pb.NewFulcrumClient(conn)
	//realizar Create
	reloj, err := ccc.DeleteCity(context.Background(), &pb.LocateCity{})
	if err != nil {
		log.Fatalf("did not delete: %s", err)
	}
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
	var linea string

	fmt.Println("Ingrese el comando: ")
	fmt.Scanln(&linea)

	s := strings.Fields(linea)
	var reloj pb.Clock
	if s[0] == "AddCity" {
		if len(s) == 4 {
			numero, err := strconv.Atoi(s[3])
			if err != nil {
				fmt.Println("NuevoValor invalido: ", err)
			} else {
				reloj = agCity(&pb.DataCity{NombrePlaneta: s[1], NombreCiudad: s[2], NuevoValor: int32(numero)})
			}
		} else {
			reloj = agCity(&pb.DataCity{NombrePlaneta: s[1], NombreCiudad: s[2], NuevoValor: 0})
		}
	} else if s[0] == "UpdateName" {
		reloj = upNCity(&pb.ChangeNameCity{NombrePlaneta: s[1], NombreCiudad: s[2], NuevoNombre: s[3]})
	} else if s[0] == "UpdateNumber" {
		numero, err := strconv.Atoi(s[3])
		if err != nil {
			fmt.Println("NuevoValor invalido: ", err)
		} else {
			reloj = upVCity(&pb.DataCity{NombrePlaneta: s[1], NombreCiudad: s[2], NuevoValor: int32(numero)})
		}
	} else if s[0] == "DeleteCity" {
		reloj = DeleteC(&pb.LocateCity{NombrePlaneta: s[1], NombreCiudad: s[2]})
	} else {
		fmt.Println("Comando Invalido.")
	}
	// Contact the server and print out its response.
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r := agCity(&pb.DataCity{NombrePlaneta: planeta, NombreCiudad: ciudad, NuevoValor: 0})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	log.Printf("El reloj es:[ %s,%s,%s ]", strconv.Itoa(int(reloj.X)), strconv.Itoa(int(reloj.Y)), strconv.Itoa(int(reloj.Z)))
}

func main() {
	//conectar al Broker
	flag := true
	var opcion int
	for flag {
		fmt.Println("Ingrese una opcion: ")
		fmt.Println("1) Ingresar Comando")
		fmt.Println("2) Salir.")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			gnr()
		} else {
			flag = false
		}
	}
}
