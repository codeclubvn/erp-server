package utils

// RoleID to role_id
func StructPascalToSnakeCase(str string) string {
	var result string
	IDs := str[len(str)-3:]
	ID := str[len(str)-2:]

	if IDs == "IDs" {
		str = str[:len(str)-3] + "Ids"
	}
	if ID == "ID" {
		str = str[:len(str)-2] + "Id"
	}

	for i, v := range str {
		if v >= 'A' && v <= 'Z' {
			if i != 0 {
				result += "_"
			}
			result += string(v + 32)
		} else {
			result += string(v)
		}
	}
	return result
}
