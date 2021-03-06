package server

import (
	"blockchain_demo/utils"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type addr struct {
	AddrList []string
}

func sendAddr(address string) {
	nodes := addr{
		AddrList: knownNodes,
	}
	nodes.AddrList = append(nodes.AddrList, nodeAddress)
	payload := utils.GobEncode(nodes)
	request := append(commandToBytes("addr"), payload...)
	sendData(address, request)
}

func handleAddr(request []byte) {
	var buff bytes.Buffer
	var payload addr

	buff.Write(extractCommand(request))
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	knownNodes = append(knownNodes, payload.AddrList...)
	fmt.Printf("There are %d known nodes now!\n", len(knownNodes))
	requestBlocks()
}

func requestBlocks() {
	for _, node := range knownNodes {
		sendGetBlocks(node)
	}
}
