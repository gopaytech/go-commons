package compress

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopaytech/go-commons/pkg/dir"
)

func Tar(sourceDirectory string, writer io.Writer) (err error) {
	if _, err = os.Stat(sourceDirectory); err != nil {
		return
	}

	tarWriter := tar.NewWriter(writer)
	defer tarWriter.Close()

	err = filepath.Walk(sourceDirectory, func(file string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.Mode().IsRegular() {
			return nil
		}

		header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())
		if err != nil {
			return err
		}

		localDirectory := strings.Replace(file, sourceDirectory, "", -1)
		header.Name = strings.TrimPrefix(localDirectory, string(filepath.Separator))

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		openFile, err := os.Open(file)
		if err != nil {
			return err
		}
		if _, err := io.Copy(tarWriter, openFile); err != nil {
			return err
		}
		err = openFile.Close()
		return err
	})

	return
}

func TarGz(sourceDirectory string, writer io.Writer) (err error) {
	if _, err = os.Stat(sourceDirectory); err != nil {
		return
	}

	gzWriter := gzip.NewWriter(writer)
	defer gzWriter.Close()
	return Tar(sourceDirectory, gzWriter)

}

func UnTarGz(src io.Reader, dst string) (totalWritten int64, err error) {
	zipReader, errReader := gzip.NewReader(src)
	if errReader != nil {
		err = errReader
		return
	}
	defer zipReader.Close()

	return UnTar(zipReader, dst)
}

func UnTar(src io.Reader, dst string) (written int64, err error) {
	tarReader := tar.NewReader(src)

	validPath := func(path string) bool {
		if path == "" ||
			strings.Contains(path, `\`) ||
			strings.HasPrefix(path, "/") ||
			strings.Contains(path, "../") {
			return false
		}
		return true
	}

	var totalWritten int64
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return totalWritten, err
		}

		if !validPath(header.Name) {
			return totalWritten, fmt.Errorf("tar contained invalid path %s\n", header.Name)
		}

		target := filepath.Join(dst, header.Name)

		switch header.Typeflag {

		case tar.TypeDir:
			if !dir.Exists(target) {
				if err := os.MkdirAll(target, 0755); err != nil {
					return totalWritten, err
				}
			}
		case tar.TypeReg:
			fileToWrite, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return totalWritten, err
			}
			written, err := io.Copy(fileToWrite, tarReader)
			if err != nil {
				return totalWritten, err
			}

			totalWritten = totalWritten + written
			_ = fileToWrite.Close()
		}
	}

	return totalWritten, nil
}
