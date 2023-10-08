package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	ID        int    `json:"id" gorm:"column:id"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width"`
	Height    int    `json:"height" gorm:"column:height"`
	CloudName string `json:"cloudName,omitempty" gorm:"-"`
	Extension string `json:"extenstion,omitempty" gorm:"-"`
}

func (Image) TableName() string { return "images" }

// DB => Application
func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return ErrInternal(errors.New(fmt.Sprint("failed to unmarshal JSONB value: ", value)))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return ErrInternal(errors.New(fmt.Sprint("failed to unmarshal JSONB value: ", value)))
	}

	*j = img
	return nil
}

// Application => DB
func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

type Images []Image

// DB => Application
func (j *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return ErrInternal(errors.New(fmt.Sprint("failed to unmarshal JSONB value: ", value)))
	}

	var img []Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return ErrInternal(errors.New(fmt.Sprint("failed to unmarshal JSONB value: ", value)))
	}

	*j = img
	return nil
}

// Application => DB
func (j *Images) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}
