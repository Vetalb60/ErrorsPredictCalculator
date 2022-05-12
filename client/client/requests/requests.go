// Package requests
//
//	________requests.go________
//
//	A package with a description of request processing.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package requests

import (
	"context"
	client "exec/client/client/operations"
	httptransport "github.com/go-openapi/runtime/client"
)

type Requests struct {
	customized_ httptransport.Runtime
	client_     client.ClientService
	ctx_        context.Context
}

// Init
//	Initialization request params(socket, newClient, context)
func (req *Requests) Init(address string, port string, proto string) {
	req.customized_ = *httptransport.New(address+":"+port, "/", []string{proto})

	req.client_ = client.New(&req.customized_, nil)

	req.ctx_, _ = context.WithCancel(context.Background())
}

func (req *Requests) GetCustomized() httptransport.Runtime {
	return req.customized_
}

func (req *Requests) GetClient() client.ClientService {
	return req.client_
}

func (req *Requests) GetContext() context.Context {
	return req.ctx_
}
