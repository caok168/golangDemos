package main

import (
	"fmt"
	"github.com/linxGnu/goseaweedfs"
	"os"
	"time"
)

var sw *goseaweedfs.Seaweed

var MediumFile, SmallFile string

func init() {
	// check master url
	if masterURL := os.Getenv("GOSWFS_MASTER_URL"); masterURL != "" {
		scheme := os.Getenv("GOSWFS_SCHEME")
		if scheme == "" {
			scheme = "http"
		}

		var filer []string
		if _filer := os.Getenv("GOSWFS_FILER_URL"); _filer != "" {
			filer = []string{_filer}
		}

		sw = goseaweedfs.NewSeaweed(scheme, masterURL, filer, 2*1024*1024, 5*time.Minute)
	}

	MediumFile = os.Getenv("GOSWFS_MEDIUM_FILE")
	SmallFile = os.Getenv("GOSWFS_SMALL_FILE")

	time.Sleep(10 * time.Second)
}

func main() {
	test1()
}

func test1() {
	if sw == nil || MediumFile == "" {
		return
	}

	for i := 1; i <= 2; i++ {
		_, _, fID, err := sw.UploadFile(MediumFile, "", "")
		if err != nil {
			fmt.Println(err)
		}

		//
		if _, err := sw.LookupServerByFileID(fID, nil, true); err != nil {
			fmt.Println(err)
		}

		//
		if fullURL, err := sw.LookupFileID(fID, nil, true); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fullURL:", fullURL)
		}

		//
		if err := sw.ReplaceFile(fID, SmallFile, false); err != nil {
			fmt.Println(err)
			return
		}

		//
		if err := sw.ReplaceFile(fID, SmallFile, true); err != nil {
			fmt.Println(err)
			return
		}

		if err = sw.DeleteFile(fID, nil); err != nil {
			fmt.Println(err)
			return
		}

		// test upload file
		fh, err := os.Open(MediumFile)
		if err != nil {
			fmt.Println(err)
		}
		defer fh.Close()

		var size int64
		if fi, fiErr := fh.Stat(); fiErr != nil {
			fmt.Println(fiErr)
		} else {
			size = fi.Size()
		}

		if _, fID, err = sw.Upload(fh, "test.txt", size, "col", ""); err != nil {
			fmt.Println(err)
		}

		// Replace with small file
		fs, err := os.Open(SmallFile)
		if err != nil {
			fmt.Println(err)
		}
		defer fs.Close()
		if fi, fiErr := fs.Stat(); fiErr != nil {
			fmt.Println(fiErr)
		} else {
			size = fi.Size()
		}

		if err := sw.Replace(fID, fs, "ta.txt", size, "", "", false); err != nil {
			fmt.Println(err)
			return
		}

		// finally delete
		if err = sw.DeleteFile(fID, nil); err != nil {
			fmt.Println(err)
		}
	}
}
