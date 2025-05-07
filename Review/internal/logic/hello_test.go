package logic_test

// black box testing : tester don't need to know about source,
//  artitechture or operation internal sytem ,need to know input and output
// white box testing : tester must access to source and know about system internal operation/

import (
	"be-ep/internal/database"
	"be-ep/internal/logic"
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var currentTime time.Time

func testSayHello(t *testing.T) {
	t.Parallel()

	output := logic.SayHello("World")
	assert.Equal(t, "Hello World", output, "Incorrect ouput")

	output = logic.SayHello("")
	assert.Equal(t, "Hello ", output, "Incorrect ouput")
}

func TestCurrentTime(t *testing.T) {
	//Create a controller to manage lifecycle of mock
	mockController := gomock.NewController(t)
	defer mockController.Finish() //Finish

	//Create mock for userDataAccessor
	userDataAccessor := database.NewMockUserDataAccessor(mockController)
	userDataAccessor.EXPECT().GetUser(gomock.Any(), uint64(1)).Return(database.User{
		ID:   1,
		Name: "Nguyen Anh Tu",
	}, nil).
		Times(1)
		// Bất kỳ giá trị nào (gomock.Any()) cho tham số context.Context
		// Giá trị cụ thể 1 (uint64(1)) cho tham số userID.
		// .Return(...): Khi được gọi, mock sẽ trả về một đối tượng User có ID 1 và Name "Nguyen Anh Tu", cùng với giá trị lỗi nil.
		// .Times(1): Phương thức GetUser chỉ được phép gọi đúng một lần trong suốt quá trình test.

	///Before mock , value will return (Return....)
	user, err := userDataAccessor.GetUser(context.Background(), 1)
	assert.Nil(t, err)
	assert.Equal(t, database.User{
		ID:   1,
		Name: "Nguyen Anh Tu",
	}, user)
}

func TestMain(m *testing.M) {
	var err error
	timeEnvVar := os.Getenv("TIME")
	if timeEnvVar == "" {
		currentTime = time.Now()
	} else {
		currentTime, err = time.Parse(time.RFC3339, timeEnvVar)
		if err != nil {
			return
		}
	}

	m.Run()
}
