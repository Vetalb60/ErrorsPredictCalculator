// Package requests
//
//	________get_files_list_req.go________
//
//	Request to get a list from the database.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package requests

import (
	client "exec/client/client/operations"
)

type FilesMetaInfo struct {
	Id_             int64
	File_name_      string
	File_size_      int64
	Date_of_insert_ string
}

func (req *Requests) GetFilesList() ([]FilesMetaInfo, error) {

	ok, err := req.GetClient().GetList(&client.GetListParams{
		Context:    req.GetContext(),
		HTTPClient: nil,
	})
	if err != nil {
		return nil, err
	}

	ret_array := make([]FilesMetaInfo, len(ok.Payload[0]))

	for index, elem := range ok.Payload[0] {
		ret_array[index].Id_ = elem.ID
		ret_array[index].File_name_ = elem.Name
		ret_array[index].File_size_ = elem.FileSize
		ret_array[index].Date_of_insert_ = elem.InsertDate
	}

	return ret_array, nil
}
