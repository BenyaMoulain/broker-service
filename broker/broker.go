// TODO: delegación de carga a los DNS, uso de random para
// elegir servidor

// TODO: Creación de servicio, broker actua como servidor intermediario

// TODO: Entrega IP de servidor DNS específico al admin en caso de conflicto

// serverAddress1 := fmt.Sprintf("%s:%s", *serverIP1, *serverPort)
// serverAddress2 := fmt.Sprintf("%s:%s", *serverIP2, *serverPort)
// serverAddress3 := fmt.Sprintf("%s:%s", *serverIP3, *serverPort)

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"

	"google.golang.org/grpc"

	broker "github.com/benyamoulain/broker-service/broker/broker_service"
	dns "github.com/benyamoulain/broker-service/dns/dns_service"
)

var (
	port      = flag.Int("port", 10001, "The server port")
	serverIP1 = flag.String("dns1_ip", "localhost", "Dirección IP del servidor DNS #1")
	serverIP2 = flag.String("dns2_ip", "localhost", "Dirección IP del servidor DNS #2")
	serverIP3 = flag.String("dns3_ip", "localhost", "Dirección IP del servidor DNS #3")
	dnsPort   = flag.String("dns_port", "10000", "Puerto que usarán los DNS")
	brokerIP  = flag.String("broker_ip", "localhost", "Dirección IP del broker")
)

type brokerServiceServer struct {
	broker.UnimplementedBrokerServiceServer
}

func newServer() *brokerServiceServer {
	s := &brokerServiceServer{}
	return s
}

// Obtiene la ip de un nombre de dominio
func (s *brokerServiceServer) Read(ctx context.Context, req *broker.ReadRequest) (*broker.ReadResponse, error) {

	domainNameArray := strings.Split(req.GetDomainName(), ".")
	// Separa el domain y name
	name, domain := domainNameArray[0], domainNameArray[1]
	log.Printf("Solicitud Read -> Domain: %s, Name: %s", domain, name)

	dnsReq := &dns.ReadRequest{
		DomainName: req.GetDomainName(),
	}
	dnsAddr := getDNSAddr()

	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(dnsAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := dns.NewDNSServiceClient(conn)
	dnsRes, err := client.Read(context.Background(), dnsReq)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Respuesta de Read: %v", dnsRes.GetVectorClock())

	res := &broker.ReadResponse{
		Ip:          dnsRes.GetIp(),
		VectorClock: dnsRes.GetVectorClock(),
		DnsIp:       strings.Split(dnsAddr, ":")[0],
	}

	return res, nil
}

// Obtiene la ip de un nombre de dominio desde un servidor DNS en particular en caso de conflicto
func (s *brokerServiceServer) ReadConflict(ctx context.Context, req *broker.ReadConflictRequest) (*broker.ReadConflictResponse, error) {

	domainNameArray := strings.Split(req.GetDomainName(), ".")
	// Separa el domain y name
	name, domain := domainNameArray[0], domainNameArray[1]
	log.Printf("Solicitud Read -> Domain: %s, Name: %s", domain, name)
	dnsIP := req.GetDnsIp()

	dnsReq := &dns.ReadRequest{
		DomainName: req.GetDomainName(),
	}
	dnsAddr := fmt.Sprintf("%s:%s", dnsIP, *dnsPort)

	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(dnsAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := dns.NewDNSServiceClient(conn)
	dnsRes, err := client.Read(context.Background(), dnsReq)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Respuesta de Read: %v", dnsRes.GetVectorClock())

	res := &broker.ReadConflictResponse{
		Ip:          dnsRes.GetIp(),
		VectorClock: dnsRes.GetVectorClock(),
		DnsIp:       strings.Split(dnsAddr, ":")[0],
	}

	return res, nil
}

func (s *brokerServiceServer) GetDNS(ctx context.Context, req *broker.GetDNSRequest) (*broker.GetDNSResponse, error) {
	dnsIP := strings.Split(getDNSAddr(), ":")[0]
	res := &broker.GetDNSResponse{
		Ip: dnsIP,
	}
	return res, nil

}

// Obtiene una ip de dns al azar
func getDNSAddr() string {
	dnsArray := []string{*serverIP1, *serverIP2, *serverIP3}
	randomIndex := rand.Intn(3)
	fmt.Printf("Se obtuvo la DNS #%d \n", randomIndex)
	dnsIP := dnsArray[randomIndex]
	return fmt.Sprintf("%s:%s", dnsIP, *dnsPort)
}

func main() {
	fmt.Println("Starting server...")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *brokerIP, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	broker.RegisterBrokerServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
