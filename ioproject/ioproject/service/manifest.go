package service

import (
	"archive/tar"
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var DefaultPathnameMaxLength int32 = 64

// Manifest represent Aurora manifest
type Manifest struct {
	hash          []byte
	size          int64
	indexDocument string
	errorDocument string
	directories   Directory
}

type Directory struct {
	name      string
	files     []File
	directory []Directory
}

// NewRandomManifest returns new pseudorandom manifest with many files
func NewRandomManifest(r *rand.Rand, fileCount int, fileSize int64) (manifest *Manifest, err error) {
	files := make([]File, fileCount)
	sz := int64(0)

	for i := 0; i < fileCount; i++ {
		pathnameLength := int64(r.Int31n(DefaultPathnameMaxLength-1)) + 1
		b := make([]byte, pathnameLength)
		_, err = r.Read(b)
		if err != nil {
			return nil, err
		}
		pathname := hex.EncodeToString(b)
		r.Seed(time.Now().UnixNano() + int64(i))
		file := NewRandomFile(r, pathname, fileSize)
		err = file.CalculateHash()
		if err != nil {
			return nil, err
		}

		files[i] = file
		sz += fileSize
	}

	manifest = &Manifest{
		size: sz,
		directories: Directory{
			name:  "",
			files: files,
		},
	}

	return
}

type FileStructure struct {
	Name  string
	Size  int64
	Input *bytes.Buffer
}

type DirStructure struct {
	Name   string
	Files  []FileStructure
	SubDir []DirStructure
}

func NewStructuredManifest(r *rand.Rand, directory DirStructure) (*Manifest, error) {
	var (
		manifest  Manifest
		iterateFn func(ds DirStructure, d *Directory) int64
	)

	iterateFn = func(ds DirStructure, d *Directory) (sz int64) {
		d.name = ds.Name
		d.files = make([]File, len(ds.Files))

		for i, f := range ds.Files {
			if f.Input != nil {
				d.files[i] = NewBufferFile(f.Name, f.Input)
			} else {
				d.files[i] = NewRandomFile(r, f.Name, f.Size)
			}

			d.files[i].CalculateHash()

			sz += d.files[i].Size()
		}

		if len(ds.SubDir) == 0 {
			return sz
		}

		d.directory = make([]Directory, len(ds.SubDir))

		for i, dir := range ds.SubDir {
			sz += iterateFn(dir, &d.directory[i])
		}

		return sz
	}

	manifest.size = iterateFn(directory, &manifest.directories)

	return &manifest, nil
}

func (m *Manifest) SetIndexDocument(pathname string) bool {
	pathname = strings.TrimPrefix(pathname, string(os.PathSeparator))
	if pathname == "" {
		return false
	}
	filename := pathname[strings.Index(pathname, string(os.PathSeparator)):]
	basepath := strings.TrimSuffix(strings.TrimSuffix(pathname, filename), string(os.PathSeparator))

	dir := m.directories

	dirnames := strings.Split(basepath, string(os.PathSeparator))
	for _, dirname := range dirnames {
		var found bool
		for _, d := range dir.directory {
			if d.name == dirname {
				found = true
				dir = d
				break
			}
		}
		if !found {
			return false
		}
	}

	for _, f := range dir.files {
		if f.name == filename {
			m.indexDocument = pathname
			return true
		}
	}

	return false
}

func (m *Manifest) SetErrorDocument(pathname string) bool {
	pathname = strings.TrimPrefix(pathname, string(os.PathSeparator))
	if pathname == "" {
		return false
	}
	filename := pathname[strings.Index(pathname, string(os.PathSeparator)):]
	basepath := strings.TrimSuffix(strings.TrimSuffix(pathname, filename), string(os.PathSeparator))

	dir := m.directories

	dirnames := strings.Split(basepath, string(os.PathSeparator))
	for _, dirname := range dirnames {
		var found bool
		for _, d := range dir.directory {
			if d.name == dirname {
				found = true
				dir = d
				break
			}
		}
		if !found {
			return false
		}
	}

	for _, f := range dir.files {
		if f.name == filename {
			m.errorDocument = pathname
			return true
		}
	}

	return false
}

func (m *Manifest) Archive() (*bytes.Buffer, error) {
	var iterateFn func(d Directory, path string) error

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	iterateFn = func(d Directory, path string) (err error) {
		for _, f := range d.files {
			// create tar header and write it
			hdr := tar.Header{
				Typeflag: tar.TypeReg,
				Name:     path + f.Name(),
				Mode:     0600,
				Size:     f.Size(),
			}
			err = tw.WriteHeader(&hdr)
			if err != nil {
				return
			}

			var b []byte

			b, err = ioutil.ReadAll(f.DataReader())
			if err != nil {
				return
			}

			// write the file data to the tar
			_, err = tw.Write(b)
			if err != nil {
				return
			}
		}

		if len(d.directory) != 0 {
			for _, dir := range d.directory {
				hdr := tar.Header{
					Typeflag: tar.TypeDir,
					Name:     path + dir.name,
					Mode:     0600,
				}
				err = tw.WriteHeader(&hdr)
				if err != nil {
					return
				}
				err = iterateFn(dir, path+dir.name+string(os.PathSeparator))
				if err != nil {
					return
				}
			}
		}

		return
	}

	if err := iterateFn(m.directories, ""); err != nil {
		return nil, err
	}

	if err := tw.Close(); err != nil {
		return nil, err
	}

	return buf, nil
}

// Address returns manifest's address

// Hash returns manifest's hash
func (m *Manifest) Hash() []byte {
	return m.hash
}

// Size returns manifest total size
func (m *Manifest) Size() int64 {
	return m.size
}

// RootDirectory returns manifest root directory
func (m *Manifest) RootDirectory() Directory {
	return m.directories
}

// IndexDocument returns manifest specify index document name
func (m *Manifest) IndexDocument() string {
	return m.indexDocument
}

// ErrorDocument returns manifest specify error document name
func (m *Manifest) ErrorDocument() string {
	return m.errorDocument
}

// Name returns directory name
func (d Directory) Name() string {
	return d.name
}

// Files returns directory contains files
func (d Directory) Files() []File {
	return d.files
}

// SubDirectories returns directory contains sub-directories
func (d Directory) SubDirectories() []Directory {
	return d.directory
}
