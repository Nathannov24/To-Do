package responses

func StatusDuplicated() map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "failed",
		"message": "Activity was inputed, try input another Activity",
	}
	return result
}

func StatusFailedInput() map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "failed",
		"message": "Failed to Post",
	}
	return result
}

func StatusSuccessInput(message string) map[string]interface{} {
	var result = map[string]interface{}{
		"status":   "success",
		"message":  "Success to add activity",
		"activity": message,
	}
	return result
}
