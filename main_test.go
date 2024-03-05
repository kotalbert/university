package main

import (
	"reflect"
	"testing"
)

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

func Test_readStudent(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    Student
		wantErr bool
	}{
		{"Test 1", args{"John Doe 3.5 Biotech Chemistry Engineering"}, Student{"John", "Doe", 3.5, Biotech, Chemistry, Engineering}, false},
		{"Test 2", args{"Jane Smith 3.8 Mathematics Physics Chemistry"}, Student{"Jane", "Smith", 3.8, Mathematics, Physics, Chemistry}, false},
		{"Test 3", args{"Alice Johnson 3.9 Physics Mathematics Engineering"}, Student{"Alice", "Johnson", 3.9, Physics, Mathematics, Engineering}, false},
		{"Test 4", args{"Bob Brown 3.7 Chemistry Biotech Mathematics"}, Student{"Bob", "Brown", 3.7, Chemistry, Biotech, Mathematics}, false},
		{"Test 5", args{"Carl White 3.6 Engineering Physics Biotech"}, Student{"Carl", "White", 3.6, Engineering, Physics, Biotech}, false},
		{"Test 6", args{"David Black 3.4 Unknown Unknown Unknown"}, Student{"David", "Black", 3.4, -1, -1, -1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readStudent(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("readStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
