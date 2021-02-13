package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func getFileHash(path string) (sum []byte, err error) {
	fp, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer fp.Close()

	h := sha256.New()
	if _, err := io.Copy(h, fp); err != nil {
		return nil, err
	}

	sum = h.Sum(nil)

	return sum, err
}

func checksumFile(path string, info os.FileInfo, err error) error {
	// Skip directories
	if info.IsDir() {
		return err
	}

	hash, err := getFileHash(path)

	// Log errors and keep going
	if err != nil {
		log.Println(err)

		return nil
	}

	log.Printf("File: %s, size: %d B, checksum: %x\n", path, info.Size(), hash)

	return err
}

func processDir(dirname string) {
	fmt.Printf("Checksumming files in \"%s\"", dirname)

	filepath.Walk(dirname, checksumFile)
}

func main() {
	var dirname string = "D:\\Photos\\Photos\\2020 Calendar"

	processDir(dirname)
}
