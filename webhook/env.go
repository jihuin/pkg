/*
Copyright 2020 The Knative Authors

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

package webhook

import (
	"os"
	"strconv"
)

const portEnvKey = "WEBHOOK_PORT"

// Port returns the port number that webhook uses.
func Port(defaultPort int) int {
	port, err := strconv.Atoi(os.Getenv(portEnvKey))
	if err != nil {
		return defaultPort
	}
	return port
}
