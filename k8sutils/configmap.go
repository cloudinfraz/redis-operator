package k8sutils

import (
	"context"
	"fmt"
	redisv1beta1 "redis-operator/api/v1beta1"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateredisConf is a method that will generate a Redis Config interface
func GeneratedefConf(cr *redisv1beta1.Redis, role string) string {
	var initConf = make(map[string][]string)
	initConf["cluster-enabled"] = []string{"yes"}
	initConf["cluster-node-timeout"] = []string{"5000"}
	initConf["cluster-require-full-coverage"] = []string{"no"}
	initConf["cluster-migration-barrier"] = []string{"1"}
	initConf["save"] = []string{"900 1", "300 10", "60 10000"}
	initConf["appendonly"] = []string{"yes"}
	initConf["cluster-config-file"] = []string{"/data/nodes.conf"}
	initConf["dir"] = []string{"/data"}
	initConf["appendfilename"] = []string{"\"appendonly.aof\""}
	switch {
	case role == "master":
		masterRediscfg := &cr.Spec.Master.RedisConfig
		if len(*masterRediscfg) == 0 {
			return AddsliceElements(initConf)
		} else {
			masterRedisconfig := MergeMap(initConf, *masterRediscfg)
			return AddsliceElements(*masterRedisconfig)
		}
	case role == "slave":
		slaveRediscfg := &cr.Spec.Slave.RedisConfig
		if len(*slaveRediscfg) == 0 {
			return AddsliceElements(initConf)
		} else {
			slaveRedisconfig := MergeMap(initConf, *slaveRediscfg)
			return AddsliceElements(*slaveRedisconfig)
		}
	default:
		return AddsliceElements(initConf)
	}
}

// AddsliceElements is a function that will generate a Redis Config String
func AddsliceElements(redisConf map[string][]string) string {
	defConf := make([]string, 0)
	for k, v := range redisConf {
		if len(v) > 1 {
			for i := range v {
				line1 := k + " " + v[i]
				defConf = append(defConf, line1)
			}
		} else {
			line2 := k + " " + v[0]
			defConf = append(defConf, line2)
		}
	}
	return fmt.Sprintf(`%s`, strings.Join(defConf, "\n"))
}

// MergeMaps is a function that will merge two maps
func MergeMap(oldConf, newConf map[string][]string) *map[string][]string {
	MergeConf := make(map[string][]string)
	for k, v := range oldConf {
		for i, j := range newConf {
			if k == i {
				MergeConf[k] = j
			} else {
				if _, ok := MergeConf[k]; !ok {
					MergeConf[k] = v
				}
				if _, ok := MergeConf[i]; !ok {
					MergeConf[i] = j
				}
			}
		}
	}
	return &MergeConf
}

// GenerateConfigMap is a method that will generate a ConfigMap interface
func GenerateConfigMap(cr *redisv1beta1.Redis, role string) *corev1.ConfigMap {

	configMapData := make(map[string]string, 0)
	configMapData["redis-config"] = GeneratedefConf(cr, role)
	labels := map[string]string{
		"app": cr.ObjectMeta.Name + role + "conf",
	}
	ConfigMap := &corev1.ConfigMap{
		TypeMeta:   GenerateMetaInformation("ConfigMap", "v1"),
		ObjectMeta: GenerateObjectMetaInformation(cr.ObjectMeta.Name+role+"conf", cr.Namespace, labels, GenerateSecretAnots()),
		Data:       configMapData,
	}
	AddOwnerRefToObject(ConfigMap, AsOwner(cr))
	return ConfigMap
}

// CreateRedisConfigMap method will create a redis ConfigMap
func CreateRedisConfigMap(cr *redisv1beta1.Redis, role string) {
	reqLogger := log.WithValues("Request.Namespace", cr.Namespace, "Request.Name", cr.ObjectMeta.Name)
	ConfigMapBody := GenerateConfigMap(cr, role)
	ConfigMapName, err := GenerateK8sClient().CoreV1().ConfigMaps(cr.Namespace).Get(context.TODO(), cr.ObjectMeta.Name, metav1.GetOptions{})
	if err != nil {
		reqLogger.Info("Creating ConfigMap for redis", "ConfigMap.Name", cr.ObjectMeta.Name)
		_, err := GenerateK8sClient().CoreV1().ConfigMaps(cr.Namespace).Create(context.TODO(), ConfigMapBody, metav1.CreateOptions{})
		if err != nil {
			reqLogger.Error(err, "Failed in creating ConfigMap for redis")
		}
	} else if ConfigMapBody != ConfigMapName {
		reqLogger.Info("Reconciling ConfigMap for redis", "ConfigMap.Name", cr.ObjectMeta.Name)
		_, err := GenerateK8sClient().CoreV1().ConfigMaps(cr.Namespace).Update(context.TODO(), ConfigMapBody, metav1.UpdateOptions{})
		if err != nil {
			reqLogger.Error(err, "Failed in updating ConfigMap for redis")
		}
	} else {
		reqLogger.Info("ConfigMap for redis are in sync", "ConfigMap.Name", cr.ObjectMeta.Name)
	}
}
