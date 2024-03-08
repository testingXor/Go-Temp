// Issue 89
// Passing tainted data into logger.Print can
// result in log injection.

// Flaky test
// Taint trace message sometimes indicates that taint is coming from username,
// sometimes indicates password as taint source.

package testdata

import (
	"log"
	"net/http"
	"os"
)

func handler(req *http.Request) {
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")
	logger := log.New(os.Stdout, "", log.LstdFlags)
	// OpenRefactory Warning:
	// Possible Sensitive Data Leak!
	// Path:
	//	File: src.go, Line: 19
	//		password := req.URL.Query().Get("password")
	//		Variable 'password' is assigned a tainted value.
	//	File: src.go, Line: 21
	//		logger.Printf("User %s logged in with password %s", username, password)
	//		Tainted information is used in a sink.
	logger.Printf("User %s logged in with password %s", username, password)
}
