package k8s_assist

import "strings"

func IsNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found")
}

func IsAlreadyExists(err error) bool {
	return strings.Contains(err.Error(), "already exists")
}
