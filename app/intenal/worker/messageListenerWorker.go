package worker

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// MessageFetcher interface represents the contract
// of object that will fetch the messages from tg bot
type MessageFetcher interface {
	Start() tgbotapi.UpdatesChannel
	Stop()
}

type MessageProcessor interface {
	HandleIncomingMessage(ctx context.Context, message *tgbotapi.Message)
}

type MessageListenerWorker struct {
	fetcher   MessageFetcher
	processor MessageProcessor
}

func NewMessageListenerWorker(messageFetcher MessageFetcher, processor MessageProcessor) *MessageListenerWorker {
	return &MessageListenerWorker{
		fetcher:   messageFetcher,
		processor: processor,
	}
}
func (w *MessageListenerWorker) Run(ctx context.Context) {
	for update := range w.fetcher.Start() {
		select {
		case <-ctx.Done():
			{
				w.fetcher.Stop()
				log.Println("canceled receiving messages from tg")
			}
		default:
			{
				message := update.Message
				w.processor.HandleIncomingMessage(
					ctx,
					message,
				)
			}
		}
	}
}
