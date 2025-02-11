package v2

import (
	"testing"

	fieldparams "github.com/prysmaticlabs/prysm/config/fieldparams"
	"github.com/prysmaticlabs/prysm/encoding/bytesutil"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
	"github.com/prysmaticlabs/prysm/testing/require"
)

func TestBeaconState_LatestBlockHeader(t *testing.T) {
	s, err := InitializeFromProto(&ethpb.BeaconStateAltair{})
	require.NoError(t, err)
	got := s.LatestBlockHeader()
	require.DeepEqual(t, (*ethpb.BeaconBlockHeader)(nil), got)

	want := &ethpb.BeaconBlockHeader{Slot: 100}
	s, err = InitializeFromProto(&ethpb.BeaconStateAltair{LatestBlockHeader: want})
	require.NoError(t, err)
	got = s.LatestBlockHeader()
	require.DeepEqual(t, want, got)

	// Test copy does not mutate.
	got.Slot = 101
	require.DeepNotEqual(t, want, got)
}

func TestBeaconState_BlockRoots(t *testing.T) {
	s, err := InitializeFromProto(&ethpb.BeaconStateAltair{})
	require.NoError(t, err)
	got := s.BlockRoots()
	require.DeepEqual(t, ([][]byte)(nil), got)

	want := [][]byte{{'a'}}
	s, err = InitializeFromProto(&ethpb.BeaconStateAltair{BlockRoots: want})
	require.NoError(t, err)
	got = s.BlockRoots()
	require.DeepEqual(t, want, got)

	// Test copy does not mutate.
	got[0][0] = 'b'
	require.DeepNotEqual(t, want, got)
}

func TestBeaconState_BlockRootAtIndex(t *testing.T) {
	s, err := InitializeFromProto(&ethpb.BeaconStateAltair{})
	require.NoError(t, err)
	got, err := s.BlockRootAtIndex(0)
	require.NoError(t, err)
	require.DeepEqual(t, ([]byte)(nil), got)

	r := [][]byte{{'a'}}
	s, err = InitializeFromProto(&ethpb.BeaconStateAltair{BlockRoots: r})
	require.NoError(t, err)
	got, err = s.BlockRootAtIndex(0)
	require.NoError(t, err)
	want := bytesutil.PadTo([]byte{'a'}, fieldparams.RootLength)
	require.DeepSSZEqual(t, want, got)
}
