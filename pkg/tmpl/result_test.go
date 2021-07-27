package tmpl

import (
	"testing"
)

func FileWalk(t *testing.T) {
	path := "/home/jasoet/Document/Template/vanilla-gcloud-compute-template"
	result, err := TemplateScanOption(path, DefaultOption())
	if err != nil {
		panic(err)
	}

	t.Logf("Result Path: %s, extension: %s \n", result.RootPath(), result.Extension())

	for key, value := range result.TemplateMap() {
		t.Logf("TEMPlATE: Key: %s, is template %v", key, value.IsTemplate)
	}

	// lookup template
	for _, value := range result.TemplateList() {
		t.Logf("template: %s, is available: %v", value, result.Template().Lookup(value) != nil)
	}

	data := map[string]string{}
	executeResult, err := result.Execute(data)
	if err != nil {
		panic(err)
	}

	for key, value := range executeResult {
		t.Logf("RESULT: Key: %s, value: %d", key, len(value))
	}

	err = result.ExecuteToPath(data, "/tmp/randomalpha")
	if err != nil {
		panic(err)
	}
}
