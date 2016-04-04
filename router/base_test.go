package router

import (
	"net/http"
	"regexp"
)

type route struct {
	URLPattern      string
	ParametersName  []string
	ParametersValue []string
	R               *http.Request
}

func setupTest(fn func(r route)) error {

	for _, r := range allRoutes {
		fn(r)
	}

	return nil
}

var acRegex = regexp.MustCompile("[:*][^/]+")

func newRoute(url string, values ...string) route {
	names := acRegex.FindAllString(url, 1000)
	for index, newName := range names {
		names[index] = newName[1:]
	}
	var i = 0
	newURL, _ := http.NewRequest("GET", acRegex.ReplaceAllStringFunc(url, func(m string) string {
		m = values[i]
		i++
		return m
	}), nil)

	return route{R: newURL, URLPattern: url, ParametersName: names, ParametersValue: values}
}

var allRoutes = []route{
	newRoute("/users"),
	newRoute("/users/:userId", "131231"),
	newRoute("/users/:userId/subscriptions", "131231"),
	newRoute("/users/:userId/subscriptions/:subscription", "131231", "12"),
	newRoute("/assets/*files", "css/styles.css"),
}

type dummy struct{}

func (d *dummy) Write(b []byte) (int, error) {
	return 0, nil
}
func (d *dummy) WriteHeader(c int) {

}
func (d *dummy) Header() http.Header {
	return nil
}

var dd = &dummy{}

func runBenchmark(s http.Handler) {
	for _, r := range allRoutes {
		s.ServeHTTP(dd, r.R)
	}
}
