package tmpl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gopaytech/go-commons/pkg/file"
	"github.com/stretchr/testify/assert"
)

func TestTemplateScan(t *testing.T) {
	unzipTarget, err := os.MkdirTemp("", "scan-test-file-*")
	assert.NoError(t, err)

	zipName := "scan_test_file.zip"
	t.Logf("unzip %s to %s", zipName, unzipTarget)

	fileList, err := file.Unzip(zipName, unzipTarget)
	assert.NoError(t, err)
	assert.NotEmpty(t, fileList)

	option := DefaultOption()
	option.StartDelim = "[["
	option.EndDelim = "]]"

	t.Logf("test with option %+v", option)

	var result ScanResult
	t.Run("template scan should works", func(t *testing.T) {
		result, err = TemplateScanOption(unzipTarget, option)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// print all template available
		for key, value := range result.TemplateMap() {
			t.Logf("TEMPlATE: Key: %s, is template %v", key, value.IsTemplate)
		}

		assert.Equal(t, unzipTarget+"/", result.RootPath())
		assert.Equal(t, option.Ext, result.Option().Ext)
		assert.Len(t, result.TemplateList(), 5)
		assert.Len(t, result.DirList(), 2)
	})

	t.Run("execute result to memory", func(t *testing.T) {
		data := map[string]interface{}{}
		executeResult, err := result.Execute(data)
		assert.NoError(t, err)

		// all non template and non dir files should exists
		for key, value := range result.TemplateMap() {
			if !value.IsTemplate && !value.Info.IsDir() {
				t.Logf("RESULT NON TEMPLATE: Key: %s, is template %v", key, value.IsTemplate)
				assert.NotEmpty(t, executeResult[key])
			}
		}

		// all template files should exists with trimSuffix filename
		for key, value := range result.TemplateMap() {
			if value.IsTemplate {
				trimmedKey := strings.TrimSuffix(key, result.Option().Ext)
				t.Logf("RESULT TEMPLATE: Key: %s, Trimmed Key %s, is template %v", key, trimmedKey, value.IsTemplate)
				assert.NotEmpty(t, executeResult[trimmedKey])
			}
		}
	})

	t.Run("validate result", func(t *testing.T) {
		data := map[string]interface{}{}
		validateResult := result.Validate(data)

		for key, value := range validateResult {
			assert.NoError(t, value.Error)
			assert.NotEmpty(t, value.InvalidLines)
			t.Logf("Validate %s, %+v", key, value)
		}

	})

	t.Run("execute result to path", func(t *testing.T) {
		templateTarget, err := os.MkdirTemp("", "scan-test-target-*")
		assert.NoError(t, err)

		data := map[string]interface{}{}
		err = result.ExecuteToPath(data, templateTarget)
		assert.NoError(t, err)

		// all non template and non dir files should exists
		for key, value := range result.TemplateMap() {
			if !value.IsTemplate && !value.Info.IsDir() {
				fullFileName := filepath.Join(templateTarget, key)
				t.Logf("RESULT NON TEMPLATE: Key: %s, target: %s, is template %v", key, fullFileName, value.IsTemplate)
				assert.True(t, file.FileExists(fullFileName))
			}
		}

		// all template files should exists with trimSuffix filename
		for key, value := range result.TemplateMap() {
			if value.IsTemplate {
				trimmedKey := strings.TrimSuffix(key, result.Option().Ext)
				fullFileName := filepath.Join(templateTarget, trimmedKey)
				t.Logf("RESULT TEMPLATE: Key: %s, Trimmed Key %s, target %s, is template %v", key, trimmedKey, fullFileName, value.IsTemplate)
				assert.True(t, file.FileExists(fullFileName))
			}
		}

		_ = os.RemoveAll(templateTarget)
	})

}
