/**
 * @Author: Seedzheng
 * 代理模式
 */

package proxy

import "fmt"

type Uploader interface {
    Upload()
}

type ImageUploader struct {
    Url string
}

func (u *ImageUploader) Upload() {
    fmt.Println(fmt.Sprintf("正在上传图片到 %s", u.Url))
}

type ImageUploaderProxy struct {
    uploader ImageUploader
}

func (u *ImageUploaderProxy) Upload() {
    fmt.Println(fmt.Sprintf("准备上传图片到 %s", u.uploader.Url))
    u.uploader.Upload()
    fmt.Println(fmt.Sprintf("上传图片到 %s 完成", u.uploader.Url))
}

func NewImageUploaderProxy(uploader ImageUploader) *ImageUploaderProxy {
    return &ImageUploaderProxy{uploader: uploader}
}
