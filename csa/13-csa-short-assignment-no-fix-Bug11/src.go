// Short assignment declaration of 'v'.
// CSA should not generate a warning.

package testdata

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v
		go func() {
			fmt.Println(v)
			// OpenRefactory Warning:
			// Possible blocking channel operation!
			//
			// In goroutine #1:
			//	File: src.go, Line: 16
			//		done <- true
			//		In function main$1, there is a send operation.
			//		But no matching receive operation is found on that channel from any other goroutine.
			done <- true
		}()
	}
	for range values {
		<-done
	}
}

//<<<<<216, 222>>>>>
