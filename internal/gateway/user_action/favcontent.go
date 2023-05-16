package user_action

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	interceptorClient "github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/interceptor/client"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/ping"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
	favContentProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/proto/favcontent"
	rateProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/proto/rating"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Gateway struct {
	logger            logging.Logger
	favContentManager favContentProto.FavoritesContentServiceClient
	ratingManager     rateProto.RatingServiceClient
}

func NewGateway(logger logging.Logger, cfg ServiceUserActionConfig) (*Gateway, error) {
	interceptor := interceptorClient.NewInterceptorClient("favorites", logger)

	grpcConn, err := grpc.Dial(
		cfg.Addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.AccessLog),
	)
	if err != nil {
		logger.Error("cant connect to grpc session service")
		return nil, err
	}

	favContentManager := favContentProto.NewFavoritesContentServiceClient(grpcConn)
	ratingManager := rateProto.NewRatingServiceClient(grpcConn)

	err = ping.Ping(grpcConn)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &Gateway{
		logger:            logger,
		favContentManager: favContentManager,
		ratingManager:     ratingManager,
	}, nil
}

func (gate *Gateway) DeleteFavContent(ctx context.Context, favorite domain.FavoriteContent) error {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	favRequest := favContentProto.Favorite{}
	err := dto.Map(&favRequest, favorite)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	_, err = gate.favContentManager.DeleteContent(ctx, &favRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	return nil
}

func (gate *Gateway) AddFavContent(ctx context.Context, favorite domain.FavoriteContent) error {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	favRequest := favContentProto.Favorite{}
	err := dto.Map(&favRequest, favorite)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	_, err = gate.favContentManager.AddContent(ctx, &favRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	return nil
}

func (gate *Gateway) HasFavContent(ctx context.Context, favorite domain.FavoriteContent) (bool, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	favRequest := favContentProto.Favorite{}
	err := dto.Map(&favRequest, favorite)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return false, err
	}
	hasFav, err := gate.favContentManager.HasFavContent(ctx, &favRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return false, err
	}
	return hasFav.HasFav, nil
}

func (gate *Gateway) GetFavContent(ctx context.Context, options domain.FavoritesOptions) (domain.FavoritesContent, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	favOptionsRequest := favContentProto.FavoritesOptions{}
	err := dto.Map(&favOptionsRequest, options)

	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.FavoritesContent{}, err
	}
	favsResponse, err := gate.favContentManager.GetContent(ctx, &favOptionsRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.FavoritesContent{}, err
	}
	favs := domain.FavoritesContent{}
	err = dto.Map(&favs, favsResponse)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.FavoritesContent{}, err
	}
	return favs, nil
}
