package main

import (
	"fmt"
	"net/http"
)

func createWallet(resposne http.ResponseWriter, request *http.Request) {
	fmt.Println("create wallet")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /wallets", createWallet)

	http.ListenAndServe(":8080", mux)
}
