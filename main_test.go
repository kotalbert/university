package main

import "testing"

func Test_translateDepartment(t *testing.T) {
	type args struct {
		department string
	}
	tests := []struct {
		name string
		args args
		want Department
	}{
		{"Test 1", args{"Biotech"}, Biotech},
		{"Test 2", args{"Chemistry"}, Chemistry},
		{"Test 3", args{"Engineering"}, Engineering},
		{"Test 4", args{"Mathematics"}, Mathematics},
		{"Test 5", args{"Physics"}, Physics},
		{"Test 6", args{"Unknown"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateDepartment(tt.args.department); got != tt.want {
				t.Errorf("translateDepartment() = %v, want %v", got, tt.want)
			}
		})
	}
}
