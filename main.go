package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type Server struct {
	Addr string
	Ln   net.Listener
}

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
	}
}

func (s *Server) Listen() error {
	ln, err := net.Listen("tcp", s.Addr)
	s.Ln = ln
	return err
}

func main() {
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Println("Error listening to port: 2000, Error: ", err.Error())
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error: %s while connecting to addr: %s\n",
				err.Error(), conn.RemoteAddr())
		} else {
			log.Printf("Connected succesfully to remote addr: %s\n",
				conn.RemoteAddr())
		}

		go handleIncoming(conn)
		go handleOutgoing(conn)
	}
}

func handleIncoming(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte(
		"Your connection to the our server is succesful!\n"))
	for {
		buf := make([]byte, 64)

		if err != nil {
			log.Printf("Writing failed")
		}

		n, err := conn.Read(buf)
		if err != nil {
			log.Fatalf("Connection lost!")
		}

		fmt.Printf("client:%s\n", string(buf[:n]))
	}
}

func handleOutgoing(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error while reading line\n")
		}

		_, err = conn.Write(fmt.Appendf(nil, "server: %s\n", line))
		if err != nil {
			log.Printf("Connection to addr: %s closed.\n", conn.RemoteAddr())
		}
	}
}
