## Install Requisites

See [the TKO documentation](https://github.com/nephio-experimental/tko/blob/main/INSTALL.md).

## Creatte TKO db in Postgres

Issue the following commands in `psql`:
```
CREATE USER tko WITH PASSWORD 'tko';
CREATE DATABASE tko WITH OWNER tko;
```

## Build TKO executables

1. Clone [TKO from git](https://github.com/nephio-experimental/tko)
2. Run `scripts/build`
3. Add `$HOME/go/bin/tko` tko to your path

## Setup Python environment for TKO

```
python3 -m venv /tmp/tko-python-env

. /tmp/tko-python-env/bin/activate

pip install --upgrade pip
pip install ruamel.yaml cbor2 grpcio-tools ansible ansible-builder awxkit
```

## Run the three servers

```
bin/run-tko-data.sh
bin/run-tko-preparer.sh
bin/run-tko-meta-scheduler.sh
```

The TKO GUI should now be available on [localhost:50051]

## Install plugin for Kind

Install the Kind plugin for TKO from wherever you have TKO cloned locally.

```
tko plugin register schedule kind ~/git/github/nephio-experimental/tko/examples/plugins/schedule_kind.py --trigger=kind.x-k8s.io,v1alpha4,Cluster -v
```

Stop and restart the Meta Scheduler (Plugins are not flushed to the meta scheduler).
```
bin/run-tko-meta-scheduler.sh
```

## Create sites from Template

```
tko template register zooby-site:v1.0.0 --url=tko-test-data/sites/zooby-site-template/ --metadata=name=zooby,type=kind-cluster
tko site register zooby/001 zooby-site:v1.0.0 --url=tko-test-data/sites/zooby-001/
tko site register zooby/002 zooby-site:v1.0.0 --url=tko-test-data/sites/zooby-002/
```

## Create a template and deploy workload from template on sites

Create the cluters.

```
tko template register ros/hello-world:v1.0.0 --url tko-test-data/deployments/hello-world/
tko deployment create ros/hello-world:v1.0.0 --site=zooby/001
tko deployment create ros/hello-world:v1.0.0 --site=zooby/002
```

Check the clusters:
```
kind get clusters
zooby-001
zooby-002

% kubectl --context=kind-zooby-001 get pods -A
NAMESPACE            NAME                                              READY   STATUS    RESTARTS   AGE
kube-system          coredns-6f6b679f8f-42ktp                          1/1     Running   0          3m11s
kube-system          coredns-6f6b679f8f-94rvz                          1/1     Running   0          3m11s
kube-system          etcd-zooby-001-control-plane                      1/1     Running   0          3m17s
kube-system          kindnet-pt48j                                     1/1     Running   0          3m11s
kube-system          kube-apiserver-zooby-001-control-plane            1/1     Running   0          3m16s
kube-system          kube-controller-manager-zooby-001-control-plane   1/1     Running   0          3m17s
kube-system          kube-proxy-6zjx4                                  1/1     Running   0          3m11s
kube-system          kube-scheduler-zooby-001-control-plane            1/1     Running   0          3m16s
local-path-storage   local-path-provisioner-57c5987fd4-gvp9v           1/1     Running   0          3m11s

% kubectl --context=kind-zooby-002 get pods -A
NAMESPACE            NAME                                              READY   STATUS    RESTARTS   AGE
kube-system          coredns-6f6b679f8f-k8p6r                          1/1     Running   0          3m2s
kube-system          coredns-6f6b679f8f-rzjgc                          1/1     Running   0          3m2s
kube-system          etcd-zooby-002-control-plane                      1/1     Running   0          3m10s
kube-system          kindnet-css84                                     1/1     Running   0          3m3s
kube-system          kube-apiserver-zooby-002-control-plane            1/1     Running   0          3m10s
kube-system          kube-controller-manager-zooby-002-control-plane   1/1     Running   0          3m9s
kube-system          kube-proxy-tdz9h                                  1/1     Running   0          3m3s
kube-system          kube-scheduler-zooby-002-control-plane            1/1     Running   0          3m9s
local-path-storage   local-path-provisioner-57c5987fd4-wdnc2           1/1     Running   0          3m2s
```

## Use the hello-world template to apply a deployment to the clusters

```
tko template register ros/hello-world:v1.0.0 --url=tko-test-data/templates/hello-world/ -v
tko deployment create ros/hello-world:v1.0.0 --site=zooby/001
tko deployment create ros/hello-world:v1.0.0 --site=zooby/002
```

Check the pods are up on the clusters
```
% kubectl --context=kind-zooby-001 get pods -n hello-world
NAME                           READY   STATUS    RESTARTS   AGE
hello-world-7475cf4c9f-b5kvn   1/1     Running   0          16s

% kubectl --context=kind-zooby-002 get pods -n hello-world
NAME                           READY   STATUS    RESTARTS   AGE
hello-world-7475cf4c9f-prpg4   1/1     Running   0          20s
```