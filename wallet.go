package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	Private *ecdsa.PrivateKey
	PubKey []byte
}

func NewWallet()*Wallet  {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	pubKeyOrig:=privateKey.PublicKey
	pubKey:=append(pubKeyOrig.X.Bytes(),pubKeyOrig.Y.Bytes()...)
	return &Wallet{Private:privateKey,PubKey:pubKey}
}

func (w *Wallet)NewAddress() string {
	pubKey := w.PubKey
	rip160HashValue := HashPubKey(pubKey)
	version:=byte(00)
	payload := append([]byte{version}, rip160HashValue...)
	checkCode := CheckSum(payload)
	payload = append(payload, checkCode...)
	address := base58.Encode(payload)
	return address
}

func HashPubKey(data []byte) []byte {
	hash := sha256.Sum256(data)
	rip160hasher := ripemd160.New()
	_, err := rip160hasher.Write(hash[:])
	if err != nil {
		panic(err)
	}
	rip160HashValue := rip160hasher.Sum(nil)
	return rip160HashValue
}

func CheckSum(data []byte) []byte {
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	checkCode := hash2[:4]
	return checkCode
}
