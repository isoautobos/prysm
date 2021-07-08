package kv

import (
	"testing"

	"github.com/prysmaticlabs/go-bitfield"
	ethpb "github.com/prysmaticlabs/prysm/proto/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/shared/testutil"
	"github.com/prysmaticlabs/prysm/shared/testutil/require"
)

func TestAttCaches_hasSeenBit(t *testing.T) {
	c := NewAttCaches()

	seenA1 := testutil.HydrateAttestation(&ethpb.Attestation{AggregationBits: bitfield.Bitlist{0b10000011}})
	seenA2 := testutil.HydrateAttestation(&ethpb.Attestation{AggregationBits: bitfield.Bitlist{0b11100000}})
	require.NoError(t, c.insertSeenBit(seenA1))
	require.NoError(t, c.insertSeenBit(seenA2))
	tests := []struct {
		att  *ethpb.Attestation
		want bool
	}{
		{att: testutil.HydrateAttestation(&ethpb.Attestation{AggregationBits: bitfield.Bitlist{0b10000000}}), want: true},
		{att: testutil.HydrateAttestation(&ethpb.Attestation{AggregationBits: bitfield.Bitlist{0b10000001}}), want: true},
		{att: testutil.HydrateAttestation(&ethpb.Attestation{AggregationBits: bitfield.Bitlist{0b11100000}}), want: true},
		{att: testutil.HydrateAttestation(&ethpb.Attestation{AggregationBits: bitfield.Bitlist{0b10000011}}), want: true},
		{att: testutil.HydrateAttestation(&ethpb.Attestation{AggregationBits: bitfield.Bitlist{0b10001000}}), want: false},
		{att: testutil.HydrateAttestation(&ethpb.Attestation{AggregationBits: bitfield.Bitlist{0b11110111}}), want: false},
	}
	for _, tt := range tests {
		got, err := c.hasSeenBit(tt.att)
		require.NoError(t, err)
		if got != tt.want {
			t.Errorf("hasSeenBit() got = %v, want %v", got, tt.want)
		}
	}
}

func TestAttCaches_insertSeenBitDuplicates(t *testing.T) {
	c := NewAttCaches()
	att1 := testutil.HydrateAttestation(&ethpb.Attestation{AggregationBits: bitfield.Bitlist{0b10000011}})
	r, err := hashFn(att1.Data)
	require.NoError(t, err)
	require.NoError(t, c.insertSeenBit(att1))
	require.Equal(t, 1, c.seenAtt.ItemCount())

	_, expirationTime1, ok := c.seenAtt.GetWithExpiration(string(r[:]))
	require.Equal(t, true, ok)

	// Make sure that duplicates are not inserted, but expiration time gets updated.
	require.NoError(t, c.insertSeenBit(att1))
	require.Equal(t, 1, c.seenAtt.ItemCount())
	_, expirationTime2, ok := c.seenAtt.GetWithExpiration(string(r[:]))
	require.Equal(t, true, ok)
	require.Equal(t, true, expirationTime2.After(expirationTime1), "Expiration time is not updated")
}
