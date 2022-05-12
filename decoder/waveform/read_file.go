// Package waveform
//
//	________read_file.go________
//
//	Components for processing .wav file data.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package waveform

import (
	"encoding/binary"
	"errors"
	"log"
	"os"
)

// WAVHandler
//	Structure with methods for processing data of WAVE type files.
type WAVHandler struct{}

func (proc *WAVHandler) readFile(log *log.Logger, file_path string) (desc *WAVFormat, err error) {
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	file_info, _ := file.Stat()
	log.Printf("Reading file %s...", file_info.Name())

	data := make([]byte, file_info.Size())
	number, err := file.Read(data)
	if err != nil {
		return nil, err
	}

	log.Printf("%d bytes was read.\n", number)
	err = proc.isWAV(log, data)
	if err != nil {
		return nil, err
	}

	desc, err = proc.fillDesc(data)
	if err != nil {
		return nil, err
	}

	return desc, nil
}

func (proc *WAVHandler) fillDesc(data []byte) (desc *WAVFormat, err error) {
	fmt_index, err := proc.goToFMT(data)
	if err != nil {
		return nil, err
	}

	desc = proc.fillMetaInfo(data, fmt_index)

	data_index, err := proc.goToData(data)
	if err != nil {
		return nil, err
	}

	desc = proc.fillDataInfo(data, data_index, desc)

	return desc, nil
}

// fillMetaInfo
//	A method for filling in meta information.
func (proc *WAVHandler) fillMetaInfo(data []byte, fmt_index int) *WAVFormat {
	desc := new(WAVFormat)

	fmt_index += 8
	desc.Meta_.AudioFormat_ = binary.LittleEndian.Uint16(data[fmt_index : fmt_index+2])
	fmt_index += 2
	desc.Meta_.NumChannels_ = binary.LittleEndian.Uint16(data[fmt_index : fmt_index+2])
	fmt_index += 2
	desc.Meta_.SampleRate_ = binary.LittleEndian.Uint32(data[fmt_index : fmt_index+4])
	fmt_index += 4
	desc.Meta_.ByteRate_ = binary.LittleEndian.Uint32(data[fmt_index : fmt_index+4])
	fmt_index += 4
	desc.Meta_.BlockAlign_ = binary.LittleEndian.Uint16(data[fmt_index : fmt_index+2])
	fmt_index += 2
	desc.Meta_.BitsPerSample_ = binary.LittleEndian.Uint16(data[fmt_index : fmt_index+2])

	return desc
}

// fillDataInfo
//	A method for filling data into descriptor.
func (proc *WAVHandler) fillDataInfo(data []byte, data_index int, desc *WAVFormat) *WAVFormat {
	data_index += 4
	desc.Meta_.FileSize_ = binary.LittleEndian.Uint32(data[data_index : data_index+4])
	data_index += 4

	desc.Data_ = data[data_index:]

	return desc
}

// goToFMT
//	Go to the "fmt" label. The method returns the position of the label.
func (proc *WAVHandler) goToFMT(file_data []byte) (position int, err error) {
	for i := int(0); i < 100; i++ {
		if int(binary.BigEndian.Uint32(file_data[i:i+4])) == FMT {
			return i, nil
		}
	}

	return 0, errors.New("invalid type of file. Is not .wav file")
}

// goToData
//	Go to the "data" label. The method returns the position of the label.
func (proc *WAVHandler) goToData(file_data []byte) (position int, err error) {
	for i := int(0); i < 100; i++ {
		if int(binary.BigEndian.Uint32(file_data[i:i+4])) == DATA {
			return i, nil
		}
	}

	return 0, errors.New("invalid type of file. Is not .wav file")
}

// isWAV
//	Check: whether the file format is .wav.
//	To do this, 2 labels are checked at the beginning of the file - "RIFF" and "WAVE".
func (proc *WAVHandler) isWAV(log *log.Logger, data []byte) error {
	if (proc.checkWAVE(data) && proc.checkRIFF(data)) == false {
		log.Printf("File is not .wav!")
		return errors.New("file is not .wav")
	}

	return nil
}

func (proc *WAVHandler) checkRIFF(data []byte) bool {
	return binary.BigEndian.Uint32(data[0:4]) == RIFF
}

func (proc *WAVHandler) checkWAVE(data []byte) bool {
	return binary.BigEndian.Uint32(data[8:12]) == WAVE
}
