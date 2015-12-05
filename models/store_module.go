package models

import "encoding/json"

type (
	// StoreModule - Module info for store record
	StoreModule struct {
		ID string `json:"id"`

		// Company name and logo
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Order       int    `json:"order"`
		Version     int    `json:"version"`

		CompanyID string `json:"company_id"`
	}
)

// ToJSON - Return json byte sequence of store object
func (store_m StoreModule) ToJSON() []byte {
	jsonStr, _ := json.Marshal(store_m)
	return jsonStr
}

// NameList: Full name list of modules, which should limit the input of name field
func (store_m StoreModule) NameList() []string {
	return []string{
		"product",
		"aboutus",
		"project",
		"certification",
		"culture",
		"nearshop",
		"navigation",
		"contactus",
	}
}
