package archive

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const ShadowTar = ".config.tar.gz"

func Tar(src, dst string, compress string, maxSize int) error {
	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fw.Close()

	tw := new(tar.Writer)
	switch compress {
	case "gz":
		gw := gzip.NewWriter(fw)
		defer gw.Close()
		tw = tar.NewWriter(gw)
	default:
		tw = tar.NewWriter(fw)
	}
	defer tw.Close()

	src = path.Clean(src)
	return filepath.Walk(src, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip large file
		if fi.Size() >= int64(maxSize) {
			return nil
		}

		if fi.Name() == ShadowTar {
			return nil
		}

		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}

		hdr.Name = strings.TrimPrefix(fileName, src)
		hdr.Name = strings.TrimPrefix(hdr.Name, string(filepath.Separator))

		err = tw.WriteHeader(hdr)
		if err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		fr, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer fr.Close()
		n, err := io.Copy(tw, fr)
		if err != nil {
			return err
		}
		log.Printf("pack %s %s, written %d bytes\n", hdr.Name, fi.Name(), n)

		return nil
	})
}

func UnTar(src, dst string, compress string) error {
	fr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fr.Close()

	tr := new(tar.Reader)
	switch compress {
	case "gz":
		gr, err := gzip.NewReader(fr)
		if err != nil {
			return err
		}
		defer gr.Close()
		tr = tar.NewReader(gr)
	default:
		tr = tar.NewReader(fr)
	}

	for {
		hdr, err := tr.Next()

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case hdr == nil:
			continue
		}
		dstFileDir := path.Join(dst, hdr.Name)

		switch hdr.Typeflag {
		case tar.TypeDir:
			if b := ExistDir(dstFileDir); !b {
				if err := os.MkdirAll(dstFileDir, 0775); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			file, err := os.OpenFile(dstFileDir, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			defer file.Close()
			n, err := io.Copy(file, tr)
			if err != nil {
				return err
			}
			log.Printf("unpack %s, handle %d bytes\n", dstFileDir, n)
		}
	}
}

func ExistDir(dirname string) bool {
	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}
