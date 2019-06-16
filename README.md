# MongoDB Atlas Kubernetes Operator

## Overview

A Kubernetes Operator for [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) with which you can manage your MongoDB Atlas projects and clusters from within Kubernetes such as you do with your containerized applications. It is built using the [Operator Framework](https://github.com/operator-framework) and [Kubernetes Custom Resource Definitions (CRDs)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#customresourcedefinitions).
This project was inspired from the [MongoDB Atlas Terraform Provider](https://github.com/akshaykarle/terraform-provider-mongodbatlas) with the goal to have Kubernetes as the single source for both (stateless) applications and MongoDB Atlas as the persistence layer. Furthermore, the Kubernetes operator ensures via Reconcile loops to have the desired state matching with the actual state and thus following the GitOps approach.

<!-- vim-markdown-toc GFM -->

* [Scope](#scope)
* [Prerequisites](#prerequisites)
* [Getting Started](#getting-started)
* [Cleanup](#cleanup)
* [Developers Build Guide](#developers-build-guide)
* [Contributing](#contributing)

<!-- vim-markdown-toc -->

## Scope

**Currently it supports**:

* Create/Delete MongoDB Atlas Project

## Prerequisites

* A running Kubernetes cluster, for example Minikube
* A MongoDB Atlas Account
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Getting Started

This example creates a MongoDB Atlas project:

First, create the MongoDB Atlas project CRD and some RBAC:

```shell
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml
kubectl create -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_crd.yaml
```

Create a Kubernetes secret containing the Private Key of the [MongoDB Atlas Programmatic API Key](https://docs.atlas.mongodb.com/configure-api-access/#programmatic-api-keys)

```shell
kubectl create secret generic example-monogdb-atlas-project \
    --from-literal=privateKey=xxxxxxxxx
```

Deploy the MongoDB Atlas Project Operator:

```shell
kubectl apply -f deploy/operator.yaml
```

Adapt the `publicKey` and `orgId` in [knappek_v1alpha1_mongodbatlasproject_cr.yaml](./deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml) accordingly and deploy your first MongoDB Atlas Project

```shell
kubectl apply -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml
```

## Cleanup

```shell
kubectl delete mongodbatlasproject example-project
kubectl delete -f deploy/
kubectl delete -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_crd.yaml
```

## Developers Build Guide

**Create MongoDB Atlas Project CRD**

Run this once:

```shell
make project
```

**Run Operator locally**

```shell
export KUBECONFIG=/path/to/config
make dev
```

**Run Operator in Kubernetes Cluster**

```shell
make deploy DOCKERHUB_USERNAME=<your-github-user>
```

**Create MongoDB Atlas Project**

```shell
make deploy-project
```

**Delete MongoDB Atlas Project**

```shell
make delete-project
```

**Run Operator Scorecard Tests**

This will run the [Operator Scorecard tests](https://github.com/operator-framework/operator-sdk/blob/master/doc/test-framework/scorecard.md):

```shell
make test
```

## Contributing

Read through the [Contributing Guidelines and Code of Conduct](./CONTRIBUTING.md).
