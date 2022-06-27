package test

import (
	"errors"
	"github.com/Thunderbirrd/ShortLinksService/internal/handler"
	"github.com/Thunderbirrd/ShortLinksService/internal/mocks/mock_repository"
	"github.com/Thunderbirrd/ShortLinksService/internal/repository"
	"github.com/Thunderbirrd/ShortLinksService/internal/service"
	"github.com/Thunderbirrd/ShortLinksService/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var createUrlTestData = models.UrlObject{
	LongUrl: "https://umbraco.com/knowledge-base/http-status-codes/",
}

var createUserRequest = `{"long_url": "https://umbraco.com/knowledge-base/http-status-codes/"}`

var testingError = errors.New("testing error")

func TestCreateShortUrlOk(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(createUserRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUrlRepository(ctrl)

	h := handler.NewHandler(service.NewService(&repository.Repository{UrlRepository: mockRepo}), ":8080")

	mockRepo.EXPECT().CheckLongUrl(createUrlTestData.LongUrl).Return("", nil).Times(1)
	mockRepo.EXPECT().SaveNewUrl(createUrlTestData.LongUrl, gomock.Any()).Return(nil).Times(1)

	if assert.NoError(t, h.CreateShortUrl(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateShortUrlCheckError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(createUserRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUrlRepository(ctrl)

	h := handler.NewHandler(service.NewService(&repository.Repository{UrlRepository: mockRepo}), ":8080")

	mockRepo.EXPECT().CheckLongUrl(createUrlTestData.LongUrl).Return("", testingError).Times(1)

	assert.Error(t, h.CreateShortUrl(c))
}

func TestCreateShortUrlSavingError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(createUserRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUrlRepository(ctrl)

	h := handler.NewHandler(service.NewService(&repository.Repository{UrlRepository: mockRepo}), ":8080")

	mockRepo.EXPECT().CheckLongUrl(createUrlTestData.LongUrl).Return("", nil).Times(1)
	mockRepo.EXPECT().SaveNewUrl(createUrlTestData.LongUrl, gomock.Any()).Return(testingError).Times(1)

	assert.Error(t, h.CreateShortUrl(c))
}
