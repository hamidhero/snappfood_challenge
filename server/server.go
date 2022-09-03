package server

import (
	"github.com/spf13/viper"
	"strings"
)

// Init initializes app server
func Init() {
	r := NewRouter()

	s := strings.Builder{}
	s.WriteString(viper.GetString("dev.app.address"))
	s.WriteString(viper.GetString("dev.app.port"))
	address := s.String()

	err := r.Run(address)
	if err != nil {
		panic(err)
	}
}
