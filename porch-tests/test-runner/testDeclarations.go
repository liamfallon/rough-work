package testRunner

import corev1 "k8s.io/api/core/v1"

type TestContext struct {
	yamlByteArrays [][]byte
	testConfigMaps [4]corev1.ConfigMap
}
