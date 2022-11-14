package main

import (
	"container/list"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
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

func decimalToBinary(data []int) []int {
	binary_value := []int{}
	for _, ascii := range data {
		remainder_array := []int{}
		for ascii >= 0 {
			numerator, remainder := ascii/2, ascii%2
			remainder_array = append(remainder_array, remainder)
			if ascii == 0 {
				break
			}
			ascii = numerator
		}
		for i := len(remainder_array) - 1; i >= 0; i-- {
			binary_value = append(binary_value, remainder_array[i])
		}
	}
	return binary_value
}

func hexToBinary(data string) []int {
	decimal := 0
	hex_numbers := make(map[string]int)
	length_of_hex := len(data) - 1
	hex_numbers["0"], hex_numbers["1"], hex_numbers["2"], hex_numbers["3"], hex_numbers["4"], hex_numbers["5"], hex_numbers["6"], hex_numbers["7"], hex_numbers["8"], hex_numbers["9"], hex_numbers["a"], hex_numbers["b"], hex_numbers["c"], hex_numbers["d"], hex_numbers["e"], hex_numbers["f"] = 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
	for i, chrac_bit := range data {
		hexmultiplier := math.Pow(16, float64(length_of_hex-i))
		decimal += hex_numbers[string(chrac_bit)] * int(hexmultiplier)
	}
	decimal_value := []int{decimal}
	binary := decimalToBinary(decimal_value)
	return binary
}

func main() {
	fmt.Println("Proof of Work : Blockchain")
	block_data := []string{"My Name is Suraj Joshi", "Learning BloakChain", "With Golang"}
	time_now := ""
	time_hash := ""
	new_hash_string := ""
	new_hash_binary := []int{}
	blocks := list.New()
	nounce := 0
	difficulty := 1
	genesis_hash := string(hashGenerator(block_data[0], time_now, strconv.Itoa(nounce), strconv.Itoa(difficulty)))
	genesis_last_hash := strings.Repeat("0", 64)
	genesisBlock := addBlock{timestamp: time_now, data: block_data[0], hash: genesis_hash, lastHash: genesis_last_hash, nounce: strconv.Itoa(nounce), difficulty: strconv.Itoa(difficulty)}
	blocks.PushBack(genesisBlock)
	for i := 1; i < len(block_data); i++ {
		time_now := time.Now().String()
		prevHash := blocks.Back()
		prevHash_value := addBlock(prevHash.Value.(addBlock))
		new_hash := string(hashGenerator(block_data[i], time_now, strconv.Itoa(nounce), strconv.Itoa(difficulty)))
		new_hash_binary = hexToBinary(new_hash)
		for _, i := range new_hash_binary {
			new_hash_string = new_hash_string + strconv.Itoa(i)
		}
		fmt.Println("outside loop ", new_hash_string)
		for strings.Index(new_hash_string, strings.Repeat("0", difficulty)) != 0 {
			time_hash = time.Now().String()
			new_hash = string(hashGenerator(block_data[i], time_hash, strconv.Itoa(nounce), strconv.Itoa(difficulty)))
			new_hash_binary = hexToBinary(new_hash)
			for _, i := range new_hash_binary {
				new_hash_string = new_hash_string + strconv.Itoa(i)
			}
			fmt.Println("inside loop", new_hash_string)
			nounce += 1
		}
		fmt.Println("Iterations Taken to add new Block: ", nounce)
		fmt.Println("New Block Hash ---->  ", new_hash_binary)
		new_block := addBlock{timestamp: time_hash, data: block_data[i], hash: new_hash, lastHash: prevHash_value.hash, nounce: strconv.Itoa(nounce), difficulty: strconv.Itoa(difficulty)}
		blocks.PushBack(new_block)

	}
	for i := blocks.Front(); i != nil; i = i.Next() {
		Blockchain := addBlock(i.Value.(addBlock))
		fmt.Println(Blockchain)
	}
}
