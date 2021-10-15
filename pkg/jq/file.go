package jq

import (
	"context"
	"os"
)

type FileQuery struct {
	Path      string    `json:"path"`  // path to file
	Query     string    `json:"query"` // jq query syntax
	QueryType QueryType `json:"query_type"`
}

func (f *FileQuery) Execute(ctx context.Context, callback QueryCallback) error {
	fileBytes, err := os.ReadFile(f.Path)
	if err != nil {
		return err
	}

	return Execute(ctx, fileBytes, f.Query, f.QueryType, callback)
}

func NewFileQuery(path string, query string, queryType QueryType) JsonQuery {
	return &FileQuery{Path: path, Query: query, QueryType: queryType}
}
