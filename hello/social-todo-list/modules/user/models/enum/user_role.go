package enum

import (
	"database/sql/driver"
	"fmt"
)

type UserRole int

const (
	RoleUser UserRole = 1 << iota
	RoleAdmin
	RoleShipper
	RoleMod
)

func (role UserRole) String() string {
	switch role {
	case RoleAdmin:
		return "admin"
	case RoleShipper:
		return "shipper"
	case RoleMod:
		return "mod"
	default:
		return "user"
	}
}

func (role *UserRole) Scan(value interface{}) error {
	var roleStr string

	switch v := value.(type) {
	case []byte:
		roleStr = string(v)
	case string:
		roleStr = v
	default:
		return fmt.Errorf("Failed to scan role: unexpected type %T", value)
	}

	switch roleStr {
	case "admin":
		*role = RoleAdmin
	case "shipper":
		*role = RoleShipper
	case "mod":
		*role = RoleMod
	default:
		*role = RoleUser
	}

	return nil
}

func (role *UserRole) Value() (driver.Value, error) {
	if role == nil {
		return nil, nil
	}
	return role.String(), nil
}

func (role *UserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", role.String())), nil
}
