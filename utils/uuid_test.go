package utils

import "testing"

func Test_UUID(t *testing.T) {
	t.Log(TimeUUID().String())
}
