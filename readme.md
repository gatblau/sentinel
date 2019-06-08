<img src="./pics/sentinel_small.png" align="right" height="200" width="200"/>

# Sentinel 

Sentinel is a Go application, which looks for state changes of [kubernetes objects](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/), and publish their associated metadata to specific endpoints for further processing by downstream systems.

The application process can either run inside or outside of a [Kubernetes](https://kubernetes.io/docs/concepts/) cluster and can publish events to either a [webhook](https://en.wikipedia.org/wiki/Webhook) or to an instance of the [Apache Kafka](https://kafka.apache.org/) message broker.

The following image shows how the application works:

<div style="text-align:center">
<img src="./pics/arc.png" align="center" height="350" width="350"/>
</div>

## Getting Started

Follow the [guide here](./docs/getting_started.md) to get started.

## API compatibility

The current version uses the [Kubernetes API 1.11.10](https://github.com/kubernetes/api/releases/tag/kubernetes-1.11.10), compatible with Openshift 3.11.

To build the solution for another api version, the dependencies in the [go module](./go.mod) file has to be updated, for example, using [go get](https://golang.org/cmd/go/#hdr-Module_aware_go_get) as follows:

```bash
# for example, to update to version 1.14.2
$ go get -u k8s.io/api@kubernetes-1.14.2
$ go get -u k8s.io/apimachinery@kubernetes-1.14.2
$ go get -u k8s.io/client-go@kubernetes-1.14.2
```

__NOTE__: the minimal required [go version is 1.12.5](https://golang.org/dl/)

## Multi Cluster publishing with Kafka

The following example shows a configuration where events are published from three Kubernetes clusters into Kafka:

<div style="text-align:center">
<img src="./pics/kafka.png" height="420" width="350"/>
</div>

## Configuration

The Sentinel process can be configured by either a [config file](./config.toml) or via environment variables.

When environment variables are set, they override the values in the config file.

The available environment variables are described below.

### General Vars

| File Var | Environment Var | Description | Default |
|---|---|---|---|
| KubeConfig | SL_KUBECONFIG | the path to the kubernetes configuration file used by the Sentinel to connect to the kubernetes API. | ~/.kube/config |
| LoginLevel | SL_LOGINLEVEL | defines the login level used by the software. Possible values are: __Trace, Debug, Info, Warning, Error, Fatal and Panic__. | Info |
| Publishers.Publisher| SL_PUBLISHERS_PUBLISHER | defines which publisher to use (i.e. webhook, broker or logger). The logger publisher is there to write to standard output. | logger |

### _Logger Publisher Variables_

| File Var | Environment Var | Description | Default |
|---|---|---|---|
| Publishers.Logger.OutputTo | SL_PUBLISHERS_LOGGER_OUTPUTTO | whether to log to the standard output (stdout) or to the file system (file) | stdout |
| Publishers.Logger.LogFolder| SL_PUBLISHERS_LOGGER_LOGFOLDER | the path to the log folder, only required if Output = "file" | logs |

### _Webhook Publisher Variables_

| File Var | Environment Var | Description | Default |
|---|---|---|---|
| Publishers.Webhook.URI | SL_PUBLISHERS_WEBHOOK_URI | the uri of the webhook | localhost:8080/sentinel |
| Publishers.Webhook.Authentication | SL_PUBLISHERS_WEBHOOK_AUTHENTICATION | authentication mode to use for posting events to the webhook endpoint (i.e. none, basic) | - |
| Publishers.Webhook.Username | SL_PUBLISHERS_WEBHOOK_USERNAME | the optional username for basic authentication | sentinel |
| Publishers.Webhook.Password | SL_PUBLISHERS_WEBHOOK_PASSWORD | the optional password for basic authentication | s3nt1nel |

### _Broker Publisher Variables_

| File Var | Environment Var | Description | Default |
|---|---|---|---|
| Publishers.Broker.Addr | SL_PUBLISHERS_BROKER_ADDR | the address to bind to | ":8080" |
| Publishers.Broker.Brokers | SL_PUBLISHERS_BROKER_BROKERS | the Kafka brokers to connect to, as a comma separated list | - |
| Publishers.Broker.Certificate | SL_PUBLISHERS_BROKER_CERTIFICATE | optional certificate file for client authentication | - |
| Publishers.Broker.Key | SL_PUBLISHERS_BROKER_KEY | optional key file for client authentication | - |
| Publishers.Broker.CA | SL_PUBLISHERS_BROKER_CA | optional certificate authority file for TLS client authentication | - |
| Publishers.Broker.Verify | SL_PUBLISHERS_BROKER_VERIFY | optional verify ssl certificates chain | false |

### _Observable Object Variables_

| File Var | Environment Var | Description | Default |
|---|---|---|---|
| Observe.Service | SL_OBSERVE_SERVICE | whether to observe create, update and delete service events | true |
| Observe.Pod | SL_OBSERVE_POD | whether to observe create, update and delete pod events | true |
| Observe.PersistentVolume | SL_OBSERVE_PERSISTENTVOLUME | whether to observe create, update and delete persistent volume events | true |
| Observe.Namespace | SL_OBSERVE_NAMESPACE | whether to observe create, update and delete namespace events | true |
| Observe.Deployment | SL_OBSERVE_DEPLOYMENT | whether to observe create, update and delete deployment events | false |
| Observe.ReplicationController | SL_OBSERVE_REPLICATIONCONTROLLER | whether to observe create, update and delete replication controller events | false |
| Observe.ReplicateSet | SL_OBSERVE_REPLICASET | whether to observe create, update and delete replica set events | false |
| Observe.DaemonSet | SL_OBSERVE_DAEMONSET | whether to observe create, update and delete daemon set events | false |
| Observe.Job | SL_OBSERVE_JOB | whether to observe create, update and delete job events | false |
| Observe.Secret | SL_OBSERVE_SECRET | whether to observe create, update and delete secret events | false |
| Observe.ConfigMap | SL_OBSERVE_CONFIGMAP | whether to observe create, update and delete config map events | false |
| Observe.Ingress | SL_OBSERVE_INGRESS | whether to observe create, update and delete ingress events | false |
| Observe.ServiceAccount | SL_OBSERVE_SERVICEACCOUNT | whether to observe create, update and delete service account events | false |
| Observe.ClusterRole | SL_OBSERVE_CLUSTERROLE | whether to observe create, update and delete cluster role events | false |



[*] _The Sentinel icon was made by [Freepik](https://www.freepik.com) from [Flaticon](https://www.flaticon.com) and is licensed by [Creative Commons BY 3.0](http://creativecommons.org/licenses/by/3.0)_