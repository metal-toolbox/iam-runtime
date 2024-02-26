package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"syscall"

	"github.com/metal-toolbox/iam-runtime/pkg/iam/runtime/authentication"
	"github.com/metal-toolbox/iam-runtime/pkg/iam/runtime/authorization"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

var (
	socket = flag.String("socket", "/tmp/runtime.sock", "Socket path")
)

type authorizationServer struct {
	authorization.UnimplementedAuthorizationServer
}

func (s *authorizationServer) CheckAccess(ctx context.Context, req *authorization.CheckAccessRequest) (*authorization.CheckAccessResponse, error) {
	tok := req.GetCredential()

	log.Printf("received token: %s", tok)
	if tok != "hello" {
		err := status.Error(codes.Unauthenticated, "who are you?")
		return nil, err
	}

	for _, action := range req.Actions {
		if action.GetAction() != "greet" || action.GetResourceId() != "world" {
			err := status.Error(codes.PermissionDenied, "what are you trying to do?")
			return nil, err
		}
	}

	return &authorization.CheckAccessResponse{}, nil
}

type authenticationServer struct {
	authentication.UnimplementedAuthenticationServer
}

func (s *authenticationServer) ValidateCredential(ctx context.Context, req *authentication.ValidateCredentialRequest) (*authentication.ValidateCredentialResponse, error) {
	if req.GetCredential() != "hello" {
		out := &authentication.ValidateCredentialResponse{
			Result: authentication.ValidateCredentialResponse_RESULT_INVALID,
		}

		return out, nil
	}

	claimsMap := map[string]any{
		"aud": "world",
	}

	claims, err := structpb.NewStruct(claimsMap)
	if err != nil {
		return nil, err
	}

	out := &authentication.ValidateCredentialResponse{
		Result: authentication.ValidateCredentialResponse_RESULT_VALID,
		Subject: &authentication.Subject{
			SubjectId: "hello",
			Claims:    claims,
		},
	}

	return out, nil
}

func main() {
	flag.Parse()

	if _, err := os.Stat(*socket); err == nil {
		log.Printf("socket found at %s, unlinking", *socket)
		if err := syscall.Unlink(*socket); err != nil {
			log.Fatalf("error unlinking socket: %v", err)
		}
	}

	listener, err := net.Listen("unix", *socket)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	authorization.RegisterAuthorizationServer(srv, &authorizationServer{})
	authentication.RegisterAuthenticationServer(srv, &authenticationServer{})

	log.Printf("runtime listening at %s", listener.Addr())

	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
