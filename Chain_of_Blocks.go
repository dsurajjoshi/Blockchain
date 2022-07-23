package main

import (
	"container/list"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type addBlock struct {
	timestamp string
	data      string
	hash      string
	lastHash  string
}

func hashGenerator(data, timestamp string) string {
	h := sha256.New()
	h.Write([]byte(string(data + timestamp)))
	created_hash := hex.EncodeToString(h.Sum(nil))
	return created_hash
}

func main() {
	fmt.Println("My First Blockchain Program")
	block_data := []string{"My Name is Suraj Joshi", "Learning BloakChain", "With Golang"}
	time := time.Now().String()
	blocks := list.New()
	genesis_hash := string(hashGenerator(block_data[0], time))
	genesis_last_hash := strings.Repeat("0", 64)
	genesisBlock := addBlock{timestamp: time, data: block_data[0], hash: genesis_hash, lastHash: genesis_last_hash}
	blocks.PushBack(genesisBlock)

	for i := 1; i < len(block_data); i++ {
		new_hash := string(hashGenerator(block_data[i], time))
		prevHash := blocks.Back()
		prevHash_value := addBlock(prevHash.Value.(addBlock))

		new_block := addBlock{timestamp: time, data: block_data[i], hash: new_hash, lastHash: prevHash_value.hash}
		blocks.PushBack(new_block)
	}
	for i := blocks.Front(); i != nil; i = i.Next() {
		Blockchain := addBlock(i.Value.(addBlock))
		fmt.Println(Blockchain.timestamp, Blockchain.data, Blockchain.hash, Blockchain.lastHash)
	}
}
