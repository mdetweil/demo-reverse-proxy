package requestRouter

import (
	"net/http"
	"net/url"
	"regexp"
	c "github.com/mdetweil/demo-reverse-proxy/proxy/calcRouter"
	w "github.com/mdetweil/demo-reverse-proxy/proxy/worldRouter"

)

var (
	m map[string]*regexp.Regexp
)

type requestRouter struct {
	Calc c.CalcDialer
	World w.WorldDialer
}

type RequestRouter interface {
	RouteRequest(r *http.Request) (res interface{}, err error)
}

func NewRequestRouter(c c.CalcDialer, world w.WorldDialer) RequestRouter {
	initalizeRegexMap()
	return &requestRouter{Calc: c, World: world}
}

func initalizeRegexMap() {
	m = make(map[string]*regexp.Regexp)
	m["calc"] = regexp.MustCompile(`^/calc`)
	m["world"] = regexp.MustCompile(`^/world`)
}

func (rr *requestRouter) RouteRequest(r *http.Request) (interface{}, error) {
	parsed, _ := url.Parse(r.URL.String())
	path := parsed.Path
	queries := parsed.Query()
	switch {
	case m["calc"].MatchString(path):
		return rr.Calc.RouteCalcRequest(path, queries)
	case m["world"].MatchString(path):
		return rr.World.RouteWorldRequest(path, queries)
	}
	return nil, nil
}
