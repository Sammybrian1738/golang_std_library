package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type myCloser interface {
	Close() error
}

// closeFile is a helper function which streamlines closing
// with error checking on different file types.
func closeFile(f myCloser) {
	err := f.Close()
	check(err)
}

// check is a helper function which streamlines error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readAll(file *zip.File) []byte {
	fc, err := file.Open()
	check(err)
	defer closeFile(fc)

	content, err := ioutil.ReadAll(fc)
	check(err)

	return content
}

func Zip() {
	// creates the archive file
	archive, err := os.Create("./std_pkg/archive/zip/archive.zip")
	if err != nil {
		log.Fatalln(err)
	}

	defer archive.Close()

	zipWriter := zip.NewWriter(archive)

	// opens the file in the root directory. file must exist
	fmt.Println("opening first file...")
	f1, err := os.Open("./std_pkg/archive/zip/test.csv")
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	fmt.Println("writing first file to archive...")
	w1, err := zipWriter.Create("csv/text.csv")
	if err != nil {
		panic(err)
	}

	// copies the contents of the file  you opened into the file in the archive
	if _, err := io.Copy(w1, f1); err != nil {
		panic(err)
	}

	// opens the file in the root directory. file must exist
	fmt.Println("opening second file")
	f2, err := os.Open("./std_pkg/archive/zip/test.txt")
	if err != nil {
		panic(err)
	}

	defer f2.Close()

	fmt.Println("writing second file to archive")
	w2, err := zipWriter.Create("txt/test.txt")
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(w2, f2); err != nil {
		panic(err)
	}

	fmt.Println("closing zip archive...")
	zipWriter.Close()

	// zip reader
	fmt.Println("reading the files in the archive")
	zf, err := zip.OpenReader("./std_pkg/archive/zip/archive.zip")
	check(err)
	defer closeFile(zf)

	for _, file := range zf.File {
		fmt.Printf("=%s\n", file.Name)
		// fmt.Printf("%s\n\n", readAll(file))

	}
}
