package main

import (
	//"crypto/tls"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"spikenet.com/core/internal/routes"
)

func main() {
	// setup router
	rootRouter := mux.NewRouter()

	// Skeleton for logging all requests later on
	rootRouter.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	})

	// http security headers
	rootRouter.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//w.Header().Set("X-Frame-Options", "SAMEORIGIN")
			//w.Header().Set("X-XSS-Protection", "1; mode=block")
			//w.Header().Set("Content-Security-Policy", "default-src 'self'; object-src 'none'; style-src 'unsafe-inline' 'self'; img-src 'self' data:")
			//w.Header().Set("Strict-Transport-Security", "max-age=31536000")

			next.ServeHTTP(w, r)
		})
	})

	// favicon.ico
	rootRouter.Path("/favicon.ico").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "C:/Users/rthsp/GolandProjects/SpikeNet/build/dist/local/spikenet/favicon.ico")
	}))

	// static path
	rootRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("C:/Users/rthsp/GolandProjects/SpikeNet/build/dist/local/spikenet"))))

	// API routes
	routes.Prepare(rootRouter.PathPrefix("/api").Subrouter())

	// Catch all path to server index.html
	rootRouter.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// prevent caching the entire app
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "-1")

		// serve
		http.ServeFile(w, r, "C:/Users/rthsp/GolandProjects/SpikeNet/build/dist/local/spikenet/index.html")
	}))

	// TLS configs for security
	//tlsConfig := &tls.Config{
	//	CipherSuites: []uint16{
	//		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	//		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	//		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
	//		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
	//		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
	//		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
	//		tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
	//		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	//		tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
	//	},
	//	MinVersion: tls.VersionTLS12,
	//}

	server := &http.Server{
		Addr:    ":8080",
		Handler: rootRouter,
		//TLSConfig: tlsConfig,
	}

	// Start!
	log.Println("Starting SpikeNet!")
	log.Fatal( // TODO: enable TLS when cert is acquired
		server.ListenAndServe().Error(),
	)
}
