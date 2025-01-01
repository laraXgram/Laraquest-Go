package updates

type BusinessIntro struct {
    Title string `json:"title"`
    Message string `json:"message"`
    Sticker *Sticker `json:"sticker"`
}