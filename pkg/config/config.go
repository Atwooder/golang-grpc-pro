package config

import (
	"github.com/spf13/viper"
)

type Setting struct {
	Vp *viper.Viper
}

func NewViper() (*Setting, error) {
	vp := viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath("configs/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}
	return &Setting{vp}, err
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.Vp.Unmarshal(v)
	if err != nil {
		return err
	}
	return nil
}
