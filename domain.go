package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"net"
)

func main() {

	// init
	server := "ie.whois-servers.net:43"
	search := "cainteach"

	// connect
	conn, err := net.Dial("tcp", server)

	// connection error
	if err != nil {
		log.Fatal(err)
	}

	// make the connection
	fmt.Fprintf(conn, "%s.ie\r\n", search)
	scanner := bufio.NewScanner(conn)

	// response
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] != '%' {
			fields := strings.Split(line, ": ") // k,v split
			fmt.Println(fields)
		}
	}

	// response error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
