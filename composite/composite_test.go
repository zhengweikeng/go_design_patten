/**
 * @Author: Seedzheng
 */

package composite

import (
	"fmt"
	"testing"
)

func TestCountFilesNum(t *testing.T) {
	CountFilesNum(t)
}

func CountFilesNum(t *testing.T) {
	root := NewDirectory("/")
	dir1 := NewDirectory("/a")
	dir2 := NewDirectory("/b")
	root.AddSubNode(dir1, dir2)

	file1 := NewFile("/a/demo1.md")
	file2 := NewFile("/a/demo2.md")
	dir1SubDir := NewDirectory("/a/aSub1")
	dir1.AddSubNode(file1, file2, dir1SubDir)

	fmt.Println(root.CountFilesNum())
}
