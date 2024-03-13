package tests

import (
	"context"
	"testing"

	"github.com/onemariovalentino/fetch-webpage/src/pkg/di"
)

func TestFetch(t *testing.T) {
	services := di.New()

	urls := []string{"https://www.okadoc.com", "https://www.wikipedia.com"}

	err := services.FetchUsecase.DownloadPage(context.Background(), urls)
	if err != nil {
		t.Fatalf("failed to download page: %s", err.Error())
	}
}

func TestMetadata(t *testing.T) {
	services := di.New()

	url := "https://www.okadoc.com"

	result, err := services.FetchUsecase.GetMetadata(context.Background(), url)
	if err != nil {
		t.Fatalf("failed to download page: %s", err.Error())
	}

	if result.NumLinks != 18 {
		t.Fatalf(`response num links is not same`)
	}
}
