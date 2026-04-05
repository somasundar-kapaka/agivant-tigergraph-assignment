package controllers

import (
	"fmt"
	"strings"
)

func ValidateClusterSpec(spec TigerGraphClusterSpec) error {

	// Existing checks
	if spec.Image == "" {
		return fmt.Errorf("image required")
	}

	if spec.Replicas == nil || *spec.Replicas == 0 {
		return fmt.Errorf("replicas must be > 0")
	}

	if spec.License.SecretName == "" || spec.License.SecretKey == "" {
		return fmt.Errorf("license required")
	}

	//Storage validation
	if spec.Storage.Type == "" {
		return fmt.Errorf("storage type required")
	}

	if spec.Storage.Type != "persistent-claim" {
		return fmt.Errorf("invalid storage type")
	}

	vct := spec.Storage.VolumeClaimTemplate

	if vct.StorageClassName == "" {
		return fmt.Errorf("storageClassName required")
	}

	if len(vct.AccessModes) == 0 {
		return fmt.Errorf("accessModes required")
	}

	if vct.Storage == "" {
		return fmt.Errorf("storage size required")
	}

	// Additional storage
	for _, s := range spec.Storage.AdditionalStorages {
		if s.Name == "" {
			return fmt.Errorf("additional storage name required")
		}
		if s.StorageSize == "" {
			return fmt.Errorf("additional storage size required")
		}
	}

	// Listener validation
	listenerType := strings.TrimSpace(spec.Listener.Type)

	switch listenerType {
	case "LoadBalancer", "NodePort", "ClusterIP":
		return nil
	default:
		return fmt.Errorf("invalid listener type")
	}
}
