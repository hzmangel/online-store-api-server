package models

import "encoding/json"

type (
	// Store - store info
	Store struct {
		ID string `json:"id"`

		// Company name and logo
		Name       string `json:"name"`
		LogoPath   string `json:"logo_path"`
		LogoWidth  int    `json:"logo_width"`
		LogoHeight int    `json:"logo_height"`

		// QRcode image for company
		QrcodePath   string `json:"qrcode_path"`
		QrcodeWidth  int    `json:"qrcode_width"`
		QrcodeHeight int    `json:"qrcode_height"`

		// Product relationship
		ProductIDs    []string `json:"product_ids"`
		ProductsCount int      `json:"products_count"`

		// Items for saving spec and record seller's history
		CompanyItemIDs []string `json:"company_item_ids"`
		SellerItemIDs  []string `json:"seller_item_ids"`
		GoodsRecordIDs []string `json:"goods_record_ids"`

		// Product category relationship
		ProductCategoryIDs     []string `json:"product_category_ids"`
		ProductCategoriesCount int      `json:"product_categories_count"`

		// Distributor relationships
		DistributorsCount      int `json:"distributors_count"`
		DistributorLevelsCount int `json:"distributor_levels_count"`
	}
)

// ToJSON - Return json byte sequence of store object
func (store Store) ToJSON() []byte {
	storeJSON, _ := json.Marshal(store)
	return storeJSON
}
