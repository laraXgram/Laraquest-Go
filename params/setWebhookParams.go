package params

type SetWebhookParams struct { 
    Url any `json:"url"`
    Certificate *any `json:"certificate"`
    Ip_address *any `json:"ip_address"`
    Max_connections *any `json:"max_connections"`
    Allowed_updates *any `json:"allowed_updates"`
    Drop_pending_updates *any `json:"drop_pending_updates"`
    Secret_token *any `json:"secret_token"`
}