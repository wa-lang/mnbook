// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed _static
var _staticFS embed.FS

func (p *BookRendor) renderStaticFile() error {
	if p.Book.Info.Giscus.Enabled {
		dst := filepath.Join(p.Book.Root, "book/static/wabook/giscus.js")
		js, err := p.genGiscusJs()
		if err != nil {
			panic(err)
		}

		os.MkdirAll(filepath.Dir(dst), 0777)
		err = os.WriteFile(dst, []byte(js), 0666)
		if err != nil {
			panic(err)
		}
	}

	return cpDir(filepath.Join(p.Book.Root, "book"), p.Book.Root, p.Ignores)
}

func (p *BookRendor) renderStaticAsset() error {
	staticFS, err := fs.Sub(_staticFS, "_static")
	if err != nil {
		return err
	}
	err = fs.WalkDir(staticFS, ".", func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}

		data, err := fs.ReadFile(staticFS, path)
		if err != nil {
			return err
		}

		dstpath := filepath.Join(p.Book.Root, "book", path)
		os.MkdirAll(filepath.Dir(dstpath), 0777)

		err = os.WriteFile(dstpath, data, 0666)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
