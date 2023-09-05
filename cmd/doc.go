//go:build doc

package main

import (
	_ "embed"
)

//go:embed swagger.yaml
var documentation string

func init() {
	headers := map[string]string{"Content-Type": "text/yaml"}
	handler = func(c Context, r Event) (Response, error) {
		return Response{StatusCode: 200, Headers: headers, Body: documentation}, nil
	}
}
