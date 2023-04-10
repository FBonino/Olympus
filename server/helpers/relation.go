package helpers

func TransformRelation(status uint8) string {
	enum := map[uint8]string{
		0: "Blocked",
		1: "Pending",
		2: "Friend",
	}

	return enum[status]
}
