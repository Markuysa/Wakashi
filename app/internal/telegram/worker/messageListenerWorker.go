package worker

import (
	"context"
	"log"
	"tgBotIntern/app/internal/telegram/controllers"
)

type MessageListenerWorker struct {
	fetcher   controllers.MessageFetcher
	processor controllers.MessageProcessor
}

func NewMessageListenerWorker(messageFetcher *controllers.FetcherWorker, processor *controllers.MessageHandler) *MessageListenerWorker {
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
