package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var err error

func main() {

	SetupDB()

	http.HandleFunc("/", index)
	http.HandleFunc("/honeypot-servers", listHoneypotServers)
	http.HandleFunc("/join", join)
	http.HandleFunc("/api/private/ssh", ssh)
	http.HandleFunc("/api/ssh_logins.json", sshLoginList)
	fmt.Println("listening...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
