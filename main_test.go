package main

import (
	"bytes"
	"testing"
)

func TestCountWrods(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	res, err := count(b)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if res.words != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2 word4\nline3 word5")
	exp := 3
	res, err := count(b)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if res.lines != exp {
		t.Errorf("Expected %d, got %d", exp, res.lines)
	}
}
func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1\n ")
	exp := b.Len()
	res, err := count(b)
	if err != nil {
		t.Errorf("Error: %v", err)

	}

	if res.bytes != exp {
		t.Errorf("Expected %d, got %d", exp, res.bytes)
	}
}
