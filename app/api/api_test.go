package api_test

import (
	"encoding/json"
	. "github/desteves/cicd-pulumi-demo/app/api"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api", func() {
	Context("When the GET spell route is set\n", func() {

		var r *gin.Engine
		var w *httptest.ResponseRecorder

		BeforeEach(func() {

			w = httptest.NewRecorder()
			_, r = gin.CreateTestContext(w)
			r.GET(Version+HealthcheckEndpoint, HealthcheckHandler)

		})

		It("can return 200 OK world response", func() {

			req, _ := http.NewRequest("GET", Version+HealthcheckEndpoint, nil)

			r.ServeHTTP(w, req)

			var response string
			err := json.Unmarshal([]byte(w.Body.Bytes()), &response)

			Expect(err).NotTo(HaveOccurred())
			Expect(w.Code).Should(Equal(http.StatusOK))
			Expect(response).Should(Equal("world"))

		})

	})

})
