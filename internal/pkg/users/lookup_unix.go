//go:build unix

/*
Copyright 2024 k0s authors

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

package users

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"os/user"
	"strconv"
)

// Lookup looks up a user's UID by username. If the user cannot be found, the
// returned error is [ErrNotExist]. If an error is returned, the returned UID
// will be [UnknownUID].
func LookupUID(name string) (int, error) {
	var uid string

	if entry, err := user.Lookup(name); err != nil {
		if !errors.Is(err, user.UnknownUserError(name)) {
			return UnknownUID, err
		}

		err = ErrNotExist

		// fallback to call external `id` in case NSS is used
		out, idErr := exec.Command("id", "-u", name).Output()
		if idErr != nil {
			var exitErr *exec.ExitError
			if errors.As(idErr, &exitErr) {
				return UnknownUID, fmt.Errorf("%w (%w: %s)", err, idErr, bytes.TrimSpace(exitErr.Stderr))
			}
			return UnknownUID, fmt.Errorf("%w (%w)", err, idErr)
		}

		uid = string(bytes.TrimSpace(out))
	} else {
		uid = entry.Uid
	}

	parsedUID, err := strconv.Atoi(uid)
	if err != nil {
		return UnknownUID, fmt.Errorf("UID %q is not a decimal integer: %w", uid, err)
	}
	if parsedUID < 0 {
		return UnknownUID, fmt.Errorf("UID is negative: %d", parsedUID)
	}

	return parsedUID, nil
}
