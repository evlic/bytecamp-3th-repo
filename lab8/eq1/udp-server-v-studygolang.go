package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	host = *flag.String("host", "", "udp serve listen host")
	port = *flag.String("port", "7777", "udp server port")
)

// server, err := net.Listen("udp", address)
// if err != nil {
// 	log.Fatalln("run server err!!", err)
// }
// log.Println("run server on >> '", address, "'")
//
// for {
// 	conn, err := server.Accept()
// 	if err != nil {
// 		log.Fatalln("accept connection err!!", err)
// 	}
// 	go handler(conn)
// }

func main() {
	flag.Parse()
	address := fmt.Sprintf("%s:%s", host, port)
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalln("can't init udp address >>", err)
	}

	serve, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln("can't init udp serve >>", err)
	}
	defer func(serve *net.UDPConn) {
		_ = serve.Close()
	}(serve)
	log.Println("serve run!! list udp connection on >> ", address)

	for {
		handler(serve)
	}
}

func handler(conn net.Conn) {
	log.Fatalln("get connection and do nothing")
}
