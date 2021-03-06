// Copyright 2018 Nimrod Shneor <nimrodshn@gmail.com>
// and other contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Block is a simple representation of a blockchain block.
type Block struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,inline"`
	Spec              BlockSpec `json:"spec"`
}

// BlockSpec provides specifications for the block.
type BlockSpec struct {
	Data          string `json:"data"`
	Timestamp     int64  `json:"timestamp,omitempty"`
	PrevBlockHash []byte `json:"prev_block_hash,omitempty"`
	Hash          []byte `json:"hash,omitempty"`
	Nonce         int    `json:"nonce,omitempty"`
}

// BlockList is a list of blocks.
type BlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,inline"`

	Items []Block `json:"items"`
}

// Process files in all the fields for our Block type.
func (b *Block) Process(successChan chan<- bool) {

	pow := NewProofOfWork(b)
	nonce, hash := pow.Run()

	b.Spec.Timestamp = time.Now().Unix()
	b.Spec.Hash = hash
	b.Spec.Nonce = nonce

	successChan <- true
}
