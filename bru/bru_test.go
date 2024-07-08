package bru_test

import (
	_ "embed"
	"testing"

	"github.com/mukezhz/bru-go/bru"
)

//go:embed example.bru
var example []byte

func TestBruStruct(t *testing.T) {
	t.Parallel()
	b := bru.Bru{}
	if err := b.Unmarshal(example); err != nil {
		t.Error(err)
	}

	if b.Auth.Token == "" {
		t.Errorf("Expected auth token, got %s", b.Auth.Token)
	}

	if b.Auth.Token != "{{token}}" {
		t.Errorf("Expected {{token}}, got %s", b.Auth.Token)
	}

	if b.Meta.Name != "name of the example" {
		t.Errorf("Expected name of the example, got %s", b.Meta.Name)
	}
}

func TestUnmarshal(t *testing.T) {
	t.Parallel()
	b, err := bru.Unmarshal(example)
	if err != nil {
		t.Error(err)
	}

	if b.Auth.Token == "" {
		t.Errorf("Expected auth token, got %s", b.Auth.Token)
	}

	if b.Auth.Token != "{{token}}" {
		t.Errorf("Expected {{token}}, got %s", b.Auth.Token)
	}

	if b.Meta.Name != "name of the example" {
		t.Errorf("Expected name of the example, got %s", b.Meta.Name)
	}
}
