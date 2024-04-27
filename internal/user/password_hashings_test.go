package user

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kopia/kopia/internal/crypto"
)

// The password hashing constants defined in this package are used as "lookup
// keys" for the register password-based key derivers in the crypto package.
// This trivial test is a change detector to ensure that the constants defined
// in the user package match those defined in the crypto package.
func TestPasswordHashingConstantMatchCryptoPackage(t *testing.T) {
	require.Equal(t, crypto.ScryptAlgorithm, scryptHashAlgorithm)
	require.Equal(t, crypto.Pbkdf2Algorithm, pbkdf2HashAlgorithm)
}

// The passwordHashSaltLength constant defines the salt length used in this
// package for password hashing. This trivial test ensures that this hash length
// meets the minimum requirement for the instantiations of the registered
// password hashers (PB key derivers in the crypto package).
func TestSaltLengthIsSupported(t *testing.T) {
	const badPwd = "password"
	var salt [passwordHashSaltLength]byte

	for _, h := range PasswordHashingAlgorithms() {
		_, err := computePasswordHash(badPwd, salt[:], h)
		require.NoError(t, err)
	}
}
