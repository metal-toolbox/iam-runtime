package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/metal-toolbox/iam-runtime/pkg/iam/runtime/authentication"
	"github.com/metal-toolbox/iam-runtime/pkg/iam/runtime/authorization"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	runtime = flag.String("runtime", "unix:///tmp/runtime.sock", "Runtime socket")
	address = flag.String("address", ":8080", "Server port")
)

func getToken(req *http.Request) string {
	return strings.TrimPrefix(req.Header.Get("Authorization"), "Bearer ")
}

type server struct {
	runtime interface {
		authorization.AuthorizationClient
		authentication.AuthenticationClient
	}
}

func writeMessage(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	if _, err := w.Write([]byte(msg + "\n")); err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func (s *server) handleWhoAmI(w http.ResponseWriter, req *http.Request) {
	validateRequest := &authentication.ValidateCredentialRequest{
		Credential: getToken(req),
	}

	resp, err := s.runtime.ValidateCredential(req.Context(), validateRequest)
	if err != nil {
		log.Printf("error getting user info: %v", err)
		writeMessage(w, http.StatusInternalServerError, err.Error())

		return
	}

	if resp.Result == authentication.ValidateCredentialResponse_RESULT_INVALID {
		writeMessage(w, http.StatusUnauthorized, "who are you?")
	}

	sub := resp.Subject.SubjectId
	msg := fmt.Sprintf("you are: %s\n", sub)

	writeMessage(w, http.StatusOK, msg)
}

func (s *server) handleCanI(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	what := query.Get("what")
	who := query.Get("who")

	validateRequest := &authentication.ValidateCredentialRequest{
		Credential: getToken(req),
	}

	credsResp, err := s.runtime.ValidateCredential(req.Context(), validateRequest)
	if err != nil {
		log.Printf("error getting user info: %v", err)
		writeMessage(w, http.StatusInternalServerError, err.Error())

		return
	}

	if credsResp.Result == authentication.ValidateCredentialResponse_RESULT_INVALID {
		writeMessage(w, http.StatusUnauthorized, "who are you?")

		return
	}

	accessRequest := &authorization.CheckAccessRequest{
		Credential: getToken(req),
		Actions: []*authorization.AccessRequestAction{
			{
				Action:     what,
				ResourceId: who,
			},
		},
	}

	accessResp, err := s.runtime.CheckAccess(req.Context(), accessRequest)

	if err != nil {
		log.Printf("error trying to do the thing: %v", err)
		writeMessage(w, http.StatusInternalServerError, err.Error())

		return
	}

	if accessResp.Result == authorization.CheckAccessResponse_RESULT_DENIED {
		writeMessage(w, http.StatusForbidden, "no!")
	}

	if _, err := w.Write([]byte("yes!\n")); err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*runtime, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error connecting to runtime: %v", err)
	}

	var runtime struct {
		authorization.AuthorizationClient
		authentication.AuthenticationClient
	}

	runtime.AuthorizationClient = authorization.NewAuthorizationClient(conn)
	runtime.AuthenticationClient = authentication.NewAuthenticationClient(conn)

	srv := &server{
		runtime: runtime,
	}

	http.HandleFunc("/whoami", srv.handleWhoAmI)
	http.HandleFunc("/can-i", srv.handleCanI)

	log.Printf("server listening at %s", *address)

	if err := http.ListenAndServe(*address, nil); err != nil {
		log.Fatalf("error serving: %v", err)
	}
}
