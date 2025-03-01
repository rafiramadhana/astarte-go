// Copyright © 2020 Ispirata Srl
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deviceid

import (
	"encoding/base64"

	"github.com/google/uuid"
)

// IsValid returns whether the provided Device ID is a valid Astarte Device ID or not.
func IsValid(deviceID string) bool {
	decoded, err := base64.RawURLEncoding.DecodeString(deviceID)
	if err != nil {
		return false
	}

	// 16 bytes == 128 bit
	if len(decoded) != 16 {
		return false
	}

	return true
}

// GenerateRandom returns a new Astarte Device ID on a fully Random basis.
// Do not use in production environments.
func GenerateRandom() (string, error) {
	randomUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	deviceID, err := randomUUID.MarshalBinary()
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(deviceID), nil
}

// Generate returns an Astarte Device ID generated from a namespaced arbitrary payload.
// It is guaranteed to be always the same for the same namespace and payload.
// This is the go-to function to generate Astarte device IDs.
func Generate(uuidNamespace string, payloadData []byte) (string, error) {
	encodedUUIDNamespace, err := uuid.Parse(uuidNamespace)
	if err != nil {
		return "", err
	}

	deviceUUID := uuid.NewSHA1(encodedUUIDNamespace, payloadData)

	deviceID, err := deviceUUID.MarshalBinary()
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(deviceID), nil
}

// ToUUID converts a Device ID from the standard Astarte representation (Base 64 Url Encoded) to
// UUID string representation. This is useful to interact directly with Cassandra, that uses that
// representation to store Device IDs.
func ToUUID(deviceID string) (string, error) {
	bytes, err := base64.RawURLEncoding.DecodeString(deviceID)
	if err != nil {
		return "", err
	}
	deviceUUID, err := uuid.FromBytes(bytes)
	if err != nil {
		return "", err
	}

	return deviceUUID.String(), nil
}

// FromUUID converts a UUID string to a Device ID in the standard Astarte representation (Base
// 64 Url Encoded)
func FromUUID(deviceUUIDString string) (string, error) {
	deviceUUID, err := uuid.Parse(deviceUUIDString)
	if err != nil {
		return "", err
	}
	deviceID, err := deviceUUID.MarshalBinary()
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(deviceID), nil
}
