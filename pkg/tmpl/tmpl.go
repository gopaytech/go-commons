package tmpl

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"text/template"
)

type FileFilter func(path string, info os.FileInfo) bool

type FileDetail struct {
	Info       os.FileInfo
	TargetPath string
	IsTemplate bool
}

type FileMap map[string]FileDetail

type ScanResult struct {
	RootPath    string
	Extension   string
	FilterName  string
	Template    *template.Template
	TemplateMap FileMap
}

func (result *ScanResult) TemplateList() (list []string) {
	list = []string{}
	for key, value := range result.TemplateMap {
		if value.IsTemplate {
			list = append(list, key)
		}
	}
	return
}

// DirList result should be sorted
func (result *ScanResult) DirList() (list []string) {
	list = []string{}
	for key, value := range result.TemplateMap {
		if value.Info.IsDir() {
			list = append(list, key)
		}
	}
	sort.Strings(list)
	return
}

func (result *ScanResult) ExecuteToPath(data interface{}, targetPath string) (err error) {

	// add / to targetPath if necessary
	if !strings.HasSuffix(targetPath, string(filepath.Separator)) {
		targetPath = targetPath + string(filepath.Separator)
	}

	err = os.MkdirAll(targetPath, os.ModePerm)
	if err != nil {
		return
	}

	// create all directory first
	for _, value := range result.DirList() {
		ierr := os.MkdirAll(targetPath+value, os.ModePerm)
		if ierr != nil {
			err = ierr
			return
		}
	}

	for key, value := range result.TemplateMap {
		if value.IsTemplate {
			var buff bytes.Buffer
			ierr := result.Template.ExecuteTemplate(&buff, key, data)
			if ierr != nil {
				err = ierr
				return
			}

			targetFile := targetPath + value.TargetPath
			ierr = ioutil.WriteFile(targetFile, buff.Bytes(), 0644)
			if ierr != nil {
				err = ierr
				return
			}
		} else {
			absoluteSourceFile := result.RootPath + key
			absoluteDestinationFile := targetPath + key

			stat, ierr := os.Stat(absoluteSourceFile)
			if ierr != nil {
				err = ierr
				return
			}

			if !stat.IsDir() {
				source, ierr := ioutil.ReadFile(absoluteSourceFile)
				if ierr != nil {
					err = ierr
					return
				}

				ierr = ioutil.WriteFile(absoluteDestinationFile, source, 0644)
				if ierr != nil {
					err = ierr
					return
				}
			}
		}
	}
	return
}

func (result *ScanResult) Execute(data interface{}) (mapResult map[string]string, err error) {
	mapResult = map[string]string{}
	for key, value := range result.TemplateMap {
		if value.IsTemplate {
			var buff bytes.Buffer
			ierr := result.Template.ExecuteTemplate(&buff, key, data)
			if ierr != nil {
				err = ierr
				return
			}

			// use target path as key
			mapResult[value.TargetPath] = buff.String()
		} else {
			absolutePath := result.RootPath + key
			stat, ierr := os.Stat(absolutePath)
			if ierr != nil {
				err = ierr
				return
			}

			if !stat.IsDir() {
				byteArray, ierr := ioutil.ReadFile(absolutePath)
				if ierr != nil {
					err = ierr
					return
				}

				mapResult[key] = string(byteArray)
			}
		}
	}

	return
}

func TemplateScan(scanPath string, filter FileFilter, tmplExt string) (result *ScanResult, err error) {

	fileStat, err := os.Stat(scanPath)
	if err != nil {
		return
	}

	if !fileStat.IsDir() {
		err = fmt.Errorf("%s is not a directory", scanPath)
		return
	}

	// add / to scanPath if necessary
	if !strings.HasSuffix(scanPath, string(filepath.Separator)) {
		scanPath = scanPath + string(filepath.Separator)
	}

	// if extension doesn't start with "." add it
	if !strings.HasSuffix(tmplExt, ".") {
		tmplExt = "." + tmplExt
	}

	result = &ScanResult{
		RootPath:  scanPath,
		Extension: tmplExt,
		Template:  template.New(scanPath),
	}

	// store filter applied for scan
	result.FilterName = runtime.FuncForPC(reflect.ValueOf(filter).Pointer()).Name()

	// start directory walk
	templateMap := FileMap{}

	err = filepath.Walk(scanPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath := strings.TrimPrefix(path, scanPath)
		// enable filter, ignore files excluded by filter
		if filter(relativePath, info) {
			templateFile := FileDetail{
				Info: info,
			}

			// if template found, mark and parse it
			if !info.IsDir() && filepath.Ext(info.Name()) == tmplExt {
				templateFile.IsTemplate = true
				templateFile.TargetPath = strings.TrimSuffix(relativePath, tmplExt)

				// load template file
				byteArray, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}

				// initiate new template with path as name
				innerTemplate := result.Template.New(relativePath).Funcs(sprig.TxtFuncMap())

				// parse template
				_, err = innerTemplate.Parse(string(byteArray))
				if err != nil {
					return err
				}
			}

			templateMap[relativePath] = templateFile
		}

		return nil
	})

	result.TemplateMap = templateMap

	return
}

func IgnoreGit() FileFilter {
	return func(path string, info os.FileInfo) bool {
		if info.IsDir() && strings.HasPrefix(info.Name(), ".git") {
			return false
		}

		if strings.Contains(path, ".git/") {
			return false
		}

		return true
	}
}
