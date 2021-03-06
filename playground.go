package gqlplay

import (
	"encoding/json"
	"html/template"
	"net/http"
)

// Ground generates http handler function for Go Playground (with CDN)
func Ground(version string, option Option) http.HandlerFunc {
	tmp := template.Must(template.New("gqlplay").Parse(gqltmp))
	return func(resp http.ResponseWriter, req *http.Request) {
		// raiseErr := func(err error) {
		// 	var buf bytes.Buffer
		// 	buf.WriteString("Cannot Open Playground: ")
		// 	buf.WriteString(err.Error())
		// 	resp.Header().Add("Content-Type", "text/plain")
		// 	resp.WriteHeader(http.StatusInternalServerError)
		// 	buf.WriteTo(resp)
		// }
		// cfgTxt, err := json.Marshal(option)
		// if err != nil {
		// 	raiseErr(err)
		// 	return
		// }
		cfgTxt, _ := json.Marshal(option)
		resp.Header().Add("Content-Type", "text/html")
		resp.WriteHeader(http.StatusOK)
		tmp.Execute(resp, map[string]interface{}{
			"setting": template.JS(cfgTxt),
			"version": version,
		})
	}
}
