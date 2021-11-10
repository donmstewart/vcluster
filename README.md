# Porter VCluster Mixin

This plugin enables the creation of virtual Kubernetes clusters that run within regular namespaces of a parent cluster.

The mixin enables the use of vcluster from [vcluster.com](https://www.vcluster.com/). The plugin supports: -

* Create
* Update
* Delete

funcions on the virtual cluters.

The mixin has been tested to work with both the [porter cli](https://github.com/getporter/porter) & 
[porter operator for kubernetes](https://github.com/getporter/operator).

## Example `porter.yaml` step

```yaml
parameters:
  - name: kubeconfig
    description: "Parent k8s cluster config"
    type: file
    path: /root/.kube/config
    default: ""
  - name: vcluster
    description: "Name of the vcluster to create"
    type: string
    default: "vcluster-1"
    
mixins:
  - vcluster

install:
  - vcluster:
      description: "Create new vcluster"
      create:
        name: "{{ bundle.parameters.vcluster }}"
        namespace: "{{ bundle.parameters.vcluster }}"
        createNamespace: true
        expose: true
        connect: true
        extraConfig: config/vcluster/values.yaml
```

These are all the configuration items you need to create a new vcluster.

### Enable the mixin

To enable the vcluster mixin and copiy the mixin to the porter bundle.

```yaml
mixins:
  - vcluster
```

### Provision a new `vcluster`

Use the parter `install:` step to provision a new vcluster: -

```yaml
install:
  - vcluster:
      description: "Create new vcluster"
      create:
```

and provide the necessary arguments to the intall/create step: -

```yaml
        name: "{{ bundle.parameters.vcluster }}"
        namespace: "{{ bundle.parameters.vcluster }}"
        createNamespace: true
        expose: true
        connect: true
        extraConfig: config/vcluster/values.yaml
```

In the example above the vlcuster name & namespace are set to the value of a bubdle parameter called `vcluster` shown
above as `"{{ bundle.parameters.vcluster }}"`.

The following two parameters are needed to successfully create a new vcluster.

```yaml
parameters:
  - name: kubeconfig
    description: "Parent k8s cluster config"
    type: file
    path: /root/.kube/config
    default: ""
  - name: vcluster
    description: "Name of the vcluster to create"
    type: string
    default: "vcluster-1"
```

and the kube config of the target host cluster is provided as an additional porter bundle parameter `kubeconfig`.

### Extra Config

vcluster allows for control of certain aspects of the vcluster. For more information on these settings see
[here](https://www.vcluster.com/docs/config-reference).

```yaml
rbac:
  clusterRole:
    create: true

syncer:
  extraArgs: ["--fake-nodes=false", "--sync-all-nodes", "--enable-storage-classes", "--fake-persistent-volumes" ]
```

## Porter CLI Bundle Invocation

To invoke a porter bundle called `porter-vcluster` from the cli the following can be used: -

```bash
porter install porter-vcluster --param kubeconfig=$HOME/.kube/config --param vcluster=vcluster-1
```

## Porter Operator Bundle Invocation

To invoke a porter bundle called `porter-vcluster` from the operator the following CR can be used: -

```yaml
apiVersion: porter.sh/v1
kind: Installation
metadata:
  labels:
    installVersion: "v0.38.7"
  namespace: porter-operator-namespace
  name: porter-vcluster
spec:
  reference: "imagerepoitory/image:tag"
  action: "install"
  parameters:
    vcluster: vcluster-1
    kubeconfig: YXBpVmVyc2lvbjogdjEKa2luZDo ...
```

* NOTE: *
The `kubeconfig` is a base64 encoded kube config yaml for the host cluster.