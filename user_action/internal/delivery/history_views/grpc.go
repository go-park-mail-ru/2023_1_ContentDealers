package history_views

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
	viewsProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/proto/history_views"
	"github.com/golang/protobuf/ptypes"
)

type Grpc struct {
	viewsProto.UnimplementedHistoryViewsServiceServer
	viewsUseCase ViewsUseCase
	logger       logging.Logger
}

func NewGrpc(viewsUseCase ViewsUseCase, logger logging.Logger) *Grpc {
	return &Grpc{viewsUseCase: viewsUseCase, logger: logger}
}

func (service *Grpc) UpdateProgressView(ctx context.Context, viewRequest *viewsProto.View) (*viewsProto.Nothing, error) {
	view := domain.View{
		UserID:    viewRequest.UserID,
		ContentID: viewRequest.ContentID,
		StopView:  viewRequest.StopView.AsDuration(),
		Duration:  viewRequest.Duration.AsDuration(),
	}
	err := service.viewsUseCase.UpdateProgressView(ctx, view)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	return &viewsProto.Nothing{}, nil
}

func (service *Grpc) HasView(ctx context.Context, viewRequest *viewsProto.View) (*viewsProto.HasViewMessage, error) {
	view := domain.View{
		UserID:    viewRequest.UserID,
		ContentID: viewRequest.ContentID,
	}
	hasView, err := service.viewsUseCase.HasView(ctx, view)
	if err != nil {
		return nil, err
	}
	hasViewResponse := viewsProto.HasViewMessage{
		HasView: hasView.HasView,
		View: &viewsProto.View{
			UserID:    hasView.View.UserID,
			ContentID: hasView.View.ContentID,
			StopView:  ptypes.DurationProto(hasView.View.StopView),
			Duration:  ptypes.DurationProto(hasView.View.Duration),
		},
	}
	return &hasViewResponse, nil
}

func (service *Grpc) GetPartiallyViewsByUser(ctx context.Context, viewsOptionsRequest *viewsProto.ViewsOptions) (*viewsProto.Views, error) {
	viewsOptions := domain.ViewsOptions{}
	err := dto.Map(&viewsOptions, viewsOptionsRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}

	views, err := service.viewsUseCase.GetPartiallyViewsByUser(ctx, viewsOptions)
	if err != nil {
		return nil, err
	}
	viewsResponse := viewsProto.Views{}
	for _, view := range views.Views {
		viewsResponse.Views = append(viewsResponse.Views, &viewsProto.View{
			UserID:    view.UserID,
			ContentID: view.ContentID,
		})
	}
	viewsResponse.IsLast = views.IsLast

	return &viewsResponse, nil
}

func (service *Grpc) GetAllViewsByUser(ctx context.Context, viewsOptionsRequest *viewsProto.ViewsOptions) (*viewsProto.Views, error) {
	viewsOptions := domain.ViewsOptions{}
	err := dto.Map(&viewsOptions, viewsOptionsRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}

	views, err := service.viewsUseCase.GetAllViewsByUser(ctx, viewsOptions)
	if err != nil {
		return nil, err
	}
	viewsResponse := viewsProto.Views{}
	for _, view := range views.Views {
		viewsResponse.Views = append(viewsResponse.Views, &viewsProto.View{
			UserID:    view.UserID,
			ContentID: view.ContentID,
		})
	}
	viewsResponse.IsLast = views.IsLast

	return &viewsResponse, nil
}
