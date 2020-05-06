package ch5

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type ConvertType string

const (
	TypeImage ConvertType = "Image"
	TypeAudio ConvertType = "Audio"
)

var fromFlag = flag.String("from", ".png", "変換したいファイルの拡張子")
var toFlag = flag.String("to", ".jpg", "変換後のファイルの拡張子")

func Convert() error {
	flag.Parse()

	var err error

	for _, v := range flag.Args() {
		object := NewObject(v)

		err = object.Transform(*fromFlag, *toFlag)
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

type Object struct {
	ConvertType
	Path     string
	Name     string
	FullPath string
}

func NewObject(path string) *Object {
	dir, name := filepath.Split(path)

	object := new(Object)
	object.Path = dir
	object.FullPath = path
	object.Name = name
	object.ConvertType = TypeImage

	return object
}

func (object *Object) Transform(from, to string) error {
	var (
		fileInfo os.FileInfo
		err      error
	)

	fileInfo, err = os.Stat(object.FullPath)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		dir := NewDirectory(object)
		err = dir.Transform(from, to)
	} else {
		file := NewFile(object)
		err = file.Transform(from, to)
	}
	if err != nil {
		return err
	}

	return nil
}

type Directory struct {
	*Object
}

func NewDirectory(object *Object) *Directory {
	return &Directory{object}
}

func (directory *Directory) Transform(from, to string) error {
	var err error

	err = filepath.Walk(directory.FullPath, func(path string, info os.FileInfo, err error) error {

		// ディレクトリの場合はスキップ
		if info.IsDir() {
			return nil
		}

		file := NewFile(NewObject(path))

		err = file.Transform(from, to)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

type File struct {
	*Object
	Ext string
}

func (file *File) Transform(from, to string) error {
	var err error

	// ファイルの拡張子がターゲットではない場合スキップ
	if file.Ext != from {
		return nil
	}

	// 変換を行う
	switch file.ConvertType {
	case TypeImage:
		img := NewImage(file)
		err = img.Transform(from, to)
	case TypeAudio:
		audio := NewAudio(file)
		err = audio.Transform(from, to)
	}

	if err != nil {
		return err
	}

	return nil
}

func (file *File) IsImage() bool {
	return file.IsMatchExt(`.(png|jpg|jpeg|gif)$`)
}

func (file *File) IsMatchExt(pattern string) bool {
	return regexp.MustCompile(pattern).Match([]byte(file.Ext))
}

type Audio struct {
	*File
}

func (audio *Audio) Transform(from, to string) error {
	log.Fatal("実装の予定は無いけどポリモーフィズムを意識するために定義しているよ")

	return nil
}

type Image struct {
	*File
}

func (img *Image) Transform(from, to string) error {
	// 対象の拡張子でない場合スキップ
	if from != img.Ext {
		return nil
	}

	var (
		originFile  *os.File
		originImage image.Image
		newFile     *os.File
		err         error
	)

	originFile, err = os.Open(img.FullPath)
	if err != nil {
		return err
	}
	defer originFile.Close()

	originImage, _, err = image.Decode(originFile)
	if err != nil {
		return err
	}

	newFile, err = os.Create(fmt.Sprintf("%s/%s%s", img.Path, img.Name, to))
	if err != nil {
		return err
	}
	defer newFile.Close()

	switch to {
	case ".png":
		err = png.Encode(newFile, originImage)
	case ".jpeg":
		err = jpeg.Encode(newFile, originImage, nil)
	case ".jpg":
		err = jpeg.Encode(newFile, originImage, nil)
	case ".gif":
		err = gif.Encode(newFile, originImage, nil)
	}
	if err != nil {
		return err
	}

	return nil
}

func NewFile(object *Object) *File {
	file := File{
		Object: object,
		Ext:    filepath.Ext(object.FullPath),
	}

	return &file
}

func NewImage(file *File) *Image {
	return &Image{file}
}

func NewAudio(file *File) *Audio {
	return &Audio{file}
}

type Target interface {
	Transform(from, to string) error
}
