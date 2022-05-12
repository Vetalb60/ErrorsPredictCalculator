// Package requests
//
//	________delete_file_id_req.go________
//
//	Request to delete a file from the database by id.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package requests

import client "exec/client/client/operations"

func (req *Requests) DeleteFileById(id int64) (string, error) {

	ok, err := req.GetClient().DeleteDeleteID(&client.DeleteDeleteIDParams{
		ID:         id,
		Context:    req.GetContext(),
		HTTPClient: nil,
	})
	if err != nil {
		return "", err
	}
	return *ok.Payload.Message, nil
}
