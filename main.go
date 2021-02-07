package main

import (
	"context"
	"fmt"
	"time"

	"enimaxpods/utils"

	"github.com/dariubs/percent"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset
var threshold float64 = 90

func main() {

	clientset, err := getClient("")
	failOnError(err, "Failed to create kubernetes clientset")

	api := clientset.CoreV1()

	// reconcile loop
	for {

		nodes, err := api.Nodes().List(context.TODO(), metav1.ListOptions{})
		failOnError(err, "Failed to get list of kubernetes nodes")

		iterateNodes(nodes)

		time.Sleep(60 * time.Second)
	}

}

func iterateNodes(nodes *v1.NodeList) {

	limitsMap := utils.Ec2Limits()

	if len(nodes.Items) > 0 {
		for _, node := range nodes.Items {

			if node.Labels["kubernetes.io/role"] == "node" {

				if podsLimits, ok := limitsMap[node.Labels["beta.kubernetes.io/instance-type"]]; ok {

					logrus.Info(
						"\n",
						"Node name: ", node.Name, "\n",
						"Instance role: ", node.Labels["kubernetes.io/role"], "\n",
						"Instance type: ", node.Labels["beta.kubernetes.io/instance-type"], "\n",
						"Current pods: ", node.Status.Capacity.Pods(), "\n",
						"Pods limit: ", podsLimits, "\n",
						"------------------------")

					var currentPods = node.Status.Capacity.Pods().Value()
					checkPodLimits(node.Name, podsLimits, currentPods)
				}
			}
		}
	}
}

// helper functions

func getClient(pathToCfg string) (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	if pathToCfg == "" {
		logrus.Info("Using in cluster config")
		config, err = rest.InClusterConfig()
		// in cluster access
	} else {
		logrus.Info("Using out of cluster config")
		config, err = clientcmd.BuildConfigFromFlags("", pathToCfg)
	}
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func failOnError(err error, msg string) {
	if err != nil {
		logrus.Error("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))

	}
}

func checkPodLimits(nodeName string, podsLimits int, currentPods int64) {

	usedPercent := percent.PercentOf(int(currentPods), int(podsLimits))

	if usedPercent >= threshold {
		logrus.Warn("WARNING! number of pods ", currentPods, " of ", podsLimits, " allowed, on Node ", nodeName, " , usage of ", usedPercent, "%")
	}
	return
}
