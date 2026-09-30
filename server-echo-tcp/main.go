package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("Gagal membuka port 9001:", err)
	}

	defer ln.Close()
	log.Println("Server Echo mendengarkan di port 9001")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Gagal menerima sambungan:", err)
			continue
		}
		go layani(conn)
	}
}

func layani(conn net.Conn) {
	defer conn.Close()
	alamat := conn.RemoteAddr().String()
	log.Println("Klien tersambung:", alamat)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		baris := scanner.Text()
		log.Printf("dari %s: %q", alamat, baris)
		if _, err := fmt.Fprintf(conn, "%s\n", baris); err != nil {
			log.Println("Gagal Mengirim ke", alamat, err)
			return
		}
	}
	log.Println("Klien terputus:", alamat)
}
