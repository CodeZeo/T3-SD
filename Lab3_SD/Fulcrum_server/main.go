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

type reloj struct {
	planeta string
	x       int
	y       int
	z       int
}

var relojesfulcrum []reloj
var posi int

//Actualiza los relojes
func updateReloj(planeta string) int {
	var relojito reloj
	var indice int
	if len(relojesfulcrum) != 0 {
		for i := 0; i < len(relojesfulcrum); i++ {
			if planeta == relojesfulcrum[i].planeta {
				if posi == 0 {
					relojesfulcrum[i].x = relojesfulcrum[i].x + 1
				} else if posi == 1 {
					relojesfulcrum[i].y = relojesfulcrum[i].y + 1
				} else {
					relojesfulcrum[i].z = relojesfulcrum[i].z + 1
				}
				indice = i
				i = len(relojesfulcrum)
			} else if i == len(relojesfulcrum)-1 {
				if posi == 0 {
					relojito.planeta = planeta
					relojito.z = 0
					relojito.x = 0
					relojito.x = 1
					relojesfulcrum = append(relojesfulcrum, relojito)
					indice = len(relojesfulcrum) - 1
				} else if posi == 1 {
					relojito.planeta = planeta
					relojito.z = 0
					relojito.x = 0
					relojito.y = 1
					relojesfulcrum = append(relojesfulcrum, relojito)
					indice = len(relojesfulcrum) - 1
				} else {
					relojito.planeta = planeta
					relojito.z = 1
					relojito.x = 0
					relojito.y = 0
					relojesfulcrum = append(relojesfulcrum, relojito)
					indice = len(relojesfulcrum) - 1
				}
			}
		}
	} else {
		if posi == 0 {
			relojito.planeta = planeta
			relojito.z = 0
			relojito.x = 0
			relojito.x = 1
			relojesfulcrum = append(relojesfulcrum, relojito)
			indice = len(relojesfulcrum) - 1
		} else if posi == 1 {
			relojito.planeta = planeta
			relojito.z = 0
			relojito.x = 0
			relojito.y = 1
			relojesfulcrum = append(relojesfulcrum, relojito)
			indice = len(relojesfulcrum) - 1
		} else {
			relojito.planeta = planeta
			relojito.z = 1
			relojito.x = 0
			relojito.y = 0
			relojesfulcrum = append(relojesfulcrum, relojito)
			indice = len(relojesfulcrum) - 1
		}
	}
	return indice
}

// Revisa si un archivo existe
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// Toma un archivo, y devuelve un slice de slice que
// son las lineas, separadas por espacio
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

// Revisa si una linea es valida, sirve para obviar los espacios en blanco
func lineValid(line []string) bool {
	for _, cosa := range line {
		if len(cosa) > 0 {
			return true
		}
	}
	return false
}

// Toma un slice de slice de strings y lo convierte en
// un archivo de nombre fileName con cada linea un substring,
// y cada subsub string una palabra separada por espacio
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

//Revisa si un archivo existe, y si no existe lo crea
func createFile(name string) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		os.Create(name)
	}
}

// Le agrega una ciudad a un archivo planeta
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

// Le actualiza el nombre a una ciudad en el archivo planeta
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

// Le actualiza la cantidad de rebeldes a una ciudad en el archivo planeta
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

// Elimina una ciudad en un archivo planeta
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

// Devuelve la cantidad de rebeldes en una ciudad
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

// Servicio que retorna la cantidad de rebeldes en una ciudad
func (s *Server) ReturnNumberRebelds(ctx context.Context, LocateCity *pb.LocateCity) (*pb.NumberRebelds, error) {
	fmt.Println("RNR invoked")
	return &pb.NumberRebelds{NR: int32(fileNumberRebelds(LocateCity.NombrePlaneta, LocateCity.NombreCiudad))}, nil
}

