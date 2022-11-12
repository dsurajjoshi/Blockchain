package main

import (
	"container/list"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type addBlock struct {
	timestamp  string
	data       string
	hash       string
	lastHash   string
	nounce     string
	difficulty string
}

func hashGenerator(data, timestamp, nounce, difficulty string) string {
	h := sha256.New()
	h.Write([]byte(string(data + timestamp + nounce + difficulty)))
	created_hash := hex.EncodeToString(h.Sum(nil))
	return created_hash
}

func main() {
	fmt.Println("My First Blockchain Program")
	block_data := []string{"My Name is Suraj Joshi", "Learning BloakChain", "With Golang"}
	time := time.Now().String()
	blocks := list.New()
	nounce := 0
	difficulty := 3
	genesis_hash := string(hashGenerator(block_data[0], time, strconv.Itoa(nounce), strconv.Itoa(difficulty)))
	genesis_last_hash := strings.Repeat("0", 64)
	genesisBlock := addBlock{timestamp: time, data: block_data[0], hash: genesis_hash, lastHash: genesis_last_hash, nounce: strconv.Itoa(nounce), difficulty: strconv.Itoa(difficulty)}
	blocks.PushBack(genesisBlock)

	for i := 1; i < len(block_data); i++ {
		prevHash := blocks.Back()
		prevHash_value := addBlock(prevHash.Value.(addBlock))
		new_hash := string(hashGenerator(block_data[i], time, strconv.Itoa(nounce), strconv.Itoa(difficulty)))
		for strings.Index(new_hash, strings.Repeat("0", difficulty)) != 0 {
			nounce += 1
			new_hash = string(hashGenerator(block_data[i], time, strconv.Itoa(nounce), strconv.Itoa(difficulty)))
			fmt.Println("count: 3: ", new_hash)
		}
		fmt.Println("New Block")
		new_block := addBlock{timestamp: time, data: block_data[i], hash: new_hash, lastHash: prevHash_value.hash, nounce: strconv.Itoa(nounce), difficulty: strconv.Itoa(difficulty)}
		blocks.PushBack(new_block)

	}
	for i := blocks.Front(); i != nil; i = i.Next() {
		Blockchain := addBlock(i.Value.(addBlock))
		fmt.Println(Blockchain.timestamp, Blockchain.data, Blockchain.hash, Blockchain.lastHash)
	}
}
