package jq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/itchyny/gojq"
)

type QueryType string

const (
	ObjectType QueryType = "object"
	ListType   QueryType = "list"
)

// QueryCallback will be executed for each item found, iteration will be terminated if callback return false
type QueryCallback func(value interface{}) (cont bool)

type JsonQuery interface {
	Execute(ctx context.Context, callback QueryCallback) error
}

func Execute(ctx context.Context, inputJsonBytes []byte, query string, queryType QueryType, callback QueryCallback) error {
	parsedQuery, err := gojq.Parse(query)
	if err != nil {
		return err
	}

	var iter gojq.Iter

	if queryType == ListType {
		var result []interface{}
		err := json.Unmarshal(inputJsonBytes, &result)
		if err != nil {
			return err
		}

		iter = parsedQuery.RunWithContext(ctx, result)
	} else if queryType == ObjectType {
		var result map[string]interface{}
		err := json.Unmarshal(inputJsonBytes, &result)
		if err != nil {
			return err
		}
		iter = parsedQuery.RunWithContext(ctx, result)
	} else {
		return fmt.Errorf("QueryType %s is not supported", queryType)
	}

	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return err
		}

		if !callback(v) {
			break
		}
	}

	return nil
}
