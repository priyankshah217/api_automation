package header

type Header map[string]string

func GetJsonHeader() Header {
	return Header{
		"content-type": "application/json",
	}
}
