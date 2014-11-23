package main

import (
	"os"
	"io"
	"os/user"
	"log"
	"io/ioutil"
)

func WriteToFile(file string, body []byte) {
	err := ioutil.WriteFile(file, body, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}
	}

	return
}

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)
	sep := string(os.PathSeparator)

	for _, obj := range objects {

		sourcefilepointer := source + sep + obj.Name()
		destinationfilepointer := dest + sep + obj.Name()

		if obj.IsDir() {
			performCopy(sourcefilepointer, destinationfilepointer)
		} else {
			performCopy(sourcefilepointer, destinationfilepointer)
		}
	}
	return
}

func performCopy(sourcefilepointer string, destinationfilepointer string) {
	err := CopyFile(sourcefilepointer, destinationfilepointer)
	if err != nil {
		log.Fatal(err)
	}
}

func IsExist(filename string) bool {
	_, err := os.Stat(filename);
	return !os.IsNotExist(err)
}

func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
