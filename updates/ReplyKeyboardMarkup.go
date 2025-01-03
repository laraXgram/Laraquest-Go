package updates

type ReplyKeyboardMarkup struct {
    Keyboard *[][]KeyboardButton `json:"keyboard"`
    Is_persistent bool `json:"is_persistent"`
    Resize_keyboard bool `json:"resize_keyboard"`
    One_time_keyboard bool `json:"one_time_keyboard"`
    Input_field_placeholder string `json:"input_field_placeholder"`
    Selective bool `json:"selective"`
}