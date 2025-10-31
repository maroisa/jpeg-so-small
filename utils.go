package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"slices"
)

var JpgExtensions []string = []string{".jpg", ".jpeg", ".JPG", ".JPEG"}

func ReadDir(dirpath string) ([]os.DirEntry, error) {
	var filteredEntries []os.DirEntry

	entries, err := os.ReadDir(dirpath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil, err
	}

	for _, entry := range entries {
		contain := slices.Contains(
			JpgExtensions,
			filepath.Ext(entry.Name()),
		)

		if contain {
			filteredEntries = append(filteredEntries, entry)
		}
	}
	return filteredEntries, nil
}

func CompressJpg(quality int, entry string, output string) error {
	file, err := os.Open(entry)
	if err != nil {
		log.Println("Failed to open file:", err)
		return err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println("Failed to decode image:", err)
		return err
	}

	out, err := os.Create(output)
	if err != nil {
		log.Println("Failed to create image:", err)
		return err
	}

	defer out.Close()

	err = jpeg.Encode(out, img, &jpeg.Options{
		Quality: quality,
	})
	if err != nil {
		log.Println("Failed to create image:", err)
		return err
	}

	return nil
}
