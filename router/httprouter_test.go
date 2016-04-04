package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
)

var httprouterRouter = httprouter.New()
var OK bool
var _ = setupTest(func(r route) {
	httprouterRouter.GET(r.URLPattern, func(_ http.ResponseWriter, _ *http.Request, p httprouter.Params) {
		for i, value := range r.ParametersName {
			OK = p.ByName(value) != r.ParametersValue[i]
		}
	})
})

func BenchmarkHttprouterRouter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runBenchmark(httprouterRouter)
	}
}
