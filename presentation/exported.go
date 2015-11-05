package main
import (
"net"
"net/http"
	"errors"
)

func main() {

}

//START OMIT
// Serve accepts incoming HTTP connections on the listener l,
// creating a new service goroutine for each.  The service goroutines
// read requests and then call handler to reply to them.
// Handler is typically nil, in which case the DefaultServeMux is used.
func Serve(l net.Listener, handler http.Handler) error {
//END OMIT
	return errors.New("not implemented")
}
