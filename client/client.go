package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"

	broker "github.com/benyamoulain/broker-service/broker/broker_service"
)

var (
	serverAddr = flag.String("broker_addr", "localhost:10001", "The server address in the format of host:port")
	zfMap      = make(map[string]zfRegister)
)

type zfRegister struct {
	VectorClock []int32
	LastDNS     string
}

func getDomain(client broker.BrokerServiceClient, domainName string) {

	domainNameArray := strings.Split(domainName, ".")
	// Separa el domain y name
	domain := domainNameArray[1]

	req := &broker.ReadRequest{
		DomainName: domainName,
	}
	log.Printf("Llamada a Read: DomainName: %v", req.GetDomainName())
	res, err := client.Read(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	ip := res.GetIp()
	dnsIP := res.GetDnsIp()
	vectorClock := res.GetVectorClock()

	// Aplica Monotonic Reads de ser necesario
	if compareVectorClocks(zfMap[domain].VectorClock, vectorClock) {
		newRes, err := client.ReadConflict(context.Background(), &broker.ReadConflictRequest{DomainName: domainName, DnsIp: zfMap[domain].LastDNS})
		if err != nil {
			log.Fatalf("error while calling Greet RPC: %v", err)
		}
		ip = newRes.GetIp()
		dnsIP = newRes.GetDnsIp()
		vectorClock = newRes.GetVectorClock()
	}

	zfMap[domain] = zfRegister{
		VectorClock: vectorClock,
		LastDNS:     dnsIP,
	}

	log.Printf("Respuesta de Read -> vector_clock: %v, ip: %v, dns_ip: %v", vectorClock, ip, dnsIP)
	if ip == "" {
		fmt.Println("No se encontró una IP para el nombre de dominio solicitado")
	}
}

// Compara los relojes de vector para saber si aplicar Monotonic Reads
func compareVectorClocks(oldVectorClock []int32, newVectorClock []int32) bool {
	for index, clock := range oldVectorClock {
		if newVectorClock[index] < clock {
			return true
		}
	}
	return false
}

func showCommands() {
	fmt.Println("---------------------")
	fmt.Println("help \t\t\t\t\t\t- Muestra los comandos disponibles")
	fmt.Println("exit \t\t\t\t\t\t- Para volver a la consola")
	fmt.Println("get nombre.dominio \t\t\t\t- Solicita la IP de un nombre de dominio al broker")
	fmt.Println("---------------------")
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := broker.NewBrokerServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bienvenido")
	showCommands()
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.ToLower(text[0 : len(text)-1])
		splitedString := strings.Split(text, " ")
		length := len(splitedString)

		if strings.Compare("exit", text) == 0 {
			return
		} else if strings.Compare("help", text) == 0 {
			showCommands()
		} else if length == 2 {
			// Verifica si se ingresó un comando válido

			action := splitedString[0]
			domainName := splitedString[1]
			if action == "get" {
				getDomain(client, domainName)
			}
		} else {
			fmt.Println("El comando ingresado es invalido o no se ingresaron los campos requeridos, porfavor ingrese help para ver los comandos disponibles.")
		}
	}
	// createDomain(client, "google.cl", "127.0.0.1")
}
