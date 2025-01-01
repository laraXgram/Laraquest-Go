package updates

type MessageReactionUpdated struct {
    Chat *Chat `json:"chat"`
    Message_id int64 `json:"message_id"`
    User *User `json:"user"`
    Actor_chat *Chat `json:"actor_chat"`
    Date int64 `json:"date"`
    Old_reaction *[]ReactionType `json:"old_reaction"`
    New_reaction *[]ReactionType `json:"new_reaction"`
}