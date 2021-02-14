package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

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

func checksumFile(dir string, info os.FileInfo) (*File, error) {
	// Skip directories
	if info.IsDir() {
		return nil, nil
	}

	file := NewFile(info)

	file.SetFullPath(path.Join(dir, info.Name()))

	file, err := getFileHash(file)

	// Log errors and keep going
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return file, err
}

func processDir(dirname string) {
	fmt.Printf("Checksumming files in \"%s\"", dirname)

	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		cs, err := checksumFile(dirname, f)

		if err != nil {
			log.Printf("Error: %s, skipping %s\\%s", err, dirname, f.Name())

			continue
		}

		log.Printf("File: %s, size: %d B, checksum: %x\n", cs.FullPath(), cs.Size(), cs.Hash())
	}

}

func main() {
	t := time.Now()

	var dirname string = "D:\\Photos\\Photos\\2020 Calendar"

	processDir(dirname)

	fmt.Printf("Run time: %.4fs", time.Since(t).Seconds())
}
