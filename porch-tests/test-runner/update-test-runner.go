package testRunner

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

func ParseTestFile(testYamlFile string) (TestContext, error) {
	ctx := TestContext{}

	var err error
	ctx.yamlByteArrays, err = splitYamlFile(testYamlFile)
	if err != nil {
		log.Fatalf("Failed to read yaml file %s : %v", testYamlFile, err)
		return ctx, err
	}

	if len(ctx.yamlByteArrays) != 4 {
		errMsg := fmt.Sprintf("Test yaml file %s must contain 4 ConfigMap entries, it contains %d ConfigMap entries", testYamlFile, len(ctx.yamlByteArrays))
		log.Fatalln(errMsg)
		return ctx, errors.New(errMsg)
	}

	testConfigMaps := []corev1.ConfigMap{}

	for i := 0; i < 4; i++ {
		err := yaml.Unmarshal(ctx.yamlByteArrays[i], &testConfigMaps[i])
		if err != nil {
			fmt.Println(err)
			return ctx, err
		}
	}

	return ctx, nil
}

func Run(ctx TestContext) {
	err := createBlueprint(ctx.testConfigMaps[0])
	if err != nil {
		return
	}

	err = pullPushPackage(ctx.testConfigMaps[0])
	if err != nil {
		return
	}

	err = proposeApprovePackage(ctx.testConfigMaps[0])
	if err != nil {
		return
	}

	err = clonePackage(ctx.testConfigMaps[0].Annotations["package-rev"], ctx.testConfigMaps[1])
	if err != nil {
		return
	}

	err = pullPushPackage(ctx.testConfigMaps[1])
	if err != nil {
		return
	}

	err = proposeApprovePackage(ctx.testConfigMaps[1])
	if err != nil {
		return
	}

	err = copyPackage(ctx.testConfigMaps[0].Annotations["package-rev"], ctx.testConfigMaps[2])
	if err != nil {
		return
	}

	err = pullPushPackage(ctx.testConfigMaps[2])
	if err != nil {
		return
	}

	err = proposeApprovePackage(ctx.testConfigMaps[2])
	if err != nil {
		return
	}

	err = copyPackage(ctx.testConfigMaps[1].Annotations["package-rev"], ctx.testConfigMaps[3])
	if err != nil {
		return
	}

	err = updatePackage(ctx.testConfigMaps[3])
	if err != nil {
		return
	}

	err = proposeApprovePackage(ctx.testConfigMaps[3])
	if err != nil {
		return
	}

	err = pullCheckPackage(ctx.testConfigMaps[3], string(ctx.yamlByteArrays[3]))
	if err != nil {
		return
	}

	DeleteAllPackages(ctx)
}

func DeleteAllPackages(ctx TestContext) {
	for i := 0; i < 4; i++ {
		pkgRevs := getPackageRevs4Package(ctx.testConfigMaps[i])

		for _, pkgRev := range pkgRevs {
			deletePackage(ctx.testConfigMaps[i].GetNamespace(), pkgRev)
			fmt.Println(pkgRev)
		}
	}
}

