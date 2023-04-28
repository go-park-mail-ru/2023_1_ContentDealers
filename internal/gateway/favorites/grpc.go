package favorites

import (
	"context"
	"time"

	"github.com/dranikpg/dto-mapper"
	favContentProto "github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/pkg/proto/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Gateway struct {
	logger            logging.Logger
	favContentManager favContentProto.FavoritesContentServiceClient
	interseptor       FavoritesInterceptor
}

func pingServer(ctx context.Context, client favContentProto.FavoritesContentServiceClient) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	_, err := client.Ping(ctx, &favContentProto.PingRequest{})
	if err != nil {
		return err
	}

	return nil
}

func NewGateway(logger logging.Logger, cfg ServiceFavoritesConfig) (*Gateway, error) {
	interseptor := FavoritesInterceptor{logger: logger}

	grcpConn, err := grpc.Dial(
		cfg.Addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interseptor.AccessLog),
	)
	if err != nil {
		logger.Error("cant connect to grpc session service")
		return nil, err
	}

	favContentManager := favContentProto.NewFavoritesContentServiceClient(grcpConn)

	err = pingServer(context.Background(), favContentManager)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &Gateway{logger: logger, favContentManager: favContentManager}, nil
}

func (gate *Gateway) Delete(ctx context.Context, favorite domain.FavoriteContent) error {
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

func (gate *Gateway) Add(ctx context.Context, favorite domain.FavoriteContent) error {
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

func (gate *Gateway) Get(ctx context.Context, options domain.FavoritesOptions) ([]domain.FavoriteContent, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	favOptionsRequest := favContentProto.FavoritesOptions{}
	err := dto.Map(&favOptionsRequest, options)

	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return []domain.FavoriteContent{}, err
	}
	favsResponse, err := gate.favContentManager.GetContent(ctx, &favOptionsRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return []domain.FavoriteContent{}, err
	}
	favs := FavoritesContentDTO{}
	err = dto.Map(&favs, favsResponse)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return []domain.FavoriteContent{}, err
	}
	return favs.Favorites, nil
}
