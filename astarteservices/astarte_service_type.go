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

package astarteservices

import (
	"errors"
)

// AstarteService represents one of Astarte's Services
type AstarteService int

const (
	// Unknown Astarte Service
	Unknown AstarteService = iota
	// Housekeeping is Astarte's service for managing Realms
	Housekeeping
	// RealmManagement is Astarte's service for managing configuration of a Realm
	RealmManagement
	// Pairing is Astarte's service for managing device provisioning and access
	Pairing
	// AppEngine is Astarte's service for interacting with Devices, Groups and more
	AppEngine
	// Channels is Astarte's service for WebSockets
	Channels
	// Flow is Astarte Flow
	Flow
)

var astarteServiceValidNames = map[string]AstarteService{
	"housekeeping":     Housekeeping,
	"hk":               Housekeeping,
	"realm-management": RealmManagement,
	"realmmanagement":  RealmManagement,
	"realm":            RealmManagement,
	"pairing":          Pairing,
	"appengine":        AppEngine,
	"app":              AppEngine,
	"channels":         Channels,
	"flow":             Flow,
}

func (service AstarteService) String() string {
	names := [...]string{
		"",
		"housekeeping",
		"realm-management",
		"pairing",
		"appengine",
		"channels",
		"flow"}

	if service < Housekeeping || service > Flow {
		return ""
	}

	return names[service]
}

// FromString returns a valid AstarteService out of a string
func FromString(astarteServiceString string) (AstarteService, error) {
	if value, exist := astarteServiceValidNames[astarteServiceString]; exist {
		return value, nil
	}

	return Unknown, errors.New("Invalid type")
}
