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

func getUserName(req *http.Request) string {
	return req.URL.Query().Get("username")
}
func getPass(req *http.Request) string {
	return req.URL.Query().Get("password")
}
func Handler(req *http.Request) {
	username := getUserName(req)
	password := getPass(req)
	logger := log.New(os.Stdout, "", log.LstdFlags)
	// OpenRefactory Warning:
	// Possible Sensitive Data Leak!
	// Path:
	//	File: src.go, Line: 25
	//		password := getPass(req)
	//		Variable 'password' is assigned a tainted value.
	//	File: src.go, Line: 27
	//		logger.Printf("User %s logged in with password %s", username, password)
	//		Tainted information is used in a sink.
	logger.Printf("User %s logged in with password %s", username, password)
}
