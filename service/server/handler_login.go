package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/karpovicham/secured-tcp-server/internal/proto"
	"github.com/karpovicham/secured-tcp-server/pkg/util/hasher"
	"github.com/karpovicham/secured-tcp-server/service/server/domain"
)

func (h *RequestHandler) HandleLoginRequest(ctx context.Context, requestMsg *proto.Message) error {
	var reqData proto.LoginRequestData
	if err := json.Unmarshal(requestMsg.Data, &reqData); err != nil {
		return h.sendErrorResponse(proto.ErrorCodeInvalidData, err.Error())
	}

	user, err := h.Repo.GetUserByUsername(ctx, reqData.Username)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return h.sendErrorResponse(proto.ErrorCodeNotFound, "User not found")
		}
		h.Log.Error("GetUserByUsername:", err)
		return h.sendErrorResponse(proto.ErrorCodeUnavailable, "")
	}

	// Check hashes of passwords
	equals, err := hasher.CompareToSha256Hash(reqData.Password, user.HashedPassword)
	if err != nil {
		h.Log.Error("Compare hash:", err)
		return h.sendErrorResponse(proto.ErrorCodeUnavailable, "")
	}

	if !equals {
		return h.sendErrorResponse(proto.ErrorCodeInvalidCredentials, "Incorrect password")
	}

	timeNow := h.NowFn()
	sessionID := uuid.New().String()

	if err := h.Repo.UpdateUserSession(ctx, reqData.Username, sessionID, timeNow); err != nil {
		h.Log.Error("UpdateUserSession:", err)
		return h.sendErrorResponse(proto.ErrorCodeUnavailable, "")
	}

	respData := proto.LoginResponseData{
		UserID:    user.ID,
		SessionID: sessionID,
		ExpireAt:  timeNow.Add(h.Cfg.SessionTTL).Unix(),
	}
	respDataJson, _ := respData.MarshalJSON()

	// Response with Success authentication
	if err := h.Msgr.Send(proto.TypeLogin, respDataJson); err != nil {
		return fmt.Errorf("send: %w", err)
	}

	return nil
}
