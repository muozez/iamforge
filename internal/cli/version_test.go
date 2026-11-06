package cli

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestWriteVersionText(t *testing.T) {
	buf := new(bytes.Buffer)
	if err := writeVersion(buf, false); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()
	if !strings.HasPrefix(got, "iamforge dev") {
		t.Errorf("expected string to start with 'iamforge dev', got: %s", got)
	}
	if !strings.Contains(got, "commit: none") {
		t.Errorf("expected commit info, got: %s", got)
	}
}

func TestWriteVersionJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	if err := writeVersion(buf, true); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var info VersionInfo
	if err := json.Unmarshal(buf.Bytes(), &info); err != nil {
		t.Fatalf("invalid json output: %v", err)
	}

	if info.Version != "dev" {
		t.Errorf("expected version dev, got: %s", info.Version)
	}
	if info.Commit != "none" {
		t.Errorf("expected commit none, got: %s", info.Commit)
	}
	if info.DataVersion != "2026-11-07" {
		t.Errorf("expected data_version 2026-11-07, got: %s", info.DataVersion)
	}
}

