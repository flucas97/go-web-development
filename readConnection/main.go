package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// 1 - listen
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		// 2 - accept that connection
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn) // goroutine
	}
}

func handle(conn net.Conn) {
	// 3 - write in connection
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { // is there anything?
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close() // never closes
}
