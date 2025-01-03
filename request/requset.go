package request

import (
	"github.com/laraXgram/Laraquest-Go/params"
	"reflect"
)

func structToMap(s interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(s)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return result
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Tag.Get("json")

		if field.IsValid() && (field.Kind() != reflect.Ptr || !field.IsNil()) {
			result[fieldName] = field.Interface()
		}
	}

	return result
}

var c *Curl = NewCurl("****", "http://127.0.0.1:8081")

func GetUpdates(params params.GetUpdatesParams) (map[string]interface{}, error) {
	return c.Endpoint("getUpdates", structToMap(params), false, true)
}

func SetWebhook(params params.SetWebhookParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setWebhook", structToMap(params), true, response)
}

func DeleteWebhook(params params.DeleteWebhookParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("deleteWebhook", structToMap(params), true, response)
}

func GetWebhookInfo() (map[string]interface{}, error) {
	return c.Endpoint("getWebhookInfo", structToMap(nil), false, true)
}

func GetMe() (map[string]interface{}, error) {
	return c.Endpoint("getMe", structToMap(nil), false, true)
}

func LogOut(response bool) (map[string]interface{}, error) {
	return c.Endpoint("logOut", structToMap(nil), true, response)
}

func Close(response bool) (map[string]interface{}, error) {
	return c.Endpoint("close", structToMap(nil), true, response)
}

func SendMessage(params params.SendMessageParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendMessage", structToMap(params), true, response)
}

func ForwardMessage(params params.ForwardMessageParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("forwardMessage", structToMap(params), true, response)
}

func ForwardMessages(params params.ForwardMessagesParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("forwardMessages", structToMap(params), true, response)
}

func CopyMessage(params params.CopyMessageParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("copyMessage", structToMap(params), true, response)
}

func CopyMessages(params params.CopyMessagesParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("copyMessages", structToMap(params), true, response)
}

func SendPhoto(params params.SendPhotoParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendPhoto", structToMap(params), true, response)
}

func SendAudio(params params.SendAudioParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendAudio", structToMap(params), true, response)
}

func SendDocument(params params.SendDocumentParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendDocument", structToMap(params), true, response)
}

func SendVideo(params params.SendVideoParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendVideo", structToMap(params), true, response)
}

func SendAnimation(params params.SendAnimationParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendAnimation", structToMap(params), true, response)
}

func SendVoice(params params.SendVoiceParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendVoice", structToMap(params), true, response)
}

func SendVideoNote(params params.SendVideoNoteParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendVideoNote", structToMap(params), true, response)
}

func SendPaidMedia(params params.SendPaidMediaParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendPaidMedia", structToMap(params), true, response)
}

func SendMediaGroup(params params.SendMediaGroupParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendMediaGroup", structToMap(params), true, response)
}

func SendLocation(params params.SendLocationParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendLocation", structToMap(params), true, response)
}

func SendVenue(params params.SendVenueParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendVenue", structToMap(params), true, response)
}

func SendContact(params params.SendContactParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendContact", structToMap(params), true, response)
}

func SendPoll(params params.SendPollParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendPoll", structToMap(params), true, response)
}

func SendDice(params params.SendDiceParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendDice", structToMap(params), true, response)
}

func SendChatAction(params params.SendChatActionParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendChatAction", structToMap(params), true, response)
}

func SetMessageReaction(params params.SetMessageReactionParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setMessageReaction", structToMap(params), true, response)
}

func GetUserProfilePhotos(params params.GetUserProfilePhotosParams) (map[string]interface{}, error) {
	return c.Endpoint("getUserProfilePhotos", structToMap(params), false, true)
}

func GetFile(params params.GetFileParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("getFile", structToMap(params), true, response)
}

func BanChatMember(params params.BanChatMemberParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("banChatMember", structToMap(params), true, response)
}

func UnbanChatMember(params params.UnbanChatMemberParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("unbanChatMember", structToMap(params), true, response)
}

func RestrictChatMember(params params.RestrictChatMemberParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("restrictChatMember", structToMap(params), true, response)
}

func PromoteChatMember(params params.PromoteChatMemberParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("promoteChatMember", structToMap(params), true, response)
}

func SetChatAdministratorCustomTitle(params params.SetChatAdministratorCustomTitleParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setChatAdministratorCustomTitle", structToMap(params), true, response)
}

func BanChatSenderChat(params params.BanChatSenderChatParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("banChatSenderChat", structToMap(params), true, response)
}

func UnbanChatSenderChat(params params.UnbanChatSenderChatParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("unbanChatSenderChat", structToMap(params), true, response)
}

