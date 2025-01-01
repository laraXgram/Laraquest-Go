package updates

type InputSticker struct {
    Sticker *InputFileOrString `json:"sticker"`
    Format string `json:"format"`
    Emoji_list []string `json:"emoji_list"`
    Mask_position *MaskPosition `json:"mask_position"`
    Keywords []string `json:"keywords"`
}