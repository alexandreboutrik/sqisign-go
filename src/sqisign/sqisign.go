package sqisign

import (
	"crypto"
	"fmt"
	"io"
	"strings"
	"unsafe"
)

/*
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/include/
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/protocols/ref/lvl1
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/gf/ref/lvl1
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/id2iso/ref/lvl1
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/ec/ref/lvl1
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/klpt/ref/lvl1
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/precomp/ref/lvl1

#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/protocols/ref/lvl3
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/gf/ref/lvl3
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/id2iso/ref/lvl3
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/ec/ref/lvl3
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/klpt/ref/lvl3
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/precomp/ref/lvl3

#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/protocols/ref/lvl5
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/gf/ref/lvl5
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/id2iso/ref/lvl5
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/ec/ref/lvl5
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/klpt/ref/lvl5
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/precomp/ref/lvl5

#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/quaternion/ref/generic
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/common/generic
#cgo LDFLAGS: -L/home/boutrik/.local/src/projetos/bbpq/the-sqisign/build/src/intbig/ref/generic


#cgo lvl1 LDFLAGS: -lsqisign_lvl1_nistapi -lsqisign_lvl1 -lsqisign_protocols_lvl1
#cgo lvl1 LDFLAGS: -lsqisign_gf_lvl1 -lsqisign_id2iso_lvl1 -lsqisign_ec_lvl1
#cgo lvl1 LDFLAGS: -lsqisign_klpt_lvl1 -lsqisign_precomp_lvl1

#cgo lvl3 LDFLAGS: -lsqisign_lvl3_nistapi -lsqisign_lvl3 -lsqisign_protocols_lvl3
#cgo lvl3 LDFLAGS: -lsqisign_gf_lvl3 -lsqisign_id2iso_lvl3 -lsqisign_ec_lvl3
#cgo lvl3 LDFLAGS: -lsqisign_klpt_lvl3 -lsqisign_precomp_lvl3

#cgo lvl5 LDFLAGS: -lsqisign_lvl5_nistapi -lsqisign_lvl5 -lsqisign_protocols_lvl5
#cgo lvl5 LDFLAGS: -lsqisign_gf_lvl5 -lsqisign_id2iso_lvl5 -lsqisign_ec_lvl5
#cgo lvl5 LDFLAGS: -lsqisign_klpt_lvl5 -lsqisign_precomp_lvl5

#cgo LDFLAGS: -lsqisign_quaternion_generic -lsqisign_common_sys
#cgo LDFLAGS: -lgmp -lm

#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/src/precomp/ref/lvl5/include
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/src/verification/ref/include
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/src/nistapi/lvl5/
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/src/ec/ref/include
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/src/gf/ref/include
#define RADIX_64
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/src/common/generic/include
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/src/signature/ref/include
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/src/quaternion/ref/generic/include
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/src/mp/ref/generic/include
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/src/hd/ref/include
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/src/id2iso/ref/include
#cgo CFLAGS: -I /home/boutrik/.local/src/projetos/bbpq/the-sqisign/src/gf/broadwell/include

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

#define ENABLE_SIGN
#include "src/verification/ref/lvlx/verify.c"
#include "src/verification/ref/lvlx/encode_verification.c"
#include "src/signature/ref/lvlx/keygen.c"
#include "src/precomp/ref/lvl5/quaternion_data.c"
#include "src/ec/ref/lvlx/ec.c"
#include "src/gf/broadwell/lvlx/fp2.c"
#include "src/nistapi/lvl5/api.c"
#include "src/sqisign.c"

// C function to print hex values
static void print_hex(const unsigned char *hex, int len) {
    for (int i = 0; i < len; ++i) {
        printf("%02x", hex[i]);
    }
    printf("\n");
}
*/
import "C"

const (
	LV1 int = 1
	LV3 int = 3
	LV5 int = 5
)

//var CRYPTO_SECRETKEYBYTES int = C.CRYPTO_SECRETKEYBYTES
//var CRYPTO_PUBLICKEYBYTES int = C.CRYPTO_PUBLICKEYBYTES
//var CRYPTO_BYTES int = C.CRYPTO_BYTES
//var CRYPTO_ALGNAME string = C.CRYPTO_ALGNAME

