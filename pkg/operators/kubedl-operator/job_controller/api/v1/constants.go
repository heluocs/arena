package v1

import (
	v1 "k8s.io/api/core/v1"
)

const (
	// ReplicaIndexLabel represents the label key for the replica-index, e.g. the value is 0, 1, 2.. etc
	ReplicaIndexLabel = "replica-index"

	// ReplicaTypeLabel represents the label key for the replica-type, e.g. the value is ps , worker etc.
	ReplicaTypeLabel = "replica-type"

	// ReplicaNameLabel represents the label key for the replica-name, the value is replica name.
	ReplicaNameLabel = "replica-name"

	// GroupNameLabel represents the label key for group name, e.g. the value is kubeflow.org
	GroupNameLabel = "group-name"

	// JobNameLabel represents the label key for the job name, the value is job name
	JobNameLabel = "job-name"

	// JobRoleLabel represents the label key for the job role, e.g. the value is master
	JobRoleLabel = "job-role"
)

// Constant label/annotation keys for job configuration.
const (
	KubeDLPrefix = "kubedl.io"

	// LabelPlatform indicates the platform running instances from.
	LabelPlatform = KubeDLPrefix + "/platform"
	// LabelGangScheduler indicates a specific gang scheduler name.
	LabelGangScheduler = KubeDLPrefix + "/gang-scheduler"
	// LabelDisableGangScheduling indicates job disables gang scheduling.
	LabelDisableGangScheduling = KubeDLPrefix + "/disable-gang-scheduling"

	// LabelTargetNode is the target node allocated by gpu coordinator.
	LabelTargetNode = "alloc-group.scheduling.sigs.k8s.io/target_node"
	// LabelAllocToken is the token of the alloc-consumer pod, which is the UID of the corresponding alloc-reserver pod.
	LabelAllocToken = "alloc-group.scheduling.sigs.k8s.io/token"

	// AnnotationGitSyncConfig annotate git sync configurations.
	AnnotationGitSyncConfig = KubeDLPrefix + "/git-sync-config"
	// AnnotationTenancyInfo annotate tenancy information.
	AnnotationTenancyInfo = KubeDLPrefix + "/tenancy"
	// AnnotationKubeflowTenancyInfo annotate tenancy information.
	AnnotationKubeflowTenancyInfo = "kubeflow.org/tenancy"
	// AnnotationGPUVisibleDevices is the visible gpu devices in view of process.
	AnnotationGPUVisibleDevices = "gpus." + KubeDLPrefix + "/visible-devices"

	// AnnotationTensorBoardConfig annotate tensorboard configurations.
	AnnotationTensorBoardConfig = KubeDLPrefix + "/tensorboard-config"
	// ReplicaTypeTensorBoard is the type for TensorBoard.
	ReplicaTypeTensorBoard ReplicaType = "TensorBoard"
	// ReplicaTypeAllReduceWorker is the type for all-reduce workers
	ReplicaTypeAllReduceWorker ReplicaType = "Worker"
	// JobReplicaTypeAIMaster means the AIMaster role for all job
	JobReplicaTypeAIMaster ReplicaType = "AIMaster"

	// AllocGroupJobName marks which job the allocgroup belongs to.
	// The value is the job name.
	AllocGroupJobName = KubeDLPrefix + "/allocgroupjobname"

	// AnnotationInternalNetworkMode annotate pod network mode
	// Only used internal.
	AnnotationInternalNetworkMode = KubeDLPrefix + "/network-mode"

	// AnnotationEnableHostNetwork annotate job enable hostnetwork mode.
	AnnotationEnableHostNetwork = KubeDLPrefix + "/enable-hostnetwork"
)

const (
	DefaultKubeDLNamespace                 = "kubedl"
	ResourceNvidiaGPU      v1.ResourceName = "nvidia.com/gpu"
)

// NetworkMode specify network mode for job.
type NetworkMode string

const (
	// HostNetworkMode means the pods of the job will use host network.
	HostNetworkMode NetworkMode = "host"
)
