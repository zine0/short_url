package utils

func TrimUrl(rawUrl string) (trimUrl string) {
	if rawUrl[:7] == "http://" {
		trimUrl = rawUrl
	} else if rawUrl[:8] == "https://" {
		trimUrl = rawUrl
	} else {
		trimUrl = "https://" + rawUrl
	}

	return

}
