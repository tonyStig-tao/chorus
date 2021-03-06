// Copyright 2017 Baptist-Publication Information Technology Services Co.,Ltd.
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

package common

import (
	"net"
	"strings"
)

// protoAddr: e.g. "tcp://127.0.0.1:8080" or "unix:///tmp/test.sock"
func Connect(protoAddr string) (net.Conn, error) {
	parts := strings.SplitN(protoAddr, "://", 2)
	proto, address := parts[0], parts[1]
	conn, err := net.Dial(proto, address)
	return conn, err
}
