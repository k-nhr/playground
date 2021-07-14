package opcua

/*
#cgo windows LDFLAGS: -lws2_32
#include "open62541.h"
*/
import "C"
import "unsafe"

type UA_Client struct {
	client C.UA_Client
}

func UA_Server_getClientConfig() C.UA_ClientConfig {
	return C.UA_Server_getClientConfig()
}

func UA_Client_new(config C.UA_ClientConfig) *UA_Client {
	cnf := *(*C.UA_ClientConfig)(unsafe.Pointer(&config))
	client := C.UA_Client_new(cnf)
	return (*UA_Client)(unsafe.Pointer(client))
}

func (c *UA_Client) Connect(endpointUrl string) (s C.UA_StatusCode) {
	endpoint := C.CString(endpointUrl)
	defer C.free(unsafe.Pointer(endpoint))

	client := (*C.UA_Client)(unsafe.Pointer(c))
	stat := C.UA_Client_connect(client, endpoint)
	s = *(*C.UA_StatusCode)(unsafe.Pointer(&stat))
	return
}

func (c *UA_Client) disconnect() {
	client := (*C.UA_Client)(unsafe.Pointer(c))
	C.UA_Client_disconnect(client)
}

func (c *UA_Client) Reset() {
	client := (*C.UA_Client)(unsafe.Pointer(c))
	C.UA_Client_reset(client)
}

func (c *UA_Client) Delete() {
	client := (*C.UA_Client)(unsafe.Pointer(c))
	C.UA_Client_delete(client)
}
