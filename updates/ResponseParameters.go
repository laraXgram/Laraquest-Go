package updates

type ResponseParameters struct {
    Migrate_to_chat_id int64 `json:"migrate_to_chat_id"`
    Retry_after int64 `json:"retry_after"`
}