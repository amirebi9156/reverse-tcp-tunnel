package client

import (
	"fmt"
	"net"
	"time"
)

func Start(address string) error {
	fmt.Println("[+] تلاش برای اتصال به سرور...", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("خطا در اتصال به سرور: %v", err)
	}
	defer conn.Close()

	fmt.Println("[✓] اتصال برقرار شد با سرور", address)

	message := "سلام از سمت کلاینت!\n"
	_, err = conn.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("خطا در ارسال پیام: %v", err)
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("خطا در دریافت پاسخ: %v", err)
	}

	echo := string(buffer[:n])
	fmt.Println("📥 پاسخ از سرور:", echo)

	time.Sleep(5 * time.Second)

	return nil
}