func SetChatPermissions(params params.SetChatPermissionsParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setChatPermissions", structToMap(params), true, response)
}

func ExportChatInviteLink(params params.ExportChatInviteLinkParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("exportChatInviteLink", structToMap(params), true, response)
}

func CreateChatInviteLink(params params.CreateChatInviteLinkParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("createChatInviteLink", structToMap(params), true, response)
}

func EditChatInviteLink(params params.EditChatInviteLinkParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("editChatInviteLink", structToMap(params), true, response)
}

func RevokeChatInviteLink(params params.RevokeChatInviteLinkParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("revokeChatInviteLink", structToMap(params), true, response)
}

func ApproveChatJoinRequest(params params.ApproveChatJoinRequestParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("approveChatJoinRequest", structToMap(params), true, response)
}

func DeclineChatJoinRequest(params params.DeclineChatJoinRequestParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("declineChatJoinRequest", structToMap(params), true, response)
}

func SetChatPhoto(params params.SetChatPhotoParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setChatPhoto", structToMap(params), true, response)
}

func DeleteChatPhoto(params params.DeleteChatPhotoParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("deleteChatPhoto", structToMap(params), true, response)
}

func SetChatTitle(params params.SetChatTitleParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setChatTitle", structToMap(params), true, response)
}

func SetChatDescription(params params.SetChatDescriptionParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setChatDescription", structToMap(params), true, response)
}

func PinChatMessage(params params.PinChatMessageParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("pinChatMessage", structToMap(params), true, response)
}

func UnpinChatMessage(params params.UnpinChatMessageParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("unpinChatMessage", structToMap(params), true, response)
}

func UnpinAllChatMessages(params params.UnpinAllChatMessagesParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("unpinAllChatMessages", structToMap(params), true, response)
}

func LeaveChat(params params.LeaveChatParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("leaveChat", structToMap(params), true, response)
}

func GetChat(params params.GetChatParams) (map[string]interface{}, error) {
	return c.Endpoint("getChat", structToMap(params), true, true)
}

func GetChatAdministrators(params params.GetChatAdministratorsParams) (map[string]interface{}, error) {
	return c.Endpoint("getChatAdministrators", structToMap(params), true, true)
}

func GetChatMemberCount(params params.GetChatMemberCountParams) (map[string]interface{}, error) {
	return c.Endpoint("getChatMemberCount", structToMap(params), true, true)
}

func GetChatMember(params params.GetChatMemberParams) (map[string]interface{}, error) {
	return c.Endpoint("getChatMember", structToMap(params), true, true)
}

func SetChatStickerSet(params params.SetChatStickerSetParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setChatStickerSet", structToMap(params), true, response)
}

func DeleteChatStickerSet(params params.DeleteChatStickerSetParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("deleteChatStickerSet", structToMap(params), true, response)
}

func GetForumTopicIconStickers() (map[string]interface{}, error) {
	return c.Endpoint("getForumTopicIconStickers", structToMap(nil), true, true)
}

func CreateForumTopic(params params.CreateForumTopicParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("createForumTopic", structToMap(params), true, response)
}

func EditForumTopic(params params.EditForumTopicParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("editForumTopic", structToMap(params), true, response)
}

func CloseForumTopic(params params.CloseForumTopicParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("closeForumTopic", structToMap(params), true, response)
}

func ReopenForumTopic(params params.ReopenForumTopicParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("reopenForumTopic", structToMap(params), true, response)
}

func DeleteForumTopic(params params.DeleteForumTopicParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("deleteForumTopic", structToMap(params), true, response)
}

func UnpinAllForumTopicMessages(params params.UnpinAllForumTopicMessagesParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("unpinAllForumTopicMessages", structToMap(params), true, response)
}

func EditGeneralForumTopic(params params.EditGeneralForumTopicParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("editGeneralForumTopic", structToMap(params), true, response)
}

func CloseGeneralForumTopic(params params.CloseGeneralForumTopicParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("closeGeneralForumTopic", structToMap(params), true, response)
}

func ReopenGeneralForumTopic(params params.ReopenGeneralForumTopicParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("reopenGeneralForumTopic", structToMap(params), true, response)
}

func HideGeneralForumTopic(params params.HideGeneralForumTopicParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("hideGeneralForumTopic", structToMap(params), true, response)
}

func UnhideGeneralForumTopic(params params.UnhideGeneralForumTopicParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("unhideGeneralForumTopic", structToMap(params), true, response)
}

