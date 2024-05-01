package handlers

import (
	"os"
	"net/http"
)

func VersionEndpoint(w http.ResponseWriter, r *http.Request) {
	version := getVersion()
	w.Write(version)
}

func getVersion() []byte {
	versionFile, err := os.ReadFile("./VersionArtifact")
	if err != nil {
		panic(err)
	}
	return versionFile
}