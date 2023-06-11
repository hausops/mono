package local

import (
	"context"
	"net/mail"

	"github.com/hausops/mono/services/auth-svc/domain/session"
)

type sessionRepository struct {
	// Email is the primary key.
	byEmail map[mail.Address]session.Session
	// AcccessToken is an index.
	byAccessToken map[session.AccessToken]mail.Address
}

func NewSessionRepository() *sessionRepository {
	return &sessionRepository{
		byEmail:       make(map[mail.Address]session.Session),
		byAccessToken: make(map[session.AccessToken]mail.Address),
	}
}

var _ session.Repository = (*sessionRepository)(nil)

func (r *sessionRepository) DeleteByAccessToken(_ context.Context, token session.AccessToken) (session.Session, error) {
	email, ok := r.byAccessToken[token]
	if !ok {
		return session.Session{}, session.ErrNotFound
	}

	sess, ok := r.byEmail[email]
	if !ok {
		return session.Session{}, session.ErrNotFound
	}

	delete(r.byEmail, email)
	delete(r.byAccessToken, token)
	return sess, nil
}

func (r *sessionRepository) FindByAccessToken(_ context.Context, token session.AccessToken) (session.Session, error) {
	email, ok := r.byAccessToken[token]
	if !ok {
		return session.Session{}, session.ErrNotFound
	}

	sess, ok := r.byEmail[email]
	if !ok {
		return session.Session{}, session.ErrNotFound
	}
	return sess, nil
}

func (r *sessionRepository) FindByEmail(_ context.Context, email mail.Address) (session.Session, error) {
	sess, ok := r.byEmail[email]
	if !ok {
		return session.Session{}, session.ErrNotFound
	}
	return sess, nil
}

func (r *sessionRepository) Upsert(_ context.Context, sess session.Session) error {
	email := sess.Email

	// If updating, remove the previous access token index for the session.
	if prev, ok := r.byEmail[email]; ok {
		delete(r.byAccessToken, prev.AccessToken)
	}

	r.byEmail[email] = sess
	r.byAccessToken[sess.AccessToken] = email
	return nil
}
