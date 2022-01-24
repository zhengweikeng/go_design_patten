/**
 * @Author: Seedzheng
 */

package proxy

import "testing"

func TestImageUploaderProxy_Upload(t *testing.T) {
    tests := []struct {
        name string
    }{
        {
            name: "proxy patten",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            imageUploader := ImageUploader{Url: "image.test.com"}
            u := NewImageUploaderProxy(imageUploader)
            u.Upload()
        })
    }
}
