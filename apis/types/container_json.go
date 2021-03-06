// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerJSON ContainerJSON contains response of Engine API:
// GET "/containers/{id}/json"
//
// swagger:model ContainerJSON
type ContainerJSON struct {

	// AppArmorProfile are specific for AppArmor to Unix platforms
	AppArmorProfile string `json:"AppArmorProfile,omitempty"`

	// The arguments to the command being run
	Args []string `json:"Args"`

	// config
	Config *ContainerConfig `json:"Config,omitempty"`

	// The time the container was created
	Created string `json:"Created,omitempty"`

	// driver
	Driver string `json:"Driver,omitempty"`

	// exec ids of container
	ExecIds []string `json:"ExecIDs"`

	// graph driver
	GraphDriver *GraphDriverData `json:"GraphDriver,omitempty"`

	// host config
	HostConfig *HostConfig `json:"HostConfig,omitempty"`

	// the path of container's hostname file on host.
	HostnamePath string `json:"HostnamePath,omitempty"`

	// the path of container's hosts file on host.
	HostsPath string `json:"HostsPath,omitempty"`

	// The ID of the container
	ID string `json:"Id,omitempty"`

	// The container's image
	Image string `json:"Image,omitempty"`

	// the path of container's log file on host.
	LogPath string `json:"LogPath,omitempty"`

	// MountLabel contains the options for the 'mount' command.
	MountLabel string `json:"MountLabel,omitempty"`

	// Set of mount point in a container.
	Mounts []MountPoint `json:"Mounts"`

	// name of the created container.
	Name string `json:"Name,omitempty"`

	// NetworkSettings exposes the network settings in the API.
	NetworkSettings *NetworkSettings `json:"NetworkSettings,omitempty"`

	// The path to the command being run
	Path string `json:"Path,omitempty"`

	// process label
	ProcessLabel string `json:"ProcessLabel,omitempty"`

	// the path of container's resolvConf file on host.
	ResolvConfPath string `json:"ResolvConfPath,omitempty"`

	// the container's restart time
	RestartCount int64 `json:"RestartCount,omitempty"`

	// The total size of all the files in this container.
	SizeRootFs *int64 `json:"SizeRootFs,omitempty"`

	// The size of files that have been created or changed by this container.
	SizeRw *int64 `json:"SizeRw,omitempty"`

	// snapshotter
	Snapshotter *SnapshotterData `json:"Snapshotter,omitempty"`

	// The state of the container.
	State *ContainerState `json:"State,omitempty"`
}

// Validate validates this container JSON
func (m *ContainerJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGraphDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerJSON) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerJSON) validateGraphDriver(formats strfmt.Registry) error {

	if swag.IsZero(m.GraphDriver) { // not required
		return nil
	}

	if m.GraphDriver != nil {
		if err := m.GraphDriver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GraphDriver")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerJSON) validateHostConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.HostConfig) { // not required
		return nil
	}

	if m.HostConfig != nil {
		if err := m.HostConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HostConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerJSON) validateMounts(formats strfmt.Registry) error {

	if swag.IsZero(m.Mounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Mounts); i++ {

		if err := m.Mounts[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Mounts" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ContainerJSON) validateNetworkSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkSettings) { // not required
		return nil
	}

	if m.NetworkSettings != nil {
		if err := m.NetworkSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkSettings")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerJSON) validateSnapshotter(formats strfmt.Registry) error {

	if swag.IsZero(m.Snapshotter) { // not required
		return nil
	}

	if m.Snapshotter != nil {
		if err := m.Snapshotter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Snapshotter")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerJSON) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("State")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerJSON) UnmarshalBinary(b []byte) error {
	var res ContainerJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
