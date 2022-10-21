package main

import (
	"net/http"

	"github.com/JensSkipr/openapi-generator-poc/generate/common"
	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = false
	// loader.ReadFromURIFunc = func(loader *openapi3.Loader, uri *url.URL) ([]byte, error) {
	// 	return fs.ReadFile(uri.Path)
	// }

	doc, err := loader.LoadFromFile("../openapi.yml")
	if err != nil {
		panic(err)
	}

	if err = doc.Validate(loader.Context); err != nil {
		panic(err)
	}

	pathsByTag := make(map[string]*openapi3.Paths, len(doc.Tags))
	for path, pathItem := range doc.Paths {
		tag := getTagForPathItem(pathItem)
		if paths, found := pathsByTag[tag]; found {
			(*paths)[path] = pathItem
		} else {
			pathsByTag[tag] = &openapi3.Paths{path: pathItem}
		}
	}

	templateData := struct {
		Doc        *openapi3.T
		PathsByTag map[string]*openapi3.Paths
	}{
		Doc:        doc,
		PathsByTag: pathsByTag,
	}

	common.EnsureDirExists("../skipr")
	common.DeleteDirContents("../skipr")
	common.ParseTemplates("../templates", "../skipr", templateData)

	// doc.Paths["/expenses"].T

	// json.Marshal(doc.Paths["/expenses"].Get.Responses["200"].Value.Content.Get("application/json").Schema)
}

func getTagForPathItem(pathItem *openapi3.PathItem) string {
	untagged := "Untagged"
	if pathItem == nil {
		return untagged
	}

	methods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete}
	for _, method := range methods {
		operation := pathItem.GetOperation(method)
		if operation == nil {
			continue
		}
		if len(operation.Tags) == 0 {
			continue
		}
		return operation.Tags[0]
	}
	return untagged
}
