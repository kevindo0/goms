package main

import (
    "fmt"
    "log"
    "time"
    "net/http"
    "math/rand"
    "github.com/opentracing/opentracing-go"
    "github.com/openzipkin/zipkin-go"
    zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
    zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`<a href="/home"> Click here to start a request </a>`))
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
    span := opentracing.StartSpan("/home")
    defer span.Finish()
    w.Write([]byte("Request started"))
    go func() {
        http.Get("http://localhost:8080/async")
    }()
    http.Get("http://localhost:8080/service")
    time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
    w.Write([]byte("Request done!"))
}

// Mocks a service endpoint that makes a DB call
func serviceHandler(w http.ResponseWriter, r *http.Request) {
    // ...
    var sp opentracing.Span
    opName := r.URL.Path
    // Attempt to join a trace by getting trace context from the headers.
    wireContext, err := opentracing.GlobalTracer().Extract(
        opentracing.TextMap,
        opentracing.HTTPHeadersCarrier(r.Header))
    if err != nil {
        // If for whatever reason we can't join, go ahead an start a new root span.
        sp = opentracing.StartSpan(opName)
    } else {
        sp = opentracing.StartSpan(opName, opentracing.ChildOf(wireContext))
    }
    defer sp.Finish()
    http.Get("http://localhost:8080/db")
    time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
}

// Mocks a DB call
func dbHandler(w http.ResponseWriter, r *http.Request) {
    var sp opentracing.Span
    opName := r.URL.Path
    // Attempt to join a trace by getting trace context from the headers.
    wireContext, err := opentracing.GlobalTracer().Extract(
        opentracing.TextMap,
        opentracing.HTTPHeadersCarrier(r.Header))
    if err != nil {
        // If for whatever reason we can't join, go ahead an start a new root span.
        sp = opentracing.StartSpan(opName)
    } else {
        sp = opentracing.StartSpan(opName, opentracing.ChildOf(wireContext))
    }
    defer sp.Finish()
    time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
}

func main() {
    reporter := zipkinhttp.NewReporter("http://10.6.124.21:9411/api/v2/spans")
    defer reporter.Close()
  
    // create our local service endpoint
    endpoint, err := zipkin.NewEndpoint("myService", "myservice.mydomain.com:80")
    if err != nil {
        log.Fatalf("unable to create local endpoint: %+v\n", err)
    }

    // initialize our tracer
    nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
    if err != nil {
        log.Fatalf("unable to create tracer: %+v\n", err)
    }

    // use zipkin-go-opentracing to wrap our tracer
    tracer := zipkinot.Wrap(nativeTracer)
  
    // optionally set as Global OpenTracing tracer instance
    opentracing.SetGlobalTracer(tracer)

    port := 8080
    addr := fmt.Sprintf(":%d", port)
    mux := http.NewServeMux()
    mux.HandleFunc("/", indexHandler)
    mux.HandleFunc("/home", homeHandler)
    mux.HandleFunc("/async", serviceHandler)
    mux.HandleFunc("/service", serviceHandler)
    mux.HandleFunc("/db", dbHandler)
    fmt.Printf("Go to http://localhost:%d/home to start a request!\n", port)
    log.Fatal(http.ListenAndServe(addr, mux))
}