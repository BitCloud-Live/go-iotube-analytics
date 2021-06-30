// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/prometheus/common/route"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/prometheus/prometheus/util/httputil"
)

type status string

const (
	statusSuccess status = "success"
	statusError   status = "error"
)

type errorType string

const (
	errorTimeout     errorType = "timeout"
	errorCanceled    errorType = "canceled"
	errorExec        errorType = "execution"
	errorBadData     errorType = "bad_data"
	errorInternal    errorType = "internal"
	errorUnavailable errorType = "unavailable"
	errorNotFound    errorType = "not_found"
)

var (
	LocalhostRepresentations = []string{"127.0.0.1", "localhost", "::1"}
)

type apiError struct {
	typ errorType
	err error
}

func (e *apiError) Error() string {
	return fmt.Sprintf("%s: %s", e.typ, e.err)
}

type response struct {
	Status    status      `json:"status"`
	Data      interface{} `json:"data,omitempty"`
	ErrorType errorType   `json:"errorType,omitempty"`
	Error     string      `json:"error,omitempty"`
}

type apiFuncResult struct {
	data interface{}
	err  *apiError
}

type apiFunc func(r *http.Request) apiFuncResult

// API can register a set of endpoints in a router and handle
// them using the provided storage and query engine.
type API struct {
	now     func() time.Time
	logger  log.Logger
	readAPI api.QueryAPI
}

// New returns an initialized API type.
func New(
	logger log.Logger,
	ctx context.Context,
	tsDB influxdb2.Client,
) *API {

	readAPI := tsDB.QueryAPI("my-org")
	a := &API{
		readAPI: readAPI,
		now:     time.Now,
		logger:  logger,
	}

	return a
}

func setUnavailStatusOnTSDBNotReady(r apiFuncResult) apiFuncResult {
	if r.err != nil && errors.Cause(r.err.err) == tsdb.ErrNotReady {
		r.err.typ = errorUnavailable
	}
	return r
}

// Register the API's endpoints in the given router.
func (api *API) Register(r *route.Router) {
	wrap := func(f apiFunc) http.HandlerFunc {
		hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			result := setUnavailStatusOnTSDBNotReady(f(r))
			if result.err != nil {
				api.respondError(w, result.err, result.data)
				return
			}

			if result.data != nil {
				api.respond(w, result.data)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		})
		return httputil.CompressionHandler{
			Handler: hf,
		}.ServeHTTP
	}

	r.Post("/query", wrap(api.query))
}

type queryData struct {
	Result []interface{} `json:"result"`
}

func invalidParamError(err error, parameter string) apiFuncResult {
	return apiFuncResult{nil, &apiError{
		errorBadData, errors.Wrapf(err, "invalid parameter %q", parameter),
	}}
}

func (api *API) query(r *http.Request) (result apiFuncResult) {
	ctx := r.Context()
	ctx = httputil.ContextFromRequest(ctx, r)

	// Get parser flux query result
	q, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return invalidParamError(err, "reading query from body")
	}
	res, err := api.readAPI.Query(ctx, string(q))
	if err != nil {
		return invalidParamError(err, "query")
	}

	if res.Err() != nil {
		return apiFuncResult{nil, returnAPIError(res.Err())}
	}
	sliced, err := toSlice(res)
	if err != nil {
		return invalidParamError(err, "parsing influxdb results")
	}
	return apiFuncResult{&queryData{
		Result: sliced,
	}, nil}
}

func returnAPIError(err error) *apiError {
	if err == nil {
		return nil
	}
	return &apiError{errorExec, err}
}

func (api *API) respond(w http.ResponseWriter, data interface{}) {
	statusMessage := statusSuccess

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&response{
		Status: statusMessage,
		Data:   data,
	})
	if err != nil {
		level.Error(api.logger).Log("msg", "error marshaling json response", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if n, err := w.Write(b); err != nil {
		level.Error(api.logger).Log("msg", "error writing response", "bytesWritten", n, "err", err)
	}
}

func (api *API) respondError(w http.ResponseWriter, apiErr *apiError, data interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&response{
		Status:    statusError,
		ErrorType: apiErr.typ,
		Error:     apiErr.err.Error(),
		Data:      data,
	})

	if err != nil {
		level.Error(api.logger).Log("msg", "error marshaling json response", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var code int
	switch apiErr.typ {
	case errorBadData:
		code = http.StatusBadRequest
	case errorExec:
		code = http.StatusUnprocessableEntity
	case errorCanceled, errorTimeout:
		code = http.StatusServiceUnavailable
	case errorInternal:
		code = http.StatusInternalServerError
	case errorNotFound:
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if n, err := w.Write(b); err != nil {
		level.Error(api.logger).Log("msg", "error writing response", "bytesWritten", n, "err", err)
	}
}

// toMap converts api.QueryTableResult to a []interface{}.
func toSlice(res *api.QueryTableResult) ([]interface{}, error) {
	out := make([]interface{}, 0)
	for res.Next() {
		if res.Err() != nil {
			return nil, errors.Wrap(res.Err(), "error while iterating over api.QueryTableResult")
		}
		out = append(out, res.Record().Values())
	}
	return out, nil
}
