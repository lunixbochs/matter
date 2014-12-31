package matter

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestRaw(t *testing.T) {
	matter := []byte("test matter")
	data := []byte("test data")
	err := WriteFile("test", matter, data, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test")
	matter2, data2, err := ReadFile("test")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(matter, matter2) {
		fmt.Println("in: ", matter)
		fmt.Println("out:", matter2)
		t.Fatal("Frontmatter was not preserved.")
	}
	if !bytes.Equal(data, data2) {
		fmt.Println("in: ", data)
		fmt.Println("out:", data2)
		t.Fatal("Data was not preserved.")
	}
}

type TestStruct struct {
	FieldOne string
	FieldTwo string
}

func TestYAML(t *testing.T) {
	matter := &TestStruct{"a", "b"}
	data := []byte("test data")
	err := WriteYAML("test.yml", matter, data, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test.yml")
	matter2 := &TestStruct{}
	data2, err := ReadYAML("test.yml", matter2)
	if err != nil {
		t.Fatal(err)
	}
	if matter.FieldOne != matter2.FieldOne || matter.FieldTwo != matter2.FieldTwo {
		fmt.Println("in: ", matter)
		fmt.Println("out:", matter2)
		t.Fatal("Frontmatter was not preserved.")
	}
	if !bytes.Equal(data, data2) {
		fmt.Println("in: ", data)
		fmt.Println("out:", data2)
		t.Fatal("Data was not preserved.")
	}
}
