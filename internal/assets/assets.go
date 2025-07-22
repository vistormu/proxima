package assets

import (
	"embed"
)

//go:embed proxima.toml
var DefaultConfig embed.FS
