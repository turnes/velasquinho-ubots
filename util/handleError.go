package util

import "log"

func handleError(err error) bool{
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}
