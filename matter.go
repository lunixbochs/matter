// matter is a library for easily reading/writing files with frontmatter.
//
// Frontmatter with YAML usually looks like this:
//
//     ---
//     key: value
//     key: value
//     ---
//     File contents
//
package matter

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var sep = []byte("---\n")

func surround(data []byte) []byte {
	return append(append(append(sep, data...), '\n'), sep...)
}

func separate(p []byte) (matter []byte, data []byte, err error) {
	result := bytes.SplitN(p, sep, 3)
	if len(result) < 3 {
		err = errors.New("Could not find frontmatter.")
		return
	}
	matter, data = result[1], result[2]
	matter = bytes.TrimRight(matter, "\r\n")
	return
}

func combine(matter []byte, data []byte) []byte {
	matter = bytes.TrimRight(matter, "\r\n")
	return append(surround(matter), data...)
}

// ReadFile is equivalent to ioutil.ReadFile, except frontmatter will be returned separately.
func ReadFile(filename string) (matter []byte, data []byte, err error) {
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	matter, data, err = separate(data)
	return
}

// WriteFile is equivalent to ioutil.WriteFile, except `matter` will be prefixed as frontmatter.
func WriteFile(filename string, matter []byte, data []byte, perm os.FileMode) error {
	data = combine(matter, data)
	return ioutil.WriteFile(filename, data, perm)
}

// ReadYAML is equivalent to ReadFile, but will also unmarshal YAML frontmatter into `out`.
func ReadYAML(filename string, out interface{}) (data []byte, err error) {
	var matter []byte
	matter, data, err = ReadFile(filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(matter, out)
	return
}

// WriteYAML is equivalent to WriteFile, but will marshal `matter` as YAML into frontmatter.
func WriteYAML(filename string, in interface{}, data []byte, perm os.FileMode) error {
	matter, err := yaml.Marshal(in)
	if err != nil {
		return err
	}
	data = combine(matter, data)
	return ioutil.WriteFile(filename, data, perm)
}
