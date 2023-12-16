package utils

func WriteJsonToFile(filename string, v any, indent bool) error {
	var jsonString string
	var err error
	if jsonString, err = JsonStringify(v, indent); err != nil {
		return err
	}
	return WriteStringToFile(filename, jsonString)
}
