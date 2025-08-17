package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		args []string
		conf Config
	}{
		{
			[]string{"-m", "[a-z]", "-f", "dummy", "-c", "1 as int"},
			Config{regex: "[a-z]", casts: []string{"1 as int"}, inputFilePath: "dummy"},
		},
		{
			[]string{"-m", "[0-9]", "-c", "2 as float"},
			Config{regex: "[0-9]", casts: []string{"2 as float"}},
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			conf, error := ParseArgs(tt.args)
			if error != nil {
				t.Errorf("Error got %v, want nil", error)
			} else if !reflect.DeepEqual(*conf, tt.conf) {
				t.Errorf("Conf got: %+v\nwant: %+v", *conf, tt.conf)
			}
		})
	}
}

func TestGetInputFile(t *testing.T) {
	dummyFile, err := os.Open("dummy")
	if err != nil {
		t.Errorf("Error opening dummy file %v", err)
	}
	dummyFileStat, err := dummyFile.Stat()
	if err != nil {
		t.Errorf("Error stat-ing dummy File file %v", err)
	}

	stdinStat, err := os.Stdin.Stat()
	if err != nil {
		t.Errorf("Error stat-ing stdin file %v", err)
	}

	tests := []struct {
		args     string
		fileInfo os.FileInfo
	}{
		{"dummy", dummyFileStat},
		{"", stdinStat},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			inputFile, err := GetInputFile(tt.args)
			if err != nil {
				t.Errorf("Error got %v, want nil", err)

				return
			}
			inputFileStat, err := inputFile.Stat()
			if err != nil {
				t.Errorf("Error got %v, want nil", err)

				return
			}

			if !reflect.DeepEqual(inputFileStat, tt.fileInfo) {
				t.Errorf("File stat got: %+v\nwant: %+v", inputFileStat, tt.fileInfo)
			}
		})
	}
}
