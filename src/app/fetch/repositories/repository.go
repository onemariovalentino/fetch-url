package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/onemariovalentino/fetch-webpage/src/app/fetch/models"
)

type (
	FetchRepositoryInterface interface {
		SaveToJSON(cx context.Context, data map[string]*models.Metadata) error
		LoadFromJSON() (map[string]*models.Metadata, error)
	}

	FetchRepository struct {
		filename string
	}
)

func New(filename string) FetchRepositoryInterface {
	return &FetchRepository{filename: filename}
}

func (r *FetchRepository) SaveToJSON(cx context.Context, data map[string]*models.Metadata) error {
	currentData := map[string]*models.Metadata{}
	_, err := os.Stat(r.filename)
	if os.IsNotExist(err) {
		file, err := os.OpenFile(r.filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
	} else {
		if err != nil {
			return err
		}

		currentData, err = r.LoadFromJSON()
		if err != nil {
			return err
		}
	}

	for k, v := range data {
		if _, ok := currentData[k]; ok {
			currentData[k] = v
		} else {
			currentData[k] = v
		}
	}

	jsonData, err := json.MarshalIndent(currentData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON:%s", err.Error())

	}

	err = os.WriteFile(r.filename, jsonData, 0644)
	if err != nil {
		fmt.Println("error write:", err)
		return err
	}

	return nil
}

func (r *FetchRepository) LoadFromJSON() (map[string]*models.Metadata, error) {
	file, err := os.ReadFile(r.filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file:%s", err.Error())
	}

	var data map[string]*models.Metadata
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON:%s", err.Error())
	}

	return data, nil
}
