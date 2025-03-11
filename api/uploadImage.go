package api

import (
	"fmt"
	"math/rand"
	"myWeb/common"
	"myWeb/config"
	"myWeb/context"
	"net/http"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 阿里云 OSS 配置
var (
	accessKeyID     = config.OSSfg.AccessKeyID     // 你的 AccessKey ID
	accessKeySecret = config.OSSfg.AccessKeySecret // 你的 AccessKey Secret
	endpoint        = config.OSSfg.Endpoint        // OSS Endpoint
	bucketName      = config.OSSfg.BucketName      // 你的 Bucket 名称
)

// uploadImage 处理文件上传到阿里云 OSS
func (*Api) UploadImage(w http.ResponseWriter, r *http.Request) {
	// 创建 OSS 客户端
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		http.Error(w, "无法创建 OSS 客户端", http.StatusInternalServerError)
		return
	}

	// 获取 OSS Bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		http.Error(w, "无法访问 OSS Bucket", http.StatusInternalServerError)
		return
	}

	// 解析请求中的表单数据
	err = r.ParseMultipartForm(10 << 20) // 限制最大文件大小为 10MB
	if err != nil {
		http.Error(w, "解析表单数据失败", http.StatusBadRequest)
		return
	}

	// 获取文件
	file, _, err := r.FormFile("file") // 获取表单中名为 "file" 的文件
	if err != nil {
		http.Error(w, "获取文件失败", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 获取文件的扩展名
	ext := ".jpg" // 默认使用 .jpg
	// 假设获取文件扩展名，这里可以根据实际情况修改
	if fileHeader, _, err := r.FormFile("file"); err == nil {
		ext = getFileExtension(fileHeader)
	}

	// 生成唯一的文件名，使用时间戳和文件扩展名
	fileName := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), generateRandomString(8)+ext)

	// 将文件上传到 OSS
	err = bucket.PutObject(fileName, file)
	if err != nil {
		http.Error(w, "文件上传到 OSS 失败", http.StatusInternalServerError)
		return
	}

	// 生成文件的 URL
	imageURL := fmt.Sprintf("https://%s.%s/%s", bucketName, endpoint, fileName)

	// 使用 SuccessResult 返回响应数据
	common.SuccessResult(w, map[string]string{"url": imageURL})
}

// 获取文件扩展名
func getFileExtension(file interface{}) string {
	// 根据实际需要从文件获取扩展名
	return ".jpg" // 示例返回 jpg 后缀
}

// 生成随机字符串，用于文件名
func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func (*Api) UploadImageNew(ctx *context.MsContext) {
	// 创建 OSS 客户端
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		http.Error(ctx.W, "无法创建 OSS 客户端", http.StatusInternalServerError)
		return
	}

	// 获取 OSS Bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		http.Error(ctx.W, "无法访问 OSS Bucket", http.StatusInternalServerError)
		return
	}

	// 解析请求中的表单数据
	err = ctx.Request.ParseMultipartForm(10 << 20) // 限制最大文件大小为 10MB
	if err != nil {
		http.Error(ctx.W, "解析表单数据失败", http.StatusBadRequest)
		return
	}

	// 获取文件
	file, _, err := ctx.Request.FormFile("file") // 获取表单中名为 "file" 的文件
	if err != nil {
		http.Error(ctx.W, "获取文件失败", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 获取文件的扩展名
	ext := ".jpg" // 默认使用 .jpg
	// 假设获取文件扩展名，这里可以根据实际情况修改
	if fileHeader, _, err := ctx.Request.FormFile("file"); err == nil {
		ext = getFileExtension(fileHeader)
	}

	// 生成唯一的文件名，使用时间戳和文件扩展名
	fileName := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), generateRandomString(8)+ext)

	// 将文件上传到 OSS
	err = bucket.PutObject(fileName, file)
	if err != nil {
		http.Error(ctx.W, "文件上传到 OSS 失败", http.StatusInternalServerError)
		return
	}

	// 生成文件的 URL
	imageURL := fmt.Sprintf("https://%s.%s/%s", bucketName, endpoint, fileName)

	// 使用 SuccessResult 返回响应数据
	common.SuccessResult(ctx.W, map[string]string{"url": imageURL})
}
