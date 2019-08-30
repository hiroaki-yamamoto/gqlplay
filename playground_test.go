package gqlplay

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
)

var _ = Describe("Playground", func() {
	templateIteration := func(opt Option) {
		It("Should load correctly", func() {
			fun := Ground(opt)
			tester := httptest.NewRecorder()
			fun(tester, httptest.NewRequest("GET", "/test", nil))

			var expBuf bytes.Buffer
			tmp := template.Must(template.New("gqlplay").Parse(gqltmp))
			cfgTxt, err := json.Marshal(opt)
			Expect(err).To(BeNil())
			err = tmp.Execute(
				&expBuf,
				map[string]interface{}{
					"setting": template.JS(cfgTxt),
				},
			)
			Expect(err).To(BeNil())
			Expect(tester.Code).To(Equal(http.StatusOK))
			Expect(tester.Header().Get("Content-Type")).To(Equal("text/html"))
			Expect(tester.Body.String()).To(Equal(expBuf.String()))
		})
	}
	Describe("Template", func() {
		Context("Without Config", func() {
			var config Option
			templateIteration(config)
		})
		Context("With Config", func() {
			var opt Option
			BeforeEach(func() {
				opt = Option{
					Settings: Settings{
						CursorShape:    CursorShapeLine,
						Credentials:    SameOriginCredentials,
						PollingEnabled: false,
					},
					Endpoint:             "/pub",
					SubscriptionEndpoint: "/pub",
				}
			})
			templateIteration(opt)
		})
		// Context("Has non json-serializable values", func() {
		// 	config := Config{
		// 		"test": 1 + 5i, // Complex number cannot be serialized into json.
		// 	}
		// 	It("Should be marshal error", func() {
		// 		fun := Ground(config)
		// 		_, err := json.Marshal(config)
		// 		var buf bytes.Buffer
		// 		buf.WriteString("Cannot Open Playground: ")
		// 		buf.WriteString(err.Error())
		// 		tester := httptest.NewRecorder()
		// 		fun(tester, httptest.NewRequest("GET", "/test", nil))
		// 		Expect(tester.Code).To(Equal(http.StatusInternalServerError))
		// 		Expect(tester.Header().Get("Content-Type")).To(Equal("text/plain"))
		// 		Expect(tester.Body.String()).To(Equal(buf.String()))
		// 	})
		// })
	})
})
