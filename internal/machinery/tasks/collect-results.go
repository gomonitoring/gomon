package tasks

func CollectResults(results ...[]byte) error {
	for _, data := range results {
		result := CallUrlResult{}
		decodeCallResult(data, result)
		// TODO notify user if neccessary and create poll object
	}
	// TODO bulk insert poll objects
	return nil
}
