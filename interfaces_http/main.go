package main

// net package is an interface for network
// http packages is within net packages and helps to do requests.
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// custom log writer type
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// no body for the response!!
	// response type is a struct
	fmt.Println("Response (no body):", resp)
	fmt.Println("--------------------------")

	// body of the Response struct
	// body is io,ReadCloser -> is an interface!!

	// Use make because we need to make sure that there is enougth space
	// bs := []byte{}
	bs := make([]byte, 99999) //Because the read function just fill the slice of bytes.
	// resp.Body.Read(bs)        // Reads response and move to bs -> commet to test io.Copy

	fmt.Println("body:", string(bs)) //If empty is because Read is commented above so we can test the Copy version
	fmt.Println("--------------------------")

	// Helps to do read
	// Uses *File* type which implements the Writer interface
	fmt.Println("Read with Copy")
	// io.Copy(os.Stdout, resp.Body) // Use an implemented writer -> commet to test custom writer
	fmt.Println("\n--------------------------")

	// Use custom implementation of writer
	lw := logWriter{}

	fmt.Println("Read with Copy and custom Writer")
	io.Copy(lw, resp.Body)
	fmt.Println("\n--------------------------")

}

// Implementing the Writer interface
func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}
