package dot

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var appName = "testapp"
var configName = "testConfig"

func targetPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, "."+appName, configName), nil
}

func removeTargetPath() error {
	target, err := targetPath()
	if err != nil {
		return err
	}

	return os.RemoveAll(filepath.Dir(target))
}

func setupTestCase(t *testing.T) func() {
	err := removeTargetPath()
	if err != nil {
		t.Fatal(err)
	}

	return func() {
		err := removeTargetPath()
		if err != nil {
			t.Fatal(err)
		}
	}

}

type testConfig struct {
	ServerName string
	UserName   string
}

func TestDot_GetConfigPath(t *testing.T) {
	wantString, err := targetPath()
	if err != nil {
		t.Fatal(err)
	}

	type testConfig struct {
		a string
	}

	type fields struct {
		AppName string
		Folder  Folder
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
			fields{
				appName,
				HomeDir,
			},
			args{&testConfig{}},
			wantString,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dot{
				AppName: tt.fields.AppName,
				Folder:  tt.fields.Folder,
			}
			got, err := d.GetConfigPath(tt.args.configuration)
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
	teardown := setupTestCase(t)
	defer teardown()

	type fields struct {
		AppName    string
		FileFormat FileFormat
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
			fields{appName, YAML},
			args{&testConfig{"b", "c"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dot{
				AppName:    tt.fields.AppName,
				FileFormat: tt.fields.FileFormat,
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

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Dot
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDot_Load(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dot{
				AppName: tt.fields.AppName,
			}
			if err := d.Load(tt.args.configuration); (err != nil) != tt.wantErr {
				t.Errorf("Dot.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDot_load(t *testing.T) {
	type fields struct {
		AppName string
	}
	type args struct {
		configuraiton interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dot{
				AppName: tt.fields.AppName,
			}
			if err := d.load(tt.args.configuraiton); (err != nil) != tt.wantErr {
				t.Errorf("Dot.loadFromYAML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDot_save(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dot{
				AppName: tt.fields.AppName,
			}
			if err := d.save(tt.args.configuration); (err != nil) != tt.wantErr {
				t.Errorf("Dot.saveToYAML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDot_Dump(t *testing.T) {
	type fields struct {
		AppName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dot{
				AppName: tt.fields.AppName,
			}
			if err := d.Dump(); (err != nil) != tt.wantErr {
				t.Errorf("Dot.Dump() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDot_GetConfigFolder(t *testing.T) {
	type fields struct {
		AppName string
		Folder  Folder
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dot{
				AppName: tt.fields.AppName,
				Folder:  tt.fields.Folder,
			}
			got, err := d.GetConfigFolder()
			if (err != nil) != tt.wantErr {
				t.Errorf("Dot.getConfigFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Dot.getConfigFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
