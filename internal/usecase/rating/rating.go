package favorites

import (
	"context"
	"fmt"

	domainContent "github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainRate "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
)

type UseCase struct {
	gate    Gateway
	session SessionGateway
	logger  logging.Logger
	content ContentGateway
}

func NewUseCase(gate Gateway, session SessionGateway, content ContentGateway, logger logging.Logger) *UseCase {
	return &UseCase{gate: gate, session: session, content: content, logger: logger}
}

func (uc *UseCase) getUserIDByContext(ctx context.Context) (uint64, error) {
	session, ok := ctx.Value("session").(domainSession.Session)
	if !ok {
		return 0, fmt.Errorf("session not found")
	}
	return session.UserID, nil
}

func (uc *UseCase) Delete(ctx context.Context, rating domainRate.Rating) error {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	rating.UserID = userID
	deletedRating, err := uc.gate.DeleteRating(ctx, rating)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	err = uc.content.DeleteRating(ctx, deletedRating.ContentID, deletedRating.Rating)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	return nil
}

func (uc *UseCase) Add(ctx context.Context, rating domainRate.Rating) (float64, uint64, error) {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return 0, 0, err
	}
	rating.UserID = userID
	err = uc.gate.AddRating(ctx, rating)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return 0, 0, err
	}
	err = uc.content.AddRating(ctx, rating.ContentID, rating.Rating)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return 0, 0, err
	}

	content, err := uc.content.GetContentByContentIDs(ctx, []uint64{rating.ContentID})
	if err != nil {
		return 0, 0, err
	}
	return content[0].Rating, content[0].CountRatings, nil
}

func (uc *UseCase) Has(ctx context.Context, rating domainRate.Rating) (domainRate.HasRating, error) {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return domainRate.HasRating{}, err
	}
	rating.UserID = userID
	return uc.gate.HasRating(ctx, rating)
}

func (uc *UseCase) GetByUser(ctx context.Context, options domainRate.RatingsOptions) ([]domainContent.Content, []domainRate.Rating, bool, error) {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return []domainContent.Content{}, []domainRate.Rating{}, false, err
	}
	options.UserID = userID
	ratings, err := uc.gate.GetRatingByUser(ctx, options)
	if err != nil {
		return []domainContent.Content{}, []domainRate.Rating{}, false, err
	}

	contentIDs := make([]uint64, 0, len(ratings.Ratings))
	for _, rate := range ratings.Ratings {
		contentIDs = append(contentIDs, rate.ContentID)
	}

	contentSliceSorted, err := uc.content.GetContentByContentIDs(ctx, contentIDs)
	if err != nil {
		return []domainContent.Content{}, []domainRate.Rating{}, false, err
	}

	// сортировка contentSliceSorted согласно порядку id-шников в contentIDs

	contentDict := make(map[uint64]domainContent.Content)
	for _, item := range contentSliceSorted {
		contentDict[item.ID] = item
	}

	ratingsSlice := []domainRate.Rating{}
	contentSlice := []domainContent.Content{}
	for idx, id := range contentIDs {
		content, ok := contentDict[id]
		// content_id может не существовать, т.к. таблицы не связаны
		if ok {
			ratingsSlice = append(ratingsSlice, ratings.Ratings[idx])
			contentSlice = append(contentSlice, content)
		}
	}

	return contentSlice, ratingsSlice, ratings.IsLast, nil
}

func (uc *UseCase) GetByContent(ctx context.Context, options domainRate.RatingsOptions) (domainRate.Ratings, error) {
	userID, err := uc.getUserIDByContext(ctx)
	if err != nil {
		uc.logger.WithRequestID(ctx).Trace(err)
		return domainRate.Ratings{}, err
	}
	options.UserID = userID
	ratings, err := uc.gate.GetRatingByContent(ctx, options)
	if err != nil {
		return domainRate.Ratings{}, err
	}
	return ratings, nil
}
