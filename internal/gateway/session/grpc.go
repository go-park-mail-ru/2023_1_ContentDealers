package session

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/proto/session"
	"google.golang.org/grpc"
)

type Gateway struct {
	logger      logging.Logger
	sessManager session.SessionServiceClient
}

func pingServer(ctx context.Context, client session.SessionServiceClient) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	_, err := client.Ping(ctx, &session.PingRequest{})
	if err != nil {
		return err
	}

	return nil
}

func NewGateway(logger logging.Logger) (Gateway, error) {
	grcpConn, err := grpc.Dial(
		"172.27.195.147:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Error("cant connect to grpc session service")
	}

	sessManager := session.NewSessionServiceClient(grcpConn)

	err = pingServer(context.Background(), sessManager)
	if err != nil {
		logger.Error(err)
		return Gateway{}, err
	}

	return Gateway{logger: logger, sessManager: sessManager}, nil
}

func (gate *Gateway) Create(ctx context.Context, user domain.User) (domain.Session, error) {
	var request session.UserID
	request.ID = user.ID
	sessionResponse, err := gate.sessManager.Create(ctx, &request)
	if err != nil {
		gate.logger.Error(err)
		return domain.Session{}, err
	}

	expireTime, err := time.Parse(time.RFC3339, sessionResponse.ExpiresAt)
	if err != nil {
		return domain.Session{}, err
	}
	sess := domain.Session{
		ID:        sessionResponse.ID,
		UserID:    sessionResponse.UserID,
		ExpiresAt: expireTime,
	}
	return sess, nil
}

func (gate *Gateway) Get(ctx context.Context, id string) (domain.Session, error) {
	var request session.SessionID
	request.ID = id
	sessionResponse, err := gate.sessManager.Get(ctx, &request)
	if err != nil {
		gate.logger.Error(err)
		return domain.Session{}, err
	}
	expireTime, err := time.Parse(time.RFC3339, sessionResponse.ExpiresAt)
	if err != nil {
		return domain.Session{}, err
	}
	sess := domain.Session{
		ID:        sessionResponse.ID,
		UserID:    sessionResponse.UserID,
		ExpiresAt: expireTime,
	}
	return sess, nil
}

func (gate *Gateway) Delete(ctx context.Context, id string) error {
	var request session.SessionID
	request.ID = id
	_, err := gate.sessManager.Delete(ctx, &request)
	if err != nil {
		gate.logger.Error(err)
		return err
	}
	return nil
}
