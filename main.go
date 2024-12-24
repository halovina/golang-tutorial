package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

func uploadFile(ctx context.Context, client *storage.Client, bucketName, objectName, filePath string) error {
	// Buka file yang akan upload.
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("os.Open: %v", err)
	}
	defer f.Close()

	wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	fmt.Printf("File %v uploaded to %v.\n", filePath, objectName)
	return nil
}

func main() {
	ctx := context.Background()

	//path key google service account.
	credentialsFile := "sac/key.json"
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credentialsFile)

	// create client Cloud Storage.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucketName := "testuploadfile123"
	objectName := "testgolang/pinguin-testgo.jpeg"
	filePath := "pinguin.jpeg"

	if err := uploadFile(ctx, client, bucketName, objectName, filePath); err != nil {
		log.Fatal(err)
	}
}
