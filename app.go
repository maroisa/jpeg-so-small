package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

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

	return dirpath
}
