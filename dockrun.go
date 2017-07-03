package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

const progname = "dockrun"
const version = "0.1.1"

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
	fmt.Print(buf.String())
}

func printHelpAndExit() {
	fmt.Println("Usage:")
	fmt.Printf("  %s server\n", progname)
	fmt.Printf("  %s client <command>\n", progname)
	fmt.Println("")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	var showHelp, showVersion bool

	flag.BoolVar(&showHelp, "help", false, "show help and exit")
	flag.BoolVar(&showVersion, "version", false, "show version and exit")

	flag.Parse()

	if showHelp {
		printHelpAndExit()
	}

	if showVersion {
		fmt.Printf("%s %s\n", progname, version)
		os.Exit(0)
	}

	switch flag.Arg(0) {
	case "server":
		runServer()
	case "client":
		runClient()
	default:
		printHelpAndExit()
	}
}
