package conversion

import (
	"testing"
)

type SourceStruct struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type DestinationStruct struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func TestStructToStruct(t *testing.T) {
	src := SourceStruct{
		Name:  "Test",
		Value: 123,
	}

	var dst DestinationStruct

	err := StructToStruct(src, &dst)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if src.Name != dst.Name || src.Value != dst.Value {
		t.Fatalf("Expected dst to be %+v, got %+v", src, dst)
	}
}

func TestStructToStructWithUnexportedFields(t *testing.T) {
	type Source struct {
		Name  string `json:"name"`
		value int    // unexported field
	}

	type Destination struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	src := Source{
		Name:  "Test",
		value: 123,
	}

	var dst Destination

	err := StructToStruct(src, &dst)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if src.Name != dst.Name || dst.Value != 0 {
		t.Fatalf("Expected dst to be %+v, got %+v", src, dst)
	}
}

func TestStructToStructWithDifferentJsonTags(t *testing.T) {
	type Source struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	type Destination struct {
		Name  string `json:"name"`
		Value int    `json:"different_value"`
	}

	src := Source{
		Name:  "Test",
		Value: 123,
	}

	var dst Destination

	err := StructToStruct(src, &dst)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if src.Name != dst.Name || dst.Value != 0 {
		t.Fatalf("Expected dst to be %+v, got %+v", src, dst)
	}
}
