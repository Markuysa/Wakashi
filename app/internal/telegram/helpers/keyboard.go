package helpers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBotIntern/app/internal/constants/roles"
)

const (
	// Default keyboard buttons
	LoginButton    = "Login"
	RegisterButton = "Register"
	Info           = "Info"
	Exit           = "Exit"
	// Administrator buttons titles
	AdminCreateEntity      = "Create entity"
	AdminCreateCard        = "Create card"
	AdminVBindCardToDaimyo = "Bind card to daimyo"
	AdminBindSlaveToMaster = "Bind slave to master"
	AdminGetEntityData     = "Get entity data"

	// Shogun buttons titles
	ShogunGetSlavesList = "Get my slaves list"
	ShogunCreateCard    = "Create new card"
	ShogunGetSlaveData  = "Get particular slave data"

	// Daimyo buttons titles
	DaimyoGetCards           = "Get my cards list"
	DaimyoCreateCardRequest  = "Create increment card request"
	DaimyoGetSamuraiList     = "Get my slave samurai list"
	DaimyoGetSamuraiTurnover = "Get samurai turnover"
	DaimyoGetCardsTotal      = "Get my cards total"
	DaimyoBindShogun         = "Bind me to shogun"

	// Samurai buttons titles
	SamuraiGetTurnover = "Get my turnover"
	SamuraiBindDaimyo  = "Bind me to daimyo"

	// Collector buttons titles
	CollectorProccessRequest  = "Process incoming request from Daimyo"
	CollectorShowTransactions = "Show all transactions (requests)"
)

// Keyboards that will be shown after success login and decision of role
var (
	adminKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(AdminCreateEntity),
			tgbotapi.NewKeyboardButton(AdminCreateCard),
			tgbotapi.NewKeyboardButton(AdminVBindCardToDaimyo),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(AdminBindSlaveToMaster),
			tgbotapi.NewKeyboardButton(AdminGetEntityData),
		),
	)
	shogunKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(ShogunGetSlavesList),
			tgbotapi.NewKeyboardButton(ShogunCreateCard),
			tgbotapi.NewKeyboardButton(ShogunGetSlaveData),
		),
	)
	daimyoKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(DaimyoGetCards),
			tgbotapi.NewKeyboardButton(DaimyoCreateCardRequest),
			tgbotapi.NewKeyboardButton(DaimyoGetSamuraiList),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(DaimyoGetSamuraiTurnover),
			tgbotapi.NewKeyboardButton(DaimyoBindShogun),
			tgbotapi.NewKeyboardButton(DaimyoGetCardsTotal),
		),
	)
	samuraiKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(SamuraiGetTurnover),
			tgbotapi.NewKeyboardButton(SamuraiBindDaimyo),
		),
	)
	collectorKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CollectorProccessRequest),
			tgbotapi.NewKeyboardButton(CollectorShowTransactions),
		),
	)
	defaultKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(LoginButton),
			tgbotapi.NewKeyboardButton(RegisterButton),
			tgbotapi.NewKeyboardButton(Info),
			tgbotapi.NewKeyboardButton(Exit),
		),
	)
)

func GetKeyboard(roleID int) tgbotapi.ReplyKeyboardMarkup {
	switch roleID {
	case roles.Administrator:
		{
			return adminKeyboard
		}
	case roles.Shogun:
		{
			return shogunKeyboard
		}
	case roles.Daimyo:
		{
			return daimyoKeyboard
		}
	case roles.Samurai:
		{
			return samuraiKeyboard
		}
	case roles.Collector:
		{
			return collectorKeyboard
		}
	default:
		return defaultKeyboard
	}
}
