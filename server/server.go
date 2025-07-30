package server

import (
	"fmt"
	"io"
	"net"
)

func Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("خطا در گوش دادن روی %s: %v", address, err)
	}
	defer listener.Close()

	fmt.Println("[+] سرور آماده است و در حال گوش دادن روی", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[-] خطا در پذیرش اتصال:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("[+] اتصال جدید از", conn.RemoteAddr())

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("[-] خطا در خواندن:", err)
			}
			break
		}

		data := buf[:n]
		fmt.Printf("📥 دریافت شده: %s\n", string(data))

		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("[-] خطا در نوشتن:", err)
			break
		}
	}
}
