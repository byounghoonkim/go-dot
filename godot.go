package dot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Dot is the main struct of dot configuration library.
// Create with the dot.New() function.
type Dot struct {
	AppName    string
	Folder     Folder
	FileFormat FileFormat
}

// Folder Enum
type Folder int

const (
	// HomeDir - user home dir
	HomeDir Folder = iota
	// CurrentDir - current dir
	CurrentDir = iota
)

// FileFormat Enum
type FileFormat string

const (
	// YAML format
	YAML = ".yml"
	//JSON format
	JSON = ".json"
)

// New creates a new dot struct with default values.
// default app name is current executable name
// default folder type is HomeDir
// default FileFormat is yaml
func New() *Dot {
	return &Dot{
		AppName:    filepath.Base(os.Args[0]),
		Folder:     HomeDir,
		FileFormat: YAML,
	}
}

// ByFileFormat set file format enum to Dot struct
func (d *Dot) ByFileFormat(ff FileFormat) *Dot {
	d.FileFormat = ff
	return d
}

// ByFolder set Folder enum to Dot struct
func (d *Dot) ByFolder(f Folder) *Dot {
	d.Folder = f
	return d
}

// Load loads Configuration from Dot files
func (d *Dot) Load(configuration interface{}) error {

	configValue := reflect.ValueOf(configuration)
	if typ := configValue.Type(); typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("configuration should be a pointer to a struct type")
	}

	return d.load(configuration)

}

// GetConfigFolder return full path of config folder
func (d *Dot) GetConfigFolder() (string, error) {
	folder := ""
	switch d.Folder {
	case HomeDir:
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		folder = homeDir
	case CurrentDir:
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		folder = wd
	default:
		return "", fmt.Errorf("unknown folder type")
	}
	return filepath.Join(folder, "."+d.AppName), nil
}

// GetConfigPath  return full path of config file
func (d *Dot) GetConfigPath(configuration interface{}) (string, error) {
	configFolder, err := d.GetConfigFolder()
	if err != nil {
		return "", err
	}

	configValue := reflect.ValueOf(configuration)
	structName := configValue.Type().String()
	fields := strings.Split(structName, ".")
	fileName := fields[len(fields)-1]
	fileName += string(d.FileFormat)

	return filepath.Join(configFolder, fileName), nil
}

func (d *Dot) load(configuration interface{}) error {

	configPath, err := d.GetConfigPath(configuration)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	switch d.FileFormat {
	case YAML:
		err = yaml.Unmarshal(data, configuration)
	case JSON:
		err = json.Unmarshal(data, configuration)
	default:
		err = fmt.Errorf("unsupport file format")

	}

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

	return d.save(configuration)
}

func (d *Dot) save(configuration interface{}) error {
	var err error
	var data []byte

	switch d.FileFormat {
	case YAML:
		data, err = yaml.Marshal(configuration)
	case JSON:
		data, err = json.Marshal(configuration)
	default:
		data, err = nil, fmt.Errorf("unsupport file format")
	}

	if err != nil {
		return err
	}

	configPath, err := d.GetConfigPath(configuration)
	if err != nil {
		return err
	}

	os.MkdirAll(filepath.Dir(configPath), 0700)

	err = ioutil.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Dump prints Dot struct for debugging
func (d *Dot) Dump() error {
	configFolder, err := d.GetConfigFolder()
	if err != nil {
		return err
	}

	fmt.Printf("config folder : %s\n", configFolder)
	fmt.Printf("%+v", d)

	return nil
}
