package archive

import "testing"

func TestTar(t *testing.T) {
	isCompress := "gz"
	src := "/home/yydfjt/test/test1"
	dst := "/home/yydfjt/test/test2.tar.gz"
	err := Tar(src, dst, isCompress, 1024)
	if err != nil {
		t.Fatal(err)
	}

	untardir := "/home/yydfjt/test/test2"
	err = UnTar(dst, untardir, isCompress)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal()
}
