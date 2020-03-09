package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Rabbit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RabbitSpec   `json:"spec"`
	Status RabbitStatus `json:"status"`
}

type RabbitSpec struct {
	NickName   string `json:"nick_name"`
	Birthday   int64  `json:"birthday"`
	Generation uint   `json:"generation"`
}

type RabbitStatus struct {
	Age int `json:"age"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RabbitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Rabbit `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RabbitWorldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RabbitWorld `json:"items"`
}
