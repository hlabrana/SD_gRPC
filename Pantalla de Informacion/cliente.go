package main

import (
	"log"
	"time"
	"fmt"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/protos"
)

func ConsultarTorre (torre Torre){
	// Set up a connection to the server.
	conn, err := grpc.Dial(torre.ipTorre, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	stub := pb.NewAPIClient(conn)

	// Contact the server and print out its response.
	mensaje1 := "departures"
	mensaje2 := "arrivals"
	if len(os.Args) > 1 {
		mensaje1 = os.Args[1]
		mensaje2 = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	departures, err := stub.Departures(ctx, &pb.Mensaje{Mensaje: mensaje1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	arrivals, err := stub.Arrivals(ctx, &pb.Mensaje{Mensaje: mensaje2})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Imprimir la información de los vuelos por pantalla.
	fmt.Print("\n\n")
	fmt.Println("[Pantalla de información - ",torre.NombreTorre,"]")
	fmt.Println("Departures								|   Arrivals")
	fmt.Println("Avion		Destino		Pista		Hr	|	Avion		Proveniente		Pista		Hr")
	/*for f:=0; f<len(departures); f++{
	fmt.Print(departures[f].codigoAvion,"		",departures[f].destino,"		",departures[f].pistaDespegue,"		",departures[f].horaDespegue," | ")
	fmt.Println(arrivals[f].codigoAvion,"		",arrivals[f].proveniente,"		",arrivals[f].pistaLlegada,"		",arrivals[f].horaLlegada)
	}*/
	fmt.Println(departures)
	fmt.Println(arrivals)
}