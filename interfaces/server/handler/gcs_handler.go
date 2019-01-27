package handler

import (
	"cloud.google.com/go/storage"
	"context"
	"io"
	"io/ioutil"
	"os"
)

var BUCKET *storage.BucketHandle
var CTX context.Context

// InitGCS is initialing the google cloud storage server
func InitGCS() error {
	bucketName := os.Getenv("GOOGLE_CLOUD_STORAGE_BUCKET_NAME")
	CTX = context.Background()
	GCD, err := storage.NewClient(CTX)
	if err != nil {
		return err
	}
	BUCKET = GCD.Bucket(bucketName)
	return nil
}

// SendImage is
func SendImage(b []byte, id string, env string) error {
	fileName := id + ".jpg"
	// Create
	file, err := os.OpenFile("/tmp"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	// Write
	err = ioutil.WriteFile("/tmp"+fileName, b, 0666)
	if err != nil {
		return err
	}
	// Save to Google Cloud Storage
	PATH := os.Getenv(env) + fileName
	wc := BUCKET.Object(PATH).NewWriter(CTX)
	// Copy
	if _, err = io.Copy(wc, file); err != nil {
		return err
	}
	// Upload
	if err = wc.Close(); err != nil {
		return err
	}
	// Change Private to Public Access
	acl := BUCKET.Object(PATH).ACL()
	if err = acl.Set(CTX, storage.AllUsers, storage.RoleReader); err != nil {
		return err
	}
	if err := os.Remove("/tmp" + fileName); err != nil {
		return err
	}
	return nil
}
