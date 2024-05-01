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
	"github.com/metal-toolbox/iam-runtime/pkg/iam/runtime/identity"
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
		err := status.Error(codes.InvalidArgument, "who are you?")
		return nil, err
	}

	result := authorization.CheckAccessResponse_RESULT_ALLOWED

	for _, action := range req.Actions {
		if action.GetAction() != "greet" || action.GetResourceId() != "world" {
			result = authorization.CheckAccessResponse_RESULT_DENIED
		}
	}

	out := &authorization.CheckAccessResponse{
		Result: result,
	}

	return out, nil
}

type authenticationServer struct {
	authentication.UnimplementedAuthenticationServer
}

func (s *authenticationServer) ValidateCredential(ctx context.Context, req *authentication.ValidateCredentialRequest) (*authentication.ValidateCredentialResponse, error) {
	if req.Credential != "hello" {
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

type identityServer struct {
	identity.UnimplementedIdentityServer
}

func (s *identityServer) GetAccessToken(ctx context.Context, req *identity.GetAccessTokenRequest) (*identity.GetAccessTokenResponse, error) {
	out := &identity.GetAccessTokenResponse{
		Token: "world",
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
	identity.RegisterIdentityServer(srv, &identityServer{})

	log.Printf("runtime listening at %s", listener.Addr())

	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
