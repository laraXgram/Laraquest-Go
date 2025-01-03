package updates

type StickerSet struct {
    Name string `json:"name"`
    Title string `json:"title"`
    Sticker_type string `json:"sticker_type"`
    Stickers *[]Sticker `json:"stickers"`
    Thumbnail *PhotoSize `json:"thumbnail"`
}