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
	"github.com/sirupsen/logrus"
)

// logs events to kafka
type BrokerPub struct {
}

func (pub *BrokerPub) Init(c Config) {
}

func (pub *BrokerPub) OnCreate(event Event, obj interface{}) {
	pub.notify(event, obj, "created")
}

func (pub *BrokerPub) OnDelete(event Event, obj interface{}) {
	pub.notify(event, obj, "deleted")
}

func (pub *BrokerPub) OnUpdate(event Event, obj interface{}) {
	pub.notify(event, obj, "updated")
}

func (pub *BrokerPub) notify(event Event, obj interface{}, action string) {
	logrus.Infof("action %s - event %+v --> %+v\n", action, event, obj)
}
