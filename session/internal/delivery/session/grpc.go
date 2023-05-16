package session

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/proto/session"
)

type Grpc struct {
	session.UnimplementedSessionServiceServer
	sessionUseCase SessionUseCase
	logger         logging.Logger
}

func NewGrpc(sessionUseCase SessionUseCase, logger logging.Logger) *Grpc {
	return &Grpc{sessionUseCase: sessionUseCase, logger: logger}
}

func (service *Grpc) Create(ctx context.Context, userID *session.UserID) (*session.Session, error) {
	sess, err := service.sessionUseCase.Create(ctx, userID.ID)
	if err != nil {
		return nil, err
	}
	response := &session.Session{
		ID:        sess.ID,
		UserID:    sess.UserID,
		ExpiresAt: sess.ExpiresAt.Format(time.RFC3339),
	}
	return response, nil

}

func (service *Grpc) Get(ctx context.Context, sessionID *session.SessionID) (*session.Session, error) {
	sess, err := service.sessionUseCase.Get(ctx, sessionID.ID)
	if err != nil {
		return nil, err
	}
	response := &session.Session{
		ID:        sess.ID,
		UserID:    sess.UserID,
		ExpiresAt: sess.ExpiresAt.Format(time.RFC3339),
	}
	return response, nil
}

func (service *Grpc) Delete(ctx context.Context, sessionID *session.SessionID) (*session.Nothing, error) {
	err := service.sessionUseCase.Delete(ctx, sessionID.ID)
	if err != nil {
		return nil, err
	}
	return &session.Nothing{}, nil
}
