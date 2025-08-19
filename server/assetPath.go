package server

import (
	"encoding/json"
	"io/fs"
	"os"
)

type ViteManifestEntry struct {
	File string `json:"file"`
	Name string `json:"name"`
	Src  string `json:"src"`
}

// this cache is used so that we don't have to parse the JSON manifest file
// every time we look something up in a prod build
var assetPathCache = map[string]string{}

func (s *server) assetPath(name string) string {
	// use the vite port to serve assets with hotreload when in dev
	if os.Getenv("APP_ENV") == "development" {
		return "http://localhost:5173/" + name
	}

	if path, ok := assetPathCache[name]; ok {
		return path
	}

	manifestBytes := []byte("")
	manifest := map[string]ViteManifestEntry{}

	if _, err := os.Stat("dist/.vite/manifest.json"); err == nil {
		manifestBytes, err := os.ReadFile("dist/.vite/manifest.json")
		if err != nil {
			return "ERROR: " + err.Error()
		}

		err = json.Unmarshal(manifestBytes, &manifest)
		if err != nil {
			return "ERROR: " + err.Error()
		}

		if _, ok := manifest[name]; !ok {
			return "ERROR: " + name + " missing from manifest file"
		}

		assetPathCache[name] = "/dist/" + manifest[name].File

		return assetPathCache[name]
	}

	if _, err := fs.Stat(s.distFS, "dist/.vite/manifest.json"); err == nil {
		manifestBytes, err := fs.ReadFile(s.distFS, "dist/.vite/manifest.json")
		if err != nil {
			return "ERROR: " + err.Error()
		}

		err = json.Unmarshal(manifestBytes, &manifest)
		if err != nil {
			return "ERROR: " + err.Error()
		}

		if _, ok := manifest[name]; !ok {
			return "ERROR: " + name + " missing from manifest file"
		}

		assetPathCache[name] = "/dist/" + manifest[name].File

		return assetPathCache[name]
	}

	return "UHOH!" + string(manifestBytes) + name
}
