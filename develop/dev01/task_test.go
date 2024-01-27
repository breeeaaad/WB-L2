package main

import (
	"errors"
	"testing"
	"time"
)

func Test_GetTime(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		if tm, _ := GetTime("0.beevik-ntp.pool.ntp.org"); time.Until(tm) > time.Second*5 {
			t.Error(errors.New("Локальное время не совпадает с хостом"))
		}
	})
}
