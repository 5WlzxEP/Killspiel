package RiotApi

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"net/http"
)

// workaround generic method
func get[T any](a Api, url string) (*T, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	req.Header.Set("X-Riot-Token", a.Key)
	req.Header.SetMethod(http.MethodGet)
	res := fasthttp.AcquireResponse()

	err := a.client.Do(req, res)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() > 300 || res.StatusCode() < 200 {
		if res.StatusCode() == 401 {
			// if unauthorized do complete check next time
			a.changed = true
		}

		return nil, errors.New(fmt.Sprintf("not a 200 status code: %d", res.StatusCode()))
	}

	ret := new(T)
	err = json.NewDecoder(res.BodyStream()).Decode(ret)
	return ret, err
}
