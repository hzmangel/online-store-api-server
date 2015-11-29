package models

import "encoding/json"

type (
	// Store - store info
	Store struct {
		ID string `json:"id"`

		Name       string `json:"name"`
		LogoPath   string `json:"logo_path"`
		LogoWidth  int    `json:"logo_width"`
		LogoHeight int    `json:"logo_height"`

		QrcodePath   string `json:"qrcode_path"`
		QrcodeWidth  int    `json:"qrcode_width"`
		QrcodeHeight int    `json:"qrcode_height"`

		ProductsCount          int
		DistributorsCount      int
		DistributorLevelsCount int
	}
)

// ToJSON - Return json byte sequence of store object
func (store Store) ToJSON() []byte {
	storeJSON, _ := json.Marshal(store)
	return storeJSON
}
