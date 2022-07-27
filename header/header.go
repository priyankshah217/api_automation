package header

type Header map[string]string

func GetJSONHeader() Header {
	return Header{
		"content-type": "application/json",
	}
}
