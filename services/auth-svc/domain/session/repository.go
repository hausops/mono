package session

import "context"

// Repository interface declares the behavior this package needs
// to perists and retrieve data related to users' login sessions.
type Repository interface {
	// DeleteByAccessToken deletes session with the given access token,
	// or an error if the session was not found.
	DeleteByAccessToken(context.Context, AccessToken) error

	// FindByToken retrieves a session based on the given access token,
	// or an error if the session was not found.
	FindByAccessToken(context.Context, AccessToken) (Session, error)

	// FindByUserID retrieves a session based on the given user ID,
	// or an error if the session was not found.
	FindByUserID(ctx context.Context, userID string) (Session, error)

	// Upsert adds the session to the repository if it does not exist,
	// or replaces the stored session with the same access token (without merging).
	Upsert(context.Context, Session) error
}
