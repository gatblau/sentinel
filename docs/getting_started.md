# Getting Started

The easiest way to getting started is to deploy Sentinel from a container image in [minikube](https://github.com/kubernetes/minikube).

Alternatively, you can try run the binary outside of an image and redirect the output to the file system.

## Installing required tools

To build it from code, make sure you have the latest version of [golang](https://golang.org/dl/) installed on your machine.

You will also need to have a [git](https://www.atlassian.com/git/tutorials/install-git) client and [make](https://www.gnu.org/software/make/) installed.

Finally, as Sentinel is meant to observe changes in a Kubernetes cluster, you need to have Kubernetes running and point Sentinel to it. The simplest way to have kubernetes running locally is by using [minikube](https://github.com/kubernetes/minikube). You need to [install it](https://kubernetes.io/docs/tasks/tools/install-minikube/) on your local machine. And to talk to minikube you also need to [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/), the Kubernetes command line tool.

## Deploying from a container image

To deploy Sentinel from the container image into K8S [see here](./k8s_deploy.md).

## Trying the binary out in the local machine

If you want to try it out on your local machine [see here](./binary_deploy.md)