func createBlueprint(pkgCm corev1.ConfigMap) error {
	stdout, stderr, err := commandInShell(
		"porchctl rpkg init " +
			"-n " + pkgCm.GetNamespace() + " " +
			pkgCm.Annotations["package-name"] +
			" --repository " + pkgCm.Annotations["package-repo"] +
			" --workspace " + pkgCm.Annotations["workspace"])
	if err != nil {
		log.Printf("error: %v", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("stdout:" + stdout)

	pkgCm.GetAnnotations()["package-rev"] = strings.Split(stdout, " ")[0]
	return nil
}

func pullPushPackage(pkgCm corev1.ConfigMap) error {
	pullDir, err := os.MkdirTemp("", "porch")
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	defer os.RemoveAll(pullDir)

	pullPackagePath := pullDir + "/" + pkgCm.Annotations["package-name"]

	_, stderr, err := commandInShell(
		"porchctl rpkg pull " +
			"-n " + pkgCm.GetNamespace() + " " +
			pkgCm.Annotations["package-rev"] + " " +
			pullPackagePath)
	if err != nil {
		log.Printf("error: %v\n", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("package pulled to " + pullPackagePath)

	if err = addConfigMapToPackage(pkgCm, pullPackagePath); err != nil {
		return err
	}

	_, stderr, err = commandInShell(
		"porchctl rpkg push " +
			"-n " + pkgCm.GetNamespace() + " " +
			pkgCm.Annotations["package-rev"] + " " +
			pullPackagePath)
	if err != nil {
		log.Printf("error: %v\n", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("package pushed from " + pullPackagePath)

	return nil
}

func clonePackage(sourcePackageRev string, pkgCm corev1.ConfigMap) error {
	stdout, stderr, err := commandInShell(
		"porchctl rpkg clone " +
			"-n " + pkgCm.GetNamespace() + " " +
			sourcePackageRev + " " +
			pkgCm.Annotations["package-name"] +
			" --repository " + pkgCm.Annotations["package-repo"] +
			" --workspace " + pkgCm.Annotations["workspace"] +
			" --strategy " + pkgCm.Annotations["clone-strategy"])
	if err != nil {
		log.Printf("error: %v", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("stdout:" + stdout)

	pkgCm.GetAnnotations()["package-rev"] = strings.Split(stdout, " ")[0]
	return nil
}

func copyPackage(sourcePackageRev string, pkgCm corev1.ConfigMap) error {
	stdout, stderr, err := commandInShell(
		"porchctl rpkg copy " +
			"-n " + pkgCm.GetNamespace() + " " +
			sourcePackageRev +
			" --workspace " + pkgCm.Annotations["workspace"] +
			" --replay-strategy=" + pkgCm.Annotations["replay-strategy"])
	if err != nil {
		log.Printf("error: %v", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("stdout:" + stdout)

	pkgCm.GetAnnotations()["package-rev"] = strings.Split(stdout, " ")[0]
	return nil
}

func updatePackage(pkgCm corev1.ConfigMap) error {
	stdout, stderr, err := commandInShell(
		"porchctl rpkg update " +
			"-n " + pkgCm.GetNamespace() + " " +
			pkgCm.Annotations["package-rev"] +
			" --revision=" + pkgCm.Annotations["revision"])
	if err != nil {
		log.Printf("error: %v", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("stdout:" + stdout)

	pkgCm.GetAnnotations()["package-rev"] = strings.Split(stdout, " ")[0]
	return nil
}

func proposeApprovePackage(pkgCm corev1.ConfigMap) error {
	stdout, stderr, err := commandInShell("porchctl rpkg propose -n " + pkgCm.GetNamespace() + " " + pkgCm.Annotations["package-rev"])
	if err != nil {
		log.Printf("error: %v", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("stdout:" + stdout)
	fmt.Println("stderr:" + stderr)

	stdout, stderr, err = commandInShell("porchctl rpkg approve -n " + pkgCm.GetNamespace() + " " + pkgCm.Annotations["package-rev"])
	if err != nil {
		log.Printf("error: %v", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("stdout:" + stdout)
	fmt.Println("stderr:" + stderr)
	return nil
}

func pullCheckPackage(pkgCm corev1.ConfigMap, expectedYaml string) error {
	pullDir, err := os.MkdirTemp("", "porch")
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	defer os.RemoveAll(pullDir)

	pullPackagePath := pullDir + "/" + pkgCm.Annotations["package-name"]

	_, stderr, err := commandInShell(
		"porchctl rpkg pull " +
			"-n " + pkgCm.GetNamespace() + " " +
			pkgCm.Annotations["package-rev"] + " " +
			pullPackagePath)
	if err != nil {
		log.Printf("error: %v\n", err)
		log.Println(stderr)
		return err
	}

	fmt.Println("package pulled to " + pullPackagePath)

	testResourcePath := pullPackagePath + "/" + pkgCm.Annotations["resource-name"] + ".yaml"

	yamlByteArrays, err := splitYamlFile(pullPackagePath + "/" + pkgCm.Annotations["resource-name"] + ".yaml")
	if err != nil {
		log.Fatalf("Failed to read yaml file %s : %v", testResourcePath, err)
		return err
	}

	if len(yamlByteArrays) != 1 {
		errMsg := fmt.Sprintf("Test resource file %s must contain 1 ConfigMap entry, it contains %d ConfigMap entries", testResourcePath, len(yamlByteArrays))
		log.Fatalf(errMsg)
		return errors.New(errMsg)
	}

	testConfigMap := corev1.ConfigMap{}

	err = yaml.Unmarshal(yamlByteArrays[0], &testConfigMap)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to unmarshal test config map in resource file %s", testResourcePath)
		log.Fatalf(errMsg)
		return errors.New(errMsg)
	}

	if packagesAreEqual(pkgCm, testConfigMap) {
		fmt.Println("Upgrade was successful, actual result matches the expected result")
		return nil
	} else {
		errMsg := fmt.Sprintln(
			"Upgrade failed, actual result does not match the expected result\n" +
				"Expected result:\n" +
				string(expectedYaml) + "\n" +
				"Actual result:\n" +
				string(yamlByteArrays[0]))

		log.Fatalf(errMsg)
		return errors.New(errMsg)
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

func packagesAreEqual(pkgCm, testConfigMap corev1.ConfigMap) bool {
	pkgCmNoAnnotation := pkgCm.DeepCopy()

	deleteTestAnnotations(pkgCmNoAnnotation)
	deleteTestAnnotations(&testConfigMap)

	if pkgCmNoAnnotation.APIVersion != testConfigMap.APIVersion {
		return false
	}

	if pkgCmNoAnnotation.Kind != testConfigMap.Kind {
		return false
	}

	if pkgCmNoAnnotation.ObjectMeta.Name != testConfigMap.ObjectMeta.Name {
		return false
	}

	if pkgCmNoAnnotation.ObjectMeta.Namespace != testConfigMap.ObjectMeta.Namespace {
		return false
	}

	if pkgCmNoAnnotation.ObjectMeta.GenerateName != testConfigMap.ObjectMeta.GenerateName {
		return false
	}

	if pkgCmNoAnnotation.ObjectMeta.ResourceVersion != testConfigMap.ObjectMeta.ResourceVersion {
		return false
	}

	if !reflect.DeepEqual(pkgCmNoAnnotation.ObjectMeta.Annotations, testConfigMap.ObjectMeta.Annotations) {
		return false
	}

	if !reflect.DeepEqual(pkgCmNoAnnotation.ObjectMeta.Finalizers, testConfigMap.ObjectMeta.Finalizers) {
		return false
	}

	if !reflect.DeepEqual(pkgCmNoAnnotation.ObjectMeta.Labels, testConfigMap.ObjectMeta.Labels) {
		return false
	}

	return reflect.DeepEqual(pkgCmNoAnnotation.Data, testConfigMap.Data)
}

func getPackageRevs4Package(cm corev1.ConfigMap) []string {
	pkgRevs := []string{}

	if len(cm.Annotations["package-name"]) == 0 {
		return pkgRevs
	}

	stdout, _, err := commandInShell("porchctl rpkg list -n " + cm.GetNamespace() + "| grep ' " + cm.Annotations["package-name"] + " '")

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

func addConfigMapToPackage(pkgCm corev1.ConfigMap, pullPackagePath string) error {
	pkgCmNoAnnotation := pkgCm.DeepCopy()
	deleteTestAnnotations(pkgCmNoAnnotation)

	cmYamlByteArray, err := yaml.Marshal(pkgCmNoAnnotation)
	if err != nil {
		log.Printf("error: %v", err)
		fmt.Printf("Could not marshal ConfigMap into yaml\n%v\n", pkgCmNoAnnotation)
		return err
	}

	err = os.WriteFile(pullPackagePath+"/"+pkgCm.Annotations["resource-name"]+".yaml", cmYamlByteArray, 0644)
	if err != nil {
		log.Printf("error: %v", err)
		fmt.Printf("Could not write ConfigMap to package at %s\n%v\n", pullPackagePath, pkgCm)
		return err
	}
	return nil
}

func deleteTestAnnotations(pkgCmNoAnnotation *corev1.ConfigMap) {
	delete(pkgCmNoAnnotation.Annotations, "package-name")
	delete(pkgCmNoAnnotation.Annotations, "package-rev")
	delete(pkgCmNoAnnotation.Annotations, "package-repo")
	delete(pkgCmNoAnnotation.Annotations, "resource-name")
	delete(pkgCmNoAnnotation.Annotations, "clone-strategy")
	delete(pkgCmNoAnnotation.Annotations, "replay-strategy")
	delete(pkgCmNoAnnotation.Annotations, "workspace")
	delete(pkgCmNoAnnotation.Annotations, "revision")
	delete(pkgCmNoAnnotation.Annotations, "internal.kpt.dev/upstream-identifier")
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
