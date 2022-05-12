// Package waveform
//
//	________waveform.go________
//
//	The module is designed for reading and decoding .wav files.
//	Files can be read from both the file system and the MYSQL database.
//	There is a logging parameter.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package waveform

import (
	mysql "CourseWork/api/databases/mysql"
	"log"
)

//	Some labels of .wav files.
const (
	RIFF = 0x52494646
	WAVE = 0x57415645
	FMT  = 0x666d7420
	DATA = 0x64617461
)

// WAVDecoder
//	Contains methods for processing .wav and descriptors .wav file.
type WAVDecoder struct {
	Desc *WAVFormat
	Proc *WAVHandler
}

// WAVFormat
//	Descriptor struct.
type WAVFormat struct {
	Meta_ MetaInfo
	Data_ []byte
}

// MetaInfo
//	MetaInformation of .wav file.
type MetaInfo struct {
	AudioFormat_   uint16 //	Audio Format. For PCM is 1.
	NumChannels_   uint16 //	Number of channels (Mono = 1, Stereo = 2).
	SampleRate_    uint32 //	Sampling rate.
	ByteRate_      uint32 //	The number of bytes transmitted per second of playback.
	BlockAlign_    uint16 //	The number of bytes per sample, including all channels.
	BitsPerSample_ uint16 //	The number of bits in the sample. Depth or accuracy of sound.
	FileSize_      uint32 //	Size of .wav file
}

// DecodeFile
//	The method of decoding a file from a file system.
func (dec *WAVDecoder) DecodeFile(log *log.Logger, file_path string) (*WAVFormat, error) {
	var err error

	dec.Desc, err = dec.Proc.readFile(log, file_path)
	if err != nil {
		return nil, err
	}

	log.Printf("%s was successful decode!", file_path)

	return dec.Desc, nil
}

// DecodeFileFromDB
//	The method of decoding a file from MYSQL database.
//	The file description must be passed to the method as a parameter.
//	The descriptor can be obtained using the
//	GetFileFromDB(*log.Logger, int64)(*FileDescriptor, error) from mysql package.
func (dec *WAVDecoder) DecodeFileFromDB(log *log.Logger, file_desc *mysql.FileDescriptor) (*WAVFormat, error) {
	var err error
	log.Printf("Decoding file %s from database...", file_desc.Meta_.File_name_)

	dec.Desc, err = dec.fillWAVDesc(file_desc)

	if err != nil {
		return nil, err
	}

	log.Printf("%s was successful decode!", file_desc.Meta_.File_name_)
	return dec.Desc, nil
}

// fillWAVDesc
//	A method that fills in a WAV descriptor using a file descriptor from DB.
func (dec *WAVDecoder) fillWAVDesc(file_desc *mysql.FileDescriptor) (*WAVFormat, error) {
	var err error

	dec.Desc, err = dec.Proc.fillDesc(file_desc.Data_)
	if err != nil {
		return nil, err
	}

	return dec.Desc, nil
}
