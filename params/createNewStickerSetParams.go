package params

type CreateNewStickerSetParams struct { 
    User_id any `json:"user_id"`
    Name any `json:"name"`
    Title any `json:"title"`
    Stickers any `json:"stickers"`
    Sticker_type *any `json:"sticker_type"`
    Needs_repainting *any `json:"needs_repainting"`
}