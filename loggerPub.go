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
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// logs events to standard output or file
type LoggerPub struct {
	number    int
	logToFile bool
	path      string
	log       *logrus.Entry
}

func (pub *LoggerPub) Init(config *Config, log *logrus.Entry) {
	pub.log = log
	if config.Publishers.Logger.OutputTo == "file" {
		// get the path to log
		pub.path = config.Publishers.Logger.LogFolder
		// ensures there is a back slash at the end of the path
		if pub.path[len(pub.path)-1:] != "/" {
			pub.path += "/"
		}
		// ensures there is a folder there
		err := os.MkdirAll(pub.path, os.ModePerm)
		if err != nil {
			log.Errorf("Cannot create folder %s: %s. Reverting to stdout.", config.Publishers.Logger.LogFolder, err)
			config.Publishers.Logger.OutputTo = "stdout"
		} else {
			// now ready to log to files
			pub.logToFile = true
		}
	}
}

func (pub *LoggerPub) OnCreate(event Event) {
	pub.notify(event)
}

func (pub *LoggerPub) OnDelete(event Event) {
	pub.notify(event)
}

func (pub *LoggerPub) OnUpdate(event Event) {
	pub.notify(event)
}

func (pub *LoggerPub) notify(event Event) {
	// if it can log to file (i.e. specified and out of cluster)
	if pub.logToFile {
		// write log entry to the file system
		pub.writeToFile(event)
	} else {
		objBytes, err := json.Marshal(event)
		if err != nil {
			pub.log.Errorf("Publisher could not marshal object to json: %s.", err)
		} else {
			pub.log.Infof("%s %s %s: %s",
				strings.ToUpper(event.Info.ObjectType),
				event.Info.ObjectId,
				event.Info.EventType,
				string(objBytes))
		}
	}
}

// writes the change to the file system
func (pub *LoggerPub) writeToFile(event Event) {
	filename := fmt.Sprintf("%s%s", pub.path, pub.getNextName(event.Info))
	err := ioutil.WriteFile(filename, pub.toJSON(event), os.ModePerm)
	if err != nil {
		pub.log.Errorf("Failed to write to file %s: %s.", filename, err)
	}
}

// gets the next incremental number
func (pub *LoggerPub) getNextName(c Change) string {
	return strings.Replace(
		fmt.Sprintf("%s_%s_%s_%s.json",
			strconv.FormatInt(int64(time.Now().UTC().UnixNano()), 10),
			c.ObjectType,
			c.EventType,
			c.ObjectId),
		"/",
		"_",
		-1)
}

// converts the specified object into a pretty looking JSON string
func (pub *LoggerPub) toJSON(obj interface{}) []byte {
	// serialises the object into JSON applying indentation to format the output
	jsonBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		pub.log.Errorf("Failed to marshall object: %s", err)
		return []byte(fmt.Sprintf("Publisher failed to marshall object: %s.", err))
	}
	return jsonBytes
}
