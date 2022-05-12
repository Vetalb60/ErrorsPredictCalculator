// Package requests
//
//	________get_req.go________
//
//	Basic path request.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package requests

import client "exec/client/client/operations"

func (req *Requests) Get() (string, error) {
	ok, err := req.GetClient().Get(&client.GetParams{
		Context:    req.GetContext(),
		HTTPClient: nil,
	})
	if err != nil {
		return "", err
	}

	return *ok.GetPayload().Message, nil
}
