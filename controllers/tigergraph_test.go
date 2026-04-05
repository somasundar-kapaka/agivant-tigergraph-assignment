package controllers

import (
	"testing"
)

//
// Helper
//

func TestListenerEmpty(t *testing.T) {
	spec := validSpec()
	spec.Listener.Type = ""

	err := ValidateClusterSpec(*spec)

	if err == nil {
		t.Error("expected error for empty listener type")
	}
}

func validSpec() *TigerGraphClusterSpec {
	replicas := int32(3)

	return &TigerGraphClusterSpec{
		Image:    "test-image",
		Replicas: &replicas,
		License: LicenseSpec{
			SecretName: "tg-license",
			SecretKey:  "license",
		},
		Storage: StorageSpec{
			Type: "persistent-claim",
			VolumeClaimTemplate: VolumeClaimTemplate{
				StorageClassName: "standard",
				AccessModes:      []string{"ReadWriteOnce"},
				Storage:          "100Gi",
				VolumeMode:       "Filesystem",
			},
			AdditionalStorages: []AdditionalStorage{
				{
					Name:        "tg-kafka",
					StorageSize: "10Gi",
				},
			},
		},
		Listener: ListenerSpec{
			Type: "NodePort",
		},
	}
}

// 1. Missing storage type
func TestValidateMissingStorageType(t *testing.T) {
	spec := validSpec()
	if spec == nil {
		t.Error("invalid spec found 'nil'")
	}
	spec.Storage.Type = ""

	err := ValidateClusterSpec(*spec)

	if err == nil {
		t.Error("expected error for missing storage type")
	}
}

// 2. Missing PVC storage
func TestValidateMissingPVCStorage(t *testing.T) {
	spec := validSpec()
	if spec == nil {
		t.Error("invalid spec found 'nil'")
	}
	spec.Storage.VolumeClaimTemplate.Storage = ""

	err := validSpec()
	if err == nil {
		t.Error("expected error for missing storage size")
	}
}

// 3. Missing additional storage name
func TestValidateAdditionalStorageName(t *testing.T) {
	spec := validSpec()
	if spec == nil {
		t.Error("invalid spec found 'nil'")
	}
	spec.Storage.AdditionalStorages[0].Name = ""

	err := ValidateClusterSpec(*spec)

	if err == nil {
		t.Error("expected error for missing additional storage name")
	}
}

// 4. Missing listener type
func TestValidateMissingListener(t *testing.T) {
	spec := validSpec()
	if spec == nil {
		t.Error("invalid spec found 'nil'")
	}
	spec.Listener.Type = ""

	err := ValidateClusterSpec(*spec)

	if err == nil {
		t.Error("expected error for missing listener type")
	}
}

// 5. Valid case
func TestValidateValidSpec(t *testing.T) {
	spec := validSpec()
	if spec == nil {
		t.Error("invalid spec found 'nil'")
	}

	err := ValidateClusterSpec(*spec)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
