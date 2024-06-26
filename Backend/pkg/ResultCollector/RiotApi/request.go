package RiotApi

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"net/http"
)

// workaround generic method
func getWithApi[T any](a *Api, url string) (*T, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	req.Header.Set("X-Riot-Token", a.LoL.ApiKey)
	req.Header.SetMethod(http.MethodGet)
	res := fasthttp.AcquireResponse()

	err := a.client.Do(req, res)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() > 300 || res.StatusCode() < 200 {
		if res.StatusCode() == http.StatusUnauthorized || res.StatusCode() == http.StatusForbidden {
			// if unauthorized do complete check next time
			a.ready = false
		}

		return nil, errors.New(fmt.Sprintf("not a 200 status code: %d", res.StatusCode()))
	}

	data := res.Body()
	ret := new(T)
	err = json.Unmarshal(data, ret)
	return ret, err
}

func get[T any](url, apiKey string, client *fasthttp.Client) (*T, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	req.Header.Set("X-Riot-Token", apiKey)
	req.Header.SetMethod(http.MethodGet)
	res := fasthttp.AcquireResponse()

	err := client.Do(req, res)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() > 300 || res.StatusCode() < 200 {
		logger.WithFields(log.Fields{
			"request": url,
			"method":  "GET",
			"status":  res.StatusCode(),
			"body":    string(res.Body()),
		}).Error()
		return nil, errors.New(fmt.Sprintf("not a 200 status code: %d", res.StatusCode()))
	}

	data := res.Body()
	ret := new(T)
	err = json.Unmarshal(data, ret)
	return ret, err
}
