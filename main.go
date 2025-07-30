package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"reverse/client"
	"reverse/server"
)

func main() {
	role := flag.String("role", "", "نقش این برنامه: client یا server")
	address := flag.String("address", "0.0.0.0:9000", "آدرس اتصال")
	flag.Parse()

	if *role == "server" {
		fmt.Println("[+] اجرا به عنوان سرور (در ایران)")
		err := server.Start(*address)
		if err != nil {
			log.Fatal("خطا در اجرای سرور:", err)
		}
	} else if *role == "client" {
		fmt.Println("[+] اجرا به عنوان کلاینت (در خارج)")
		err := client.Start(*address)
		if err != nil {
			log.Fatal("خطا در اجرای کلاینت:", err)
		}
	} else {
		fmt.Println("لطفاً با پارامتر --role=client یا --role=server اجرا کنید")
		os.Exit(1)
	}
}
