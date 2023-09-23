package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"time"
)

func Run(dirs []string) error {
	t := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %.4fs\n", time.Since(t).Seconds())
	}()

	for _, d := range dirs {
		if err := processDir(d); err != nil {
			return fmt.Errorf("Error processing dir %q\n: %v", d, err)
		}
	}

	return nil
}

func processDir(dirname string) error {
	log.Printf("Generating file checksums in %q\n", dirname)

	files, err := os.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, f := range files {
		cs, err := checksumFile(dirname, f)
		if err != nil {
			log.Printf("Error: %s, skipping %s\\%s", err, dirname, f.Name())
			continue
		}

		// Skip dirs.
		if cs == nil {
			continue
		}

		log.Printf("File: %s, size: %d B, checksum: %x\n", cs.FullPath(), cs.Size(), cs.Hash())
	}

	return nil
}

func checksumFile(dir string, node fs.DirEntry) (*File, error) {
	// Skip directories
	if node.IsDir() {
		return nil, nil
	}

	file, err := NewFile(node)
	if err != nil {
		return nil, err
	}

	file.SetFullPath(path.Join(dir, node.Name()))

	hash, err := getFileHash(file)
	if err != nil {
		return nil, err
	}

	return hash, err
}

func getFileHash(file *File) (*File, error) {
	fp, err := os.Open(file.FullPath())

	if err != nil {
		return nil, err
	}

	defer fp.Close()

	h := sha256.New()
	if _, err := io.Copy(h, fp); err != nil {
		return nil, err
	}

	file.SetHashFunc(h)
	file.SetHash(h.Sum(nil))

	return file, nil
}
