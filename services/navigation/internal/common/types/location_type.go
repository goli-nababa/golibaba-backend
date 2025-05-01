package types

import (
	"database/sql/driver"
	"fmt"
)

type LocationType string

const (
	LocationTypeAirport      LocationType = "AIRPORT"
	LocationTypeBusTerminal  LocationType = "BUS_TERMINAL"
	LocationTypeTrainStation LocationType = "TRAIN_STATION"
	LocationTypePort         LocationType = "PORT"
)

var validLocationTypes = map[LocationType]bool{
	LocationTypeAirport:      true,
	LocationTypeBusTerminal:  true,
	LocationTypeTrainStation: true,
	LocationTypePort:         true,
}

func (lt *LocationType) Scan(value interface{}) error {
	if value == nil {
		*lt = ""
		return nil
	}

	strVal, ok := value.(string)
	if !ok {
		return fmt.Errorf("unsupported type for LocationType: %T", value)
	}

	*lt = LocationType(strVal)
	return lt.Validate()
}

func (lt LocationType) Value() (driver.Value, error) {
	if err := lt.Validate(); err != nil {
		return nil, err
	}
	return string(lt), nil
}

func (lt LocationType) Validate() error {
	if !validLocationTypes[lt] {
		return fmt.Errorf("invalid location type: %s", lt)
	}
	return nil
}

func (lt LocationType) String() string {
	return string(lt)
}

func GetValidLocationTypes() []LocationType {
	types := make([]LocationType, 0, len(validLocationTypes))
	for t := range validLocationTypes {
		types = append(types, t)
	}
	return types
}
