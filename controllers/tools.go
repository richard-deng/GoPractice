package controllers

import (
	"fmt"
	"strings"
)

func BuildWhere(whereMap map[string]string) string {
	var whereArr []string
	for key := range whereMap {
		if whereMap[key] != "" {
			whereArr = append(whereArr, fmt.Sprintf("%s=\"%s\"", key, whereMap[key]))
		}
	}
	whereStr := strings.Join(whereArr, " and ")
	return whereStr
}
