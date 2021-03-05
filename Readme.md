## Setup

- A working kubernetes instance
- kubectl command line from local env

### Obtain the node IP address

Go to your kube cluster's UI, navigate to `Worker nodes` tab, then copy the public IP field for the worker node.

### Obtain the node exposed port

```bash
# This needs to run after each time service is deployed.
kubectl get service/deployment-lab -o jsonpath='{.spec.ports[0].nodePort}'
```


## Rolling Deployment

```shell
# deploy v1 app
kubectl apply -f kube/rolling/deployment.v1.yaml
# deploy the service
kubectl apply -f kube/rolling/service.yaml

# Now you should see all web response are from v1
# The IP address and port needs to be updated based on actual service
while true; curl 169.57.112.152:30661; sleep 1; end
# deployment-lab-dp-green-5cc64d976f-hdd9b: hello v1
# deployment-lab-dp-green-5cc64d976f-89pnh: hello v1

# deploy v2 app
kubectl apply -f kube/rolling/deployment.v2.yaml

# at this stage, you will notice pods started to terminate and becoming v2
kubectl get pods

# Now you should see some web response are from v1 while others are from v2
# Once all deployments are done, all traffic should be from v2
# The IP address and port needs to be updated based on actual service
while true; curl 169.57.112.152:30661; sleep 1; end
# deployment-lab-dp-green-5cc64d976f-hdd9b: hello v1
# deployment-lab-dp-green-5cc64d976f-89pnh: hello v2
```

### Cleanup

```shell
kubectl apply -f kube/rolling
```

## Blue / Green Deployment

Switch from `blue` version to `green` version

```shell
# deploy blue app
kubectl apply -f kube/bluegreen/deployment.blue.yaml
# deploy blue service
kubectl apply -f kube/bluegreen/service.blue.yaml

# Now you should see all web response are from blue
# The IP address and port needs to be updated based on actual service
while true; curl 169.57.112.152:30661; sleep 1; end
# deployment-lab-dp-green-5cc64d976f-hdd9b: hello blue
# deployment-lab-dp-green-5cc64d976f-89pnh: hello blue

# deploy green app
kubectl apply -f kube/bluegreen/deployment.green.yaml

# at this stage, all traffic are still returning blue.
# We can wait for all deployments to green to finish
kubectl get pods

# OK, about time to swtich to green version deploy blue service
kubectl apply -f kube/bluegreen/service.green.yaml

# Now you should see all web response are from green now
# The IP address and port needs to be updated based on actual service
while true; curl 169.57.112.152:30661; sleep 1; end
# deployment-lab-dp-green-5cc64d976f-hdd9b: hello green
# deployment-lab-dp-green-5cc64d976f-89pnh: hello green
```

### Cleanup

```shell
kubectl apply -f kube/bluegreen
```

## Canary Deployment

We have `v1` of app running, we'd like to canary deploy a new version of `v3`

```shell
# deploy v1 app
kubectl apply -f kube/canary/deployment.v1.yaml
# deploy service
kubectl apply -f kube/canary/service.yaml

# Now you should see all web response are from v1
# The IP address and port needs to be updated based on actual service
while true; curl 169.57.112.152:30661; sleep 1; end
# deployment-lab-dp-green-5cc64d976f-hdd9b: hello v1
# deployment-lab-dp-green-5cc64d976f-89pnh: hello v1

# deploy canary app v3
kubectl apply -f kube/canary/deployment.v3.yaml

# at this stage, wait for v3 app to finish deploying
kubectl get pods

# Now we should see traffic alternating between v1 and v3. Although v3 should have a much smaller hit ratio
# The IP address and port needs to be updated based on actual service
while true; curl 169.57.112.152:30661; sleep 1; end
# deployment-lab-dp-green-5cc64d976f-hdd9b: hello v1
# deployment-lab-dp-green-5cc64d976f-89pnh: hello v3
```

### Cleanup

```shell
kubectl apply -f kube/canary
```