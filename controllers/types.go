package controllers

// Core Spec

type TigerGraphClusterSpec struct {
	Image    string
	Replicas *int32
	License  LicenseSpec

	Storage  StorageSpec
	Listener ListenerSpec
}

type LicenseSpec struct {
	SecretName string
	SecretKey  string
}

type StorageSpec struct {
	Type                string
	VolumeClaimTemplate VolumeClaimTemplate
	AdditionalStorages  []AdditionalStorage
}

type VolumeClaimTemplate struct {
	StorageClassName string
	AccessModes      []string
	Storage          string
	VolumeMode       string
}

type AdditionalStorage struct {
	Name             string
	StorageSize      string
	StorageClassName string
}

type ListenerSpec struct {
	Type string
}
