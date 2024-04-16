package simplefs

import (
	"bytes"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestFs(t *testing.T) {
	rdir := "~/test/data"
	rdir, _ = homedir.Expand(rdir)
	fs, err := New(rdir)
	if err != nil {
		t.Fatal(err)
	}

	key := []byte("testa")
	val := []byte("abcdefg")
	err = fs.Put(key, val)
	if err != nil {
		t.Fatal(err)
	}

	nval, err := fs.Get(key)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(val, nval) {
		t.Fatal("unequal val")
	}

	pval, err := fs.Get(key, 1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(val[1:3], pval) {
		t.Fatal("unequal partial val")
	}

	wval, err := fs.Get(key, 2, 8)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(val[2:], wval) {
		t.Fatal("unequal partial out val")
	}
}
