package client

import (
	"context"
	"github.com/google/uuid"
	"github.com/karpovicham/secured-tcp-server/internal/proto"
	"time"

	"github.com/karpovicham/secured-tcp-server/service/client/resolver"
)

// DemoProcess Makes some requests to test the flow
func (c *Client) DemoProcess(ctx context.Context, r resolver.Resolver) error {
	// Test case 1 - Login with unknown user
	c.Log.Info("--------------------------------------------")
	reqData := proto.LoginRequestData{
		Username: "qwerty",
		Password: "windscribe",
	}

	c.Log.Info("Login with none existing user:", reqData.Username)
	respData, err := r.Login(reqData)
	if err != nil {
		c.Log.Info("Got error:", err)
	} else {
		c.Log.Info("Got response:", respData)
	}

	time.Sleep(time.Second)

	// Test case 2 - Login with existing user
	c.Log.Info("--------------------------------------------")
	reqData = proto.LoginRequestData{
		Username: "alex",
		Password: "windscribe",
	}

	c.Log.Info("Login with existing user:", reqData.Username)
	respData, err = r.Login(reqData)
	if err != nil {
		c.Log.Info("Got error:", err)
	} else {
		c.Log.Info("Got session ID:", respData.SessionID)
	}

	time.Sleep(time.Second)

	// Test case 3 - Logout with invalid session
	c.Log.Info("--------------------------------------------")
	const userIDAlex = "d77308d3-a0f7-4ce3-993c-00502754789c"
	logoutReqData := proto.LogoutRequestData{
		UserID:    userIDAlex,
		SessionID: uuid.New().String(),
	}

	c.Log.Info("Logout with invalid session")
	if _, err = r.Logout(logoutReqData); err != nil {
		c.Log.Info("Got error:", err)
	} else {
		c.Log.Info("Got session ID:", respData.SessionID)
	}

	time.Sleep(time.Second)

	// Test case 4 - Logout with valid session
	c.Log.Info("--------------------------------------------")
	logoutReqData = proto.LogoutRequestData{
		UserID:    userIDAlex,
		SessionID: respData.SessionID,
	}

	c.Log.Info("Logout with valid session")
	if _, err = r.Logout(logoutReqData); err != nil {
		c.Log.Info("Got error:", err)
	} else {
		c.Log.Info("Got succeed result")
	}

	return nil
}
