// Copyright 2017 Teralytics.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "testing"

func TestLabelsSetDefault(t *testing.T) {
	testCases := []struct {
		Name          string
		DefaultLabels string
		LabelSet      map[string]string
		Expected      map[string]string
	}{
		{
			Name: "no labels",
			LabelSet: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			Name:          "default labels",
			DefaultLabels: "key1,key3",
			LabelSet: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
			Expected: map[string]string{
				"key1": "value1",
				"key3": "value3",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := NewLabels(tc.DefaultLabels)
			for k, v := range tc.LabelSet {
				l.SetDefault(k, v)
			}

			result := l.labels
			for ekey, evalue := range tc.Expected {

				value, ok := result[ekey]
				if !ok {
					t.Errorf("label %s, not found", ekey)
					return
				}

				if value != evalue {
					t.Errorf("label %s, expected value %s, got %s", ekey, evalue, value)
					return
				}

				delete(result, ekey)
			}

			if len(result) != 0 {
				t.Errorf("there are unexpected labels: %s", result)
			}
		})
	}
}

func TestForEachLabel(t *testing.T) {
	testCases := []struct {
		Name     string
		Labels   string
		Expected map[string]string
	}{
		{
			Name:   "no labels",
			Labels: "",
		},
		{
			Name:   "simple string",
			Labels: "invalid",
		},
		{
			Name:   "invalid spec",
			Labels: " = ",
		},
		{
			Name:   "empty value",
			Labels: "key=",
		},
		{
			Name:   "empty key",
			Labels: "=value",
		},
		{
			Name:   "simple valid",
			Labels: "name=value",
			Expected: map[string]string{
				"name": "value",
			},
		},
		{
			Name:   "simple valid with ws",
			Labels: " name = value ",
			Expected: map[string]string{
				"name": "value",
			},
		},
		{
			Name:   "partially valid, ignore invalid labels",
			Labels: "=, name = value , empty=",
			Expected: map[string]string{
				"name": "value",
			},
		},
		{
			Name:   "complex lables spec",
			Labels: "name1=value1,name2=value2",
			Expected: map[string]string{
				"name1": "value1",
				"name2": "value2",
			},
		},
		{
			Name:   "complex with invalid items",
			Labels: " = ,name1=value1,name2=value2, empty=",
			Expected: map[string]string{
				"name1": "value1",
				"name2": "value2",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			forEachLabel(tc.Labels, func(key, value string) {
				if tc.Expected == nil {
					t.Errorf("unexpected callback call")
					return
				}

				v, ok := tc.Expected[key]
				if !ok {
					t.Errorf("unexpected label %s/%s", key, value)
					return
				}

				if v != value {
					t.Errorf("label %s, expected value: %s, got: %s", key, v, value)
				}

				delete(tc.Expected, key)
			})

			if len(tc.Expected) != 0 {
				t.Errorf("expected labels not found: %s", tc.Expected)
			}
		})
	}
}
