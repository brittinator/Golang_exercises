/*Server a randomly slow/fast server. It is used by the
worker program.
*/
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// RandomlySlowServer is sometimes fast, and other times
// much slower
func RandomlySlowServer(w http.ResponseWriter, req *http.Request) {
	fiftyFiftyShot := rand.Intn(2)
	fmt.Println(fiftyFiftyShot)

	if fiftyFiftyShot == 0 {
		time.Sleep(6 * time.Second)
		fmt.Fprintf(w, "GO sloooowly %v\n", fiftyFiftyShot)
		fmt.Printf("GO sloooowly %v\n", fiftyFiftyShot)
		return
	}
	fmt.Fprintf(w, "GO quickly %v\n", fiftyFiftyShot)
	fmt.Printf("GO quickly %v\n", fiftyFiftyShot)
	return
}

func main() {
	http.HandleFunc("/", RandomlySlowServer)
	http.ListenAndServe(":1234", nil)
}
