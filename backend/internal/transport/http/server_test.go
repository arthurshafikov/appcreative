package http

import (
	"context"
	"testing"
	"time"

	mock_server "github.com/arthurshafikov/appcreative/backend/internal/transport/http/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestServe(t *testing.T) {
	ctrl := gomock.NewController(t)
	handlerMock := mock_server.NewMockHandler(ctrl)
	loggerMock := mock_server.NewMockLogger(ctrl)
	server := NewServer(handlerMock, loggerMock)
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)
	gomock.InOrder(
		handlerMock.EXPECT().InitRoutes(gomock.Any()),
		loggerMock.EXPECT().Info("Shutdown Server ..."),
	)
	group.Go(func() error {
		time.Sleep(time.Second / 100)
		cancel()

		return nil
	})

	server.Serve(ctx, group, "9999")

	require.Equal(t, ":9999", server.httpSrv.Addr)
	require.NoError(t, group.Wait())
}
