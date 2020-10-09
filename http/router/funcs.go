package router

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/toolkits/pkg/errors"
	"github.com/toolkits/pkg/param"
	"github.com/toolkits/pkg/str"
)

func isDangerous(key, val string) {
	if str.Dangerous(val) {
		errors.Bomb("arg[%s] is dangerous", key)
	}
}

func isBlank(key, val string) {
	v := strings.TrimSpace(val)
	if v == "" {
		errors.Bomb("arg[%s] is blank", key)
	}
}

func bindJSON(r *http.Request, val interface{}) {
	errors.Dangerous(param.BindJson(r, val))
}

func urlParamStr(r *http.Request, field string) string {
	val, found := mux.Vars(r)[field]
	if !found {
		errors.Bomb("[%s] not found in url", field)
	}

	if val == "" {
		errors.Bomb("[%s] is blank", field)
	}

	return val
}

func urlParamInt64(r *http.Request, field string) int64 {
	strval := urlParamStr(r, field)
	intval, err := strconv.ParseInt(strval, 10, 64)
	if err != nil {
		errors.Bomb("cannot convert %s to int64", strval)
	}

	return intval
}

func urlParamInt(r *http.Request, field string) int {
	return int(urlParamInt64(r, field))
}
