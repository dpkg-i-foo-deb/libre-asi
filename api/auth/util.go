package auth

func IsAdmin(tk string) bool {

	role, err := RoleFromToken(tk)

	if err != nil {
		return false
	}

	if role != "admin" {
		return false
	}

	return true
}
