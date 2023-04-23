package user

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (service *Grpc) LogInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	var reqID string
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists || len(md.Get("requestId")) == 0 {
		reqID = "unknown"
	} else {
		reqID = md.Get("requestId")[0]
	}

	ctx = context.WithValue(ctx, "requestID", reqID)

	reply, err := handler(ctx, req)

	service.logger.WithFields(logrus.Fields{
		"info_full_method": info.FullMethod,
		"request":          req,
		"reply":            reply,
		"time":             time.Since(start),
		"request_id":       reqID,
		"metadata":         md,
		"err":              err,
	}).Debug("The request arrived at the user microservice")
	return reply, err
}
