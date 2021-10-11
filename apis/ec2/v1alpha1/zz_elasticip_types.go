/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ElasticIPObservation struct {
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	AssociationID *string `json:"associationId,omitempty" tf:"association_id,omitempty"`

	CarrierIP *string `json:"carrierIp,omitempty" tf:"carrier_ip,omitempty"`

	CustomerOwnedIP *string `json:"customerOwnedIp,omitempty" tf:"customer_owned_ip,omitempty"`

	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	PrivateDNS *string `json:"privateDns,omitempty" tf:"private_dns,omitempty"`

	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	PublicDNS *string `json:"publicDns,omitempty" tf:"public_dns,omitempty"`

	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ElasticIPParameters struct {

	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// +kubebuilder:validation:Optional
	AssociateWithPrivateIP *string `json:"associateWithPrivateIp,omitempty" tf:"associate_with_private_ip,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerOwnedIPv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool,omitempty"`

	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkBorderGroup *string `json:"networkBorderGroup,omitempty" tf:"network_border_group,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterface *string `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIPv4Pool *string `json:"publicIpv4Pool,omitempty" tf:"public_ipv4_pool,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Vpc *bool `json:"vpc,omitempty" tf:"vpc,omitempty"`
}

// ElasticIPSpec defines the desired state of ElasticIP
type ElasticIPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ElasticIPParameters `json:"forProvider"`
}

// ElasticIPStatus defines the observed state of ElasticIP.
type ElasticIPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ElasticIPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticIP is the Schema for the ElasticIPs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ElasticIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticIPSpec   `json:"spec"`
	Status            ElasticIPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticIPList contains a list of ElasticIPs
type ElasticIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticIP `json:"items"`
}

// Repository type metadata.
var (
	ElasticIPKind             = "ElasticIP"
	ElasticIPGroupKind        = schema.GroupKind{Group: Group, Kind: ElasticIPKind}.String()
	ElasticIPKindAPIVersion   = ElasticIPKind + "." + GroupVersion.String()
	ElasticIPGroupVersionKind = GroupVersion.WithKind(ElasticIPKind)
)

func init() {
	SchemeBuilder.Register(&ElasticIP{}, &ElasticIPList{})
}
