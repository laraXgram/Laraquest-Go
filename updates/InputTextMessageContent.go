package updates

type InputTextMessageContent struct {
    Message_text string `json:"message_text"`
    Parse_mode string `json:"parse_mode"`
    Entities *[]MessageEntity `json:"entities"`
    Link_preview_options *LinkPreviewOptions `json:"link_preview_options"`
}