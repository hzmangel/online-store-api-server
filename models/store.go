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

		// Distributor and related relationships
		DistributorLevelsCount int      `json:"distributor_levels_count"`
		DistributorIDs         []string `json:"distributor_ids"`
		DistributorInfoIDs     []string `json:"distributor_info_ids"`
		DistributorsCount      int      `json:"distributors_count"`
		CartIDs                []string `json:"cart_ids"`

		// Manager and retailer
		ManagerIDs []string `json:"manager_ids"`

		// Orders
		OrderIDs   []string `json:"order_ids"`
		OrderSteps []string `json:"order_steps"`

		// Modules
		ModuleIDs []string `json:"module_ids"`

		// Materials
		MaterialIDs         []string `json:"material_ids"`
		MaterialCategoryIDs []string `json:"material_category_ids"`

		// Geo areas
		GeoAreaIDs []string `json:"geo_area_ids"`
	}
)

// ToJSON - Return json byte sequence of store object
func (store Store) ToJSON() []byte {
	storeJSON, _ := json.Marshal(store)
	return storeJSON
}
