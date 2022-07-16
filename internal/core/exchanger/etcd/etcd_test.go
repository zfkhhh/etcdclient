/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limit
 */

package etcd

import (
	"encoding/json"
	"github.com/acmestack/envcd/internal/envcd"
	"github.com/acmestack/envcd/internal/pkg/config"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	str := "{\"exchanger\": \"etcd://root:root@localhost:2379\"}"
	envcdConfig := &config.Config{}
	if e := json.Unmarshal([]byte(str), envcdConfig); e != nil {
		log.Fatalf("error %v", e)
	}
	envcd.Start(envcdConfig)
	log.Printf("config = %v",envcd.EnvcdConfig.ExchangerConnMetadata)
	tests := []struct {
		name string
		want *Etcd
	}{
		{
			want: New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want == nil || tt.want.client == nil {
				t.Errorf("failed to create client, want %v", tt.want)
			}
		})
	}
}

func TestEtcd_Get(t *testing.T) {
	str := "{\"exchanger\": \"etcd://root:root@localhost:2379\"}"
	envcdConfig := &config.Config{}
	if e := json.Unmarshal([]byte(str), envcdConfig); e != nil {
		log.Fatalf("error %v", e)
	}
	envcd.Start(envcdConfig)
	tests := []struct {
		name string
		want *Etcd
	}{
		{
			want: New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want == nil || tt.want.client == nil {
				t.Errorf("failed to create client, want %v", tt.want)
			}
			get, err := tt.want.Get("/test/a")
			if err != nil {
				t.Errorf("failed to get value,key = /test/a , err = %v", err)
			}
			t.Logf("value = %s", get)
		})
	}
}

func TestEtcd_Put(t *testing.T) {
	str := "{\"exchanger\": \"etcd://root:root@localhost:2379\"}"
	envcdConfig := &config.Config{}
	if e := json.Unmarshal([]byte(str), envcdConfig); e != nil {
		log.Fatalf("error %v", e)
	}
	envcd.Start(envcdConfig)
	tests := []struct {
		name string
		want *Etcd
	}{
		{
			want: New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want == nil || tt.want.client == nil {
				t.Errorf("failed to create client, want %v", tt.want)
			}
			err := tt.want.Put("/test/a", "a")
			if err != nil {
				t.Errorf("failed to put, err = %v", err)
			}
			get, _ := tt.want.Get("/test/a")
			t.Logf("value = %s", get)
		})
	}
}

func TestEtcd_Find(t *testing.T) {
	str := "{\"exchanger\": \"etcd://root:root@localhost:2379\"}"
	envcdConfig := &config.Config{}
	if e := json.Unmarshal([]byte(str), envcdConfig); e != nil {
		log.Fatalf("error %v", e)
	}
	envcd.Start(envcdConfig)
	tests := []struct {
		name string
		want *Etcd
	}{
		{
			want: New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want == nil || tt.want.client == nil {
				t.Errorf("failed to create client, want %v", tt.want)
			}
			values, err := tt.want.Find("/test/a")
			if err != nil {
				t.Errorf("failed to find key /test/a, err = %v", err)
			}
			t.Logf("values = %v", values)
		})
	}
}

func TestEtcd_Remove(t *testing.T) {
	str := "{\"exchanger\": \"etcd://root:root@localhost:2379\"}"
	envcdConfig := &config.Config{}
	if e := json.Unmarshal([]byte(str), envcdConfig); e != nil {
		log.Fatalf("error %v", e)
	}
	envcd.Start(envcdConfig)
	tests := []struct {
		name string
		want *Etcd
	}{
		{
			want: New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want == nil || tt.want.client == nil {
				t.Errorf("failed to create client, want %v", tt.want)
			}
			err := tt.want.Remove("/test/a")
			if err != nil {
				t.Errorf("failed to remove key /test/a, err = %v", err)
			}
		})
	}
}
