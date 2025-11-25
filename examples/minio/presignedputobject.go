package main

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
        "strings"
        "net/http"
)

func main() {
	endpoint := "www.xxx.yyy.zzz:port" // Replace with your MinIO endpoint

	accessKeyID := "accessKeyID"
	secretAccessKey := "secretAccessKey"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		fmt.Println("1", err)
		return
	}

	bucketName := "mybucket"
	objectName := "my-file.txt"
	expires := time.Second * 360 // URL valid for 60 seconds

	presignedURL, err := minioClient.PresignedPutObject(context.Background(), bucketName, objectName, expires)
	if err != nil {
		fmt.Println("2", err)
		return
	}

	fmt.Printf("Presigned URL for upload: %s\n", presignedURL)
	// Provide this URL to the client for uploading the file
	
	// 将 content 写入 presignedURL
	content := "Hello, MinIO!"

	// 上传 content 到 presignedURL
	// 创建PUT请求
	req, err := http.NewRequest(http.MethodPut, presignedURL.String(), strings.NewReader(content))
	if err != nil {
	    fmt.Println("创建PUT请求失败:", err)
	    return
	}
	
	// 设置Content-Type头
	req.Header.Set("Content-Type", "text/plain")
	
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
	    fmt.Println("发送PUT请求失败:", err)
	    return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("4", resp.Status)
	} else {
		fmt.Println("5", resp.Status)
	}
}