func UnpinAllGeneralForumTopicMessages(params params.UnpinAllGeneralForumTopicMessagesParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("unpinAllGeneralForumTopicMessages", structToMap(params), true, response)
}

func AnswerCallbackQuery(params params.AnswerCallbackQueryParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("answerCallbackQuery", structToMap(params), true, response)
}

func GetUserChatBoosts(params params.GetUserChatBoostsParams) (map[string]interface{}, error) {
	return c.Endpoint("getUserChatBoosts", structToMap(params), true, true)
}

func GetBusinessConnection(params params.GetBusinessConnectionParams) (map[string]interface{}, error) {
	return c.Endpoint("getBusinessConnection", structToMap(params), true, true)
}

func SetMyCommands(params params.SetMyCommandsParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setMyCommands", structToMap(params), true, response)
}

func DeleteMyCommands(params params.DeleteMyCommandsParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("deleteMyCommands", structToMap(params), true, response)
}

func GetMyCommands(params params.GetMyCommandsParams) (map[string]interface{}, error) {
	return c.Endpoint("getMyCommands", structToMap(params), true, true)
}

func SetMyName(params params.SetMyNameParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setMyName", structToMap(params), true, response)
}

func GetMyName(params params.GetMyNameParams) (map[string]interface{}, error) {
	return c.Endpoint("getMyName", structToMap(params), true, true)
}

func SetMyDescription(params params.SetMyDescriptionParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setMyDescription", structToMap(params), true, response)
}

func GetMyDescription(params params.GetMyDescriptionParams) (map[string]interface{}, error) {
	return c.Endpoint("getMyDescription", structToMap(params), true, true)
}

func SetMyShortDescription(params params.SetMyShortDescriptionParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setMyShortDescription", structToMap(params), true, response)
}

func GetMyShortDescription(params params.GetMyShortDescriptionParams) (map[string]interface{}, error) {
	return c.Endpoint("getMyShortDescription", structToMap(params), true, true)
}

func SetChatMenuButton(params params.SetChatMenuButtonParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setChatMenuButton", structToMap(params), true, response)
}

func GetChatMenuButton(params params.GetChatMenuButtonParams) (map[string]interface{}, error) {
	return c.Endpoint("getChatMenuButton", structToMap(params), true, true)
}

func SetMyDefaultAdministratorRights(params params.SetMyDefaultAdministratorRightsParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setMyDefaultAdministratorRights", structToMap(params), true, response)
}

func GetMyDefaultAdministratorRights(params params.GetMyDefaultAdministratorRightsParams) (map[string]interface{}, error) {
	return c.Endpoint("getMyDefaultAdministratorRights", structToMap(params), true, true)
}

func EditMessageText(params params.EditMessageTextParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("editMessageText", structToMap(params), true, response)
}

func EditMessageCaption(params params.EditMessageCaptionParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("editMessageCaption", structToMap(params), true, response)
}

func EditMessageMedia(params params.EditMessageMediaParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("editMessageMedia", structToMap(params), true, response)
}

func EditMessageLiveLocation(params params.EditMessageLiveLocationParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("editMessageLiveLocation", structToMap(params), true, response)
}

func StopMessageLiveLocation(params params.StopMessageLiveLocationParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("stopMessageLiveLocation", structToMap(params), true, response)
}

func EditMessageReplyMarkup(params params.EditMessageReplyMarkupParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("editMessageReplyMarkup", structToMap(params), true, response)
}

func StopPoll(params params.StopPollParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("stopPoll", structToMap(params), true, response)
}

func DeleteMessage(params params.DeleteMessageParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("deleteMessage", structToMap(params), true, response)
}

func DeleteMessages(params params.DeleteMessagesParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("deleteMessages", structToMap(params), true, response)
}

func SendSticker(params params.SendStickerParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendSticker", structToMap(params), true, response)
}

func GetStickerSet(params params.GetStickerSetParams) (map[string]interface{}, error) {
	return c.Endpoint("getStickerSet", structToMap(params), true, true)
}

func GetCustomEmojiStickers(params params.GetCustomEmojiStickersParams) (map[string]interface{}, error) {
	return c.Endpoint("getCustomEmojiStickers", structToMap(params), true, true)
}

func UploadStickerFile(params params.UploadStickerFileParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("uploadStickerFile", structToMap(params), true, response)
}

func CreateNewStickerSet(params params.CreateNewStickerSetParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("createNewStickerSet", structToMap(params), true, response)
}

func AddStickerToSet(params params.AddStickerToSetParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("addStickerToSet", structToMap(params), true, response)
}

