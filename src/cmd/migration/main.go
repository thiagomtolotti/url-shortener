package main

import (
	"fmt"
	"os"
	"path"

	"urlshortener.com/src/infra"
)

func main() {
	if err := infra.LoadEnvironment(); err != nil {
		panic(err)
	}

	db, err := infra.ConnectToDB()
	if err != nil {
		panic(err)
	}

	files, err := readAllFilesInDir("migration")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		res, err := db.Exec(string(file))
		if err != nil {
			panic(err)
		}

		fmt.Println(res.RowsAffected())
	}
}

func readAllFilesInDir(p ...string) ([][]byte, error) {
	var res [][]byte
	pathString := path.Join(p...)

	dir, err := os.ReadDir(pathString)
	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		content, err := os.ReadFile(path.Join(pathString, file.Name()))
		if err != nil {
			return nil, err
		}

		res = append(res, content)
	}

	return res, nil
}
