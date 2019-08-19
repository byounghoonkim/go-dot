package dot

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Dot is the main structure of go dot configuration library.
// Create with the dot.New() function.
type Dot struct {
	AppName string
}

// New creates a new dot structure with default Name.
func New() *Dot {
	return &Dot{
		AppName: filepath.Base(os.Args[0]),
	}
}

// Load loads Configuration from Dot files
func (d *Dot) Load(configuration interface{}) error {

	configValue := reflect.ValueOf(configuration)
	if typ := configValue.Type(); typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("configuration should be a pointer to a struct type")
	}

	return d.loadFromYAML(configuration)

}

func (d *Dot) getConfigPath(configuration interface{}) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	configValue := reflect.ValueOf(configuration)
	structName := configValue.Type().String()
	fields := strings.Split(structName, ".")
	fileName := "." + fields[len(fields)-1]

	return filepath.Join(usr.HomeDir, "."+d.AppName, fileName), nil
}

func (d *Dot) loadFromYAML(configuraiton interface{}) error {

	configPath, err := d.getConfigPath(configuraiton)
	if err != nil {
		return err
	}

	configPath += ".yml"
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &configuraiton)
	if err != nil {
		return err
	}

	return nil
}

// Save saves Confituraion to Dot files
func (d *Dot) Save(configuration interface{}) error {
	configValue := reflect.ValueOf(configuration)
	if typ := configValue.Type(); typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("configuration should be a pointer to a struct type")
	}

	return d.saveToYAML(configuration)
}

func (d *Dot) saveToYAML(configuration interface{}) error {
	data, err := yaml.Marshal(configuration)
	if err != nil {
		return err
	}

	configPath, err := d.getConfigPath(configuration)
	if err != nil {
		return err
	}

	configPath += ".yml"

	os.MkdirAll(filepath.Dir(configPath), 0700)

	err = ioutil.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Dump prints Dot structure informaion.
func (d *Dot) Dump() error {
	return nil

}