func SetStickerPositionInSet(params params.SetStickerPositionInSetParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setStickerPositionInSet", structToMap(params), true, response)
}

func DeleteStickerFromSet(params params.DeleteStickerFromSetParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("deleteStickerFromSet", structToMap(params), true, response)
}

func ReplaceStickerInSet(params params.ReplaceStickerInSetParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("replaceStickerInSet", structToMap(params), true, response)
}

func SetStickerEmojiList(params params.SetStickerEmojiListParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setStickerEmojiList", structToMap(params), true, response)
}

func SetStickerKeywords(params params.SetStickerKeywordsParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setStickerKeywords", structToMap(params), true, response)
}

func SetStickerMaskPosition(params params.SetStickerMaskPositionParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setStickerMaskPosition", structToMap(params), true, response)
}

func SetStickerSetTitle(params params.SetStickerSetTitleParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setStickerSetTitle", structToMap(params), true, response)
}

func SetStickerSetThumbnail(params params.SetStickerSetThumbnailParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setStickerSetThumbnail", structToMap(params), true, response)
}

func SetCustomEmojiStickerSetThumbnail(params params.SetCustomEmojiStickerSetThumbnailParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setCustomEmojiStickerSetThumbnail", structToMap(params), true, response)
}

func DeleteStickerSet(params params.DeleteStickerSetParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("deleteStickerSet", structToMap(params), true, response)
}

func AnswerInlineQuery(params params.AnswerInlineQueryParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("answerInlineQuery", structToMap(params), true, response)
}

func AnswerWebAppQuery(params params.AnswerWebAppQueryParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("answerWebAppQuery", structToMap(params), true, response)
}

func SendInvoice(params params.SendInvoiceParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendInvoice", structToMap(params), true, response)
}

func CreateInvoiceLink(params params.CreateInvoiceLinkParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("createInvoiceLink", structToMap(params), true, response)
}

func AnswerShippingQuery(params params.AnswerShippingQueryParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("answerShippingQuery", structToMap(params), true, response)
}

func AnswerPreCheckoutQuery(params params.AnswerPreCheckoutQueryParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("answerPreCheckoutQuery", structToMap(params), true, response)
}

func GetStarTransactions(params params.GetStarTransactionsParams) (map[string]interface{}, error) {
	return c.Endpoint("getStarTransactions", structToMap(params), true, true)
}

func RefundStarPayment(params params.RefundStarPaymentParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("refundStarPayment", structToMap(params), true, response)
}

func SetPassportDataErrors(params params.SetPassportDataErrorsParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setPassportDataErrors", structToMap(params), true, response)
}

func SendGame(params params.SendGameParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendGame", structToMap(params), true, response)
}

func SetGameScore(params params.SetGameScoreParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setGameScore", structToMap(params), true, response)
}

func GetGameHighScores(params params.GetGameHighScoresParams) (map[string]interface{}, error) {
	return c.Endpoint("getGameHighScores", structToMap(params), true, true)
}

func CreateChatSubscriptionInviteLink(params params.CreateChatSubscriptionInviteLinkParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("createChatSubscriptionInviteLink", structToMap(params), true, response)
}

func EditChatSubscriptionInviteLink(params params.EditChatSubscriptionInviteLinkParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("editChatSubscriptionInviteLink", structToMap(params), true, response)
}

func EditUserStarSubscription(params params.EditUserStarSubscriptionParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("editUserStarSubscription", structToMap(params), true, response)
}

func SetUserEmojiStatus(params params.SetUserEmojiStatusParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("setUserEmojiStatus", structToMap(params), true, response)
}

func SavePreparedInlineMessage(params params.SavePreparedInlineMessageParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("savePreparedInlineMessage", structToMap(params), true, response)
}

func GetAvailableGifts() (map[string]interface{}, error) {
	return c.Endpoint("getAvailableGifts", structToMap(nil), false, true)
}

func SendGift(params params.SendGiftParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("sendGift", structToMap(params), true, response)
}

func VerifyUser(params params.VerifyUserParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("verifyUser", structToMap(params), true, response)
}

func VerifyChat(params params.VerifyChatParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("verifyChat", structToMap(params), true, response)
}

func RemoveUserVerification(params params.RemoveUserVerificationParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("removeUserVerification", structToMap(params), true, response)
}

func VerifremoveChatVerificationyUser(params params.VerifremoveChatVerificationyUserParams, response bool) (map[string]interface{}, error) {
	return c.Endpoint("verifremoveChatVerificationyUser", structToMap(params), true, response)
}
