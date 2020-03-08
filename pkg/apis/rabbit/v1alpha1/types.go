package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Rabbit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RabbitSpec `json:"spec"`
}

type RabbitSpec struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type RabbitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Rabbit `json:"items"`
}

type RabbitWorld struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RabbitWorldSpec `json:"spec"`
}

type RabbitWorldSpec struct {
	Name   string `json:"name"`
	MaxNum int    `json:"max_num"`
	MaxAge int    `json:"max_age"`
}

type RabbitWorldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RabbitWorld `json:"items"`
}
