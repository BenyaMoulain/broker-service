// TODO: Create nombre.dominio : crea registro ZF

// TODO: Update nombre.dominio <option> <parameter> : option puede ser name o IP

// TODO: Delete nombre.dominio : deja en blanco la linea correspondiente

// TODO: Verifica la validez de los comandos a enviar

// TODO: Envia comando al broker para recibir IP de un DNS

// TODO: En caso de conflicto solicita IP al broker

// TODO: Envia comando al servidor DNS correspondiente a la IP recibida por el broker

// TODO: Guarda en memoria relojes vector para cada registro ZF que ha cambiado

// TODO: Guarda en memoria la IP del servidor DNS al que se conectó por última
// vez a cada registro ZF

// TODO: Consistencia -> Read your Writes

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	broker "github.com/benyamoulain/broker-service/broker/broker_service"
	dns "github.com/benyamoulain/broker-service/dns/dns_service"

	"google.golang.org/grpc"
)

var (
	brokerAddr = flag.String("server_addr", "localhost:10001", "The server address in the format of host:port")
	dnsPort    = flag.String("dns_port", "10000", "Puerto a ser usado por los servidores DNS")
	zfMap      = make(map[string]zfRegister)
)

type zfRegister struct {
	VectorClock []int32
	LastDNS     string
}

func createDomain(client broker.BrokerServiceClient, domainName string, ip string) {
	domainNameArray := strings.Split(domainName, ".")
	// Separa el domain y name
	domain := domainNameArray[1]

	req := &dns.CreateRequest{
		DomainName: domainName,
		Ip:         ip,
	}
	log.Printf("Llamada a Create: DomainName: %v, IP: %v", req.GetDomainName(), req.GetIp())

	// Obtiene IP y abre conexión con servidor DNS
	brokerRes, err := client.GetDNS(context.Background(), &broker.GetDNSRequest{})
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	dnsIP := brokerRes.GetIp()

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", dnsIP, *dnsPort), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	dnsClient := dns.NewDNSServiceClient(conn)

	// Envia comando al servidor DNS
	res, err := dnsClient.Create(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	// Guarda el reloj vector y la ip del servidor dns
	vectorClock := res.GetVectorClock()
	zfMap[domain] = zfRegister{
		VectorClock: vectorClock,
		LastDNS:     dnsIP,
	}
	log.Printf("Respuesta de Create: %v", vectorClock)
}

func updateDomain(client broker.BrokerServiceClient, domainName string, option string, parameter string) {
	var boolOption bool
	domainNameArray := strings.Split(domainName, ".")
	// Separa el domain y name
	domain := domainNameArray[1]

	if option == "name" {
		boolOption = false
	} else if option == "ip" {
		boolOption = true
	} else {
		log.Printf("la opción %s no es válida.", option)
		return
	}

	req := &dns.UpdateRequest{
		DomainName: domainName,
		Option:     boolOption,
		Parameter:  parameter,
	}
	log.Printf("Llamada a Update: DomainName: %v, Option: %v, Parameter: %v", req.GetDomainName(), req.GetOption(), req.GetParameter())

	// Obtiene IP y abre conexión con servidor DNS
	brokerRes, err := client.GetDNS(context.Background(), &broker.GetDNSRequest{})
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	dnsIP := brokerRes.GetIp()

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", dnsIP, *dnsPort), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	dnsClient := dns.NewDNSServiceClient(conn)

	res, err := dnsClient.Update(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	// Guarda el reloj vector y la ip del servidor dns
	vectorClock := res.GetVectorClock()
	zfMap[domain] = zfRegister{
		VectorClock: vectorClock,
		LastDNS:     dnsIP,
	}
	log.Printf("Respuesta de Update: %v", vectorClock)
}

func deleteDomain(client broker.BrokerServiceClient, domainName string) {
	domainNameArray := strings.Split(domainName, ".")
	// Separa el domain y name
	domain := domainNameArray[1]

	req := &dns.DeleteRequest{
		DomainName: domainName,
	}
	log.Printf("Llamada a Delete: DomainName: %v", req.GetDomainName())

	// Obtiene IP y abre conexión con servidor DNS
	brokerRes, err := client.GetDNS(context.Background(), &broker.GetDNSRequest{})
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	dnsIP := brokerRes.GetIp()

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", dnsIP, *dnsPort), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	dnsClient := dns.NewDNSServiceClient(conn)

	res, err := dnsClient.Delete(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	// Guarda el reloj vector y la ip del servidor dns
	vectorClock := res.GetVectorClock()
	zfMap[domain] = zfRegister{
		VectorClock: vectorClock,
		LastDNS:     dnsIP,
	}
	log.Printf("Respuesta de Delete: %v", vectorClock)
}

func showCommands() {
	fmt.Println("---------------------")
	fmt.Println("help \t\t\t\t\t\t- Muestra los comandos disponibles")
	fmt.Println("exit \t\t\t\t\t\t- Para volver a la consola")
	fmt.Println("create nombre.dominio ip \t\t\t\t- Registra un nuevo nombre de dominio en los servidores DNS con la ip ingresada")
	fmt.Println("update nombre.dominio <option> <parameter> \t- <option> puede ser name o ip, parameter es el nuevo valor a reemplazar")
	fmt.Println("delete nombre.dominio \t\t\t\t- Elimina el nombre en el registro ZF del dominio")
	fmt.Println("---------------------")
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*brokerAddr, opts...)
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
		} else if length >= 2 {
			// Verifica si se ingresó un comando válido

			if length >= 2 && length <= 4 {
				action := splitedString[0]
				domainName := splitedString[1]
				if action == "create" {
					if length < 3 {
						fmt.Println("No ingresó los argumentos necesarios para ejecutar el comando create. Ingrese help para más información.")
					} else {
						IP := splitedString[2]
						createDomain(client, domainName, IP)
					}
				} else if action == "update" {
					if length < 4 {
						fmt.Println("No ingresó los argumentos necesarios para ejecutar el comando update. Ingrese help para más información.")
					} else {
						option := splitedString[2]
						parameter := splitedString[3]
						updateDomain(client, domainName, option, parameter)
					}
				} else if action == "delete" {
					deleteDomain(client, domainName)
				} else {
					fmt.Println("El comando ingresado es invalido, porfavor ingrese help para ver los comandos disponibles.")
				}
			}
		} else {
			fmt.Println("El comando ingresado es invalido o no se ingresaron los campos requeridos, porfavor ingrese help para ver los comandos disponibles.")
		}
	}
	// createDomain(client, "google.cl", "127.0.0.1")
}
