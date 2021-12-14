package main

import (
	"context"
	"errors"
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

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
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

func lineValid(line []string) bool {
	for _, cosa := range line {
		if len(cosa) > 0 {
			return true
		}
	}
	return false
}

func sliceToFile(slice [][]string, fileName string) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range slice {
		if lineValid(line) {
			for _, word := range line {
				if _, err := f.WriteString(word + " "); err != nil {
					log.Fatal(err)
				}
			}
			if _, err := f.WriteString("\n"); err != nil {
				log.Fatal(err)
			}
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
	fmt.Println("started file add city")
	fmt.Println(nombrePlaneta) // esto esta vacio
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
	//Archivo Log
	fmt.Println("started file Log")
	fmt.Println(nombrePlaneta) // esto esta vacio
	fileName2 := nombrePlaneta + "_Log.txt"
	append2 := "AddCity " + nombrePlaneta + " " + nombreCiudad + " " + strconv.Itoa(valor) + "\n"
	createFile(fileName2)
	f2, err2 := os.OpenFile(fileName2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
	if _, err := f2.Write([]byte(append2)); err != nil {
		log.Fatal(err)
	}
	if err := f2.Close(); err != nil {
		log.Fatal(err)
	}

}

func fileUpdateName(nombrePlaneta string, nombreCiudad string, nuevoNombre string) {
	content, err := fileToSlice(nombrePlaneta + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	for i, line := range content {
		if len(line) > 1 {
			if line[1] == nombreCiudad {
				content[i][1] = nuevoNombre
			}
		}
	}
	sliceToFile(content, nombrePlaneta+".txt")

	//Archivo Log
	fmt.Println("started file Log")
	fmt.Println(nombrePlaneta) // esto esta vacio
	fileName2 := nombrePlaneta + "_Log.txt"
	append2 := "UpdateName " + nombrePlaneta + " " + nombreCiudad + " " + nuevoNombre + "\n"
	createFile(fileName2)
	f2, err2 := os.OpenFile(fileName2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
	if _, err := f2.Write([]byte(append2)); err != nil {
		log.Fatal(err)
	}
	if err := f2.Close(); err != nil {
		log.Fatal(err)
	}
}

func fileUpdateNumber(nombrePlaneta string, nombreCiudad string, valor int) {
	content, err := fileToSlice(nombrePlaneta + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	for i, line := range content {
		if len(line) > 2 {
			if line[1] == nombreCiudad {
				content[i][2] = strconv.Itoa(valor)
			}
		}
	}
	sliceToFile(content, nombrePlaneta+".txt")

	//Archivo Log
	fmt.Println("started file Log")
	fmt.Println(nombrePlaneta) // esto esta vacio
	fileName2 := nombrePlaneta + "_Log.txt"
	append2 := "UpdateNumber " + nombrePlaneta + " " + nombreCiudad + " " + strconv.Itoa(valor) + "\n"
	createFile(fileName2)
	f2, err2 := os.OpenFile(fileName2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
	if _, err := f2.Write([]byte(append2)); err != nil {
		log.Fatal(err)
	}
	if err := f2.Close(); err != nil {
		log.Fatal(err)
	}
}

func fileDeleteCity(nombrePlaneta string, nombreCiudad string) {
	content, err := fileToSlice(nombrePlaneta + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	var result [][]string
	for _, line := range content {
		if len(line) > 1 {
			if line[1] != nombreCiudad {
				result = append(result, line)
			}
		}
	}
	sliceToFile(result, nombrePlaneta+".txt")

	//Archivo Log
	fmt.Println("started file Log")
	fmt.Println(nombrePlaneta) // esto esta vacio
	fileName2 := nombrePlaneta + "_Log.txt"
	append2 := "DeleteCity " + nombrePlaneta + " " + nombreCiudad + "\n"
	createFile(fileName2)
	f2, err2 := os.OpenFile(fileName2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
	if _, err := f2.Write([]byte(append2)); err != nil {
		log.Fatal(err)
	}
	if err := f2.Close(); err != nil {
		log.Fatal(err)
	}
}

// retorna -1 si el planeta no existe
// retorna -2 si la ciudad no existe
func fileNumberRebelds(planet string, city string) int {
	fileName := planet + ".txt"
	if !fileExists(fileName) {
		return -1
	}
	lineas, err := fileToSlice(fileName)
	if err != nil {
		log.Fatal(err)
	}
	for _, linea := range lineas {
		if lineValid(linea) {
			if linea[1] == city {
				num, err := strconv.Atoi(linea[2])
				if err != nil {
					log.Fatal(err)
				}
				return num
			}
		}
	}
	return -2
}

func (s *Server) ReturnNumberRebelds(ctx context.Context, LocateCity *pb.LocateCity) (*pb.NumberRebelds, error) {
	fmt.Println("RNR invoked")
	return &pb.NumberRebelds{NR: int32(fileNumberRebelds(LocateCity.NombrePlaneta, LocateCity.NombreCiudad))}, nil
}

func (s *Server) AddCity(ctx context.Context, dataCity *pb.DataCity) (*pb.Clock, error) {
	fmt.Println("add city invoked")
	fileAddCity(dataCity.NombrePlaneta, dataCity.NombreCiudad, int(dataCity.NuevoValor))
	return &pb.Clock{X: 1, Y: 2, Z: 3}, nil
}

func (s *Server) UpdateName(ctx context.Context, ChangeNameCity *pb.ChangeNameCity) (*pb.Clock, error) {
	fmt.Println("update name invoked")
	fileUpdateName(ChangeNameCity.NombrePlaneta, ChangeNameCity.NombreCiudad, ChangeNameCity.NuevoNombre)
	return &pb.Clock{X: 2, Y: 2, Z: 3}, nil
}

func (s *Server) UpdateNumber(ctx context.Context, dataCity *pb.DataCity) (*pb.Clock, error) {
	fmt.Println("update number invoked")
	fileUpdateNumber(dataCity.NombrePlaneta, dataCity.NombreCiudad, int(dataCity.NuevoValor))
	return &pb.Clock{X: 3, Y: 2, Z: 3}, nil
}

func (s *Server) DeleteCity(ctx context.Context, locateCity *pb.LocateCity) (*pb.Clock, error) {
	fmt.Println("delete city invoked")
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
