package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

func main() {
	testYamlFile := "Test1.yaml"

	yamlByteArrays, err := splitYamlFile(testYamlFile)
	if err != nil {
		log.Fatalf("Failed to read yaml file %s : %v", testYamlFile, err)
		return
	}

	if len(yamlByteArrays) != 4 {
		log.Fatalf("Test yaml file %s must contain 4 ConfigMap entries, it contains %d ConfigMap entries", testYamlFile, len(yamlByteArrays))
		return
	}

	testConfigMaps := [4]corev1.ConfigMap{}

	for i := 0; i < 4; i++ {
		err := yaml.Unmarshal(yamlByteArrays[i], &testConfigMaps[i])
		if err != nil {
			log.Fatalf("Failed to unmarshal ConfigMap %d", i)
			return
		}
	}

	if len(os.Args) > 1 {
		deleteAllPackages(testConfigMaps)
		return
	}

	err = createBlueprint(testConfigMaps[0])
	if err != nil {
		return
	}

	err = pullPushPackage(testConfigMaps[0])
	if err != nil {
		return
	}

	err = proposeApprovePackage(testConfigMaps[0])
	if err != nil {
		return
	}

	err = clonePackage(testConfigMaps[0].Annotations["package-rev"], testConfigMaps[1])
	if err != nil {
		return
	}

	err = pullPushPackage(testConfigMaps[1])
	if err != nil {
		return
	}

	err = proposeApprovePackage(testConfigMaps[1])
	if err != nil {
		return
	}

	err = copyPackage(testConfigMaps[0].Annotations["package-rev"], testConfigMaps[2])
	if err != nil {
		return
	}

	err = pullPushPackage(testConfigMaps[2])
	if err != nil {
		return
	}

	err = proposeApprovePackage(testConfigMaps[2])
	if err != nil {
		return
	}
}

func splitYamlFile(file string) ([][]byte, error) {
	yamlFileByteArray, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	yamlStringArray := strings.Split(string(yamlFileByteArray), "\n---")

	yamlByteArrays := [][]byte{}
	for _, yamlString := range yamlStringArray {
		yamlByteArray := []byte(strings.TrimSpace(yamlString))
		if len(yamlByteArray) > 0 {
			yamlByteArrays = append(yamlByteArrays, yamlByteArray)
		}
	}
	return yamlByteArrays, nil
}

