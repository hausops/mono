package local

import (
	"context"

	"github.com/hausops/mono/services/auth-svc/domain/session"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

type sessionRepository struct {
	// AcccessToken is the primary key.
	byAccessToken map[session.AccessToken]session.Session
	// UserID is an index.
	byUserID map[user.ID]session.AccessToken
}

func NewSessionRepository() *sessionRepository {
	return &sessionRepository{
		byAccessToken: make(map[session.AccessToken]session.Session),
		byUserID:      make(map[user.ID]session.AccessToken),
	}
}

var _ session.Repository = (*sessionRepository)(nil)

func (r *sessionRepository) DeleteByAccessToken(_ context.Context, token session.AccessToken) error {
	sess, ok := r.byAccessToken[token]
	if !ok {
		return session.ErrNotFound
	}

	delete(r.byAccessToken, token)
	delete(r.byUserID, sess.UserID)
	return nil
}

func (r *sessionRepository) FindByAccessToken(_ context.Context, token session.AccessToken) (session.Session, error) {
	sess, ok := r.byAccessToken[token]
	if !ok {
		return session.Session{}, session.ErrNotFound
	}
	return sess, nil
}

func (r *sessionRepository) FindByUserID(_ context.Context, uid user.ID) (session.Session, error) {
	token, ok := r.byUserID[uid]
	if !ok {
		return session.Session{}, session.ErrNotFound
	}

	sess, ok := r.byAccessToken[token]
	if !ok {
		return session.Session{}, session.ErrNotFound
	}
	return sess, nil
}

func (r *sessionRepository) Upsert(_ context.Context, sess session.Session) error {
	token := sess.AccessToken

	// If updating, remove the previous user ID index for the session.
	if prev, ok := r.byAccessToken[token]; ok {
		delete(r.byUserID, prev.UserID)
	}

	// If a token already exists for the user, delete the old session
	// to ensure one active token per user.
	if prevToken, ok := r.byUserID[sess.UserID]; ok {
		delete(r.byAccessToken, prevToken)
	}

	r.byAccessToken[token] = sess
	r.byUserID[sess.UserID] = token
	return nil
}
