package handler

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// var BUCKET *storage.BucketHandle
// var CTX context.Context

// InitGCS is initialing the google cloud storage server
// func InitGCS() error {
// 	bucketName := os.Getenv("GOOGLE_CLOUD_STORAGE_BUCKET_NAME")
// 	CTX = context.Background()
// 	GCD, err := storage.NewClient(CTX)
// 	if err != nil {
// 		return err
// 	}
// 	BUCKET = GCD.Bucket(bucketName)
// 	if BUCKET == nil {
// 		fmt.Println("BUCKET ready")
// 	}
// 	return nil
// }

// SendImage is
func SendImage(b []byte, id string, env string) error {
	bucketName := os.Getenv("GOOGLE_CLOUD_STORAGE_BUCKET_NAME")
	ctx := context.Background()
	gcsClient, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	bucket := gcsClient.Bucket(bucketName)
	if bucket == nil {
		fmt.Println("BUCKET ready")
	}
	fileName := id + ".jpg"
	fmt.Println(fileName)
	// Create
	path := os.Getenv("FILE_TMP")
	fmt.Println(path)
	file, err := os.OpenFile(path+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	// Write
	err = ioutil.WriteFile(path+fileName, b, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Save to Google Cloud Storage
	PATH := os.Getenv(env) + fileName
	fmt.Println(PATH)
	wc := bucket.Object(PATH).NewWriter(ctx)
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	// Copy
	if _, err = io.Copy(wc, file); err != nil {
		fmt.Println(err)
		return err
	}
	// Upload
	if err = wc.Close(); err != nil {
		fmt.Println(err)
		return err
	}
	// Change Private to Public Access
	// acl := BUCKET.Object(PATH).ACL()
	// if err = acl.Set(CTX, storage.AllUsers, storage.RoleReader); err != nil {
	// 	return err
	// }
	// if err := os.Remove(path + fileName); err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	gcsClient.Close()
	return nil
}
