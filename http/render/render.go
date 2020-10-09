package render

import (
	"html/template"
	"net/http"

	"github.com/gorilla/context"
	"github.com/unrolled/render"
)

var Render *render.Render
var funcMap = template.FuncMap{}

func Init() {
	Render = render.New(render.Options{
		Directory:     "web",
		Extensions:    []string{".html"},
		Delims:        render.Delims{"{{", "}}"},
		Funcs:         []template.FuncMap{funcMap},
		IndentJSON:    false,
		IsDevelopment: true,
	})
}

func Put(r *http.Request, key string, val interface{}) {
	m, ok := context.GetOk(r, "_DATA_MAP_")
	if ok {
		mm := m.(map[string]interface{})
		mm[key] = val
		context.Set(r, "_DATA_MAP_", mm)
	} else {
		context.Set(r, "_DATA_MAP_", map[string]interface{}{key: val})
	}
}

func HTML(r *http.Request, w http.ResponseWriter, name string, htmlOpt ...render.HTMLOptions) {
	Render.HTML(w, http.StatusOK, name, context.Get(r, "_DATA_MAP_"), htmlOpt...)
}

func Text(w http.ResponseWriter, v string, codes ...int) {
	code := http.StatusOK
	if len(codes) > 0 {
		code = codes[0]
	}
	Render.Text(w, code, v)
}

func Message(w http.ResponseWriter, v interface{}) {
	if v == nil {
		Render.JSON(w, http.StatusOK, map[string]string{"err": ""})
		return
	}

	switch t := v.(type) {
	case string:
		Render.JSON(w, http.StatusOK, map[string]string{"err": t})
	case error:
		Render.JSON(w, http.StatusOK, map[string]string{"err": t.Error()})
	}
}

func Data(w http.ResponseWriter, v interface{}, err error) {
	if err != nil {
		Render.JSON(w, http.StatusOK, map[string]interface{}{"err": err.Error(), "dat": v})
	} else {
		Render.JSON(w, http.StatusOK, map[string]interface{}{"err": "", "dat": v})
	}
}
