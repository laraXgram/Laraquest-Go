package updates

type ReplyKeyboardRemove struct {
    Remove_keyboard bool `json:"remove_keyboard"`
    Selective bool `json:"selective"`
}