func createBlueprint(bpCm corev1.ConfigMap) error {
	stdout, stderr, err := commandInShell("porchctl rpkg init -n " + bpCm.GetNamespace() + " " + bpCm.Annotations["package-name"] + " --workspace v1 --repository management")
	if err != nil {
		log.Printf("error: %v", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("stdout:" + stdout)

	bpCm.GetAnnotations()["package-rev"] = strings.Split(stdout, " ")[0]
	return nil
}

func pullPushPackage(bpCm corev1.ConfigMap) error {
	pullDir, err := os.MkdirTemp("", "porch")
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	defer os.RemoveAll(pullDir)

	pullPackagePath := pullDir + "/" + bpCm.Annotations["package-name"]

	_, stderr, err := commandInShell("porchctl rpkg pull -n " + bpCm.GetNamespace() + " " + bpCm.Annotations["package-rev"] + " " + pullPackagePath)
	if err != nil {
		log.Printf("error: %v\n", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("package pulled to " + pullPackagePath)

	if err = addConfigMapToPackage(bpCm, pullPackagePath); err != nil {
		return err
	}

	_, stderr, err = commandInShell("porchctl rpkg push -n " + bpCm.GetNamespace() + " " + bpCm.Annotations["package-rev"] + " " + pullPackagePath)
	if err != nil {
		log.Printf("error: %v\n", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("package pushed from " + pullPackagePath)

	return nil
}

func clonePackage(sourcePackageRev string, bpCm corev1.ConfigMap) error {
	stdout, stderr, err := commandInShell(
		"porchctl rpkg clone -n " + bpCm.GetNamespace() + " " + sourcePackageRev + " " + bpCm.Annotations["package-name"] + " --repository edge1 --workspace v1 --strategy " + bpCm.Annotations["clone-strategy"])
	if err != nil {
		log.Printf("error: %v", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("stdout:" + stdout)

	bpCm.GetAnnotations()["package-rev"] = strings.Split(stdout, " ")[0]
	return nil
}

func copyPackage(sourcePackageRev string, bpCm corev1.ConfigMap) error {
	stdout, stderr, err := commandInShell(
		"porchctl rpkg copy -n " + bpCm.GetNamespace() + " " + sourcePackageRev + " --workspace v2 --replay-strategy=" + bpCm.Annotations["replay-strategy"])
	if err != nil {
		log.Printf("error: %v", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("stdout:" + stdout)

	bpCm.GetAnnotations()["package-rev"] = strings.Split(stdout, " ")[0]
	return nil
}

func proposeApprovePackage(bpCm corev1.ConfigMap) error {
	stdout, stderr, err := commandInShell("porchctl rpkg propose -n " + bpCm.GetNamespace() + " " + bpCm.Annotations["package-rev"])
	if err != nil {
		log.Printf("error: %v", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("stdout:" + stdout)
	fmt.Println("stderr:" + stderr)

	stdout, stderr, err = commandInShell("porchctl rpkg approve -n " + bpCm.GetNamespace() + " " + bpCm.Annotations["package-rev"])
	if err != nil {
		log.Printf("error: %v", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("stdout:" + stdout)
	fmt.Println("stderr:" + stderr)
	return nil
}

func deleteAllPackages(testConfigMaps [4]corev1.ConfigMap) {
	for i := 0; i < 4; i++ {
		pkgRevs := getPackageRevs4Package(testConfigMaps[i])

		for _, pkgRev := range pkgRevs {
			deletePackage(testConfigMaps[i].GetNamespace(), pkgRev)
		}
	}
}

func deletePackage(ns string, packageRev string) {
	stdout, stderr, _ := commandInShell("porchctl rpkg propose-delete -n " + ns + " " + packageRev)
	fmt.Println("stdout:" + stdout)
	fmt.Println("stderr:" + stderr)

	stdout, stderr, _ = commandInShell("porchctl rpkg delete -n " + ns + " " + packageRev)
	fmt.Println("stdout:" + stdout)
	fmt.Println("stderr:" + stderr)
}

func getPackageRevs4Package(cm corev1.ConfigMap) []string {
	pkgRevs := []string{}
	stdout, _, err := commandInShell("porchctl rpkg list -n " + cm.GetNamespace() + "| grep ' " + cm.Annotations["package-name"] + "'")

	if err != nil {
		fmt.Println("No packageRevs found")
		return pkgRevs
	}

	pkgRevLines := strings.Split(stdout, "\n")

	for _, pkgRevLine := range pkgRevLines {
		pkgRev := strings.Split(pkgRevLine, " ")[0]
		if len(pkgRev) > 0 {
			pkgRevs = append(pkgRevs, pkgRev)
		}
	}
	fmt.Println(pkgRevs)

	return pkgRevs
}

func addConfigMapToPackage(bpCm corev1.ConfigMap, pullPackagePath string) error {
	cmYamlByteArray, err := yaml.Marshal(bpCm)
	if err != nil {
		log.Printf("error: %v", err)
		fmt.Printf("Could not marshal ConfigMap into yaml\n%v\n", bpCm)
		return err
	}

	err = os.WriteFile(pullPackagePath+"/"+bpCm.Annotations["package-name"]+".yaml", cmYamlByteArray, 0644)
	if err != nil {
		log.Printf("error: %v", err)
		fmt.Printf("Could not write ConfigMap to package at %s\n%v\n", pullPackagePath, bpCm)
		return err
	}
	return nil
}

func commandInShell(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	fmt.Println(command)
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
