package main

import (
	"github.com/naftalimurgor/go-bgld"
	"log"
)

const (
	SERVER_HOST        = ""
	SERVER_PORT        = 8334
	USER               = "localuser"
	PASSWD             = "3xaO3o/i-]G4"
	USESSL             = false
	WALLET_PASSPHRASE  = "p1"
	WALLET_PASSPHRASE2 = "p2"
)

func main() {
	bc, err := bgld.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	if err != nil {
		log.Fatalln(err)
	}

	transaction, err := bc.GetTransaction("a1b7093d041bc1b763ba1ad894d2bd5376b38e6c7369613684e7140e8d9f7515")
	log.Println(err, transaction)

}