package source

import (
	"embed"
	"io/fs"
)

// embedded/builtin is generated from the sibling iasi-core repository before every build.
//
//go:embed embedded/builtin
var embedded embed.FS

func Builtin() fs.FS {
	root, err := fs.Sub(embedded, "embedded/builtin")
	if err != nil {
		panic(err)
	}
	return root
}
