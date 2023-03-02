package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/karpovicham/secured-tcp-server/internal/proto"
	"github.com/karpovicham/secured-tcp-server/service/server/domain"
)

func (h *RequestHandler) HandleLogoutRequest(ctx context.Context, requestMsg *proto.Message) error {
	var reqData proto.LogoutRequestData
	if err := json.Unmarshal(requestMsg.Data, &reqData); err != nil {
		return h.sendErrorResponse(proto.ErrorCodeInvalidData, err.Error())
	}

	user, err := h.Repo.GetUserByUserID(ctx, reqData.UserID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return h.sendErrorResponse(proto.ErrorCodeNotFound, "User not found")
		}
		h.Log.Error("GetUserByUserID:", err)
		return h.sendErrorResponse(proto.ErrorCodeUnavailable, "")
	}

	if err := h.validateUserSession(reqData.SessionID, user); err != nil {
		return h.sendErrorResponse(proto.ErrorCodeUnauthenticated, err.Error())
	}

	if err := h.Repo.ClearUserSession(ctx, reqData.UserID); err != nil {
		h.Log.Error("ClearUserSession:", err)
		return h.sendErrorResponse(proto.ErrorCodeUnavailable, "")
	}

	// Response with Success
	if err := h.Msgr.Send(proto.TypeLogout, nil); err != nil {
		h.Log.Error("send:", err)
		return fmt.Errorf("send: %w", err)
	}

	return nil
}
