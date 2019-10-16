package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := ioutil.TempFile("", "sample")
	check(err)
	defer os.Remove(f.Name())
	fmt.Println("Temp file name:", f.Name())
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)
	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	defer os.RemoveAll(dname)
	fmt.Println("Temp dir name:", dname)
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
