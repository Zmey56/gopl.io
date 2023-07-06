//“Exercise 10.2:
//Define a generic archive file-reading function capable of reading ZIP
//files (archive/zip) and POSIX tar files (archive/tar).
//
//Use a registration mechanism similar to the one described above so
//that support for each file format can be plugged in using blank
//imports.”

package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ArchiveFile represents a file in an archive.
type ArchiveFile interface {
	io.Reader
	io.Closer
}

// ArchiveReader represents an archive reader capable of reading different file formats.
type ArchiveReader interface {
	Open(path string) (ArchiveFile, error)
}

// ZipReader is an implementation of ArchiveReader for ZIP files.
type ZipReader struct{}

func (z ZipReader) Open(path string) (ArchiveFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	zipFile, err := zip.NewReader(file, stat.Size())
	if err != nil {
		return nil, err
	}
	return &zipArchiveFile{zipFile, file}, nil
}

// TarReader is an implementation of ArchiveReader for POSIX tar files.
type TarReader struct{}

func (t TarReader) Open(path string) (ArchiveFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &tarArchiveFile{tar.NewReader(file), file}, nil
}

// zipArchiveFile is a wrapper type that embeds zip.Reader and *os.File and implements ArchiveFile interface.
type zipArchiveFile struct {
	*zip.Reader
	File *os.File
}

func (zf *zipArchiveFile) Read(p []byte) (int, error) {
	return zf.File.Read(p)
}

func (zf *zipArchiveFile) Close() error {
	err := zf.File.Close()
	if err != nil {
		return err
	}
	return nil
}

// tarArchiveFile is a wrapper type that embeds tar.Reader and *os.File and implements ArchiveFile interface.
type tarArchiveFile struct {
	*tar.Reader
	File *os.File
}

func (tf *tarArchiveFile) Read(p []byte) (int, error) {
	return tf.Reader.Read(p)
}

func (tf *tarArchiveFile) Close() error {
	err := tf.File.Close()
	if err != nil {
		return err
	}
	return nil
}

// ArchiveFileType is a map that associates file extensions with their corresponding ArchiveReader implementations.
var ArchiveFileType = map[string]ArchiveReader{
	".zip": ZipReader{},
	".tar": TarReader{},
	// Add more file formats here as needed
}

// ReadArchiveFile reads the contents of an archive file using the appropriate ArchiveReader based on the file extension.
func ReadArchiveFile(path string) error {
	extension := filepath.Ext(path)
	reader, ok := ArchiveFileType[extension]
	if !ok {
		return fmt.Errorf("unsupported file format: %s", extension)
	}

	archiveFile, err := reader.Open(path)
	if err != nil {
		return err
	}
	defer archiveFile.Close()

	// Process the contents of the archive file
	// Here, you can read and handle individual files within the archive

	return nil
}

func main() {
	err := ReadArchiveFile("example.zip")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
