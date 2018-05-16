package main

import (
	"os"
	"path"
	"testing"

	"github.com/ingmardrewing/fs"
	store "github.com/ingmardrewing/fsKeyValueStore"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {
	store.Initialize(path.Join(fs.Pwd(), "testResources/db"))
}

func tearDown() {
	dir := path.Join(fs.Pwd(), "testResources/db")
	fs.RemoveDirContents(dir)
}

func TestExtractFBAccessToken(t *testing.T) {
	stringifiedResponse := `"access_token":"the-token-123"`

	expected := "the-token-123"
	actual := extractFBAccessToken(stringifiedResponse)

	if actual != expected {
		t.Error("Expected extracted token to be", expected, "but it was", actual)
	}
}

func TestStoreFBAccessToken(t *testing.T) {
	storeFBAccessToken("drewingde", "token-value")

	pth := path.Join(fs.Pwd(), "testResources/db/record/drewingde_fb_token.json")
	exists, _ := fs.PathExists(pth)

	if !exists {
		t.Error("Expected to find json file at", pth)
	}
}

func TestRetrieveTokenFor(t *testing.T) {
	actual := retrieveTokenFor("drewingde")
	expected := "token-value"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
