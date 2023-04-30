package session

import (
	"context"
	"time"

	interceptorClient "github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/interceptor/client"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/proto/session"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

func NewGateway(logger logging.Logger, cfg ServiceSessionConfig) (*Gateway, error) {
	interceptor := interceptorClient.NewInterceptorClient("session", logger)

	grcpConn, err := grpc.Dial(
		cfg.Addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.AccessLog),
	)
	if err != nil {
		logger.Error("cant connect to grpc session service")
		return nil, err
	}

	sessManager := session.NewSessionServiceClient(grcpConn)

	err = pingServer(context.Background(), sessManager)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &Gateway{logger: logger, sessManager: sessManager}, nil
}

func (gate *Gateway) Create(ctx context.Context, user domainUser.User) (domainSession.Session, error) {
	var request session.UserID
	request.ID = user.ID
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	sessionResponse, err := gate.sessManager.Create(ctx, &request)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domainSession.Session{}, err
	}

	expireTime, err := time.Parse(time.RFC3339, sessionResponse.ExpiresAt)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domainSession.Session{}, err
	}
	sess := domainSession.Session{
		ID:        sessionResponse.ID,
		UserID:    sessionResponse.UserID,
		ExpiresAt: expireTime,
	}
	return sess, nil
}

func (gate *Gateway) Get(ctx context.Context, id string) (domainSession.Session, error) {
	var request session.SessionID
	request.ID = id
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	sessionResponse, err := gate.sessManager.Get(ctx, &request)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domainSession.Session{}, err
	}
	expireTime, err := time.Parse(time.RFC3339, sessionResponse.ExpiresAt)
	if err != nil {
		return domainSession.Session{}, err
	}
	sess := domainSession.Session{
		ID:        sessionResponse.ID,
		UserID:    sessionResponse.UserID,
		ExpiresAt: expireTime,
	}
	return sess, nil
}

func (gate *Gateway) Delete(ctx context.Context, id string) error {
	var request session.SessionID
	request.ID = id
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := gate.sessManager.Delete(ctx, &request)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	return nil
}
