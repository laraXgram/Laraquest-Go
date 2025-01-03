package updates

type ExternalReplyInfo struct {
    Origin *MessageOrigin `json:"origin"`
    Chat *Chat `json:"chat"`
    Message_id int64 `json:"message_id"`
    Link_preview_options *LinkPreviewOptions `json:"link_preview_options"`
    Animation *Animation `json:"animation"`
    Audio *Audio `json:"audio"`
    Document *Document `json:"document"`
    Paid_media *PaidMediaInfo `json:"paid_media"`
    Photo *[]PhotoSize `json:"photo"`
    Sticker *Sticker `json:"sticker"`
    Story *Story `json:"story"`
    Video *Video `json:"video"`
    Video_note *VideoNote `json:"video_note"`
    Voice *Voice `json:"voice"`
    Has_media_spoiler bool `json:"has_media_spoiler"`
    Contact *Contact `json:"contact"`
    Dice *Dice `json:"dice"`
    Game *Game `json:"game"`
    Giveaway *Giveaway `json:"giveaway"`
    Giveaway_winners *GiveawayWinners `json:"giveaway_winners"`
    Invoice *Invoice `json:"invoice"`
    Location *Location `json:"location"`
    Poll *Poll `json:"poll"`
    Venue *Venue `json:"venue"`
}