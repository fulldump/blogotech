package mongo

import (
	"context"
	"time"

	"github.com/fulldump/box"
	"github.com/globalsign/mgo"
)

// NewSession return a new session
func NewSession(mongoURL string) (*mgo.Session, error) {

	info, err := mgo.ParseURL(mongoURL)
	if err != nil {
		return nil, err
	}

	info.Timeout = 15 * time.Second // TODO: make that configurable
	info.FailFast = true

	s, err := mgo.DialWithInfo(info)
	if nil != err {
		return nil, err
	}

	return s, nil
}

// SessionInterceptor put a mongo session into context at a certain resource node
func SessionInterceptor(session *mgo.Session) box.I {
	return func(next box.H) box.H {
		return func(ctx context.Context) {

			// Ensure a fresh session for each request
			s := session.Clone()
			defer s.Close()

			ctx = SetSession(ctx, s)
			next(ctx)
		}
	}
}

const contextKeySession = "17255078-fb7d-11e9-b7e9-6bd8b9d45680"

// SetSession put a mongo session into context
func SetSession(ctx context.Context, s *mgo.Session) context.Context {
	return context.WithValue(ctx, contextKeySession, s)
}

// GetSession retrieve mongo session from context
func GetSession(ctx context.Context) *mgo.Session {
	s := ctx.Value(contextKeySession)
	if nil == s {
		panic("mongo session should be in context")
	}

	return s.(*mgo.Session)
}
