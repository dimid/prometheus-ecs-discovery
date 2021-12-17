package main

import "testing"

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
