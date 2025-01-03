package updates

type MessageReactionCountUpdated struct {
    Chat *Chat `json:"chat"`
    Message_id int64 `json:"message_id"`
    Date int64 `json:"date"`
    Reactions *[]ReactionCount `json:"reactions"`
}