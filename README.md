# Crossplane Packet Provider
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fpackethost%2Fcrossplane-provider-packet.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fpackethost%2Fcrossplane-provider-packet?ref=badge_shield)


## Overview

[From Crossplane's Provider documentation](https://crossplane.io/docs/v0.12/introduction/providers.html):

> Providers extend Crossplane to enable infrastructure resource provisioning. In order to provision a resource, a Custom Resource Definition(CRD) needs to be registered in your Kubernetes cluster and its controller should be watching the Custom Resources those CRDs define. Provider packages contain many Custom Resource Definitions and their controllers.

This is the Crossplane Provider package for [Packet](https://www.packet.com)
infrastructure. The provider that is built from this repository can be installed
into a Crossplane control plane.

## Getting Started and Documentation

For getting started guides, installation, deployment, and administration, see the [Crossplane Documentation](https://crossplane.io/docs/latest).

## Pre-requisites

* Kubernetes cluster
  * For example Minikube, minimum version v0.28+
* Helm, minimum version v2.9.1+.

## Installing Crossplane

For the most up to date, detailed, instructions, check [Crossplane's documentation](https://crossplane.io/docs/v0.12/getting-started/install-configure.html).

The following instructions are provided for convenience.

```console
kubectl create namespace crossplane-system
helm repo add crossplane-alpha https://charts.crossplane.io/alpha
helm install crossplane --namespace crossplane-system crossplane-alpha/crossplane
```

### Install the Crossplane CLI

```console
curl -sL https://raw.githubusercontent.com/crossplane/crossplane-cli/master/bootstrap.sh | bash
```

## Install the Packet Provider

```console
kubectl crossplane package install --cluster \
  --namespace crossplane-system \
  packethost/crossplane-provider-packet:v0.0.2 provider-packet
```

The following commands will require your [Packet API key and a project ID](https://www.packet.com/developers/docs/API/getting-started/). Entering your API key and project ID when prompted:

```console
read -s -p "API Key: " APIKEY; echo
read -p "Project ID: " PROJECT_ID; echo
```

### Create a Provider Secret

Create a [Packet Project and a project level API key](https://www.packet.com/developers/docs/API/getting-started/).

Create a Kubernetes secret with the API Key and Project ID.

```bash
kubectl create -n crossplane-system secret generic packet-creds --from-file=key=<(echo '{"apiKey":"'$APIKEY'", "projectID":"'$PROJECT_ID'"}')
```

### Create a Provider record

Get the project id from the Packet Portal or using the Packet CLI (`packet project get`). With `PROJECT_ID` in your environemnt, run the command below:

```bash
cat << EOS | kubectl apply -f -
apiVersion: packet.crossplane.io/v1alpha2
kind: Provider
metadata:
  name: packet-provider
spec:
  projectID: $PROJECT_ID
  credentialsSecretRef:
    namespace: crossplane-system
    name: packet-creds
    key: key
EOS
```

## Provision a Packet Device

Save the following as `device.yaml`:

```yaml
apiVersion: server.packet.crossplane.io/v1alpha2
kind: Device
metadata:
  name: devices
spec:
  forProvider:
    hostname: crossplane
    plan: c1.small.x86
    facility: any
    operatingSystem: centos_7
    billingCycle: hourly
    hardware_reservation_id: next_available
    locked: false
    tags:
    - crossplane
    - development
  providerRef:
    name: packet-provider
  writeConnectionSecretToRef:
    name: devices-creds
    namespace: crossplane-system
  reclaimPolicy: Delete
```

```bash
$ kubectl create -f device.yaml
device.server.packet.crossplane.io/devices created
secret/devices-creds created
```

To view the device in the cluster:

```bash
$ kubectl get packet -o wide
NAME                                            PROJECT-ID                             AGE   SECRET-NAME
provider.packet.crossplane.io/packet-provider   0ac84673-b679-40c1-9de9-8a8792675515   38m   packet-creds

NAME                                         READY   SYNCED   STATE    ID                                     HOSTNAME     FACILITY   IPV4            RECLAIM-POLICY   AGE
device.server.packet.crossplane.io/devices   True    True     active   1c73767a-e16a-485c-89b4-4b553e1458b3   crossplane   sjc1       139.178.88.35   Delete           19m
```

SSH Connection credentials (including IP address, username, and password) can be found in the provider managed secret defined by `writeConnectionSecretToRef`.

**Caution** - Secret data is Base64 encoded, access to the namespace where this secret is stored offers `root` access to the provisioned device.

```bash
$ kubectl get secret -n crossplane-system devices-creds -o jsonpath='{.data}'; echo
map[endpoint:MTM5LjE3OC44OC41Nw== password:cGFzc3dvcmQ== port:MjI= username:cm9vdA==]
```

To delete the device:

```bash
$ kubectl delete -f device.yaml
device.server.packet.crossplane.io/devices deleted
secret/devices-creds deleted
```

## Roadmap and Stability

This Crossplane provider is alpha quality, not officially supported, and not intended for production use.

Packet `Devices` can be managed through this provider, which provides basic integration.  Advanced features like BGP, VPN, Volume, and SpotMarketRequests are not currently planned.  If you are interested in these features, please let us know by [opening issues](#report-a-bug) and [reaching out](#contact).

## Contributing

crossplane-provider-packet is a community driven project and we welcome contributions. See the Crossplane [Contributing](https://github.com/crossplane/crossplane/blob/master/CONTRIBUTING.md) guidelines to get started.

<!-- TODO(displague) Packet specific contribution pointers -->

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/packethost/crossplane-provider-packet/issues).

## Contact

Please use the following Slack channels to reach members of the community:

* Join the [Crossplane slack #general channel](https://slack.crossplane.io/)
* Join the [Packet slack #community channel](https://packetcommunity.slack.com/)


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fpackethost%2Fcrossplane-provider-packet.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fpackethost%2Fcrossplane-provider-packet?ref=badge_large)