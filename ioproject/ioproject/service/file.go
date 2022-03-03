package service

import (
	"bytes"
	"hash"
	"io"
	"math/rand"
	"time"

	"golang.org/x/crypto/sha3"
)

// File represents Aurora file
type File struct {
	id         string
	name       string
	hash       []byte
	dataReader io.Reader
	size       int64
	buffer     *bytes.Buffer
	startTime  time.Time
}

// NewRandomFile returns new pseudorandom file
func NewRandomFile(r *rand.Rand, name string, size int64) File {
	return File{
		name:       name,
		dataReader: io.LimitReader(r, size),
		size:       size,
	}
}

// NewBufferFile returns new file with specified buffer
func NewBufferFile(name string, buffer *bytes.Buffer) File {
	return File{
		name:       name,
		dataReader: buffer,
		size:       int64(buffer.Len()),
	}
}

// CalculateHash calculates hash from dataReader.
// It replaces dataReader with another that will contain the data.
func (f *File) CalculateHash() error {
	h := fileHasher()

	var buf bytes.Buffer
	tee := io.TeeReader(f.DataReader(), &buf)

	_, err := io.Copy(h, tee)
	if err != nil {
		return err
	}

	f.hash = h.Sum(nil)
	f.dataReader = &buf

	return nil
}

// Name returns file's name
func (f *File) Name() string {
	return f.name
}

// Hash returns file's hash
func (f *File) Hash() []byte {
	return f.hash
}

// DataReader returns file's data reader
func (f *File) DataReader() io.Reader {
	return f.dataReader
}

// Size returns file size
func (f *File) Size() int64 {
	return f.size
}

func fileHasher() hash.Hash {
	return sha3.New256()
}
