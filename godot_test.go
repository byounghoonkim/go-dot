package dot

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

var appName = "testapp"
var configName = "testConfig"

func targetPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(usr.HomeDir, "."+appName, "."+configName), nil
}

func removeTargetPath() error {
	target, err := targetPath()
	if err != nil {
		return err
	}

	return os.RemoveAll(filepath.Dir(target))
}

func setupTestCase(t *testing.T) func() {
	t.Log("called setup test case")
	err := removeTargetPath()
	if err != nil {
		t.Fatal(err)
	}

	return func() {
		err := removeTargetPath()
		if err != nil {
			t.Fatal(err)
		}

		t.Log("called teardown test case")
	}

}

type testConfig struct {
	a string
}

func TestDot_getConfigPath(t *testing.T) {
	wantString, err := targetPath()
	if err != nil {
		t.Fatal(err)
	}

	type testConfig struct {
		a string
	}

	type fields struct {
		AppName string
	}
	type args struct {
		configuration interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"succuess test",
			fields{appName},
			args{&testConfig{}},
			wantString,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dot{
				AppName: tt.fields.AppName,
			}
			got, err := d.getConfigPath(tt.args.configuration)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dot.getConfigPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Dot.getConfigPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDot_Save(t *testing.T) {

	// TODO clean up test files...

	teardown := setupTestCase(t)
	defer teardown()

	type fields struct {
		AppName string
	}
	type args struct {
		configuration interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"succuess test",
			fields{appName},
			args{&testConfig{"bbbb"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dot{
				AppName: tt.fields.AppName,
			}
			if err := d.Save(tt.args.configuration); (err != nil) != tt.wantErr {
				t.Errorf("Dot.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := d.Load(tt.args.configuration); (err != nil) != tt.wantErr {
				t.Errorf("Dot.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
