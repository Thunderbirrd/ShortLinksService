package test

import (
	"errors"
	"github.com/Thunderbirrd/ShortLinksService/internal/mocks/mock_repository"
	"github.com/Thunderbirrd/ShortLinksService/pkg/models"
	"github.com/golang/mock/gomock"
	"testing"
)

var createUrlTestData = models.UrlObject{
	LongUrl:  "https://umbraco.com/knowledge-base/http-status-codes/",
	ShortUrl: "PmtkjpQYBS",
}

var testingError = errors.New("testing error")

func TestCreateShortUrl(t *testing.T) {
	//e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUrlRepository(ctrl)

	mockRepo.EXPECT().CheckLongUrl(createUrlTestData.LongUrl).Return("", nil).Times(1)
	mockRepo.EXPECT().SaveNewUrl(createUrlTestData).Return(nil).Times(1)

	mockRepo.EXPECT().CheckLongUrl(createUrlTestData.LongUrl).Return("", testingError).Times(1)

	mockRepo.EXPECT().SaveNewUrl(createUrlTestData).Return(testingError).Times(1)

}
