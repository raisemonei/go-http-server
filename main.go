package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// dsn := os.Getenv("DB_DSN")
	// // dsn := "host=postgres user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Connect db success")
	// }
	server := http.NewServeMux()
	server.HandleFunc("/health_check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("health check running %s\n", time.Now().Format(time.RFC3339))
		w.Write([]byte("Health Check OK!!!"))
	})
	fmt.Println("Application running port: 8080")
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		panic(err)
	}
}
