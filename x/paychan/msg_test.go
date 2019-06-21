package paychan

import (
	"testing"

	coin "github.com/iov-one/weave/coin"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/weavetest/assert"
)

func TestCreatePaymentChannelMsgValidate(t *testing.T) {
	msg := &CreatePaymentChannelMsg{
		Total: coin.NewCoinp(1, 0, "IOV"),
	}
	err := msg.Validate()

	assert.FieldError(t, err, "Metadata", errors.ErrMetadata)
	assert.FieldError(t, err, "Src", errors.ErrEmpty)
	assert.FieldError(t, err, "Recipient", errors.ErrEmpty)
	assert.FieldError(t, err, "Timeout", errors.ErrInput)

	assert.FieldError(t, err, "Total", nil)
	assert.FieldError(t, err, "Memo", nil)
}
