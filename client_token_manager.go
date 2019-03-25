package apns2

import (
	"errors"

	"github.com/riftbit/apns2/token"
)

// Possible errors when parsing a .p8 file.
var (
	ErrTokenNotFound = errors.New("tokenManager: token not found in TokenManager storage")
)

// ClientManager is a way to manage multiple connections to the APNs.
type ClientTokenManager struct {
	client       *Client
	TokenManager *token.TokenManager
}

// NewClientTokenManager returns a new ClientManager for prolonged, concurrent usage
// of multiple APNs clients. ClientManager is flexible enough to work best for
// your use case. When a client is not found in the manager, Get will return
// the result of calling Factory, which can be a Client or nil.
//
// Having multiple clients per certificate in the manager is not allowed.
//
// By default, MaxSize is 64, MaxAge is 10 minutes, and Factory always returns
// a Client with default options.
func NewClientTokenManager() *ClientTokenManager {
	manager := &ClientTokenManager{
		client:       NewTokenClient(nil),
		TokenManager: token.NewTokenManager(),
	}

	return manager
}

// Push ...
func (ctm *ClientTokenManager) Push(tokenKey interface{}, n *Notification) (*Response, error) {
	var token *token.Token
	token, ok := ctm.TokenManager.Get(tokenKey)
	if !ok {
		return nil, ErrTokenNotFound
	}

	return ctm.client.PushWithContextAndToken(nil, token, n)
}

// PushWithContext ...
func (ctm *ClientTokenManager) PushWithContext(ctx Context, tokenKey interface{}, n *Notification) (*Response, error) {
	var token *token.Token
	token, ok := ctm.TokenManager.Get(tokenKey)
	if !ok {
		return nil, ErrTokenNotFound
	}

	return ctm.client.PushWithContextAndToken(ctx, token, n)
}
