package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var difficultyLevel = 3

// Block struct to map a block
type Block struct {
	Data      string `json:"data"`
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	PrevHash  string `json:"prev_hash"`
	CurHash   string `json:"cur_hash"`
	Nounce    string `json:"nounce"`
}

//CreateBlock create a new block to be add to blockchain
func CreateBlock(chain []Block, data string) Block {

	block := Block{}
	block.Data = data
	block.Index = len(chain)
	block.Timestamp = time.Now().UTC().Format(time.RFC3339Nano)
	block.PrevHash = func() string {
		if len(chain) != 0 {
			return chain[len(chain)-1].CurHash
		} else {
			return "null"
		}
	}()

	block.CurHash, block.Nounce = miner(block.Data + strconv.Itoa(block.Index) + block.Timestamp + block.PrevHash)

	return block
}

func miner(dataString string) (string, string) {

	for true {
		nounce := strconv.Itoa(rand.Intn(1e10))
		hash := sha256.Sum256([]byte(dataString + nounce))
		hash_str := base64.StdEncoding.EncodeToString(hash[:])

		if hash_str[:difficultyLevel] == strings.Repeat("0", difficultyLevel) {
			return hash_str, nounce
		}
	}

	return "-", "-"

}

// addBlock add a new block to blockchain
func addBlock(chain []Block, data string) []Block {

	chain = append(chain, []Block{CreateBlock(chain, data)}...)
	fmt.Println(chain)
	return chain
}

func init() {

	rand.Seed(time.Now().UnixNano())
}

func main() {

	blockchain := []Block{}

	if len(blockchain) == 0 {

		blockchain = addBlock(blockchain, "GENESIS BLOCK")
	}

	scan := bufio.NewScanner(os.Stdin)

	for true {

		fmt.Printf("Please enter a new transaction ")
		scan.Scan()
		blockchain = addBlock(blockchain, scan.Text())
	}
}
