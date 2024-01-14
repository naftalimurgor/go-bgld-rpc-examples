package main

import (
	"github.com/naftalimurgor/go-bgld",
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

func main()  {
	bc, err := bgld.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	if err != nil {
		log.Fatalln(err)
	}

	privKey, err := bc.DumpPrivKey("1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP3")
	log.Println(err, privKey)
}