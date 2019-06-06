/*
   Sentinel - Copyright (c) 2019 by www.gatblau.org

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software distributed under
   the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
   either express or implied.
   See the License for the specific language governing permissions and limitations under the License.

   Contributors to this project, hereby assign copyright in this code to the project,
   to be licensed under the same terms as the rest of the code.
*/
package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strings"
)

// logs events to standard output or file
type LoggerPub struct {
}

func (pub *LoggerPub) Init(config *Config) {
}

func (pub *LoggerPub) OnCreate(change Change, obj interface{}) {
	pub.notify(change, obj)
}

func (pub *LoggerPub) OnDelete(change Change, obj interface{}) {
	pub.notify(change, obj)
}

func (pub *LoggerPub) OnUpdate(change Change, obj interface{}) {
	pub.notify(change, obj)
}

func (pub *LoggerPub) notify(change Change, obj interface{}) {
	objBytes, err := json.Marshal(obj)
	if err != nil {
		logrus.Errorf("Can't serialise change: %+v", change)
	}
	logrus.Infof("%s %s %s: %s",
		strings.ToUpper(change.objectType),
		change.key,
		change.changeType,
		strings.Replace(string(objBytes), "\\", "", -1))
}
