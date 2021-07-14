package tmpl

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"text/template"
)

type FileFilter func(path string, info os.FileInfo) bool

type TemplateScanFunc func(scanPath string, filter FileFilter, tmplExt string) (ScanResult, error)

func TemplateScan(scanPath string, filter FileFilter, tmplExt string) (ScanResult, error) {

	fileStat, err := os.Stat(scanPath)
	if err != nil {
		return nil, err
	}

	if !fileStat.IsDir() {
		err = fmt.Errorf("%s is not a directory", scanPath)
		return nil, err
	}

	// add / to scanPath if necessary
	if !strings.HasSuffix(scanPath, string(filepath.Separator)) {
		scanPath = scanPath + string(filepath.Separator)
	}

	// if extension doesn't start with "." add it
	if !strings.HasSuffix(tmplExt, ".") {
		tmplExt = "." + tmplExt
	}

	result := &scanResult{
		rootPath:  scanPath,
		extension: tmplExt,
		template:  template.New(scanPath),
	}

	// store filter applied for scan
	result.filterName = runtime.FuncForPC(reflect.ValueOf(filter).Pointer()).Name()

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
