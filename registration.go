package registration

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var l *zap.Logger

// our main function
func StartListener() {
	l, _ = zap.NewDevelopment()
	defer l.Sync()
	router := mux.NewRouter()
	router.HandleFunc("/payload", HandleGitPayload).Methods("POST")
	log.Fatal(http.ListenAndServe(":4567", router))
}

func HandleGitPayload(w http.ResponseWriter, r *http.Request) {
	l.Info(fmt.Sprintf("recieved POST at /payload from %v", r.RemoteAddr),
		zap.String("source", r.RemoteArddr))

	if _, ok := r.Header["X-Github-Event"]; !ok {
		fmt.Fprintf(w, "Invalid request headers")
		return
	}

	if r.Header["X-Github-Event"][0] != "push" {
		l.Info(fmt.Sprintf("skipping request of type: %s", r.Header["X-Github-Event"][0]))
		fmt.Fprintf(w, "Processed request of type: %s", r.Header["X-Github-Event"][0])
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		l.Error(err.Error())
	}

	strbody, _ := url.QueryUnescape(string(body))
	strbody  = strings.TrimPrefix(strbody, "payload=")

	hook := new(Push); err = json.Unmarshal([]byte(strbody), hook)
	if err != nil {
		l.Error(err.Error())
		if e, ok := err.(*json.SyntaxError); ok {
			l.Error(fmt.Sprintf("Syntax error at byte offset %d", e.Offset))
		}
	}
	fmt.Printf("%+v\n", hook)
}





