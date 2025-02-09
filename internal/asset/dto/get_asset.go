package dto

import "time"

type GetAssetOutput struct {
	Name      string    `json:"name"`
	Data      []byte    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAssetInput struct {
	Name string
}

func (i *GetAssetInput) Validate() error {

	return nil
}
