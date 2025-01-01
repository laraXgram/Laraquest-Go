package updates

type Sticker struct {
    File_id string `json:"file_id"`
    File_unique_id string `json:"file_unique_id"`
    Type string `json:"type"`
    Width int64 `json:"width"`
    Height int64 `json:"height"`
    Is_animated bool `json:"is_animated"`
    Is_video bool `json:"is_video"`
    Thumbnail *PhotoSize `json:"thumbnail"`
    Emoji string `json:"emoji"`
    Set_name string `json:"set_name"`
    Premium_animation *File `json:"premium_animation"`
    Mask_position *MaskPosition `json:"mask_position"`
    Custom_emoji_id string `json:"custom_emoji_id"`
    Needs_repainting bool `json:"needs_repainting"`
    File_size int64 `json:"file_size"`
}