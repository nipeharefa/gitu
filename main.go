package main

import (
	"flag"
	"net/http"
	"regexp"

	"go.uber.org/zap"
)

var logger *zap.Logger

func middleware(c *Config, next http.Handler) http.Handler {

	routes := c.Routes
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range routes {

			if v.Src == r.RequestURI {
				next.ServeHTTP(w, r)
				return
			}

			if v.Src == "/" {
				continue
			}

			re := regexp.MustCompile(v.Src)
			match := re.MatchString(r.RequestURI)
			if match {
				for h, v := range v.Headers {
					w.Header().Add(h, v)
				}

				hasRewrite := v.Rewrite != ""
				if hasRewrite {

					r.URL.Path = v.Rewrite
				}
				next.ServeHTTP(w, r)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func StripSlashes(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var path string
		if len(path) > 1 && path[len(path)-1] == '/' {
			newPath := path[:len(path)-1]
			r.URL.Path = newPath
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func main() {

	// ss := zap.Config{}

	ss := zap.NewProductionConfig()
	ss.DisableCaller = true
	ss.DisableStacktrace = true
	logger, _ = ss.Build()

	defer logger.Sync()

	logger.Info("S")
	var nFlag = flag.String("c", "now.json", "help message for flag n")

	flag.Parse()

	c := ReadConfig(*nFlag)

	strip := StripSlashes(
		middleware(c, http.FileServer(http.Dir("./static"))),
	)
	http.Handle("/", strip)

	logger.Info("Started at :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		logger.Error(err.Error())
	}
}
