// Package requests
//
//	________get_healthz_req.go________
//
//	Request to get a heathz.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package requests

import (
	client "exec/client/client/operations"
)

func (req *Requests) GetHealthz() (string, error) {
	ok, err := req.GetClient().GetHealthz(&client.GetHealthzParams{
		Context:    req.GetContext(),
		HTTPClient: nil,
	})
	if err != nil {
		return "", err
	}

	return *ok.GetPayload().Message, nil
}
