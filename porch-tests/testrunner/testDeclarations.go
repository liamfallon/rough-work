package testrunner

import corev1 "k8s.io/api/core/v1"

type TestContext struct {
	testFile       string
	yamlByteArrays [][]byte
	testConfigMaps []corev1.ConfigMap
}
