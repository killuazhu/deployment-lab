## Setup

- A working kubernetes instance
- kubectl command line from local env

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