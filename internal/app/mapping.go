package app

import "github.com/ThrynSec/aes_implementation_poc/internal/control"

func mapUrls() {
	router.GET("/ping", control.Ping)

	router.GET("/dns/secure/:nameserver", control.Nameserver)
	router.GET("/dns/unsecure/:nameserver", control.NameserverAES)
	router.GET("/decrypt/:message", control.DecryptAES)
}
