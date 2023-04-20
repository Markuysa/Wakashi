package worker

import (
	"context"
	"log"
	"tgBotIntern/app/internal/telegram/infrastructure/processors"
)

type MessageListenerWorker struct {
	fetcher   processors.MessageFetcher
	processor processors.MessageProcessor
}

func NewMessageListenerWorker(messageFetcher *processors.FetcherWorker, processor *processors.MessageHandler) *MessageListenerWorker {
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
				err := w.processor.HandleIncomingMessage(
					ctx,
					message,
				)
				if err != nil {
					return
				}
			}
		}
	}
}
