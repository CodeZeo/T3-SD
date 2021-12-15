package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	pb "github.com/CodeZeo/T3-SD/Lab3_SD/comms"
	"google.golang.org/grpc"
)

//Struct con los relojes
type relojes struct {
	planeta string
	x       int
	y       int
	z       int
}

var listarelojes []relojes

//comparar relojes
func compareClock(planeta string, relojj *pb.Clock) bool {
	consistencia := true
	for i := 0; i < len(listarelojes); i++ {
		if listarelojes[i].planeta == planeta {
			if int32(listarelojes[i].x) > relojj.X || int32(listarelojes[i].y) > relojj.Y || int32(listarelojes[i].z) > relojj.Z {
				consistencia = false
			}
		}
	}
	return consistencia
}

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
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	//conectar al Fulcrum
	conn, err = grpc.Dial(response.Ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	ccc := pb.NewFulcrumClient(conn)
	//realizar Create
	reloj, err := ccc.AddCity(context.Background(), &pb.DataCity{NombrePlaneta: Data.NombrePlaneta, NombreCiudad: Data.NombreCiudad, NuevoValor: Data.NuevoValor})
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	return pb.Clock{X: int32(reloj.X), Y: int32(reloj.Y), Z: int32(reloj.Z)}
}

//Función que realiza consultas al Broker para obtener la dirección de algún Fulcrum al que conectarse a realiza el update
func upNCity(Data *pb.ChangeNameCity) pb.Clock {
	//Consultar por la ip de un Server Fulcrum
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9003", grpc.WithInsecure()) // deberia ser siempre 9003
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
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	//conectar al Fulcrum
	conn, err = grpc.Dial(response.Ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	ccc := pb.NewFulcrumClient(conn)
	//realizar Create
	reloj, err := ccc.UpdateName(context.Background(), &pb.ChangeNameCity{NombrePlaneta: Data.NombrePlaneta, NombreCiudad: Data.NombreCiudad, NuevoNombre: Data.NuevoNombre})
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
	comando := "UpdateNumber " + Data.NombrePlaneta + " " + Data.NombreCiudad + " " + strconv.Itoa(int(Data.NuevoValor))
	response, err := cc.GetIP(context.Background(), &pb.Command{C: comando})
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	//conectar al Fulcrum
	conn, err = grpc.Dial(response.Ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	ccc := pb.NewFulcrumClient(conn)
	//realizar Create
	reloj, err := ccc.UpdateNumber(context.Background(), &pb.DataCity{NombrePlaneta: Data.NombrePlaneta, NombreCiudad: Data.NombreCiudad, NuevoValor: Data.NuevoValor})
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
	comando := "DeleteCity " + Data.NombrePlaneta + " " + Data.NombreCiudad //broker
	//DeleteCity Planeta Ciudad
	response, err := cc.GetIP(context.Background(), &pb.Command{C: comando})
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	//conectar al Fulcrum
	conn, err = grpc.Dial(response.Ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	ccc := pb.NewFulcrumClient(conn)
	//comparar reloj
	//relojito, err := ccc.GetClock(context.Background(), &pb.Planet{Planet: Data.NombrePlaneta})
	//if err != nil {
	//	log.Fatalf("did not delete: %s", err)
	//}
	//consistencia:=compareClock(Data.NombrePlaneta,relojito)
	//if consistencia==false{
	//
	//}
	//realizar Create
	reloj, err := ccc.DeleteCity(context.Background(), &pb.LocateCity{NombrePlaneta: Data.NombrePlaneta, NombreCiudad: Data.NombreCiudad})
	if err != nil {
		log.Fatalf("did not delete: %s", err)
	}
	return pb.Clock{X: int32(reloj.X), Y: int32(reloj.Y), Z: int32(reloj.Z)}
}

func gnr() relojes {
	flag.Parse() // no se que hace esto pero me da miedo sacarlo
	//c := pb.NewBrokerClient(conn)

	fmt.Println("Ingrese el comando: ")

	in := bufio.NewReader(os.Stdin)
	linea, err := in.ReadString('\n')
	if err != nil {
		log.Fatalf("err read: %v", err)
	}
	s := strings.Fields(linea)
	var reloj pb.Clock
	var planeta string
	if s[0] == "AddCity" {
		if len(s) == 4 {
			numero, err := strconv.Atoi(s[3])
			if err != nil {
				fmt.Println("NuevoValor invalido: ", err)
			} else {
				reloj = agCity(&pb.DataCity{NombrePlaneta: s[1], NombreCiudad: s[2], NuevoValor: int32(numero)})
				planeta = s[1]
			}
		} else {
			reloj = agCity(&pb.DataCity{NombrePlaneta: s[1], NombreCiudad: s[2], NuevoValor: 0})
			planeta = s[1]
		}
	} else if s[0] == "UpdateName" {
		reloj = upNCity(&pb.ChangeNameCity{NombrePlaneta: s[1], NombreCiudad: s[2], NuevoNombre: s[3]})
		planeta = s[1]
	} else if s[0] == "UpdateNumber" {
		numero, err := strconv.Atoi(s[3])
		if err != nil {
			fmt.Println("NuevoValor invalido: ", err)
		} else {
			reloj = upVCity(&pb.DataCity{NombrePlaneta: s[1], NombreCiudad: s[2], NuevoValor: int32(numero)})
			planeta = s[1]
		}
	} else if s[0] == "DeleteCity" {
		reloj = DeleteC(&pb.LocateCity{NombrePlaneta: s[1], NombreCiudad: s[2]})
		planeta = s[1]
	} else {
		fmt.Println("Comando Invalido.")
	}
	// Contact the server and print out its response.
	log.Printf("El reloj es:[ %s,%s,%s ]", strconv.Itoa(int(reloj.X)), strconv.Itoa(int(reloj.Y)), strconv.Itoa(int(reloj.Z)))
	return relojes{planeta: planeta, x: int(reloj.X), y: int(reloj.Y), z: int(reloj.Z)}
}

func main() {
	//conectar al Broker
	flag := true
	var opcion int

	//listarelojes := make([]relojes, 0)
	for flag {
		fmt.Println("Ingrese una opcion: ")
		fmt.Println("1) Ingresar Comando")
		fmt.Println("2) Salir.")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			reloj := gnr()
			if len(listarelojes) == 0 {
				listarelojes = append(listarelojes, reloj)
			} else {
				for i := 0; i < len(listarelojes); i++ {
					if listarelojes[i].planeta == reloj.planeta {
						listarelojes[i].x = reloj.x
						listarelojes[i].y = reloj.y
						listarelojes[i].z = reloj.z
						i = len(listarelojes)
					} else if i == len(listarelojes)-1 {
						listarelojes = append(listarelojes, reloj)
					}
				}
			}
		} else {
			flag = false
		}
		fmt.Println(listarelojes)
	}
}
