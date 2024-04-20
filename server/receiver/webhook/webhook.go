package webhook

import (
	"thsr/m/server/receiver"
)

type Receiver interface {
	receiver.Receiver
}
