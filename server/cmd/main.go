package main

import (
	"github.com/Xanonymous-GitHub/sxcctw/server/internal/router"
	"github.com/Xanonymous-GitHub/sxcctw/server/pkg/vp"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// Restrict using `gin`, so we won't choose other router/framework packages.
	r := router.NewRouter()

	// Restrict using official `net/http` package.
	s := http.Server{
		Addr:           ":" + strconv.Itoa(vp.Cvp.GetInt("apiServerPort")),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	defer func(s *http.Server) {
		err := s.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(&s)

	err := s.ListenAndServe()
	if err != nil {
		log.Panicln(err)
	}
}
