package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

func main() {

	// type the node URL here
	nodeURL := "<type_node_url>"

	rpcReq := RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		ID:      1,
	}

	//transforming struc into bytes
	requestbody, err := json.Marshal(rpcReq)
	if err != nil {
		log.Fatal("Fail to marshal JSON-RPC request", err)
	}

	//implementing io.Reader interface
	resp, err := http.Post(nodeURL, "application/json", bytes.NewBuffer(requestbody))
	if err != nil {
		log.Fatal("Fail to send the request", err)
	}
	defer resp.Body.Close()

	//reading the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response", err)
	}

	//print the response
	fmt.Println(string(body))

}
