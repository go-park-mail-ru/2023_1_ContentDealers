package session

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type SessionInterceptor struct {
	logger logging.Logger
}

func (si *SessionInterceptor) AccessLog(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	// клиент
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	start := time.Now()

	reqID, ok := ctx.Value("requestID").(string)
	if !ok {
		reqID = "unknown"
	}

	err := invoker(ctx, method, req, reply, cc, opts...)
	si.logger.WithFields(logrus.Fields{
		"method":     method,
		"request":    req,
		"reply":      reply,
		"time":       time.Since(start),
		"err":        err,
		"request_id": reqID,
	}).Debug("The request will be sent to the sessions microservice")
	return err
}
