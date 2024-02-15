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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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

func errorToHTTPStatus(err error) (int, string) {
	switch status.Code(err) {
	case codes.PermissionDenied:
		return http.StatusForbidden, "can't do that!"
	case codes.Unauthenticated:
		return http.StatusUnauthorized, "who are you?"
	default:
		return http.StatusInternalServerError, "something broke, sorry"
	}
}

func writeError(w http.ResponseWriter, err error) {
	status, msg := errorToHTTPStatus(err)

	w.WriteHeader(status)
	w.Write([]byte(msg + "\n"))
}

func (s *server) handleWhoAmI(w http.ResponseWriter, req *http.Request) {
	authRequest := &authentication.ValidateCredentialRequest{
		Credential: getToken(req),
	}

	resp, err := s.runtime.ValidateCredential(req.Context(), authRequest)
	if err != nil {
		log.Printf("error getting user info: %v", err)
		writeError(w, err)
		return
	}

	sub := resp.SubjectId
	msg := fmt.Sprintf("you are: %s\n", sub)

	if _, err := w.Write([]byte(msg)); err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func (s *server) handleCanI(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	what := query.Get("what")
	who := query.Get("who")

	accessRequest := &authorization.CheckAccessRequest{
		Credential: getToken(req),
		Actions: []*authorization.AccessRequestAction{
			{
				Action:     what,
				ResourceId: who,
			},
		},
	}

	if _, err := s.runtime.CheckAccess(req.Context(), accessRequest); err != nil {
		log.Printf("error trying to do the thing: %v", err)
		writeError(w, err)
		return
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
