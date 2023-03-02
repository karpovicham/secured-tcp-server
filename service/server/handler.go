package server

import (
	"errors"
	"fmt"
	"github.com/karpovicham/secured-tcp-server/internal/messenger"
	"github.com/karpovicham/secured-tcp-server/internal/proto"
	"github.com/karpovicham/secured-tcp-server/service/server/domain"
	"net"
)

type RequestHandler struct {
	*Server
	Msgr messenger.Messenger
}

func NewRequestsHandler(server *Server, conn net.Conn) *RequestHandler {
	return &RequestHandler{
		Server: server,
		Msgr:   server.MsgrFn(conn),
	}
}

// validateUserSession return error if session ID is incorrect or expired
func (h *RequestHandler) validateUserSession(requestSessionID string, user domain.User) error {
	if requestSessionID != user.SessionID {
		return errors.New("invalid session ID")
	}

	// Check if session is not expired
	if user.LastLoginAt == nil || user.LastLoginAt.Add(h.Cfg.SessionTTL).Before(h.NowFn()) {
		return errors.New("expired session")
	}

	return nil
}

// sendErrorResponse helper function to prepare and send Error response
func (h *RequestHandler) sendErrorResponse(errCode proto.ErrorCode, errContext string) error {
	data := proto.ErrorResponseData{
		Code: errCode,
		Msg:  errContext,
	}

	dataJson, err := data.MarshalJSON()
	if err != nil {
		return fmt.Errorf("marshal JSON: %w", err)
	}

	if err := h.Msgr.Send(proto.TypeError, dataJson); err != nil {
		return fmt.Errorf("send: %w", err)
	}

	return nil
}
