// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wabook

import (
	"path/filepath"
)

type Book struct {
	Root    string
	Info    *BookToml
	Summary *Summary
	Talks   []string
}

func LoadBook(path string) (book *Book, err error) {
	book = &Book{Root: path}
	book.Info, err = LoadConfig(filepath.Join(path, "book.ini"))
	if err != nil {
		if info, errx := LoadConfig(filepath.Join(path, "book.toml")); errx != nil {
			return nil, err
		} else {
			book.Info = info
		}
	}
	book.Summary, err = LoadSummary(filepath.Join(path, "SUMMARY.md"))
	if err != nil {
		return nil, err
	}
	book.Talks = loadTalks(book.Root)
	return
}
