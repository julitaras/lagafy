package helpers

import (
	"context"
	"net/http"
	"time"
)

//Curl make request /  response return
func Curl(url string, httpMethod string, delaySeconds int) (*http.Response, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(delaySeconds))
	defer cancel()

	req, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return res, nil
}
