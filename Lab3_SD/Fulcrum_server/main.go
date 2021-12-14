package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/CodeZeo/T3-SD/Lab3_SD/comms"
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

func fileToSlice(fileName string) ([][]string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	var result [][]string
	for _, line := range lines {
		result = append(result, strings.Split(line, " "))
	}
	return result, nil
}

func sliceToFile(slice [][]string, fileName string) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range slice {
		for _, word := range line {
			if _, err := f.WriteString(word + " "); err != nil {
				log.Fatal(err)
			}
		}
		if _, err := f.WriteString("\n"); err != nil {
			log.Fatal(err)
		}
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func createFile(name string) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		os.Create(name)
	}
}

func fileAddCity(nombrePlaneta string, nombreCiudad string, valor int) {
	fileName := nombrePlaneta + ".txt"
	append := nombrePlaneta + " " + nombreCiudad + " " + strconv.Itoa(valor) + "\n"
	createFile(fileName)
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(append)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func fileUpdateName(nombrePlaneta string, nombreCiudad string, nuevoNombre string) {
	content, err := fileToSlice(nombrePlaneta + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	for i, line := range content {
		if line[1] == nombreCiudad {
			content[i][1] = nuevoNombre
		}
	}
	sliceToFile(content, nombrePlaneta+".txt")
}

func fileUpdateNumber(nombrePlaneta string, nombreCiudad string, valor int) {
	content, err := fileToSlice(nombrePlaneta + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	for i, line := range content {
		if line[1] == nombreCiudad {
			content[i][2] = strconv.Itoa(valor)
		}
	}
	sliceToFile(content, nombrePlaneta+".txt")
}

func fileDeleteCity(nombrePlaneta string, nombreCiudad string) {
	content, err := fileToSlice(nombrePlaneta + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	var result [][]string
	for _, line := range content {
		if line[1] != nombreCiudad {
			result = append(result, line)
		}
	}
	sliceToFile(result, nombrePlaneta+".txt")
}

func (s *Server) returnNumberRebelds(ctx context.Context, numberRebelds *pb.Empty) (*pb.NumberRebelds, error) {

	return &pb.NumberRebelds{NR: int32(2)}, nil
}

func (s *Server) addCity(ctx context.Context, dataCity *pb.DataCity) (*pb.Clock, error) {

	fileAddCity(dataCity.NombrePlaneta, dataCity.NombreCiudad, int(dataCity.NuevoValor))
	return &pb.Clock{X: 1, Y: 2, Z: 3}, nil
}

func (s *Server) updateName(ctx context.Context, ChangeNameCity *pb.ChangeNameCity) (*pb.Clock, error) {

	fileUpdateName(ChangeNameCity.NombrePlaneta, ChangeNameCity.NombreCiudad, ChangeNameCity.NuevoNombre)
	return &pb.Clock{X: 2, Y: 2, Z: 3}, nil
}

func (s *Server) updateNumber(ctx context.Context, dataCity *pb.DataCity) (*pb.Clock, error) {

	fileUpdateNumber(dataCity.NombrePlaneta, dataCity.NombreCiudad, int(dataCity.NuevoValor))
	return &pb.Clock{X: 3, Y: 2, Z: 3}, nil
}

func (s *Server) deleteCity(ctx context.Context, locateCity *pb.LocateCity) (*pb.Clock, error) {

	fileDeleteCity(locateCity.NombrePlaneta, locateCity.NombreCiudad)
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
