/**
 * @Author: Seedzheng
 * 组合模式
 */

package composite

import "strings"

type FileNodeSystem interface {
    GetPath() string
    CountFilesNum() int
}

type File struct {
    path string
}

func (f *File) CountFilesNum() int {
    return 1
}

func (f *File) GetPath() string {
    return f.path
}

func NewFile(path string) *File {
    return &File{path: path}
}

type Directory struct {
    path     string
    subNodes []FileNodeSystem
}

func NewDirectory(path string) *Directory {
    return &Directory{
        path: path,
    }
}

func (node *Directory) CountFilesNum() int {
    fileNums := 0
    for _, subNode := range node.subNodes {
        fileNums += subNode.CountFilesNum()
    }

    return fileNums
}

func (node *Directory) GetPath() string {
    return node.path
}

func (node *Directory) AddSubNode(subNode ...FileNodeSystem) {
    if len(subNode) == 0 {
        return
    }

    node.subNodes = append(node.subNodes, subNode...)
}

func (node *Directory) RemoveSubNode(subNode FileNodeSystem) {
    nodeSize := len(node.subNodes)
    i := 0
    for ; i < nodeSize; i++ {
        currPath := strings.ToLower(node.subNodes[i].GetPath())
        subNodePath := strings.ToLower(subNode.GetPath())
        if currPath == subNodePath {
            break
        }
    }

    if i >= nodeSize {
        return
    }

    newSubNodes := node.subNodes[0:i]
    newSubNodes = append(newSubNodes, node.subNodes[i+1:]...)
    node.subNodes = newSubNodes
}
