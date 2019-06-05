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
	v "github.com/spf13/viper"
	"strings"
)

// Sentinel configuration
type Config struct {
	KubeConfig string
	Publishers Publishers
	Observe    Observe
}

// the configuration for the event publishers
type Publishers struct {
	Mode    string
	Webhook Webhook
	Broker  Broker
}

// the configuration for the web hook publisher
type Webhook struct {
	URI            string
	Authentication string
	Username       string
	Password       string
}

// the configuration for the message broker publisher
type Broker struct {
	Addr        string
	Brokers     string
	Verbose     bool
	Certificate string
	Key         string
	CA          string
	Verify      bool
}

// the type of objects that can be observed by the controller
type Observe struct {
	Service               bool
	Pod                   bool
	PersistentVolume      bool
	Namespace             bool
	Deployment            bool
	ReplicationController bool
	ReplicaSet            bool
	DaemonSet             bool
	Job                   bool
	Secret                bool
	ConfigMap             bool
	Ingress               bool
}

// creates a new configuration file passed by value
// to avoid thread sync issues
func NewConfig() (Config, error) {
	//v := viper.New()
	// loads the configuration file
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	err := v.ReadInConfig() // find and read the config file
	if err != nil {         // handle errors reading the config file
		logrus.Errorf("Fatal error config file: %s \n", err)
		return Config{}, err
	}

	// binds all environment variables to make it container friendly
	v.AutomaticEnv()
	v.SetEnvPrefix("SL")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	_ = v.BindEnv("KubeConfig")
	_ = v.BindEnv("Publishers.Mode")
	_ = v.BindEnv("Publishers.Webhook.URI")
	_ = v.BindEnv("Publishers.Webhook.Authentication")
	_ = v.BindEnv("Publishers.Webhook.Username")
	_ = v.BindEnv("Publishers.Webhook.Password")
	_ = v.BindEnv("Publishers.Broker.Addr")
	_ = v.BindEnv("Publishers.Broker.Brokers")
	_ = v.BindEnv("Publishers.Broker.Verbose")
	_ = v.BindEnv("Publishers.Broker.Certificate")
	_ = v.BindEnv("Publishers.Broker.Key")
	_ = v.BindEnv("Publishers.Broker.CA")
	_ = v.BindEnv("Publishers.Broker.Verify")
	_ = v.BindEnv("Observe.Service")
	_ = v.BindEnv("Observe.Pod")
	_ = v.BindEnv("Observe.PersistentVolume")
	_ = v.BindEnv("Observe.Namespace")
	_ = v.BindEnv("Observe.Deployment")
	_ = v.BindEnv("Observe.ReplicationController")
	_ = v.BindEnv("Observe.ReplicaSet")
	_ = v.BindEnv("Observe.DaemonSet")
	_ = v.BindEnv("Observe.Job")
	_ = v.BindEnv("Observe.Secret")
	_ = v.BindEnv("Observe.ConfigMap")
	_ = v.BindEnv("Observe.Ingress")

	// creates a config struct and populate it with values
	c := new(Config)

	// general configuration
	c.KubeConfig = v.GetString("KubeConfig")
	c.Publishers.Webhook.URI = v.GetString("Publishers.Webhook.URI")

	// webhook publisher configuration
	c.Publishers.Webhook.Authentication = v.GetString("Publishers.Webhook.Authentication")
	c.Publishers.Webhook.Username = v.GetString("Publishers.Webhook.Username")
	c.Publishers.Webhook.Password = v.GetString("Publishers.Webhook.Password")

	// broker publisher configuration
	c.Publishers.Broker.Addr = v.GetString("")
	c.Publishers.Broker.Brokers = v.GetString("")
	c.Publishers.Broker.Certificate = v.GetString("")
	c.Publishers.Broker.Key = v.GetString("")
	c.Publishers.Broker.CA = v.GetString("")
	c.Publishers.Broker.Verbose = v.GetBool("")
	c.Publishers.Broker.Verify = v.GetBool("")

	// observable objects configuration
	c.Observe.Service = v.GetBool("Observe.Service")
	c.Observe.Pod = v.GetBool("Observe.Pod")
	c.Observe.PersistentVolume = v.GetBool("Observe.PersistentVolume")
	c.Observe.Namespace = v.GetBool("Observe.Namespace")
	c.Observe.ConfigMap = v.GetBool("Observe.ConfigMap")
	c.Observe.DaemonSet = v.GetBool("Observe.DaemonSet")
	c.Observe.Deployment = v.GetBool("Observe.Deployment")
	c.Observe.Ingress = v.GetBool("Observe.Ingress")
	c.Observe.Job = v.GetBool("Observe.Job")
	c.Observe.ReplicaSet = v.GetBool("Observe.ReplicaSet")
	c.Observe.ReplicationController = v.GetBool("Observe.ReplicationController")
	c.Observe.Secret = v.GetBool("Observe.Secret")

	// return configuration as value to avoid thread issues
	// note: does not refresh after config has been loaded though
	return *c, nil
}
