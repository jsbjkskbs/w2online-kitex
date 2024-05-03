package pprof

import (
	"log"
	"net/http"
	"os"
	"runtime"
)

func Load(port string) {
	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(`:`+port, nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
}
