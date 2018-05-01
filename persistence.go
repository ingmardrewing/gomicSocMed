package main

import (
	"errors"
	"os"

	scribble "github.com/nanobox-io/golang-scribble"
)

type Persistence interface {
	CreateIfNonExistentElseUpdate(key, value string) error
	Create(key string, value string) error
	Update(key string, value string) error
	Delete(key string) error
	Read(key string) (string, error)
}

func NewFilePersistence(dir string) (Persistence, error) {
	if isExistent, _ := pathExists(dir); isExistent != true {
		return nil, errors.New("Given dir for scribble db doesn't exist")
	}

	db, err := scribble.New(dir, nil)
	if err != nil {
		return nil, err
	}

	fp := new(filePersistence)
	fp.dir = dir
	fp.db = db
	return fp, nil
}

type Record struct {
	Value string
}

type filePersistence struct {
	dir string
	db  *scribble.Driver
}

func (fp *filePersistence) CreateIfNonExistentElseUpdate(key, value string) error {
	if fp.exists(key) {
		if err := fp.Update(key, value); err != nil {
			return err
		}
		return nil
	}
	if err := fp.Create(key, value); err != nil {
		return err
	}
	return nil
}

func (fp *filePersistence) Create(key, value string) error {
	if fp.exists(key) {
		return errors.New("Already exists")
	}
	if key == "" {
		return errors.New("Empty string given as key")
	}
	if err := fp.db.Write("record", key, Record{Value: value}); err != nil {
		return err
	}
	return nil
}

func (fp *filePersistence) Read(key string) (string, error) {
	if key == "" {
		return "", errors.New("Emtpy string given as key")
	}
	r := Record{}
	if err := fp.db.Read("record", key, &r); err != nil {
		return "", err
	}
	return r.Value, nil
}

func (fp *filePersistence) Update(key, value string) error {
	if !fp.exists(key) {
		return errors.New("Can't update, key doesn't exist")
	}
	if key == "" {
		return errors.New("Empty string given as key")
	}
	if err := fp.db.Write("record", key, Record{Value: value}); err != nil {
		return err
	}
	return nil
}

func (fp *filePersistence) Delete(key string) error {
	if !fp.exists(key) {
		return errors.New("Can't delete, key doesn't exist")
	}
	if key == "" {
		return errors.New("Empty string given as key")
	}
	if err := fp.db.Delete("record", key); err != nil {
		return err
	}
	return nil
}

func (fp *filePersistence) exists(key string) bool {
	_, err := fp.Read(key)
	if err != nil {
		return false
	}
	return true
}

func pathExists(pth string) (bool, error) {
	_, err := os.Stat(pth)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
