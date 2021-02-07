package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func openFile(path string) (fp *os.File) {
	fp, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return fp
}

func getHash(fp *os.File) []byte {
	h := sha256.New()
	if _, err := io.Copy(h, fp); err != nil {
		log.Fatal(err)
	}
	sum := h.Sum(nil)

	return sum
}

func processFile(path string, info os.FileInfo, err error) error {

	// Skip directories
	if info.IsDir() {
		return err
	}

	fp := openFile(path)
	defer fp.Close()

	hash := getHash(fp)

	fmt.Printf("File: %s, size: %d B, checksum: %x\n", path, info.Size(), hash)

	return err
}

func scandir(dirname string) {
	fmt.Println("Scanning directory: ", dirname)

	filepath.Walk(dirname, processFile)
}

func main() {
	var dirname string = "D:\\Photos\\Photos\\2020 Calendar"

	scandir(dirname)
}
