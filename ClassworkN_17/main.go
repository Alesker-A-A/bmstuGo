// package main

// import (
// 	"net/http"
// 	_ "net/http/pprof"
// )

// const sliceSize = 10_000_000

// func foo() {
// 	for {
// 		var s = make([]int, 0, sliceSize)
// 		for i := 0; i < sliceSize; i++ {
// 			s = append(s, i)
// 		}
// 	}
// }

//	func main() {
//		go foo()
//		http.ListenAndServe(":8080", nil)
//	}
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequestsCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "total count of requsts",
	},
	[]string{"path"},
)

var httpRequestsDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_requests_duration_seconds",
		Help: "Duration of HTTP requests",
	},
	[]string{"method"},
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		httpRequestsCounter.WithLabelValues(r.URL.Path).Inc()
		lambda := time.Since(start).Seconds()
		httpRequestsDuration.WithLabelValues(r.Method).Observe(lambda)
	}()

	w.Write([]byte("index page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		httpRequestsCounter.WithLabelValues(r.URL.Path).Inc()
		lambda := time.Since(start).Seconds()
		httpRequestsDuration.WithLabelValues(r.Method).Observe(lambda)
	}()

	w.Write([]byte("about page"))
}

func main() {

	prometheus.MustRegister(httpRequestsCounter, httpRequestsDuration)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Server starting...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
