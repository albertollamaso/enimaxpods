# Kubernetes Max Pods checker per Node

When you create an Amazon EKS node, it has one network interface. All Amazon EC2 instance types support more than one network interface. The network interface attached to the instance when the instance is created is called the primary network interface. Any additional network interface attached to the instance is called a secondary network interface. Each network interface can be assigned multiple private IP addresses. One of the private IP addresses is the primary IP address, whereas all other addresses assigned to the network interface are secondary IP addresses.

Each pod that you deploy is assigned one secondary private IP address from one of the network interfaces attached to the instance. For instance an m5.large instance supports three network interfaces and ten private IP addresses for each network interface. Even though an m5.large instance supports 30 private IP addresses, you can’t deploy 30 pods to that node. To determine how many pods you can deploy to a node, use the following formula:

`(Number of network interfaces for the instance type × (the number of IP addressess per network interface - 1)) + 2`


This client that works as a deployment in Kubernetes check the amount of pods of each worker node and print the output in the logs for easy check of pods usage on Kubernetes nodes.

### How to use it?

Within your kubernetes cluster execute below command to deploy to default namespace:
`kubectl apply -f kubernetes/ -n default`


### FAQS


### todo: 

- Add alerts when it passed the threashold defined.
- At this moment the threshold is defined as a global variable in main.go `var threshold float64 = 90` it should be passed as an environment variable to the docker container.