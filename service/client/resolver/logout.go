package resolver

import (
	"fmt"

	"github.com/karpovicham/secured-tcp-server/internal/proto"
)

// Logout make logout request and validate response
func (r *resolver) Logout(reqData proto.LogoutRequestData) (proto.LoginResponseData, error) {
	dataJson, err := reqData.MarshalJSON()
	if err != nil {
		return proto.LoginResponseData{}, fmt.Errorf("marshal JSON: %w", err)
	}

	if err := r.Msgr.Send(proto.TypeLogout, dataJson); err != nil {
		return proto.LoginResponseData{}, fmt.Errorf("send: %w", err)
	}

	resp, err := r.Msgr.Receive()
	if err != nil {
		return proto.LoginResponseData{}, fmt.Errorf("receive: %w", err)
	}

	// Check errors
	if err := r.validateResponse(resp, proto.TypeLogout); err != nil {
		return proto.LoginResponseData{}, err
	}

	return proto.LoginResponseData{}, nil
}
