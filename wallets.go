package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
)

const walletFile = "wallet.dat"

type Wallets struct {
	WalletsMap map[string]*Wallet
}

func NewWallets() *Wallets {
	var ws Wallets
	ws.WalletsMap = make(map[string]*Wallet)
	ws.loadFile()
	return &ws
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address:=wallet.NewAddress()
	ws.WalletsMap[address]=wallet
	ws.saveToFile()
	return address
}

func (ws *Wallets)saveToFile()  {
	var buffer bytes.Buffer
	//gob.Register()
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(ws)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(walletFile,buffer.Bytes(),0600)
}

func (ws *Wallets) loadFile()  {
	_, err := os.Stat(walletFile)
	if os.IsNotExist(err){
		fmt.Println("暂无地址,请添加后再查看")
		return
	}
	content, err := ioutil.ReadFile(walletFile)
	if err != nil {
		panic(err)
	}
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(content))
	var wsLocal Wallets
	err = decoder.Decode(&wsLocal)
	if err != nil {
		panic(err)
	}
	ws.WalletsMap = wsLocal.WalletsMap
}

func (ws *Wallets)ListAllAddresses()[]string  {
	var addresses []string
	for address := range ws.WalletsMap {
		addresses = append(addresses, address)
	}
	return addresses
}
