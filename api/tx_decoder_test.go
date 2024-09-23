package api

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeTx(t *testing.T) {
	const base64tx = "CqABCpMBChwvZGVjaW1hbC5jb2luLnYxLk1zZ1NlbmRDb2luEnMKKWQwMWRnZnNldWQ4Z2R3N3gyZ2QwYXo0NXY3ZnFtcHZxYTJnODkwN3A0EilkMDFuc2RxeTVwd2RzZDQ0OTNxZ216YWZ1a2swOGRjcGxhazA0d25jOBobCgNkZWwSFDEwMDAwMDAwMDAwMDAwMDAwMDAwEggxMjM0NTY2NhJdClkKTwooL2V0aGVybWludC5jcnlwdG8udjEuZXRoc2VjcDI1NmsxLlB1YktleRIjCiED2Qwozry71RiS96Xws6Pn2AVXPZ/cfk4NpDeCeF6WjKESBAoCCAEYHRIAGkEU8qFCSvy0nCiztZCz3Nk7SEWzChoRZuzdHDZM3OZ0yiZKOnpGKq0kFw4wCycEaPf5YnxXAJNU7SfRDt3PXhiCAQ=="
	bz, err := base64.StdEncoding.DecodeString(base64tx)
	require.NoError(t, err)
	txdec, err := decodeTransaction(bz)
	require.NoError(t, err)
	require.Equal(t, "12345666", txdec.Memo)
}
