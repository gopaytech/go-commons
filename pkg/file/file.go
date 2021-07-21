package file

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type JoinPathsFunc func(paths ...string) (path string)

func JoinPaths(paths ...string) (path string) {
	return filepath.Join(paths...)
}

type MoveDirectoryContentsFunc func(sourceDirectory string, targetDirectory string) (err error)

func MoveDirectoryContents(sourceDirectory string, targetDirectory string) (err error) {
	files, err := GetFilesInDirectory(sourceDirectory)
	if err != nil {
		return
	}
	folders, err := GetFoldersInDirectory(sourceDirectory)
	if err != nil {
		return
	}

	for _, folder := range folders {
		newDirectory := fmt.Sprintf("%s/%s", targetDirectory, folder)
		err = CreateFolder(newDirectory)
		if err != nil {
			return
		}
	}

	for _, file := range files {
		sourcePath := fmt.Sprintf("%s/%s", sourceDirectory, file)
		targetPath := fmt.Sprintf("%s/%s", targetDirectory, file)
		err = Move(sourcePath, targetPath)
		if err != nil {
			return
		}
	}
	return
}

type FileExistsFunc func(filename string) bool

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	}
	return !info.IsDir()
}

type DirExistsFunc func(dirName string) bool

func DirExists(dirName string) bool {
	info, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	}
	return info.IsDir()
}

type IsParentDirFunc func(filePath string, parent string) (result bool, err error)

func IsParentDir(filePath string, parent string) (result bool, err error) {
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return
	}

	if !FileExists(absFilePath) {
		err = fmt.Errorf("file %s does not not exists (absolute: %s)", filePath, absFilePath)
		return
	}

	absParent, err := filepath.Abs(parent)
	if err != nil {
		return
	}

	if !DirExists(absParent) {
		err = fmt.Errorf("directory %s does not not exists (absolute: %s)", parent, absParent)
		return
	}

	result = filepath.Dir(absFilePath) == absParent
	return
}

type CreateFolderFunc func(folder string) (err error)

func CreateFolder(folder string) (err error) {
	err = os.Mkdir(folder, os.ModeDir)
	return
}

type MoveFunc func(sourceFile string, destinationFile string) (err error)

func Move(sourceFile string, destinationFile string) (err error) {
	err = os.Rename(sourceFile, destinationFile)
	return
}

type CopyFunc func(sourceFile string, destinationFile string) (err error)

func Copy(sourceFile string, destinationFile string) (err error) {
	sourceFileStat, err := os.Stat(sourceFile)
	if err != nil {
		return
	}

	if !sourceFileStat.Mode().IsRegular() {
		return
	}

	source, err := os.Open(sourceFile)
	if err != nil {
		return
	}
	defer source.Close()

	destination, err := os.Create(destinationFile)
	if err != nil {
		return
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)

	return
}

type GetBasenameFunc func(filename string) (basename string)

func GetBasename(filename string) (basename string) {
	return filepath.Base(filename)
}

type GetFilesInDirectoryFunc func(directory string) (files []string, err error)

func GetFilesInDirectory(directory string) (files []string, err error) {
	baseDirectory := filepath.Clean(directory)
	err = filepath.Walk(baseDirectory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			cleanPath := strings.Replace(path, fmt.Sprintf("%s/", baseDirectory), "", 1)
			files = append(files, cleanPath)
			return nil
		})
	return
}

type GetFoldersInDirectoryFunc func(directory string) (directories []string, err error)

func GetFoldersInDirectory(directory string) (directories []string, err error) {
	baseDirectory := filepath.Clean(directory)
	err = filepath.Walk(baseDirectory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				return nil
			}
			if baseDirectory == path {
				return nil
			}
			cleanPath := strings.Replace(path, fmt.Sprintf("%s/", baseDirectory), "", 1)
			directories = append(directories, cleanPath)
			return nil
		})
	return
}

type UnzipFunc func(filename string, directory string) (filenames []string, err error)

func Unzip(filename string, directory string) (filenames []string, err error) {
	directory = filepath.Clean(directory)
	if directory == "." {
		pwd, ierr := os.Getwd()
		if ierr != nil {
			err = ierr
			return
		}
		directory = pwd
	}

	r, err := zip.OpenReader(filename)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(directory, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(directory)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			_ = os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		_ = outFile.Close()
		_ = rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

type SaveFunc func(content []byte, filename string) (err error)

func Save(content []byte, filename string) (err error) {
	return ioutil.WriteFile(filename, content, 0644)
}

type TarFunc func(sourceDirectory string, writer io.Writer) (err error)

func Tar(sourceDirectory string, writer io.Writer) (err error) {
	sourceDirectory, err = filepath.Abs(sourceDirectory)
	if err != nil {
		return err
	}

	if _, err = os.Stat(sourceDirectory); err != nil {
		return
	}
	gzWriter := gzip.NewWriter(writer)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
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
