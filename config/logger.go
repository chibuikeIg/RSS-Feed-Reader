package config

import (
	"os"
	"path/filepath"
)

func Log(s string) {

	fi, err := os.OpenFile(filepath.Join("./storage/", "go.log"), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

	if err != nil {

		panic(err)

	}

	// make a buffer to keep chunks that are read

	fi.WriteString(s + "\n====== NEW LOG START ======\n")

}
