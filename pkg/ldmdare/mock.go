package ldmdare

import (
	"net/url"

	"github.com/stretchr/testify/mock"
)

type LudumApiMock struct {
	mock.Mock
}

func (l *LudumApiMock) GetGameRank(game *LDGame) (int, error) {
	args := l.Called(game)
	return args.Int(0), args.Error(1)
}

func (l *LudumApiMock) GetEventStatsFromGame(game *LDGame) (*LDEvent, error) {
	args := l.Called(game)
	return args.Get(0).(*LDEvent), args.Error(1)
}

func (l *LudumApiMock) GetGameFromURL(userUrl *url.URL) (*LDGame, error) {
	args := l.Called(userUrl)
	return args.Get(0).(*LDGame), args.Error(1)
}
