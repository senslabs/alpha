package httpclient

import (
	"io/ioutil"
	"net/http"
	"net/url"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

func prepare(req *retryablehttp.Request, params url.Values, headers map[string]string) {
	query := req.URL.Query()
	for k, v := range params {
		for _, v := range v {
			query.Add(k, v)
		}
	}
	req.URL.RawQuery = query.Encode()

	for k, v := range headers {
		req.Header.Add(k, v)
	}
}

func PerformR(req *retryablehttp.Request, params url.Values, headers map[string]string) (int, []byte, error) {
	prepare(req, params, headers)
	client := retryablehttp.NewClient()
	if res, err := client.Do(req); err != nil {
		return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
	} else if b, err := ioutil.ReadAll(res.Body); err != nil {
		return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
	} else {
		return res.StatusCode, b, nil
	}
}

func GetR(url string, params url.Values, headers map[string]string) (int, []byte, error) {
	if req, err := retryablehttp.NewRequest("GET", url, nil); err != nil {
		logger.Error(err)
		return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
	} else {
		return PerformR(req, params, headers)
	}
}

func PostR(url string, params url.Values, headers map[string]string, rawBody interface{}) (int, []byte, error) {
	if req, err := retryablehttp.NewRequest("POST", url, rawBody); err != nil {
		logger.Error(err)
		return http.StatusInternalServerError, nil, errors.FromError(errors.GO_ERROR, err)
	} else {
		return PerformR(req, params, headers)
	}
}

func Perform(req *retryablehttp.Request, params url.Values, headers map[string]string, response interface{}) (int, error) {
	prepare(req, params, headers)
	client := retryablehttp.NewClient()
	if res, err := client.Do(req); err != nil {
		return http.StatusInternalServerError, errors.FromError(errors.GO_ERROR, err)
	} else if err := types.JsonUnmarshalFromReader(res.Body, response); err != nil {
		return http.StatusInternalServerError, err
	} else {
		return res.StatusCode, nil
	}
}

func Get(url string, params url.Values, headers map[string]string, response interface{}) (int, error) {
	if req, err := retryablehttp.NewRequest("GET", url, nil); err != nil {
		logger.Error(err)
		return http.StatusInternalServerError, errors.FromError(errors.GO_ERROR, err)
	} else {
		return Perform(req, params, headers, response)
	}
}

func Post(url string, params url.Values, headers map[string]string, rawBody interface{}, response interface{}) (int, error) {
	if req, err := retryablehttp.NewRequest("POST", url, rawBody); err != nil {
		logger.Error(err)
		return http.StatusInternalServerError, errors.FromError(errors.GO_ERROR, err)
	} else {
		return Perform(req, params, headers, response)
	}
}
