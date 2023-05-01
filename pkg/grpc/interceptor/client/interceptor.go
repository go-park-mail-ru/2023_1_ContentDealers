package user

import (
	"context"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/metrics"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type InterceptorClient struct {
	logger      logging.Logger
	serviceName string
}

func NewInterceptorClient(serviceName string, logger logging.Logger) *InterceptorClient {
	return &InterceptorClient{
		logger:      logger,
		serviceName: serviceName,
	}
}

func (inter *InterceptorClient) AccessLog(
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

	inter.logger.WithFields(logrus.Fields{
		"method":     method,
		"request":    req,
		"request_id": reqID,
	}).Debug(fmt.Sprintf("sent_to_%s_service", inter.serviceName))

	err := invoker(ctx, method, req, reply, cc, opts...)

	elapsedSeconds := time.Since(start)

	inter.logger.WithFields(logrus.Fields{
		// "reply":      reply,
		"time":       fmt.Sprintf("%d mcs", elapsedSeconds.Microseconds()),
		"err":        err,
		"request_id": reqID,
	}).Debug(fmt.Sprintf("returned_from_%s_service", inter.serviceName))

	procedure := method
	status := "OK"
	if err != nil {
		status = "FAIL"
	}

	metrics.GrpcRequestsTotal.WithLabelValues(inter.serviceName, procedure, status).Inc()
	metrics.GrpcRequestsDurationHistorgram.WithLabelValues(inter.serviceName, procedure).Observe(elapsedSeconds.Seconds())

	return err
}
