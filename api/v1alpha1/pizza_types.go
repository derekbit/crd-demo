/*
Copyright 2022.

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

// PizzaSpec defines the desired state of Pizza
type PizzaSpec struct {
	Size     string   `json:"size"`
	Price    int      `json:"price"`
	Toppings []string `json:"toppings"`
}

// PizzaStatus defines the observed state of Pizza
type PizzaStatus struct {
	Price     int `json:"price"`
	PriceRise int `json:"price_rise"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Pizza is the Schema for the pizzas API
type Pizza struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PizzaSpec   `json:"spec,omitempty"`
	Status PizzaStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PizzaList contains a list of Pizza
type PizzaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pizza `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pizza{}, &PizzaList{})
}
