package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/karpovicham/secured-tcp-server/internal/proto"
	"github.com/karpovicham/secured-tcp-server/service/server/domain"
	"net/mail"
)

func (h *RequestHandler) HandleModifyAccountSettingsRequest(ctx context.Context, requestMsg *proto.Message) error {
	var reqData proto.ModifyAccountSettingsRequestData
	if err := json.Unmarshal(requestMsg.Data, &reqData); err != nil {
		return h.sendErrorResponse(proto.ErrorCodeInvalidData, err.Error())
	}

	// Validate request parameters
	if reqData.NewEmail == "" || reqData.NewUsername == "" {
		return h.sendErrorResponse(proto.ErrorCodeInvalidParameter, "NewEmail or NewUsername must be set")
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

	// Validate new account settings
	username := user.Username
	if reqData.NewUsername != user.Username {
		username = reqData.NewUsername
	}

	if username == "" {
		return h.sendErrorResponse(proto.ErrorCodeInvalidParameter, "username must not be empty")
	}

	// Email could be empty
	email := user.Email
	if reqData.NewEmail != "" {
		if _, err := mail.ParseAddress(reqData.NewEmail); err != nil {
			return h.sendErrorResponse(proto.ErrorCodeInvalidParameter, "invalid email format")
		}
		email = &reqData.NewEmail
	} else {
		email = nil
	}

	if username == user.Username && email == user.Email {
		return h.sendErrorResponse(proto.ErrorCodeInvalidParameter, "got nothing to update")
	}

	if err := h.Repo.UpdateUserData(ctx, reqData.UserID, username, email); err != nil {
		h.Log.Error("UpdateUserData:", err)
		return h.sendErrorResponse(proto.ErrorCodeUnavailable, "")
	}

	// Response with Success
	if err := h.Msgr.Send(proto.TypeDeactivateUser, nil); err != nil {
		h.Log.Error("send:", err)
		return fmt.Errorf("send: %w", err)
	}

	return nil
}
