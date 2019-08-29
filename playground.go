package gqlplay

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
)

// Config represents the configuration of playground that depends on:
// https://github.com/prisma/graphql-playground
type Config map[string]interface{}

// Ground generates http handler function for Go Playground (with CDN)
func Ground(cfg Config) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		raiseErr := func(err error) {
			var buf bytes.Buffer
			buf.WriteString("Cannot Open Playground: ")
			buf.WriteString(err.Error())
			resp.Header().Add("Content-Type", "text/plain")
			resp.WriteHeader(http.StatusInternalServerError)
			buf.WriteTo(resp)
		}
		cfgTxt, err := json.Marshal(cfg)
		if err != nil {
			raiseErr(err)
			return
		}
		tmp := template.Must(template.New("gqlplay").Parse(gqltmp))
		resp.Header().Add("Content-Type", "text/html")
		resp.WriteHeader(http.StatusOK)
		tmp.Execute(resp, map[string]interface{}{"setting": template.JS(cfgTxt)})
	}
}
