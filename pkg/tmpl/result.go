package tmpl

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type FileDetail struct {
	Info       os.FileInfo
	TargetPath string
	IsTemplate bool
}

const InvalidMark = "<no value>"

type Validation struct {
	Valid        bool
	TemplateName string
	InvalidLines []string
	Error        error
}

type ValidationMap map[string]Validation

type FileMap map[string]FileDetail

type ScanResult interface {
	TemplateList() (list []string)
	DirList() (list []string)
	ExecuteToPath(data interface{}, targetPath string) (err error)
	Execute(data interface{}) (mapResult map[string]string, err error)
	Validate(data interface{}) (resultMap ValidationMap)
	RootPath() string
	Template() *template.Template
	TemplateMap() FileMap
	Option() *ScanOption
}

type scanResult struct {
	rootPath    string
	template    *template.Template
	templateMap FileMap
	option      *ScanOption
}

func (result *scanResult) Option() *ScanOption {
	return result.option
}

func (result *scanResult) RootPath() string {
	return result.rootPath
}

func (result *scanResult) Template() *template.Template {
	return result.template
}

func (result *scanResult) TemplateMap() FileMap {
	return result.templateMap
}

func (result *scanResult) TemplateList() (list []string) {
	list = []string{}
	for key, value := range result.templateMap {
		if value.IsTemplate {
			list = append(list, key)
		}
	}
	return
}

// DirList result should be sorted
func (result *scanResult) DirList() (list []string) {
	list = []string{}
	for key, value := range result.templateMap {
		if value.Info.IsDir() {
			list = append(list, key)
		}
	}
	sort.Strings(list)
	return
}

func (result *scanResult) ExecuteToPath(data interface{}, targetPath string) (err error) {
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

	for key, value := range result.templateMap {
		if value.IsTemplate {
			var buff bytes.Buffer
			ierr := result.template.ExecuteTemplate(&buff, key, data)
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
			absoluteSourceFile := result.rootPath + key
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

func (result *scanResult) Validate(data interface{}) (resultMap ValidationMap) {
	resultMap = ValidationMap{}
	for key, value := range result.templateMap {
		if value.IsTemplate {
			validation := Validation{
				TemplateName: key,
				Valid:        true,
			}

			var buff bytes.Buffer
			ierr := result.template.ExecuteTemplate(&buff, key, data)
			if ierr != nil {
				validation.Valid = false
				validation.Error = ierr
			} else {
				reader := strings.NewReader(buff.String())
				scanner := bufio.NewScanner(reader)
				scanner.Split(bufio.ScanLines)
				for scanner.Scan() {
					line := scanner.Text()
					if strings.Contains(line, InvalidMark) {
						validation.Valid = false
						validation.InvalidLines = append(validation.InvalidLines, line)
					}
				}

			}

			resultMap[value.TargetPath] = validation
		}
	}

	return
}

func (result *scanResult) Execute(data interface{}) (mapResult map[string]string, err error) {
	mapResult = map[string]string{}
	for key, value := range result.templateMap {
		if value.IsTemplate {
			var buff bytes.Buffer
			ierr := result.template.ExecuteTemplate(&buff, key, data)
			if ierr != nil {
				err = ierr
				return
			}

			mapResult[value.TargetPath] = buff.String()
		} else {
			absolutePath := result.rootPath + key
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
