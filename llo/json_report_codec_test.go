package llo

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/goplugin/plugin-libocr/offchainreporting2/types"

	"github.com/goplugin/plugin-common/pkg/types/llo"
	llotypes "github.com/goplugin/plugin-common/pkg/types/llo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_JSONCodec(t *testing.T) {
	t.Run("Encode=>Decode", func(t *testing.T) {
		r := Report{
			ConfigDigest:                types.ConfigDigest([32]byte{1, 2, 3}),
			SeqNr:                       43,
			ChannelID:                   llotypes.ChannelID(46),
			ValidAfterSeconds:           44,
			ObservationTimestampSeconds: 45,
			Values:                      []StreamValue{ToDecimal(decimal.NewFromInt(1)), ToDecimal(decimal.NewFromInt(2)), &Quote{Bid: decimal.NewFromFloat(3.13), Benchmark: decimal.NewFromFloat(4.4), Ask: decimal.NewFromFloat(5.12)}},
			Specimen:                    true,
		}

		cdc := JSONReportCodec{}

		encoded, err := cdc.Encode(r, llo.ChannelDefinition{})
		require.NoError(t, err)

		fmt.Println("encoded", string(encoded))
		assert.Equal(t, `{"ConfigDigest":"0102030000000000000000000000000000000000000000000000000000000000","SeqNr":43,"ChannelID":46,"ValidAfterSeconds":44,"ObservationTimestampSeconds":45,"Values":[{"Type":0,"Value":"1"},{"Type":0,"Value":"2"},{"Type":1,"Value":"Q{Bid: 3.13, Benchmark: 4.4, Ask: 5.12}"}],"Specimen":true}`, string(encoded))

		decoded, err := cdc.Decode(encoded)
		require.NoError(t, err)

		assert.Equal(t, r, decoded)
	})
}
