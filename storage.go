package main

// clarify Marshall stuff and comapre with JS as I have never heard about it before
// In Go, "marshalling" refers to the process of converting a data structure (like a struct) into a format that can be easily stored or transmitted, such as JSON. The `json.Marshal` function is used to convert Go data structures into JSON format, while `json.Unmarshal` is used to convert JSON data back into Go data structures.

// In JavaScript, the equivalent process is often referred to as "serialization" or "stringifying". The `JSON.stringify` method is used to convert JavaScript objects into JSON strings, and `JSON.parse` is used to convert JSON strings back into JavaScript objects.

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
