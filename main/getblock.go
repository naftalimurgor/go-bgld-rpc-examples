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

	block, err := bc.GetBlock("00000000000000003f8d1861d035e44d4297c49bd2517dc0a44ad73c7091926c")
	log.Println(err, block)
}