//var C_CRYPTO_SECRETKEYBYTES C.int = C.CRYPTO_SECRETKEYBYTES
//var C_CRYPTO_PUBLICKEYBYTES C.int = C.CRYPTO_PUBLICKEYBYTES
//var C_CRYPTO_BYTES C.int = C.CRYPTO_BYTES

func CryptoSignKeyPair(pk *C.uchar, sk *C.uchar) int {
	return int(C.crypto_sign_keypair(pk, sk))
}

func CryptoSign(sm *C.uchar, smlen *C.ulonglong, m *C.uchar,
	mlen C.ulonglong, sk *C.uchar) int {
	return int(C.crypto_sign(sm, smlen, m, mlen, sk))
}

func CryptoSignOpen(m *C.uchar, mlen *C.ulonglong, sm *C.uchar,
	smlen C.ulonglong, pk *C.uchar) int {
	return int(C.crypto_sign_open(m, mlen, sm, smlen, pk))
}

type PublicKey struct {
	CRYPTO_PUBLICKEYBYTES int
	CPublicKey            *C.uchar
}

type PrivateKey struct {
	CRYPTO_SECRETKEYBYTES int
	CRYPTO_BYTES          int
	CSecretKey            *C.uchar
}

var public_key *PublicKey
var private_key *PrivateKey

func GenerateKey(lvl int) (pk *PublicKey, sk *PrivateKey, err error) {
	var cskb int
	var cpkb int
	var cb int
	switch lvl {
	case LV1:
		cskb = 353
		cpkb = 65
		cb = 148
	case LV3:
		cskb = 529
		cpkb = 97
		cb = 224
	case LV5:
		cskb = 701
		cpkb = 129
		cb = 292
	}
	pk_c := (*C.uchar)(unsafe.Pointer(C.CString(strings.Repeat("0", cpkb))))
	sk_c := (*C.uchar)(unsafe.Pointer(C.CString(strings.Repeat("0", cskb))))
	ok := CryptoSignKeyPair(pk_c, sk_c)
	public_key = &PublicKey{CPublicKey: pk_c, CRYPTO_PUBLICKEYBYTES: cpkb}
	private_key = &PrivateKey{CSecretKey: sk_c, CRYPTO_SECRETKEYBYTES: cskb, CRYPTO_BYTES: cb}
	if err = nil; ok != 0 {
		err = fmt.Errorf("error during key generation process")
	}
	return public_key, private_key, err
}

func (priv *PrivateKey) Public() PublicKey {
	return *public_key
}

func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	m := C.CString(string(digest[:]))
	defer C.free(unsafe.Pointer(m))
	mlen := len(digest)
	sm := C.CString(strings.Repeat("0", (priv.CRYPTO_BYTES)+mlen))
	defer C.free(unsafe.Pointer(sm))
	smlen := priv.CRYPTO_BYTES + mlen
	ok := CryptoSign((*C.uchar)(unsafe.Pointer(sm)), (*C.ulonglong)(unsafe.Pointer(&smlen)),
		(*C.uchar)(unsafe.Pointer(m)), (C.ulonglong)(mlen), priv.CSecretKey)
	signature = C.GoBytes((unsafe.Pointer(sm)), (C.int)((priv.CRYPTO_BYTES)+mlen))
	if err = nil; ok != 0 {
		err = fmt.Errorf("error during signing process")
	}
	return
}

func (pub *PublicKey) Verify(mlen int, signature []byte) (msg []byte, err error) {
	m := C.CString(strings.Repeat("0", mlen))
	defer C.free(unsafe.Pointer(m))
	smlen := len(signature)
	sm := C.CString(string(signature))
	defer C.free(unsafe.Pointer(sm))
	ok := CryptoSignOpen((*C.uchar)(unsafe.Pointer(m)), (*C.ulonglong)(unsafe.Pointer(&mlen)),
		(*C.uchar)(unsafe.Pointer(sm)), (C.ulonglong)(smlen), pub.CPublicKey)
	msg = []byte(C.GoString(m))
	if err = nil; ok != 0 {
		err = fmt.Errorf("error during verification process")
	}
	return
}

func PrintHex(hex *C.uchar, length int) {
	C.print_hex(hex, C.int(length))
}
