package http

import (
	"bytes"
	"io/ioutil"
	"net/http"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

func Get(url string, params types.Map, headers types.Map) (int, []byte, *errors.SensError) {
	if req, err := retryablehttp.NewRequest("GET", url, nil); err != nil {
		logger.Error(err)
		return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
	} else {
		query := req.URL.Query()
		for k, v := range params {
			query.Add(k, v.(string))
		}
		req.URL.RawQuery = query.Encode()

		for k, v := range headers {
			req.Header.Add(k, v.(string))
		}

		client := retryablehttp.Client{}
		if res, err := client.Do(req); err != nil {
			return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
		} else if b, err := ioutil.ReadAll(res.Body); err != nil {
			return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
		} else {
			return res.StatusCode, b, nil
		}
	}
}

func Post(url string, params map[string]interface{}, headers map[string]interface{}, body []byte) (int, []byte, *errors.SensError) {
	if req, err := retryablehttp.NewRequest("POST", url, bytes.NewBuffer(body)); err != nil {
		logger.Error(err)
		return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
	} else {
		query := req.URL.Query()
		for k, v := range params {
			query.Add(k, v.(string))
		}
		req.URL.RawQuery = query.Encode()

		for k, v := range headers {
			req.Header.Add(k, v.(string))
		}

		client := retryablehttp.NewClient()
		logger.Debug(*req)
		if res, err := client.Do(req); err != nil {
			return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
		} else if b, err := ioutil.ReadAll(res.Body); err != nil {
			return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
		} else {
			return res.StatusCode, b, nil
		}
	}
}
