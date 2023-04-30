package server

import (
	"context"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type InterceptorServer struct {
	logger      logging.Logger
	serviceName string
}

func NewInterceptorServer(serviceName string, logger logging.Logger) *InterceptorServer {
	return &InterceptorServer{
		logger:      logger,
		serviceName: serviceName,
	}
}

func (inter *InterceptorServer) AccessLog(
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

	inter.logger.WithFields(logrus.Fields{
		"info_full_method": info.FullMethod,
		"request":          req,
		"request_id":       reqID,
		"metadata":         md,
	}).Debug(fmt.Sprintf("accepted_by_%s_service", inter.serviceName))

	reply, err := handler(ctx, req)

	inter.logger.WithFields(logrus.Fields{
		// "reply":      reply,
		"time":       fmt.Sprintf("%d mcs", time.Since(start).Microseconds()),
		"request_id": reqID,
		"err":        err,
	}).Debug(fmt.Sprintf("released_by_%s_service", inter.serviceName))
	return reply, err
}
