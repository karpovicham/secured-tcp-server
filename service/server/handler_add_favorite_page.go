package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/karpovicham/secured-tcp-server/internal/proto"
	"github.com/karpovicham/secured-tcp-server/service/server/domain"
	"net/url"
)

func (h *RequestHandler) HandleAddFavoritePageRequest(ctx context.Context, requestMsg *proto.Message) error {
	var reqData proto.AddFavoritePageRequestData
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

	if reqData.PageURL == "" {
		return h.sendErrorResponse(proto.ErrorCodeInvalidParameter, "PageURL must not be empty")
	}

	if _, err := url.Parse(reqData.PageURL); err != nil {
		return h.sendErrorResponse(proto.ErrorCodeInvalidParameter, "URL is not valid")
	}

	if err := h.Repo.AddNewUserPage(ctx, reqData.UserID, reqData.PageURL); err != nil {
		h.Log.Error("DeactivateUser:", err)
		return h.sendErrorResponse(proto.ErrorCodeUnavailable, "")
	}

	// Response with Success
	if err := h.Msgr.Send(proto.TypeAddFavoritePage, nil); err != nil {
		h.Log.Error("send:", err)
		return fmt.Errorf("send: %w", err)
	}

	return nil
}
