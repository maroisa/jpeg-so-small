package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx     context.Context
	tempDir string
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OpenDirectoryDialog() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Println("Error:", err)
	}

	if a.tempDir == "" {
		a.tempDir = homedir
	}

	dirpath, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: a.tempDir,
	})
	if err != nil {
		log.Println("Error:", err)
		return ""
	}

	a.tempDir, err = filepath.EvalSymlinks(dirpath)
	if err != nil {
		log.Println("Error:", err)
		return ""
	}

	return a.tempDir
}

func (a *App) Compress(inputDir string, outputDir string, quality int, nameFixType string, nameFix string) {
	fmt.Println(strings.Repeat("-", 32))
	fmt.Println("Compressing")
	fmt.Println("-- from:", inputDir)
	fmt.Println("-- to:", outputDir)
	fmt.Println("-- with quality:", quality)
	fmt.Println("-- and", nameFixType+": '"+nameFix+"'")

	entries, err := ReadDir(inputDir)
	if err != nil {
		log.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		var output string = entry.Name()

		if nameFixType == "prefix" {
			output = nameFix + entry.Name()
		}
		if nameFixType == "suffix" {
			trimmed := strings.TrimSuffix(
				entry.Name(),
				filepath.Ext(entry.Name()),
			)
			output = trimmed + nameFix + filepath.Ext(entry.Name())
		}

		for _, ext := range JpgExtensions {
			output = strings.TrimSuffix(output, ext)
		}

		inputPath := filepath.Join(inputDir, entry.Name())
		outputPath := filepath.Join(outputDir, output)

		err = CompressJpg(
			quality,
			filepath.Join(inputPath),
			filepath.Join(outputPath),
		)
		if err != nil {
			log.Println("Error compressing jpeg:", err)
			continue
		}

		runtime.EventsEmit(a.ctx, "onCompressed", inputPath, outputPath)
	}
}
