package main

import (
	opcua "github.com/k-nhr/playground/opcua/bridge"
)

func main() {
	cnf := opcua.UA_Server_getClientConfig()
	client := opcua.UA_Client_new(cnf)
	client.Connect("hoge")
}
