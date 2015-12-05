package models

import (
	"encoding/json"
	"time"
)

type (
	// User - all the user infos
	User struct {
		ID string `json:"id"`

		// Login related
		Email              string `json:"email"`
		Mobile             string `json:"mobile"`
		LoginToken         string `json:"login_token"`
		ThirdPartyID       string `json:"third_party_id"`
		ThirdPartyProvider string `json:"third_party_provider"`

		// Authentication related
		EncryptedPassword   string    `json:"-"`
		Salt                string    `json:"logo_path"`
		ResetPasswordToken  string    `json:"-"`
		ResetPasswordSentAt time.Time `json:"-"`

		// Statistic related
		SigninCount     int       `json:"signin_count"`
		CurrentSigninAt time.Time `json:"current_signin_at"`
		LastSigninAt    time.Time `json:"last_signin_at"`
		CurrentSigninIP string    `json:"current_signin_ip"`
		LastSigninIP    string    `json:"last_signin_ip"`

		// Push related
		DeviceTokens  []string `json:"-"`
		DeviceAliases []string `json:"-"`

		// Notifications
		NotificationIDs          []string `json:"notification_ids"`
		AllNotificationCount     int      `json:"all_notification_count"`
		UnreadNotificationCount  int      `json:"unread_notification_count"`
		CheckedNotificationCount int      `json:"checked_notification_count"`

		// Distributor infos
		RetailCompanyIDs   []string `json:"retail_company_ids"`
		DistributorInfoIDs []string `json:"distributor_info_ids"`
		AddressIDs         []string `json:"address_ids"`
		CartIDs            []string `json:"cart_ids"`

		// Manage info
		ManageCompanyID  string   `json:"manage_company_id"`
		ManageRoles      []string `json:"manage_roles"`
		OrderAuditLogIDs []string `json:"order_audit_log_ids"`

		// Some other info
		Timezone string `json:"timezone"`
	}
)

// ToJSON - Return json byte sequence of store object
func (user User) ToJSON() []byte {
	jsonString, _ := json.Marshal(user)
	return jsonString
}
