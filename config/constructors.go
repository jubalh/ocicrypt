/*
   Copyright The ocicrypt Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package config

import (
	"github.com/pkg/errors"
)

// EncryptWithJwe returns a CryptoConfig to encrypt with jwe public keys
func EncryptWithJwe(pubKeys [][]byte) (CryptoConfig, error) {
	dc := DecryptConfig{}
	ep := map[string][][]byte{
		"pubkeys": pubKeys,
	}

	return CryptoConfig{
		EncryptConfig: &EncryptConfig{
			Parameters:    ep,
			DecryptConfig: dc,
		},
		DecryptConfig: &dc,
	}, nil
}

// EncryptWithPkcs7 returns a CryptoConfig to encrypt with pkcs7 x509 certs
func EncryptWithPkcs7(x509s [][]byte) (CryptoConfig, error) {
	dc := DecryptConfig{}

	ep := map[string][][]byte{
		"x509s": x509s,
	}

	return CryptoConfig{
		EncryptConfig: &EncryptConfig{
			Parameters:    ep,
			DecryptConfig: dc,
		},
		DecryptConfig: &dc,
	}, nil
}

// EncryptWithGpg returns a CryptoConfig to encrypt with configured gpg parameters
func EncryptWithGpg(gpgRecipients [][]byte, gpgPubRingFile []byte) (CryptoConfig, error) {
	dc := DecryptConfig{}
	ep := map[string][][]byte{
		"gpg-recipients":     gpgRecipients,
		"gpg-pubkeyringfile": {gpgPubRingFile},
	}

	return CryptoConfig{
		EncryptConfig: &EncryptConfig{
			Parameters:    ep,
			DecryptConfig: dc,
		},
		DecryptConfig: &dc,
	}, nil
}

// EncryptWithPkcs11 returns a CryptoConfig to encrypt with configured PKCS11 parameters
//
// WARNING: will encrypt with modules
//
// Encrypt Config, As an example:
//
//  modulePath := "/usr/local/lib/softhsm/libsofthsm2.so"
//	modulePin := "1234"
//  modules := [][]byte{}
//  pins := [][]byte{}
//  modules = append(modules, []byte(modulePath))
//  pins = append(pins, []byte(modulePin))
//
//  validPkcs11Ccs := &config.CryptoConfig{
//		EncryptConfig: &config.EncryptConfig{
//			Parameters: map[string][][]byte{
//				"modules": modules,
//				"pins":    pins,
//			},
//		},
//		DecryptConfig: &DecryptConfig{},
//	}
func EncryptWithPkcs11(modules [][]byte, pins [][]byte) (CryptoConfig, error) {
	if len(modules) != len(pins) {
		return CryptoConfig{}, errors.New("Length of modules should match length of pins")
	}

	// TODO: experimental, just support single module
	if len(modules) != 1 {
		return CryptoConfig{}, errors.New("experimental, just support single module")
	}

	dc := DecryptConfig{}
	ep := map[string][][]byte{
		"modules": modules,
		"pins":    pins,
	}

	return CryptoConfig{
		EncryptConfig: &EncryptConfig{
			Parameters:    ep,
			DecryptConfig: dc,
		},
		DecryptConfig: &dc,
	}, nil
}

// DecryptWithPrivKeys returns a CryptoConfig to decrypt with configured private keys
func DecryptWithPrivKeys(privKeys [][]byte, privKeysPasswords [][]byte) (CryptoConfig, error) {
	if len(privKeys) != len(privKeysPasswords) {
		return CryptoConfig{}, errors.New("Length of privKeys should match length of privKeysPasswords")
	}

	dc := DecryptConfig{
		Parameters: map[string][][]byte{
			"privkeys":           privKeys,
			"privkeys-passwords": privKeysPasswords,
		},
	}

	ep := map[string][][]byte{}

	return CryptoConfig{
		EncryptConfig: &EncryptConfig{
			Parameters:    ep,
			DecryptConfig: dc,
		},
		DecryptConfig: &dc,
	}, nil
}

// DecryptWithX509s returns a CryptoConfig to decrypt with configured x509 certs
func DecryptWithX509s(x509s [][]byte) (CryptoConfig, error) {
	dc := DecryptConfig{
		Parameters: map[string][][]byte{
			"x509s": x509s,
		},
	}

	ep := map[string][][]byte{}

	return CryptoConfig{
		EncryptConfig: &EncryptConfig{
			Parameters:    ep,
			DecryptConfig: dc,
		},
		DecryptConfig: &dc,
	}, nil
}

// DecryptWithGpgPrivKeys returns a CryptoConfig to decrypt with configured gpg private keys
func DecryptWithGpgPrivKeys(gpgPrivKeys, gpgPrivKeysPwds [][]byte) (CryptoConfig, error) {
	dc := DecryptConfig{
		Parameters: map[string][][]byte{
			"gpg-privatekeys":           gpgPrivKeys,
			"gpg-privatekeys-passwords": gpgPrivKeysPwds,
		},
	}

	ep := map[string][][]byte{}

	return CryptoConfig{
		EncryptConfig: &EncryptConfig{
			Parameters:    ep,
			DecryptConfig: dc,
		},
		DecryptConfig: &dc,
	}, nil
}

// DecryptWithPkcs11 returns a CryptoConfig to decrypt with configured pkcs11 modules
func DecryptWithPkcs11(modules [][]byte, pins [][]byte) (CryptoConfig, error) {
	dc := DecryptConfig{
		Parameters: map[string][][]byte{
			"modules": modules,
			"pins":    pins,
		},
	}
	ep := map[string][][]byte{}
	return CryptoConfig{
		EncryptConfig: &EncryptConfig{
			Parameters:    ep,
			DecryptConfig: dc,
		},
		DecryptConfig: &dc,
	}, nil
}
