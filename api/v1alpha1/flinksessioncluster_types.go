/*
Copyright 2019 Google LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterState defines states for a cluster.
var ClusterState = struct {
	Reconciling string
	Running     string
	Stopping    string
	Stopped     string
}{
	Reconciling: "Reconciling",
	Running:     "Running",
	Stopping:    "Stopping",
	Stopped:     "Stopped",
}

// ClusterComponentState defines states for a cluster component.
var ClusterComponentState = struct {
	NotReady string
	Ready    string
}{
	NotReady: "NotReady",
	Ready:    "Ready",
}

// JobState defines states for a Flink job.
var JobState = struct {
	Running   string
	Succeeded string
	Failed    string
	Unknown   string
}{
	Running:   "Running",
	Succeeded: "Succeeded",
	Failed:    "Failed",
	Unknown:   "Unknown",
}

// JobRestartPolicy defines the policy for job restart.
var JobRestartPolicy = struct {
	OnFailure string
	Never     string
}{
	OnFailure: "OnFailure",
	Never:     "Never",
}

// ImageSpec defines Flink image of JobManager and TaskManager containers.
type ImageSpec struct {
	// Flink image name.
	Name string `json:"name"`

	// Flink image pull policy.
	PullPolicy *string `json:"pullPolicy,omitempty"`
}

// JobManagerPorts defines ports of JobManager.
type JobManagerPorts struct {
	// RPC port.
	RPC *int32 `json:"rpc,omitempty"`

	// Blob port.
	Blob *int32 `json:"blob,omitempty"`

	// Query port.
	Query *int32 `json:"query,omitempty"`

	// UI port.
	UI *int32 `json:"ui,omitempty"`
}

// JobManagerSpec defines properties of JobManager.
type JobManagerSpec struct {
	// The number of replicas.
	Replicas *int32 `json:"replicas,omitempty"`

	// Ports.
	Ports JobManagerPorts `json:"ports,omitempty"`
}

// TaskManagerPorts defines ports of TaskManager.
type TaskManagerPorts struct {
	// Data port.
	Data *int32 `json:"data,omitempty"`

	// RPC port.
	RPC *int32 `json:"rpc,omitempty"`

	// Query port.
	Query *int32 `json:"query,omitempty"`
}

// TaskManagerSpec defines properties of TaskManager.
type TaskManagerSpec struct {
	// The number of replicas.
	Replicas int32 `json:"replicas"`

	// Ports.
	Ports TaskManagerPorts `json:"ports,omitempty"`
}

// JobSpec defines properties of a Flink job.
type JobSpec struct {
	// JAR file of the job.
	JarFile string `json:"jarFile"`

	// Fully qualified Java class name of the job.
	ClassName *string `json:"className,omitempty"`

	// Args of the job.
	Args []string `json:"args,omitempty"`

	// Savepoint where to restore the job from (e.g., gs://my-savepoint/1234).
	Savepoint *string `json:"savepoint,omitempty"`

	// Allow non-restored state.
	AllowNonRestoredState *bool `json:"allowNonRestoredState,omitempty"`

	// Job parallelism.
	Parallelism *int32 `json:"parallelism,omitempty"`

	// No logging output to STDOUT.
	NoLoggingToStdout *bool `json:"noLoggingToStdout,omitempty"`

	// Restart policy, "OnFailure" or "Never".
	RestartPolicy string `json:"restartPolicy"`

	// TODO(dagang): support volumes and volumeMounts.
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FlinkSessionClusterSpec defines the desired state of FlinkSessionCluster
type FlinkSessionClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The name of the Flink cluster.
	Name string `json:"name"`

	// Flink image spec for the cluster's components.
	ImageSpec ImageSpec `json:"image"`

	// Flink JobManager spec.
	JobManagerSpec JobManagerSpec `json:"jobManager"`

	// Flink TaskManager spec.
	TaskManagerSpec TaskManagerSpec `json:"taskManager"`

	// Optional job spec. If specified, this cluster is an ephemeral Job
	// Cluster, which will be automatically terminated after the job finishes;
	// otherwise, it is a long-running Session Cluster.
	JobSpec *JobSpec `json:"job,omitempty"`
}

// FlinkSessionClusterComponentState defines the observed state of a component
// of a FlinkSessionCluster.
type FlinkSessionClusterComponentState struct {
	// The resource name of the component.
	Name string `json:"name"`

	// The state of the component.
	State string `json:"state"`
}

// FlinkSessionClusterComponentsStatus defines the observed status of the
// components of a FlinkSessionCluster.
type FlinkSessionClusterComponentsStatus struct {
	// The state of JobManager deployment.
	JobManagerDeployment FlinkSessionClusterComponentState `json:"jobManagerDeployment"`

	// The state of JobManager service.
	JobManagerService FlinkSessionClusterComponentState `json:"jobManagerService"`

	// The state of TaskManager deployment.
	TaskManagerDeployment FlinkSessionClusterComponentState `json:"taskManagerDeployment"`
}

type JobStatus struct {
	// The name of the job resource.
	Name string `json:"name"`

	// The state of the job.
	State string `json:"state"`
}

// FlinkSessionClusterStatus defines the observed state of FlinkSessionCluster
type FlinkSessionClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The overall state of the Flink cluster.
	State string `json:"state"`

	// The status of the components.
	Components FlinkSessionClusterComponentsStatus `json:"components"`

	// The status of the (optional) job.
	Job *JobStatus `json:"job,omitempty"`

	// Last update timestamp for this status.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
}

// +kubebuilder:object:root=true

// FlinkSessionCluster is the Schema for the flinksessionclusters API
type FlinkSessionCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FlinkSessionClusterSpec   `json:"spec"`
	Status FlinkSessionClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlinkSessionClusterList contains a list of FlinkSessionCluster
type FlinkSessionClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlinkSessionCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FlinkSessionCluster{}, &FlinkSessionClusterList{})
}
