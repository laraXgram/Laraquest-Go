package params

type SetMessageReactionParams struct { 
    Chat_id any `json:"chat_id"`
    Message_id any `json:"message_id"`
    Reaction *any `json:"reaction"`
    Is_big *any `json:"is_big"`
}