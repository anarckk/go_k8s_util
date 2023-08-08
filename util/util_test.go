package util

import (
	"testing"
)

func TestComposeMap(t *testing.T) {
	result := ComposeMap(map[string]string{
		"metadata.name": "pod1",
		"node":          "node1",
		"ttl":           "13",
	})
	if result != "metadata.name=pod1,node=node1,ttl=13" {
		t.Error(result)
	}
}
