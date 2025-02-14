package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// creating the JSON struct
type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type RPCResponse struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
}

func main() {

	// type the node URL here
	nodeURL := ""

	rpcReq := RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		ID:      1,
	}

	// transforming struc into bytes
	requestbody, err := json.Marshal(rpcReq)
	if err != nil {
		log.Fatal("Fail to marshal JSON-RPC request", err)
	}

	// sending the post request
	resp, err := http.Post(nodeURL, "application/json", bytes.NewBuffer(requestbody))
	if err != nil {
		log.Fatal("Fail to send the request", err)
	}
	defer resp.Body.Close()

	// reading the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response", err)
	}

	// creating the data variable as type RPCResponse
	var data RPCResponse

	//Extracing the JSON result in data struct
	err = json.Unmarshal(body, &data)

	//Extract and process the "result" field
	hexBlockNumber := data.Result
	cleanHex := strings.TrimPrefix(hexBlockNumber, "0x")

	//Printing the block number in decimals
	blockNumer, err := strconv.ParseUint(cleanHex, 16, 64)
	fmt.Println("The block number is:", blockNumer)

}
