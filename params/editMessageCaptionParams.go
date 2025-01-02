package params

type EditMessageCaptionParams struct { 
    Chat_id *any `json:"chat_id"`
    Message_id *any `json:"message_id"`
    Inline_message_id *any `json:"inline_message_id"`
    Caption *any `json:"caption"`
    Parse_mode *any `json:"parse_mode"`
    Caption_entities *any `json:"caption_entities"`
    Reply_markup *any `json:"reply_markup"`
    Show_caption_above_media *any `json:"show_caption_above_media"`
    Business_connection_id *any `json:"business_connection_id"`
}