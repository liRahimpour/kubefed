/*
Copyright 2018 The Kubernetes Authors.

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
	apiextv1b1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KubefedConfigSpec defines the desired state of KubefedConfig
type KubefedConfigSpec struct {
	// The scope of the kubefed control plane should be either `Namespaced` or `Cluster`.
	// `Namespaced` indicates that the kubefed namespace will be the only target for federation.
	Scope              apiextv1b1.ResourceScope `json:"scope"`
	ControllerDuration DurationConfig           `json:"controllerDuration"`
	LeaderElect        LeaderElectConfig        `json:"leaderElect"`
	FeatureGates       []FeatureGatesConfig     `json:"featureGates"`
	ClusterHealthCheck ClusterHealthCheckConfig `json:"clusterHealthCheck"`
	SyncController     SyncControllerConfig     `json:"syncController"`
}

type DurationConfig struct {
	// Time to wait before reconciling on a healthy cluster.
	AvailableDelay metav1.Duration `json:"availableDelay"`
	// Time to wait before giving up on an unhealthy cluster.
	UnavailableDelay metav1.Duration `json:"unavailableDelay"`
}
type LeaderElectConfig struct {
	// The duration that non-leader candidates will wait after observing a leadership
	// renewal until attempting to acquire leadership of a led but unrenewed leader
	// slot. This is effectively the maximum duration that a leader can be stopped
	// before it is replaced by another candidate. This is only applicable if leader
	// election is enabled.
	LeaseDuration metav1.Duration `json:"leaseDuration"`
	// The interval between attempts by the acting master to renew a leadership slot
	// before it stops leading. This must be less than or equal to the lease duration.
	// This is only applicable if leader election is enabled.
	RenewDeadline metav1.Duration `json:"renewDeadline"`
	// The duration the clients should wait between attempting acquisition and renewal
	// of a leadership. This is only applicable if leader election is enabled.
	RetryPeriod metav1.Duration `json:"retryPeriod"`
	// The type of resource object that is used for locking during
	// leader election. Supported options are `configmaps` (default) and `endpoints`.
	ResourceLock ResourceLockType `json:"resourceLock"`
}

type ResourceLockType string

const (
	ConfigMapsResourceLock ResourceLockType = "configmaps"
	EndpointsResourceLock  ResourceLockType = "endpoints"
)

type FeatureGatesConfig struct {
	Name          string            `json:"name"`
	Configuration ConfigurationMode `json:"configuration"`
}

type ConfigurationMode string

const (
	ConfigurationEnabled  ConfigurationMode = "Enabled"
	ConfigurationDisabled ConfigurationMode = "Disabled"
)

type ClusterHealthCheckConfig struct {
	// How often to monitor the cluster health (in seconds).
	PeriodSeconds int64 `json:"periodSeconds"`
	// Minimum consecutive failures for the cluster health to be considered failed after having succeeded.
	FailureThreshold int64 `json:"failureThreshold"`
	// Minimum consecutive successes for the cluster health to be considered successful after having failed.
	SuccessThreshold int64 `json:"successThreshold"`
	// Number of seconds after which the cluster health check times out.
	TimeoutSeconds int64 `json:"timeoutSeconds"`
}

type SyncControllerConfig struct {
	// Whether to adopt pre-existing resources in member clusters. Defaults to
	// "Enabled".
	AdoptResources ResourceAdoption `json:"adoptResources"`
}

type ResourceAdoption string

const (
	AdoptResourcesEnabled  ResourceAdoption = "Enabled"
	AdoptResourcesDisabled ResourceAdoption = "Disabled"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubefedConfig
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=kubefedconfigs
type KubefedConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KubefedConfigSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubefedConfigList contains a list of KubefedConfig
type KubefedConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubefedConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubefedConfig{}, &KubefedConfigList{})
}
