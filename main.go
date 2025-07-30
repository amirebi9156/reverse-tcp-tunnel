// فایل main.go — اجرای یک ریورس تونل ساده TCP در Go با منوی تعاملی

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"reverse/client"
	"reverse/server"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(`
  **** ushkayanet utunnel Reverse tunnel management console  ****

  1) Server mode
  2) Client mode
  3) Tunnel status (غیرفعال)
  4) Restart service (غیرفعال)
  5) Stop service (غیرفعال)
  6) Delete service (غیرفعال)
  7) Reset timer (غیرفعال)
  8) Exit
`)
	fmt.Print("Enter your choice (1-8): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("[+] اجرا به عنوان سرور (در ایران)")
		err := server.Start("0.0.0.0:9000")
		if err != nil {
			log.Fatal("خطا در اجرای سرور:", err)
		}
	case "2":
		fmt.Print("لطفاً IP سرور ایران را وارد کنید (مثلاً 1.2.3.4:9000): ")
		ip, _ := reader.ReadString('\n')
		ip = strings.TrimSpace(ip)
		fmt.Println("[+] اجرا به عنوان کلاینت (در خارج)")
		err := client.Start(ip)
		if err != nil {
			log.Fatal("خطا در اجرای کلاینت:", err)
		}
	case "8":
		fmt.Println("خروج از برنامه.")
		os.Exit(0)
	default:
		fmt.Println("گزینه نامعتبر است. لطفاً یکی از اعداد 1 تا 8 را وارد کنید.")
		os.Exit(1)
	}
}
