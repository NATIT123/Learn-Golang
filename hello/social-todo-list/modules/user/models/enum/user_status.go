package enum

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type UserStatus int

const (
	UserStatusActive UserStatus = iota
	UserStatusInactive
	UserStatusDeleted
)

var allUserStatuses = [3]string{"Active", "Inactive", "Deleted"}

func (User UserStatus) String() string {
	return allUserStatuses[User]
}

func parseStr2UserStatus(s string) (UserStatus, error) {
	for i := range allUserStatuses {
		if allUserStatuses[i] == s {
			return UserStatus(i), nil
		}
	}

	return UserStatus(0), errors.New("invalid status string")
}

func (User *UserStatus) Scan(value interface{}) error {
	if value == nil {
		return errors.New("nil value provided for UserStatus")
	}

	var strValue string

	switch v := value.(type) {
	case []byte:
		strValue = string(v)
	case string:
		strValue = v
	default:
		return fmt.Errorf("fail to scan data from sql: %v", value)
	}

	status, err := parseStr2UserStatus(strValue)
	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %v", err)
	}

	*User = status
	return nil
}

func (User *UserStatus) Value() (driver.Value, error) {
	if User == nil {
		return nil, nil
	}

	return User.String(), nil
}

func (User *UserStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	UserValue, err := parseStr2UserStatus(str)

	if err != nil {
		return err
	}

	*User = UserValue

	return nil
}

func (User *UserStatus) MarshalJSON() ([]byte, error) {
	if User == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", User.String())), nil
}
