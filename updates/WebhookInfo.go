package updates

type WebhookInfo struct {
    Url string `json:"url"`
    Has_custom_certificate bool `json:"has_custom_certificate"`
    Pending_update_count int64 `json:"pending_update_count"`
    Ip_address string `json:"ip_address"`
    Last_error_date int64 `json:"last_error_date"`
    Last_error_message string `json:"last_error_message"`
    Last_synchronization_error_date int64 `json:"last_synchronization_error_date"`
    Max_connections int64 `json:"max_connections"`
    Allowed_updates []string `json:"allowed_updates"`
}