package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	alamat := "localhost:9001"
	if len(os.Args) > 1 {
		alamat = os.Args[1]
	}

	conn, err := net.Dial("tcp", alamat)
	if err != nil {
		fmt.Println("Gagal tersambung ke", alamat, ":", err)
		return
	}
	defer conn.Close()
	fmt.Println("tersambung ke", alamat)
	fmt.Println("ketik pesan lalu enter. Ketik KELUAR untuk berhenti")

	keyboard := bufio.NewScanner(os.Stdin)
	dariServer := bufio.NewScanner(conn)
	for {
		fmt.Print("> ")
		if !keyboard.Scan() {
			return
		}
		baris := strings.TrimSpace(keyboard.Text())
		if strings.ToUpper(baris) == "KELUAR" {
			return
		}
		if _, err := fmt.Fprintf(conn, "%s\n", baris); err != nil {
			fmt.Println("Gagal mengirim:", err)
			return
		}
		if !dariServer.Scan() {
			fmt.Println("Server menutup Sambungan")
			return
		}
		fmt.Println("Balasan:", dariServer.Text())
	}
}
