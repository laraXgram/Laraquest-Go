package updates

type BusinessMessagesDeleted struct {
    Business_connection_id string `json:"business_connection_id"`
    Chat *Chat `json:"chat"`
    Message_ids []int64 `json:"message_ids"`
}