package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"strings"
)

func run(cmd string) string {
	words := strings.Split(cmd, " ")
	path := words[0]
	args := words[1:]

	out, err := exec.Command(path, args...).CombinedOutput()

	if err != nil {
		log.Println(err)
	}

	return string(out)
}

func handleConnection(conn net.Conn) {
	r := bufio.NewReader(conn)
	cmd, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	output := run(strings.TrimSpace(string(cmd)))
	io.Copy(conn, strings.NewReader(output))
	conn.Close()
}

func runServer() {
	ln, err := net.Listen("tcp", ":9178")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server listening")

	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func runClient() {
	conn, err := net.Dial("tcp", "localhost:9178")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s\n", strings.Join(flag.Args()[1:], " "))

	var buf bytes.Buffer
	io.Copy(&buf, conn)
	fmt.Printf("Read %d bytes\n", buf.Len())
	fmt.Println(buf.String())
}

func main() {

	flag.Parse()

	switch flag.Arg(0) {
	case "server":
		runServer()
	case "client":
		runClient()
	default:
		log.Fatal("no")
	}
}
