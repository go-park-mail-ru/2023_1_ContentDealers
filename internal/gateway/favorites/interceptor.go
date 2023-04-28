package favorites

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type FavoritesInterceptor struct {
	logger logging.Logger
}

func (si *FavoritesInterceptor) AccessLog(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	start := time.Now()

	reqID, ok := ctx.Value("requestID").(string)
	if !ok {
		reqID = "unknown"
	}

	si.logger.WithFields(logrus.Fields{
		"method":     method,
		"request":    req,
		"request_id": reqID,
	}).Debug("sent_to_user_service")

	err := invoker(ctx, method, req, reply, cc, opts...)

	si.logger.WithFields(logrus.Fields{
		"reply":      reply,
		"time":       time.Since(start),
		"err":        err,
		"request_id": reqID,
	}).Debug("returned_from_user_service")
	return err
}
