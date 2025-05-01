package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type VehicleType string

const (
	VehicleTypeAirplane VehicleType = "AIRPLANE"
	VehicleTypeBus      VehicleType = "BUS"
	VehicleTypeTrain    VehicleType = "TRAIN"
	VehicleTypeShip     VehicleType = "SHIP"
)

var validVehicleTypes = map[VehicleType]bool{
	VehicleTypeAirplane: true,
	VehicleTypeBus:      true,
	VehicleTypeTrain:    true,
	VehicleTypeShip:     true,
}

func (vt VehicleType) Validate() error {
	if vt == "" {
		return nil // Empty is valid for filtering
	}
	if !validVehicleTypes[vt] {
		return fmt.Errorf("invalid vehicle type: %s", vt)
	}
	return nil
}

type VehicleTypes []VehicleType

func (vt *VehicleTypes) Scan(value interface{}) error {
	if value == nil {
		*vt = VehicleTypes{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, vt)
	case string:
		return json.Unmarshal([]byte(v), vt)
	default:
		return fmt.Errorf("unsupported type for VehicleTypes: %T", value)
	}
}

func (vt VehicleTypes) Value() (driver.Value, error) {
	if len(vt) == 0 {
		return "[]", nil
	}
	return json.Marshal(vt)
}

func (vt VehicleTypes) Validate() error {
	if len(vt) == 0 {
		return errors.New("at least one vehicle type is required")
	}

	for _, t := range vt {
		if !validVehicleTypes[t] {
			return fmt.Errorf("invalid vehicle type: %s", t)
		}
	}
	return nil
}

func (vt VehicleTypes) Contains(t VehicleType) bool {
	for _, v := range vt {
		if v == t {
			return true
		}
	}
	return false
}

func (vt VehicleTypes) String() string {
	b, _ := json.Marshal(vt)
	return string(b)
}

func GetValidVehicleTypes() []VehicleType {
	types := make([]VehicleType, 0, len(validVehicleTypes))
	for t := range validVehicleTypes {
		types = append(types, t)
	}
	return types
}
