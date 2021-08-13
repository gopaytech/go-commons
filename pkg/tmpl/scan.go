package tmpl

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/gopaytech/go-commons/pkg/file"
)

type ScanOption struct {
	StartDelim     string
	EndDelim       string
	Ext            string
	IgnoreFileName string
	Filters        []FileFilter
}

func (s *ScanOption) AddFilter(filter ...FileFilter) {
	s.Filters = append(s.Filters, filter...)
}

func (s *ScanOption) HasCustomDelim() bool {
	startDelim := strings.TrimSpace(s.StartDelim)
	endDelim := strings.TrimSpace(s.EndDelim)

	return len(startDelim) > 0 && len(endDelim) > 0
}

//Evaluate if on of filter return true, file will be skipped (Evaluate return false)
func (s *ScanOption) Evaluate(path string, info os.FileInfo) bool {
	path = strings.TrimSpace(path)
	if len(path) == 0 {
		return false
	}

	for _, filter := range s.Filters {
		if filter(path, info) {
			return false
		}
	}
	return true
}

func DefaultOption() *ScanOption {
	return &ScanOption{
		StartDelim:     "",
		EndDelim:       "",
		Ext:            "tmpl",
		IgnoreFileName: ".tmplignore",
		Filters:        []FileFilter{IgnoreGit()},
	}
}

// FileFilter if filter return true, file will be ignored
type FileFilter func(path string, info os.FileInfo) bool

type TemplateScanOptionFunc func(scanPath string, option ScanOption) (ScanResult, error)

func TemplateScanOption(scanPath string, option *ScanOption) (ScanResult, error) {
	if !file.DirExists(scanPath) {
		return nil, fmt.Errorf("directory %s is not exists", scanPath)
	}

	// add / to scanPath if necessary
	if !strings.HasSuffix(scanPath, string(filepath.Separator)) {
		scanPath = scanPath + string(filepath.Separator)
	}

	// load ignore file
	ignoreFile := filepath.Join(scanPath, option.IgnoreFileName)
	if file.FileExists(ignoreFile) {
		stringByte, _ := ioutil.ReadFile(ignoreFile)
		ignoreFilters := Ignore(string(stringByte))
		option.AddFilter(ignoreFilters...)
	}

	// if extension doesnt start with "." add it
	tmplExt := option.Ext
	if !strings.HasPrefix(tmplExt, ".") {
		tmplExt = "." + tmplExt
		option.Ext = tmplExt
	}

	rootTemplate := template.New(scanPath)
	if option.HasCustomDelim() {
		rootTemplate = rootTemplate.Delims(option.StartDelim, option.EndDelim)
	}

	result := &scanResult{
		rootPath: scanPath,
		template: rootTemplate,
		option:   option,
	}

	// start directory walk
	templateMap := FileMap{}
	err := filepath.Walk(scanPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath := strings.TrimPrefix(path, scanPath)
		// enable filter, ignore files excluded by filter
		if option.Evaluate(relativePath, info) {
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
				innerTemplate := result.template.New(relativePath).Funcs(sprig.TxtFuncMap())

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

	result.templateMap = templateMap

	return result, err
}

type TemplateScanDelimFunc func(scanPath string, tmplExt string, startDelims string, endDelims string) (ScanResult, error)

func TemplateScanDelim(scanPath string, tmplExt string, startDelims string, endDelims string) (ScanResult, error) {
	option := DefaultOption()
	option.Ext = tmplExt
	option.StartDelim = startDelims
	option.EndDelim = endDelims

	return TemplateScanOption(scanPath, option)
}

type TemplateScanFunc func(scanPath string, tmplExt string) (ScanResult, error)

func TemplateScan(scanPath string, tmplExt string) (ScanResult, error) {
	option := DefaultOption()
	option.Ext = tmplExt

	return TemplateScanOption(scanPath, option)
}

func IgnoreGit() FileFilter {
	return func(path string, info os.FileInfo) bool {
		if info.IsDir() && strings.HasPrefix(info.Name(), ".git") {
			return true
		}

		if strings.Contains(path, ".git/") {
			return true
		}

		return false
	}
}

func Ignore(ignoreFileContent string) []FileFilter {
	reader := strings.NewReader(ignoreFileContent)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var filters []FileFilter
	for scanner.Scan() {
		pattern := scanner.Text()
		fun := func(path string, _ os.FileInfo) bool {
			return Match(pattern, path)
		}

		filters = append(filters, fun)
	}
	return filters
}
