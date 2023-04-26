package helpers

import (
	"fmt"
	"tgBotIntern/app/internal/constants/roles"
	"tgBotIntern/app/internal/entity"
)

func FormTransaction(transaction entity.Transaction) string {
	return fmt.Sprintf("\nId:%v,\nCard number:%v,\nOwner id:%v,\nOperation value:%v,\nDate:%v,\nFrom:%v\n",
		transaction.ID,
		transaction.CardNumber,
		transaction.OwnerID,
		transaction.OperationValue,
		transaction.TransactionDate,
		transaction.RequestFromID,
	)
}
func FormReport(report *entity.Report) string {
	result := fmt.Sprintf(
		"Username:%v,\nRoleID:%v,\nTransactions:\n",
		report.Username,
		report.RoleID,
	)
	for _, transaction := range report.Transactions {
		result += FormTransaction(transaction)
	}
	return result
}
func FormSlavesList(users []entity.User) string {
	var result string
	for _, user := range users {
		result += FormUser(&user)
	}
	if len(result) == 0 {
		return "Empty list of slaves"
	}
	return result
}
func FormTransactions(transactions []entity.Transaction) string {
	var result string
	for _, transaction := range transactions {
		result += FormTransaction(transaction)
	}
	if len(result) == 0 {
		return "Empty list of transactions"
	}
	return result
}
func FormUser(user *entity.User) string {
	return fmt.Sprintf(
		"\nUsername:%s\nNickname:%s\nRole:%s\n",
		user.Username, user.Nickname, roles.GetRoleString(user.Role),
	)
}
func FormCardsList(cards []entity.Card) string {
	var result string
	for _, card := range cards {
		result += fmt.Sprintf(
			"\nOwner:%v,\nIssuerBankId:%v,\nCardNumber:%v,\nDailyLimit:%v,\nTotal:%v\n",
			card.DaimyoID,
			card.IssuerBankID,
			card.CardNumber,
			card.DailyLimit,
			card.Total,
		)
	}
	return result
}
