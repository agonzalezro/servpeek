package examples

import (
	"regexp"
	"testing"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/resource/file"
)

func Test_File(t *testing.T) {
	f := &resource.File{
		Path: "/etc/hosts",
	}

	if err := file.IsRegular(f); err != nil {
		t.Errorf("Error: %s", err)
	}

	owner := "root"
	group := "wheel"
	md5 := "YOUR_MD5SUM_HERE"
	if err := file.IsOwnedBy(f, owner); err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := file.IsGrupedInto(f, group); err != nil {
		t.Errorf("Error: %s", err)
	}

	if err := file.Md5(f, md5); err != nil {
		t.Errorf("Error: %s", err)
	}

	content := regexp.MustCompile(`localhost`)
	if err := file.Contains(f, content); err != nil {
		t.Errorf("Error: %s", err)
	}
}
