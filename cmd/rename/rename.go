package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/essentialbooks/books/pkg/common"
	"github.com/kjk/u"
)

func getMdFiles(dir string) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, fi := range fileInfos {
		if fi.IsDir() || !fi.Mode().IsRegular() {
			continue
		}
		name := fi.Name()
		if !strings.HasSuffix(name, ".md") {
			continue
		}
		res = append(res, name)
	}
	return res, nil
}

// FileRenameInfo keeps info about renaming operation
type FileRenameInfo struct {
	Name     string
	No       int // extracted from name
	NameRest string
	NewName  string
}

// break name '0010-foo.md' into (10, 'foo.md')
func getNoFromName(name string) (int, string) {
	parts := strings.SplitN(name, "-", 2)
	u.PanicIf(len(parts) != 2, "invalid name '%s'", name)
	n, err := strconv.Atoi(parts[0])
	u.PanicIfErr(err, "invalid name '%s'", name)
	return n, parts[1]
}

func renameFilesInChapter(chapterDir string) error {
	files, err := getMdFiles(chapterDir)
	if err != nil {
		return err
	}
	var ri []*FileRenameInfo
	for _, file := range files {
		if file == "index.md" {
			continue
		}
		i := &FileRenameInfo{
			Name: file,
		}
		i.No, i.NameRest = getNoFromName(i.Name)
		ri = append(ri, i)
	}

	sort.Slice(ri, func(i, j int) bool {
		return ri[i].No < ri[j].No
	})
	for i, info := range ri {
		n := 10 * (i + 1)
		info.NewName = fmt.Sprintf("%03d-%s", n, info.NameRest)
	}
	for _, info := range ri {
		if info.NewName != info.Name {
			fmt.Printf("Renaming %s = > %s\n", info.Name, info.NewName)
		}
	}
	return nil
}

func renameChapters(bookDir string, chapterDirs []string) error {
	for _, dir := range chapterDirs {
		chapterDir := filepath.Join(bookDir, dir)
		err := renameFilesInChapter(chapterDir)
		u.PanicIfErr(err)
	}
	// TODO: rename chapters
	return nil
}

func renameBook(book string) {
	bookDir := filepath.Join("books", book)
	chapters, err := common.GetDirs(bookDir)
	u.PanicIfErr(err)
	renameChapters(bookDir, chapters)
}

func main() {
	books, err := common.GetDirs("books")
	u.PanicIfErr(err)
	for _, book := range books {
		renameBook(book)
	}
}
