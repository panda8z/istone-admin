package initialize

import (
	"fmt"
	"testing"
)

func TestViper(t *testing.T) {
	v := Viper()
	list := v.AllKeys()
	fmt.Println(list)
}
