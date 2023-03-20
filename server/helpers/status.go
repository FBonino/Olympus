package helpers

func TransformStatus(status uint8) string {
	enum := map[uint8]string{
		0: "Offline",
		1: "Do Not Disturb",
		2: "Idle",
		3: "Online",
	}

	return enum[status]
}
