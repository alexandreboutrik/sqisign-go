package main

import (
	"encoding/hex"
	"sqisign-go/sqisign"
	"unsafe"
)

//#include <stdlib.h>
import "C"

func main() {
	message := "message to be signed"

	pk, sk, _ := sqisign.GenerateKey(sqisign.LV3)
	defer C.free(unsafe.Pointer(sk.CSecretKey))
	defer C.free(unsafe.Pointer(pk.CPublicKey))
	println("=================SECRET KEY===================")
	sqisign.PrintHex(sk.CSecretKey, sk.CRYPTO_SECRETKEYBYTES)
	println("=================PUBLIC KEY===================")
	sqisign.PrintHex(pk.CPublicKey, pk.CRYPTO_PUBLICKEYBYTES)
	sig, _ := sk.Sign(nil, []byte(message), nil)
	println("=================SIGNATURE===================")
	println(hex.EncodeToString(sig))
	msg, _ := pk.Verify(len(message), sig)
	println("=================VERIFY===================")
	println(hex.EncodeToString(msg))
	println("=================PUBLIC===================")
	sqisign.PrintHex(sk.Public().CPublicKey, pk.CRYPTO_PUBLICKEYBYTES)
}
