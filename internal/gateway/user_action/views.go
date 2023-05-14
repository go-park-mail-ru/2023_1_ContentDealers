package user_action

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
	viewsProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/proto/history_views"
	"google.golang.org/grpc/metadata"

	"github.com/golang/protobuf/ptypes"
)

func (gate *Gateway) UpdateProgressView(ctx context.Context, view domain.View) error {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	viewRequest := viewsProto.View{
		UserID:    view.UserID,
		ContentID: view.ContentID,
		StopView:  ptypes.DurationProto(view.StopView),
		Duration:  ptypes.DurationProto(view.Duration),
	}
	_, err := gate.viewsManager.UpdateProgressView(ctx, &viewRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	return nil
}

func (gate *Gateway) HasView(ctx context.Context, view domain.View) (domain.HasView, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	viewRequest := viewsProto.View{
		UserID:    view.UserID,
		ContentID: view.ContentID,
	}
	hasView, err := gate.viewsManager.HasView(ctx, &viewRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.HasView{}, err
	}
	hasViewResponse := domain.HasView{
		HasView: hasView.HasView,
		View:    view,
	}
	hasViewResponse.View.StopView = hasView.View.StopView.AsDuration()
	hasViewResponse.View.Duration = hasView.View.Duration.AsDuration()

	return hasViewResponse, nil
}

func (gate *Gateway) GetPartiallyViewsByUser(ctx context.Context, options domain.ViewsOptions) (domain.Views, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	viewsOptionsRequest := viewsProto.ViewsOptions{}
	err := dto.Map(&viewsOptionsRequest, options)

	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.Views{}, err
	}
	viewsResponse, err := gate.viewsManager.GetPartiallyViewsByUser(ctx, &viewsOptionsRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.Views{}, err
	}
	views := domain.Views{}
	err = dto.Map(&views, viewsResponse)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		// return domain.Views{}, err
	}
	return views, nil
}

func (gate *Gateway) GetAllViewsByUser(ctx context.Context, options domain.ViewsOptions) (domain.Views, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	viewsOptionsRequest := viewsProto.ViewsOptions{}
	err := dto.Map(&viewsOptionsRequest, options)

	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.Views{}, err
	}
	viewsResponse, err := gate.viewsManager.GetAllViewsByUser(ctx, &viewsOptionsRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.Views{}, err
	}
	views := domain.Views{}
	err = dto.Map(&views, viewsResponse)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		// return domain.Views{}, err
	}
	return views, nil
}
