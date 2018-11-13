package main

import (
	"fmt"
	"github.com/tus/tusd"
	"github.com/tus/tusd/filestore"
	"net/http"
)

func main() {
	// Create a new FileStore instance which is responsible for
	// storing the uploaded file on disk in the specified directory.
	// This path _must_ exist before tusd will store uploads in it.
	// If you want to save them on a different medium, for example
	// a remote FTP server, you can implement your own storage backend
	// by implementing the tusd.DataStore interface.
	store := filestore.FileStore{
		Path: "./uploads",
	}

	// A storage backend for tusd may consist of multiple different parts which
	// handle upload creation, locking, termination and so on. The composer is a
	// place where all those seperated pieces are joined together. In this example
	// we only use the file store but you may plug in multiple.
	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	// Create a new HTTP handler for the tusd server by providing a configuration.
	// The StoreComposer property must be set to allow the handler to function.
	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:              "/files/",
		StoreComposer:         composer,
		NotifyCreatedUploads:  true,
		NotifyCompleteUploads: true,
	})

	handler.CreatedUploads = make(chan tusd.FileInfo, 1)
	handler.CompleteUploads = make(chan tusd.FileInfo, 1)

	//handler, err := tusd.NewUnroutedHandler(tusd.Config{
	//	BasePath:      "/files/",
	//	StoreComposer: composer,
	//})

	if err != nil {
		panic(fmt.Errorf("Unable to create handler: %s", err))
	}

	// Right now, nothing has happened since we need to start the HTTP server on
	// our own. In the end, tusd will start listening on and accept request at
	// http://localhost:8080/files
	http.Handle("/files/", http.StripPrefix("/files/", handler))

	go createUpload(handler.CreatedUploads)
	go complete(handler.CompleteUploads)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(fmt.Errorf("Unable to listen: %s", err))
	}

	select {}
}

func createUpload(file chan tusd.FileInfo) {
	for {
		fileInfo := <-file
		fmt.Println("createUpload,fileInfo: ", fileInfo)
	}
}

func complete(file chan tusd.FileInfo) {
	for {
		fileInfo := <-file

		fmt.Println("fileInfo: ", fileInfo)
	}
}
