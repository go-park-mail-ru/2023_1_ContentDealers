package user

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type UseCase struct {
	contentGateway ContentGateway
	userGateway    UserGateway
	sessionGateway SessionGateway
	logger         logging.Logger
}

func NewUseCase(contentGateway ContentGateway, userGateway UserGateway, sessionGateway SessionGateway, logger logging.Logger) *UseCase {
	return &UseCase{
		contentGateway: contentGateway,
		userGateway:    userGateway,
		sessionGateway: sessionGateway,
		logger:         logger,
	}
}

func (uc *UseCase) HasAccessContent(ctx context.Context, originalURI string, cookie string) error {
	if strings.Contains(originalURI, "trailers") {
		return nil
	}
	if cookie == "" {
		return fmt.Errorf("content access only for authorized users")
	}

	session, err := uc.sessionGateway.Get(ctx, cookie)
	if err != nil || session.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf("content access only for authorized users: %w", err)
	}
	ctx = context.WithValue(ctx, "session", session)

	path := strings.Split(originalURI, "/")
	contentIDString := strings.TrimSuffix(path[len(path)-1], ".mp4")
	contentID, err := strconv.Atoi(contentIDString)
	if err != nil {
		seriesAndEpisodes := strings.Split(contentIDString, "_")
		// nolint:gomnd
		if len(seriesAndEpisodes) != 2 {
			return fmt.Errorf("contentID undefined")
		}
		contentID, err = strconv.Atoi(seriesAndEpisodes[0])
		if err != nil {
			return fmt.Errorf("contentID is not numeric")
		}
	}
	content, err := uc.contentGateway.GetContentByContentIDs(ctx, []uint64{uint64(contentID)})
	if err != nil {
		return err
	}
	if len(content) == 0 {
		return fmt.Errorf("content with ID %d is not found", contentID)
	}

	if content[0].IsFree {
		return nil
	}

	user, err := uc.userGateway.GetByID(ctx, session.UserID)
	if err != nil {
		return err
	}
	if user.SubscriptionExpiryDate.Before(time.Now()) {
		return fmt.Errorf("no access to paid content with ID %d", contentID)
	}
	return nil
}
