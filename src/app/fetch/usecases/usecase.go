package usecases

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/onemariovalentino/fetch-webpage/src/app/fetch/models"
	"github.com/onemariovalentino/fetch-webpage/src/app/fetch/repositories"
	"github.com/onemariovalentino/fetch-webpage/src/pkg/utils/filename"
	"github.com/onemariovalentino/fetch-webpage/src/pkg/utils/htmlparser"
)

type (
	FetchUsecaseInterface interface {
		DownloadPage(cx context.Context, urls []string) error
		GetMetadata(cx context.Context, url string) (*models.Metadata, error)
	}

	FetchUsecase struct {
		repository repositories.FetchRepositoryInterface
	}
)

func New(repository repositories.FetchRepositoryInterface) FetchUsecaseInterface {
	return &FetchUsecase{repository: repository}
}

func (u *FetchUsecase) DownloadPage(cx context.Context, urls []string) error {
	wg := sync.WaitGroup{}
	errCh := make(chan error, len(urls))
	metadataCh := make(chan *models.Metadata, len(urls))

	for _, url := range urls {
		wg.Add(1)

		go func(url string, metadataCh chan<- *models.Metadata, errCh chan<- error) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				errCh <- fmt.Errorf("error fetching URL: %v", err)
				return
			}
			if resp.StatusCode != http.StatusOK {
				errCh <- fmt.Errorf("failed to fetch URL: %s", resp.Status)
				return
			}
			defer resp.Body.Close()
			page, err := io.ReadAll(resp.Body)
			if err != nil {
				errCh <- fmt.Errorf("error read response body: %s", err.Error())
				return
			}

			getFileName := filename.GetFileName(url)
			file, err := os.OpenFile("files/html/"+getFileName, os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				errCh <- fmt.Errorf("error creating file:%s", err.Error())
				return
			}
			defer file.Close()
			_, err = io.Copy(file, strings.NewReader(string(page)))
			if err != nil {
				errCh <- fmt.Errorf("error saving response body: %s", err.Error())
				return
			}

			htmlParser := htmlparser.New(string(page))
			numLinks, numImage, err := htmlParser.GetNumLinksAndImages()
			if err != nil {
				errCh <- fmt.Errorf("error parsing html content: %s", err.Error())
				return
			}
			metadata := &models.Metadata{
				URL:       url,
				NumLinks:  numLinks,
				NumImages: numImage,
				LastFetch: time.Now().UTC(),
			}

			metadataCh <- metadata

		}(url, metadataCh, errCh)
	}
	wg.Wait()

	close(errCh)
	for err := range errCh {
		if err != nil {
			return err
		}
	}

	close(metadataCh)
	result := map[string]*models.Metadata{}
	for metadata := range metadataCh {
		result[metadata.URL] = metadata
	}

	err := u.repository.SaveToJSON(cx, result)
	if err != nil {
		return err
	}

	return nil
}

func (u *FetchUsecase) GetMetadata(cx context.Context, url string) (*models.Metadata, error) {
	currentData, err := u.repository.LoadFromJSON()
	if err != nil {
		return nil, err
	}

	if _, ok := currentData[url]; !ok {
		return nil, errors.New(`data not found`)
	}

	return currentData[url], nil
}
