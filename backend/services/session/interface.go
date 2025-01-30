package session

// this is an interface for the session service
// this is used to define the methods that the session service should implement

type ISession interface {
	// this method is used to create a new session
	CreateSession(userID string) (string, error)
	// this method is used to get the user id from the session id
	GetUserID(sessionID string) (string, error)
	// this method is used to delete a session
	DeleteSession(sessionID string) error
	// this method is used to refresh a session
	RefreshSession(sessionID string) (string, error)
}