// Servicio que crea la ciudad, y retorna el reloj
func (s *Server) AddCity(ctx context.Context, dataCity *pb.DataCity) (*pb.Clock, error) {
	fmt.Println("add city invoked")
	fileAddCity(dataCity.NombrePlaneta, dataCity.NombreCiudad, int(dataCity.NuevoValor))
	i := updateReloj(dataCity.NombrePlaneta)
	return &pb.Clock{X: int32(relojesfulcrum[i].x), Y: int32(relojesfulcrum[i].y), Z: int32(relojesfulcrum[i].z)}, nil
}

// Servicio que actualiza el nombre de la ciudad, y retorna el reloj
func (s *Server) UpdateName(ctx context.Context, ChangeNameCity *pb.ChangeNameCity) (*pb.Clock, error) {
	fmt.Println("update name invoked")
	fileUpdateName(ChangeNameCity.NombrePlaneta, ChangeNameCity.NombreCiudad, ChangeNameCity.NuevoNombre)
	i := updateReloj(ChangeNameCity.NombrePlaneta)
	return &pb.Clock{X: int32(relojesfulcrum[i].x), Y: int32(relojesfulcrum[i].y), Z: int32(relojesfulcrum[i].z)}, nil
}

// Servicio que actualiza la cantidad de rebeldes de la ciudad, y retorna el reloj
func (s *Server) UpdateNumber(ctx context.Context, dataCity *pb.DataCity) (*pb.Clock, error) {
	fmt.Println("update number invoked")
	fileUpdateNumber(dataCity.NombrePlaneta, dataCity.NombreCiudad, int(dataCity.NuevoValor))
	i := updateReloj(dataCity.NombrePlaneta)
	return &pb.Clock{X: int32(relojesfulcrum[i].x), Y: int32(relojesfulcrum[i].y), Z: int32(relojesfulcrum[i].z)}, nil
}

// Servicio que elimina la ciudad, y retorna el reloj
func (s *Server) DeleteCity(ctx context.Context, locateCity *pb.LocateCity) (*pb.Clock, error) {
	fmt.Println("delete city invoked")
	fileDeleteCity(locateCity.NombrePlaneta, locateCity.NombreCiudad)
	i := updateReloj(locateCity.NombrePlaneta)
	return &pb.Clock{X: int32(relojesfulcrum[i].x), Y: int32(relojesfulcrum[i].y), Z: int32(relojesfulcrum[i].z)}, nil
}

// Servicio que y retorna el reloj de un planeta
func (s *Server) GetClock(ctx context.Context, planeta *pb.Planet) (*pb.Clock, error) {
	var x1 int
	var y1 int
	var z1 int
	if len(relojesfulcrum) != 0 {
		for i := 0; i < len(relojesfulcrum); i++ {
			if planeta.Planet == relojesfulcrum[i].planeta {
				x1 = relojesfulcrum[i].x
				y1 = relojesfulcrum[i].y
				z1 = relojesfulcrum[i].z
				i = len(relojesfulcrum)
			} else if i == len(relojesfulcrum)-1 {
				x1 = 0
				y1 = 0
				z1 = 0
			}
		}
	} else {
		x1 = 0
		y1 = 0
		z1 = 0
	}
	return &pb.Clock{X: int32(x1), Y: int32(y1), Z: int32(z1)}, nil
}

// Main, basicamente corre todo
func main() {

	if len(os.Args) < 2 {
		log.Fatalln("no se especifico correctmente el fullcrum a correr, deberia ser go run main.go 0 o go run main.go 1 o go run main.go 2")
	}
	if os.Args[1] == "0" {
		posi = 0
	} else if os.Args[1] == "1" {
		posi = 1
	} else if os.Args[1] == "2" {
		posi = 2
	} else {
		log.Fatalln("no se especifico correctmente el fullcrum a correr, deberia ser go run main.go 0 o go run main.go 1 o go run main.go 2")
	}
	fmt.Println("Soy el Fulcrum!")
	//q, errr, ch := openRMQ()
	// Estas variables se usan cada vez que se elimina alguien
	// Se debe llamar a sendJugadorEliminadoPozo()

	//parte cliente Lider-nameNode
	//parte Servidor Lider-Jugadores
	//cantRondasJuego1 := 1
	puerto := []int{9005, 9006, 9007}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", puerto[posi])) //9005 o 9006 o 9007
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterFulcrumServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
