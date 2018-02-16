package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/essentialbooks/books/pkg/common"
	"github.com/kjk/u"
)

func gitRename(dst, src string) (string, error) {
	cmd := exec.Command("git", "mv", src, dst)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// git rename sometimes fails with "fatal: Unable to create '*/.git/index.lock': File exists." error
// retry in this case
// it might be caused by vscode doing git operations at the same time
func gitRenameRetryMust(dst, src string) {
	defer fmt.Printf("git mv %s => %s\n", src, dst)
	out, err := gitRename(dst, src)
	if err == nil {
		return
	}
	if strings.Contains(out, "index.lock") {
		time.Sleep(time.Millisecond * 500)
		out, err = gitRename(dst, src)
	}
	if err == nil {
		return
	}
	fmt.Printf("'git mv %s %s' failed with err '%s'. Output:\n%s\n", src, dst, err, out)
	os.Exit(1)
}

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

func sortNamesByNo(names []string, prec int) []*FileRenameInfo {
	var ri []*FileRenameInfo
	for _, name := range names {
		i := &FileRenameInfo{
			Name: name,
		}
		i.No, i.NameRest = getNoFromName(i.Name)
		ri = append(ri, i)
	}

	sort.Slice(ri, func(i, j int) bool {
		return ri[i].No < ri[j].No
	})
	for i, info := range ri {
		n := 10 * (i + 1)
		fmtStr := fmt.Sprintf("%%0%dd-%%s", prec)
		info.NewName = fmt.Sprintf(fmtStr, n, info.NameRest)
	}
	return ri
}

func renameFilesInChapter(chapterDir string) error {
	files, err := getMdFiles(chapterDir)
	if err != nil {
		return err
	}
	var names []string
	for _, file := range files {
		if file == "000-index.md" {
			continue
		}
		names = append(names, file)
	}

	ri := sortNamesByNo(names, 3)
	for _, info := range ri {
		if info.NewName != info.Name {
			src := filepath.Join(chapterDir, info.Name)
			dst := filepath.Join(chapterDir, info.NewName)
			gitRenameRetryMust(dst, src)
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

	ri := sortNamesByNo(chapterDirs, 4)
	for _, info := range ri {
		if info.NewName != info.Name {
			src := filepath.Join(bookDir, info.Name)
			dst := filepath.Join(bookDir, info.NewName)
			gitRenameRetryMust(dst, src)
		}
	}
	return nil
}

func renameBook(book string) {
	bookDir := filepath.Join("books", book)
	chapters, err := common.GetDirs(bookDir)
	u.PanicIfErr(err)
	renameChapters(bookDir, chapters)
}

// TODO: remove as this is likely one shot thing
func renameIndexFilesAndExit() {
	dir := filepath.Join("books", "go")
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		name := filepath.Base(path)
		if name != "index.md" {
			return nil
		}
		dir := filepath.Dir(path)
		newPath := filepath.Join(dir, "000-"+name)
		out, err := gitRename(newPath, path)
		u.PanicIfErr(err, "gitRename failed with '%s', output: '%s'", err, out)
		return nil
	})
	os.Exit(0)
}

func main() {
	books, err := common.GetDirs("books")
	u.PanicIfErr(err)
	for _, book := range books {
		renameBook(book)
	}
}
