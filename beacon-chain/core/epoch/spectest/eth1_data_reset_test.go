package spectest

import (
	"path"
	"testing"

	"github.com/prysmaticlabs/prysm/beacon-chain/core/epoch"
	iface "github.com/prysmaticlabs/prysm/beacon-chain/state/interface"
	"github.com/prysmaticlabs/prysm/shared/params/spectest"
	"github.com/prysmaticlabs/prysm/shared/testutil"
	"github.com/prysmaticlabs/prysm/shared/testutil/require"
)

func runEth1DataResetTests(t *testing.T, config string) {
	require.NoError(t, spectest.SetConfig(t, config))

	testFolders, testsFolderPath := testutil.TestFolders(t, config, "phase0", "epoch_processing/eth1_data_reset/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			testutil.RunEpochOperationTest(t, folderPath, processEth1DataResetWrapper)
		})
	}
}

func processEth1DataResetWrapper(t *testing.T, state iface.BeaconState) (iface.BeaconState, error) {
	state, err := epoch.ProcessEth1DataReset(state)
	require.NoError(t, err, "Could not process final updates")
	return state, nil
}