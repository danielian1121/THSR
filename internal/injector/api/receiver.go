package api

import (
	"github.com/google/wire"
	"thsr/m/server/receiver"
	"thsr/m/server/receiver/webhook"
)

var ReceiverSet = wire.NewSet(
	webhook.ProvideReceiver,
)

func ProvideReceiverList(webhook webhook.Receiver) []receiver.Receiver {
	return []receiver.Receiver{webhook}
}
