// TODO: Escribir el de los cambios realizados por el admin

// TODO: Almacenar registros ZF en disco y vectores de reloj en memoria

// TODO: Propagación cada 5 min.

// TODO: Merge en caso de conflictos -> Se conservará la ultima acción

package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"google.golang.org/grpc"

	dns "github.com/benyamoulain/broker-service/dns/dns_service"
)

var (
	serverIP     = flag.String("ip", "127.0.0.1", "IP del servidor")
	port         = flag.Int("port", 10000, "Puerto del servidor")
	serverNumber = flag.Int("number", 1, "Número del servidor DNS asignado, puede variar entre 1 y 3")
	zfMap        = make(map[string]zfRegister)
)

type zfRegister struct {
	VectorClock []int32
	FileMutex   *sync.Mutex
}

type dNSServiceServer struct {
	dns.UnimplementedDNSServiceServer
}

func newServer() *dNSServiceServer {
	s := &dNSServiceServer{}
	return s
}

// Lee la ip de un nombre de dominio en los registros ZF, además entrega el vector reloj
func (s *dNSServiceServer) Read(ctx context.Context, req *dns.ReadRequest) (*dns.ReadResponse, error) {
	domainName := req.GetDomainName()
	domainNameArray := strings.Split(req.GetDomainName(), ".")
	ip := ""

	// Separa el domain y name
	name, domain := domainNameArray[0], domainNameArray[1]
	log.Printf("Solicitud Read -> Domain: %s, Name: %s", domain, name)

	dirPath := fmt.Sprintf("domains/%s/", domain)
	zfPath := dirPath + "zf.data"

	// Busca en el registro zf
	zfFile, err := ioutil.ReadFile(zfPath)
	if err != nil {
		log.Println(err)
	}
	lines := strings.Split(string(zfFile), "\n")

	for _, line := range lines {
		if strings.Contains(line, domainName) {
			ip = strings.Split(line, " ")[3]
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(zfPath, []byte(output), 0644)
	if err != nil {
		log.Println(err)
	}
	res := &dns.ReadResponse{
		Ip:          ip,
		VectorClock: zfMap[domain].VectorClock,
	}

	return res, nil
}

// Crea un nuevo nombre de dominio
func (s *dNSServiceServer) Create(ctx context.Context, req *dns.CreateRequest) (*dns.CreateResponse, error) {

	domainNameArray := strings.Split(req.GetDomainName(), ".")
	ip := req.GetIp()

	// Separa el domain y name
	name, domain := domainNameArray[0], domainNameArray[1]
	log.Printf("Solicitud Create -> Domain: %s, Name: %s", domain, name)

	dirPath := fmt.Sprintf("domains/%s/", domain)

	// Guarda el reloj vector en memoria
	var ok bool
	if _, ok = zfMap[domain]; ok {
		zfMap[domain].VectorClock[*serverNumber-1]++
	} else {
		os.RemoveAll(dirPath)
		zfMap[domain] = zfRegister{
			VectorClock: []int32{0, 0, 0},
			FileMutex:   &sync.Mutex{},
		}
	}

	// Crea el directorio del dominio
	// Crea archivos del registro ZF y log
	// Escribe en el registro ZF y en el log del dominio
	zfMap[domain].FileMutex.Lock()
	defer zfMap[domain].FileMutex.Unlock()
	writeFile(domain, name, ip, "append")

	// Envia respuesta a admin
	res := &dns.CreateResponse{
		VectorClock: zfMap[domain].VectorClock,
	}
	log.Printf("Respuesta Create -> VectorClock: %v", zfMap[domain].VectorClock)
	return res, nil
}

// Cambia el parametro de un nombre de dominio
func (s *dNSServiceServer) Update(ctx context.Context, req *dns.UpdateRequest) (*dns.UpdateResponse, error) {

	domainNameArray := strings.Split(req.GetDomainName(), ".")
	option := req.GetOption()
	parameter := req.GetParameter()

	// Separa el domain y name
	name, domain := domainNameArray[0], domainNameArray[1]
	log.Printf("Solicitud Update -> Domain: %s, Name: %s, Option: %t, Parameter: %s", domain, name, option, parameter)

	dirPath := fmt.Sprintf("domains/%s/", domain)

	// Crea el directorio del dominio
	// Crea archivos del registro ZF y log
	// Escribe en el registro ZF y en el log del dominio

	found := updateFile(domain, name, option, parameter, "update")
	if found == true {
		// Guarda el reloj vector en memoria
		var ok bool
		if _, ok = zfMap[domain]; ok {
			zfMap[domain].FileMutex.Lock()
			defer zfMap[domain].FileMutex.Unlock()
			zfMap[domain].VectorClock[*serverNumber-1]++
		} else {
			os.RemoveAll(dirPath)
			zfMap[domain] = zfRegister{
				VectorClock: []int32{0, 0, 0},
				FileMutex:   &sync.Mutex{},
			}
		}
	}

	// Envia respuesta a admin
	res := &dns.UpdateResponse{
		VectorClock: zfMap[domain].VectorClock,
	}
	log.Printf("Respuesta Update -> VectorClock: %v", zfMap)
	return res, nil
}

func (s *dNSServiceServer) Delete(ctx context.Context, req *dns.DeleteRequest) (*dns.DeleteResponse, error) {

	domainNameArray := strings.Split(req.GetDomainName(), ".")

	// Separa el domain y name
	name, domain := domainNameArray[0], domainNameArray[1]
	log.Printf("Solicitud Delete -> Domain: %s, Name: %s", domain, name)

	dirPath := fmt.Sprintf("domains/%s/", domain)

	// Crea el directorio del dominio
	// Crea archivos del registro ZF y log
	// Escribe en el registro ZF y en el log del dominio
	zfMap[domain].FileMutex.Lock()
	defer zfMap[domain].FileMutex.Unlock()
	found := deleteFile(domain, name, "delete")
	if found == true {
		// Guarda el reloj vector en memoria
		var ok bool
		if _, ok = zfMap[domain]; ok {
			zfMap[domain].VectorClock[*serverNumber-1]++
		} else {
			os.RemoveAll(dirPath)
			zfMap[domain] = zfRegister{
				VectorClock: []int32{0, 0, 0},
				FileMutex:   &sync.Mutex{},
			}
		}
	}

	// Envia respuesta a admin
	res := &dns.DeleteResponse{
		VectorClock: zfMap[domain].VectorClock,
	}
	log.Printf("Respuesta Delete -> VectorClock: %v", zfMap)
	return res, nil
}

// Crea el directorio del dominio, los archivos zf y log, y luego escribe en ellos
func writeFile(domain string, name string, ip string, action string) {
	dirPath := fmt.Sprintf("domains/%s/", domain)
	zfPath := dirPath + "zf.data"
	logPath := dirPath + "log.data"

	// Crea directorios y archivos
	os.MkdirAll(dirPath, 0700)
	os.OpenFile(zfPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	// Agrega el nuevo nombre al registro zf
	zfFile, err := os.OpenFile(zfPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer zfFile.Close()
	newZFLine := fmt.Sprintf("www.%s.%s IN A %s \n", name, domain, ip)
	if _, err := zfFile.WriteString(newZFLine); err != nil {
		log.Println(err)
	}

	// Escribe en el log
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer logFile.Close()
	newLogLine := fmt.Sprintf("%s %s.%s %s \n", action, name, domain, ip)
	if _, err := logFile.WriteString(newLogLine); err != nil {
		log.Println(err)
	}
}

// Crea el directorio del dominio, los archivos zf y log, y luego sobreescribe en ellos
func updateFile(domain string, name string, option bool, parameter string, action string) bool {
	var newLine string
	found := false
	dirPath := fmt.Sprintf("domains/%s/", domain)
	zfPath := dirPath + "zf.data"
	logPath := dirPath + "log.data"

	// Crea directorios y archivos
	os.MkdirAll(dirPath, 0700)
	os.OpenFile(zfPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	// Busca y reemplaza en el registro zf

	zfFile, err := ioutil.ReadFile(zfPath)
	if err != nil {
		log.Println(err)
	}
	lines := strings.Split(string(zfFile), "\n")

	for i, line := range lines {
		if strings.Contains(line, fmt.Sprintf("%s.%s", name, domain)) {
			if option == true {
				// Cambia nombre de dominio
				fmt.Println("Update -> option: IP")
				newLine = fmt.Sprintf("www.%s.%s IN A %s \n", name, domain, parameter)
			} else {
				// Cambia IP del nombre de dominio
				fmt.Println("Update -> option: domain_name")
				ip := strings.Split(line, " ")[3]
				newLine = fmt.Sprintf("www.%s.%s IN A %s \n", parameter, domain, ip)
			}
			lines[i] = newLine
			found = true
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(zfPath, []byte(output), 0644)
	if err != nil {
		log.Println(err)
	}

	if found == false {
		return false
	}

	// Escribe en el log
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer logFile.Close()
	newLogLine := fmt.Sprintf("%s %s.%s %s \n", action, name, domain, parameter)
	if _, err := logFile.WriteString(newLogLine); err != nil {
		log.Println(err)
	}

	return true
}

// Crea el directorio del dominio, los archivos zf y log, y luego sobreescribe en ellos
func deleteFile(domain string, name string, action string) bool {
	dirPath := fmt.Sprintf("domains/%s/", domain)
	zfPath := dirPath + "zf.data"
	logPath := dirPath + "log.data"
	found := false

	// Crea directorios y archivos
	os.MkdirAll(dirPath, 0700)
	os.OpenFile(zfPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	// Busca y reemplaza en el registro zf

	zfFile, err := ioutil.ReadFile(zfPath)
	if err != nil {
		log.Println(err)
	}
	lines := strings.Split(string(zfFile), "\n")

	for i, line := range lines {
		if strings.Contains(line, fmt.Sprintf("%s.%s", name, domain)) {
			lines[i] = ""
			found = true
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(zfPath, []byte(output), 0644)
	if err != nil {
		log.Println(err)
	}
	if found == false {
		return false
	}

	// Escribe en el log
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer logFile.Close()
	newLogLine := fmt.Sprintf("%s %s.%s \n", action, name, domain)
	if _, err := logFile.WriteString(newLogLine); err != nil {
		log.Println(err)
	}
	return true
}

func deleteLog(logPath string) {
	// Elimina el log
	os.Remove(logPath)
}

func main() {
	domainsPath := "domains/"
	os.RemoveAll(domainsPath)
	os.MkdirAll(domainsPath, 0700)

	fmt.Println("Iniciando servidor...")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *serverIP, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	dns.RegisterDNSServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
