package utils

import "github.com/jinzhu/copier"

func Copy(to interface{}, from interface{}) error {
	err := copier.Copy(to, from)
	return err
}

func CopyWithOption(to interface{}, from interface{}, opt copier.Option) error {
	err := copier.CopyWithOption(to, from, opt)
	return err
}
