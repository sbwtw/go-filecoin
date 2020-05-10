package types_test

import (
	"testing"

	"github.com/sbwtw/go-filecoin/internal/pkg/encoding"
	. "github.com/sbwtw/go-filecoin/internal/pkg/types"
	"github.com/sbwtw/go-filecoin/internal/pkg/util/convert"

	"github.com/stretchr/testify/assert"
)

func TestEncodingZeroVal(t *testing.T) {
	t.Skip("cbor fix")
	comms := Commitments{}
	data, err := encoding.Encode(comms)
	assert.NoError(t, err)
	var newComms Commitments
	err = encoding.Decode(data, &newComms)
	assert.NoError(t, err)
}

func TestEncoding(t *testing.T) {
	t.Skip("cbor fix")
	var comms Commitments

	commR := CommR(convert.To32ByteArray([]byte{0xf}))
	commD := CommD(convert.To32ByteArray([]byte{0xa}))
	commRStar := CommRStar(convert.To32ByteArray([]byte{0xc}))

	comms.CommR = &commR
	comms.CommD = &commD
	comms.CommRStar = &commRStar

	data, err := encoding.Encode(comms)
	assert.NoError(t, err)
	var newComms Commitments
	err = encoding.Decode(data, &newComms)
	assert.NoError(t, err)
}
