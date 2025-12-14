// mautrix-signal - A Matrix-signal puppeting bridge.
// Copyright (C) 2024 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package libsignalgo

/*
#include "./libsignal-ffi.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

const UsernameHashLength = 32

type UsernameHash [UsernameHashLength]byte

// UsernameHash computes the hash of a Signal username.
// The username should be in the format "nickname.discriminator" (e.g., "alice.42").
func HashUsername(username string) (UsernameHash, error) {
	var hash UsernameHash
	cUsername := C.CString(username)
	defer C.free(unsafe.Pointer(cUsername))

	signalFfiError := C.signal_username_hash((*[UsernameHashLength]C.uint8_t)(unsafe.Pointer(&hash[0])), cUsername)
	runtime.KeepAlive(username)

	if signalFfiError != nil {
		return UsernameHash{}, wrapError(signalFfiError)
	}
	return hash, nil
}
