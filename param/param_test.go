package param

import "testing"

func TestFetch(t *testing.T) {
	pn := "kzgsrs-bw6_761-lagrange-131072"

	err := Fetch(pn)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheck(t *testing.T) {
	err := CheckAll()
	if err != nil {
		t.Fatal(err)
	}
}
