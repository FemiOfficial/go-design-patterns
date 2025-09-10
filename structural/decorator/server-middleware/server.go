package servermiddleware

import (
	"fmt"
	"io"
	"net/http"
)

type PlainServer struct {}

type Server struct {
	LogWriter io.Writer
	Handler http.Handler
	Username string
	Password string
}

func (p* PlainServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "plain server implemented")
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// log request
	fmt.Fprintf(s.LogWriter, "Reques URL %s \n", r.RequestURI)
	fmt.Fprintf(s.LogWriter, "Reques ContentLength %d \n", r.ContentLength)
	fmt.Fprintf(s.LogWriter, "Reques Host %s \n", r.Host)
	fmt.Fprintf(s.LogWriter, "Reques Method %s \n", r.Method)

	user, pass, ok := r.BasicAuth()
	if ok {
		if user == s.Username && pass == s.Password {
			s.Handler.ServeHTTP(w, r)
			return
		} else {
			fmt.Fprintf(s.LogWriter, "Invalid auth credentials\n")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}

	fmt.Fprintf(s.LogWriter, "No basic auth provided\n")
	http.Error(w, "Unauthorized", http.StatusUnauthorized)

}