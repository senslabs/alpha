package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
)

func Perform(req *retryablehttp.Request) (int, []byte, error) {
	query := req.URL.Query()
	for k, v := range params {
		for _, v := range v {
			query.Add(k, v)
		}
	}
	req.URL.RawQuery = query.Encode()

	for k, v := range headers {
		req.Header.Add(k, v.(string))
	}

	client := retryablehttp.NewClient()
	if res, err := client.Do(req); err != nil {
		return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
	} else if b, err := ioutil.ReadAll(res.Body); err != nil {
		return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
	} else {
		return res.StatusCode, b, nil
	}
}

func Get(url string, params url.Values, headers map[string]string) (int, []byte, *errors.SensError) {
	if req, err := retryablehttp.NewRequest("GET", url, nil); err != nil {
		logger.Error(err)
		return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
	} else {
		return Perform(req)
	}
}

func Post(url string, params map[string]interface{}, headers map[string]interface{}, body []byte) (int, []byte, *errors.SensError) {
	if req, err := retryablehttp.NewRequest("POST", url, bytes.NewBuffer(body)); err != nil {
		logger.Error(err)
		return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
	} else {
		return Perform(req)
	}
}
