package k8s

import (
	"context"
	"fmt"
	"time"

	"github.com/orkarstoft/kscale/pkg/config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ScaleNamespaceUp(namespace string, additionalTime time.Duration) {
	c := client()

	// Create timestamps for annotations
	t := time.Now()
	now := t.Format("2006-01-02T15:04:05+00:00")
	end := t.Add(additionalTime).Format("2006-01-02T15:04:05+00:00")
	timeString := fmt.Sprintf("%s-%s", now, end)

	if config.Config.Debug {
		fmt.Printf("[DEBUG]: timeString is set to: %s\n", timeString)
	}

	// Get namespace
	result, err := c.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	// Set annotation on namespace
	result.Annotations["downscaler/force-uptime"] = timeString

	// Update namespace
	_, err = c.CoreV1().Namespaces().Update(context.TODO(), result, metav1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("[INFO]: Namespace %s is now up for %v hours\n", namespace, additionalTime.Hours())
}
