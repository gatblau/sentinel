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
	log *logrus.Entry
}

func (pub *BrokerPub) Init(c *Config, log *logrus.Entry) {
	pub.log = log
}

func (pub *BrokerPub) OnCreate(event Event) {
	pub.notify(event)
}

func (pub *BrokerPub) OnDelete(event Event) {
	pub.notify(event)
}

func (pub *BrokerPub) OnUpdate(event Event) {
	pub.notify(event)
}

func (pub *BrokerPub) notify(event Event) {
	pub.log.Warnf(
		"BROKER PUBLISHER NOT IMPLEMENTED! Trying to publish change %s for object %s\n",
		event.Info.EventType,
		event.Info.ObjectType)
}
