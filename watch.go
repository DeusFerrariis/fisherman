package main

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	gitignore "github.com/monochromegane/go-gitignore"
)

type (
	WatchDirectory struct {
		Command      string
		Path         string
		LastChecked  time.Time
		SleepSeconds int
	}
)

func (w *WatchDirectory) CheckLastModified() (bool, error) {
	return RecursivelyCheckLastModified(w.LastChecked, w.Path)
}

func (w *WatchDirectory) Watch(onChange func()) error {
	for {
		changed, err := w.CheckLastModified()

		if err != nil {
			return err
		}

		if changed == true {
			bgBlue := color.New(color.BgBlue).Add(color.FgBlack)
			bgBlue.Println("- CHANGE DETECTED -")
			w.LastChecked = time.Now()
			onChange()
		}

		// TODO: make sleep time configurable
		time.Sleep(time.Duration(w.SleepSeconds) * time.Second)
	}
}

func RecursivelyCheckLastModified(since time.Time, path string) (bool, error) {
	dir, err := os.ReadDir(path)

	if err != nil {
		return false, errors.Join(err, errors.New("could not read directory/file: "+path))
	}

	for _, file := range dir {
		info, _ := file.Info()
		fAbsPath := filepath.Join(path, info.Name())

		ignoreFile, err := gitignore.NewGitIgnore("./.fmignore")

		if err != nil {
			return false, err
		}

		// Recurse on directory
		if file.IsDir() {
			if ignoreFile.Match(fAbsPath, true) {
				continue
			}

			modified, err := RecursivelyCheckLastModified(since, fAbsPath)

			if err != nil {
				return false, err
			}

			if modified {
				return true, nil
			}

			continue
		}

		if ignoreFile.Match(fAbsPath, false) {
			continue
		}

		if info.ModTime().After(since) {
			return true, nil
		}
	}

	return false, nil
}
