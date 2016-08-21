package data

import (
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

//HTTPAsset retruns http.FileSystem that can be used to serve the data as
//sstatic assets.
func HTTPAsset() http.FileSystem {
	return &assetfs.AssetFS{
		Asset:    Asset,
		AssetDir: AssetDir,
	}
}
