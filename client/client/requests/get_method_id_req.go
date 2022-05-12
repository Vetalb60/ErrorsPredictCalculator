// Package requests
//
//	________get_method_id_req.go________
//
//	Request for calculate errors of predict.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package requests

import (
	client "exec/client/client/operations"
	"exec/client/models"
	"strconv"
)

func (req *Requests) GetErrorByMethod(method string, file_id int64) ([]float64, error) {

	var ok, err = req.GetClient().PostMethodID(&client.PostMethodIDParams{
		ID:         file_id,
		Method:     &models.Method{Method: &method},
		Context:    req.GetContext(),
		HTTPClient: nil,
	})
	if err != nil {
		return nil, err
	}

	errors := make([]float64, len(ok.Payload[0]))

	for index, elem := range ok.Payload[0] {
		errors[index], _ = strconv.ParseFloat(elem.Energy, 20)
	}
	return errors, nil
}
