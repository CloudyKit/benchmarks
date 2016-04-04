package router

import (
	"github.com/CloudyKit/router"
	"net/http"
	"testing"
)

var cloudyKitRouter = router.New()

var _ = setupTest(func(r route) {
	cloudyKitRouter.AddRoute("GET", r.URLPattern, func(_ http.ResponseWriter, _ *http.Request, p router.Parameter) {
		//for i, value := range r.ParametersName {
		//	if p.Get(value) != r.ParametersValue[i] {
		//		panic("Fail")
		//	}
		//}
	})
})

func BenchmarkCloudyKitRouter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runBenchmark(cloudyKitRouter)
	}
}
