package monitor

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	deprecatedAPIMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "k8s_deprecated_api",
			Help: "Tracks deprecated Kubernetes APIs in the cluster",
		},
		[]string{"group_version", "resource"},
	)
)

func init() {
	prometheus.MustRegister(deprecatedAPIMetric)
}

func getK8sClient() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error

	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func CheckDeprecatedAPIs(clientset *kubernetes.Clientset) {
	serverGroups, err := clientset.Discovery().ServerGroups()
	if err != nil {
		fmt.Println("Error fetching API groups:", err)
		return
	}

	for _, group := range serverGroups.Groups {
		for _, version := range group.Versions {
			gv := version.GroupVersion
			if strings.Contains(gv, "v1beta") || strings.Contains(gv, "v1alpha") {
				deprecatedAPIMetric.WithLabelValues(gv, "*").Set(1)
				fmt.Println("Deprecated API detected:", gv)
			}
		}
	}
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	clientset, err := getK8sClient()
	if err != nil {
		http.Error(w, "Error getting Kubernetes client", http.StatusInternalServerError)
		return
	}
	CheckDeprecatedAPIs(clientset)
	promhttp.Handler().ServeHTTP(w, r)
}
