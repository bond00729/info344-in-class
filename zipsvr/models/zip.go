package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Zip struct {
	Code  string `json:"code,omitempty"`
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// Slice of pointers to zip structs
type ZipSlice []*Zip

type ZipIndex map[string]ZipSlice

func LoadZips(filePath string) (ZipSlice, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	reader := csv.NewReader(f)
	// _ means i dont care about it, please ignore
	// Read only reads 1 line at a time and splits it into a slice of strings
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v", err)
	}

	zips := make(ZipSlice, 0, 43000)
	for {
		fields, err := reader.Read()
		if err == io.EOF {
			return zips, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record: %v", err)
		}

		z := &Zip{
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}

		zips = append(zips, z)
	}
}
