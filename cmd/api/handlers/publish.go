package handlers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"kitexdousheng/cmd/api/rpc"
	"kitexdousheng/kitex_gen/publish"
	"kitexdousheng/pkg/errno"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func PublishAction(c *gin.Context) {
	uid, _ := c.Get("uid")
	userId := uid.(int64)

	data, err := c.FormFile("data")
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	filename := filepath.Base(data.Filename) //返回路径最后的文件名
	finalName := fmt.Sprintf("%d_%d_%s", time.Now().UnixNano(), userId, filename)
	b := []byte(finalName)
	imgName := string(b[:len(b)-3]) + "jpg"
	// *********************************************************上传COS存储桶**********************************************
	err = uploadVideo(data, finalName)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	coverData, err := readFrameAsJpeg(viper.GetString("cos.uriVideoPath") + finalName)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	err = uploadImg(coverData, imgName)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// *********************************************************上传COS存储桶**********************************************

	resp, err := rpc.PublishAction(c, &publish.DouyinPublishActionRequest{
		UserId: userId,
		Title:  finalName,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func PublishList(c *gin.Context) {
	uid, _ := c.Get("uid")
	userId := uid.(int64)

	resp, err := rpc.PublishList(c, &publish.DouyinPublishListRequest{
		UserId: userId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func uploadVideo(data *multipart.FileHeader, finalName string) error {
	fd, err := data.Open()
	if err != nil {
		fmt.Println("error")
		return err
	}

	u, _ := url.Parse(viper.GetString("cos.url"))
	b := &cos.BaseURL{BucketURL: u}
	cosClient := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  viper.GetString("cos.SecretID"),  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: viper.GetString("cos.SecretKey"), // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	name := viper.GetString("cos.videoDir") + finalName

	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = cosClient.Object.Put(context.Background(), name, fd, nil)
	if err != nil {
		return err
	}
	return nil
}

func uploadImg(data []byte, finalName string) error {
	f := bytes.NewReader(data)

	u, _ := url.Parse(viper.GetString("cos.url"))
	b := &cos.BaseURL{BucketURL: u}
	cosClient := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  viper.GetString("cos.SecretID"),  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: viper.GetString("cos.SecretKey"), // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	name := viper.GetString("cos.imgDir") + finalName

	_, err := cosClient.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		return err
	}
	return nil
}

func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), nil
}
