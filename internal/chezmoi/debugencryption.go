package chezmoi

import (
	"github.com/rs/zerolog/log"

	"github.com/twpayne/chezmoi/v2/internal/chezmoilog"
)

// A DebugEncryption logs all calls to an Encryption.
type DebugEncryption struct {
	encryption Encryption
}

// NewDebugEncryption returns a new DebugEncryption.
func NewDebugEncryption(encryption Encryption) *DebugEncryption {
	return &DebugEncryption{
		encryption: encryption,
	}
}

// Decrypt implements Encryption.Decrypt.
func (e *DebugEncryption) Decrypt(ciphertext []byte) ([]byte, error) {
	plaintext, err := e.encryption.Decrypt(ciphertext)
	log.Err(err).
		Bytes("ciphertext", chezmoilog.Output(ciphertext, err)).
		Bytes("plaintext", chezmoilog.Output(plaintext, err)).
		Msg("Decrypt")
	return plaintext, err
}

// DecryptToFile implements Encryption.DecryptToFile.
func (e *DebugEncryption) DecryptToFile(plaintextAbsPath AbsPath, ciphertext []byte) error {
	err := e.encryption.DecryptToFile(plaintextAbsPath, ciphertext)
	log.Err(err).
		Stringer("plaintextAbsPath", plaintextAbsPath).
		Bytes("ciphertext", chezmoilog.Output(ciphertext, err)).
		Msg("DecryptToFile")
	return err
}

// Encrypt implements Encryption.Encrypt.
func (e *DebugEncryption) Encrypt(plaintext []byte) ([]byte, error) {
	ciphertext, err := e.encryption.Encrypt(plaintext)
	log.Err(err).
		Bytes("plaintext", chezmoilog.Output(plaintext, err)).
		Bytes("ciphertext", chezmoilog.Output(ciphertext, err)).
		Msg("Encrypt")
	return ciphertext, err
}

// EncryptFile implements Encryption.EncryptFile.
func (e *DebugEncryption) EncryptFile(plaintextAbsPath AbsPath) ([]byte, error) {
	ciphertext, err := e.encryption.EncryptFile(plaintextAbsPath)
	log.Err(err).
		Stringer("plaintextAbsPath", plaintextAbsPath).
		Bytes("ciphertext", chezmoilog.Output(ciphertext, err)).
		Msg("EncryptFile")
	return ciphertext, err
}

// EncryptedSuffix implements Encryption.EncryptedSuffix.
func (e *DebugEncryption) EncryptedSuffix() string {
	return e.encryption.EncryptedSuffix()
}
