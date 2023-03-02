package resolver

import (
	"errors"
	"fmt"

	"github.com/karpovicham/secured-tcp-server/internal/proto"
)

var (
	ErrUnavailable      = errors.New("unavailable response")
	ErrInvalidReqParams = errors.New("invalid request parameters")
	ErrInvalidReqType   = errors.New("not supported request type")
	ErrInvalidRespType  = errors.New("invalid response type")
	ErrUnauthenticated  = errors.New("not authenticated")
	ErrInvalidCreds     = errors.New("invalid credentials")
	ErrNotFound         = errors.New("not found")
	ErrUnknownRespError = errors.New("unknown response error")
)

// Login make login request and validate response
func (r *resolver) Login(reqData proto.LoginRequestData) (proto.LoginResponseData, error) {
	dataJson, err := reqData.MarshalJSON()
	if err != nil {
		return proto.LoginResponseData{}, fmt.Errorf("marshal JSON: %w", err)
	}

	if err := r.Msgr.Send(proto.TypeLogin, dataJson); err != nil {
		return proto.LoginResponseData{}, fmt.Errorf("send: %w", err)
	}

	resp, err := r.Msgr.Receive()
	if err != nil {
		return proto.LoginResponseData{}, fmt.Errorf("receive: %w", err)
	}

	// Check errors
	if err := r.validateResponse(resp, proto.TypeLogin); err != nil {
		return proto.LoginResponseData{}, err
	}

	var respData proto.LoginResponseData
	if err := respData.UnmarshalJSON(resp.Data); err != nil {
		return proto.LoginResponseData{}, fmt.Errorf("unmarshal JSON: %w", err)
	}
	return respData, nil
}

// validateResponse checks if the response message contains an error
func (r *resolver) validateResponse(message *proto.Message, expectedType proto.Type) error {
	// Got OK response
	if message.Type == expectedType {
		return nil
	}

	// Response does not correspond to the request
	if message.Type != proto.TypeError {
		return ErrInvalidRespType
	}

	// Got Error response
	var respData proto.ErrorResponseData
	if err := respData.UnmarshalJSON(message.Data); err != nil {
		return fmt.Errorf("unmarshal JSON: %w", err)
	}

	var err error

	// Convert API errors into the client domain ones
	switch respData.Code {
	case proto.ErrorCodeUnavailable:
		err = ErrUnavailable
	case proto.ErrorCodeUnknownRequestType:
		err = ErrInvalidReqType
	case proto.ErrorCodeInvalidParameter:
		err = ErrInvalidReqParams
	case proto.ErrorCodeUnauthenticated:
		err = ErrUnauthenticated
	case proto.ErrorCodeInvalidCredentials:
		err = ErrInvalidCreds
	case proto.ErrorCodeNotFound:
		err = ErrNotFound
	default:
		err = ErrUnknownRespError
	}

	// Add context to the error if it's provided
	if respData.Msg != "" {
		err = fmt.Errorf("%w: %s", err, respData.Msg)
	}

	return err
}
