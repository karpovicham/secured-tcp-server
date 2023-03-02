package resolver

import (
	"github.com/karpovicham/secured-tcp-server/internal/messenger"
	"github.com/karpovicham/secured-tcp-server/internal/proto"
)

// Resolver represents basic API Resolver for Client-Server requests
type Resolver interface {
	// Login - Authentication request
	Login(reqData proto.LoginRequestData) (proto.LoginResponseData, error)

	// Logout - Clear user session
	Logout(reqData proto.LogoutRequestData) (proto.LoginResponseData, error)
}

type resolver struct {
	Msgr messenger.Messenger
}

// NewClientAPIResolver returns implemented Resolver
func NewClientAPIResolver(msgr messenger.Messenger) Resolver {
	return &resolver{
		Msgr: msgr,
	}
}
