package jq

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type RemoteQuery struct {
	SourceUrl string            `json:"source_url"`
	Headers   map[string]string `json:"headers"`
	Query     string            `json:"query"` // jq query syntax
	QueryType QueryType         `json:"query_type"`
}

func (v *RemoteQuery) Execute(ctx context.Context, callback QueryCallback) error {
	client := resty.New()
	request := client.
		R().
		SetContext(ctx)

	for k, v := range v.Headers {
		request = request.SetHeader(k, v)
	}

	response, err := request.Get(v.SourceUrl)
	if err != nil {
		return err
	}

	if !response.IsSuccess() {
		return fmt.Errorf("request return %v, error: %s", response.StatusCode(), response.Body())
	}

	return Execute(ctx, response.Body(), v.Query, v.QueryType, callback)
}

func NewRemoteQuery(sourceUrl string, headers map[string]string, query string, queryType QueryType) JsonQuery {
	return &RemoteQuery{SourceUrl: sourceUrl, Headers: headers, Query: query, QueryType: queryType}
}
