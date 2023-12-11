package runner

func calculateChecksum(arr []byte) byte {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result ^= arr[i]
	}
	return result
}

func verifyChecksum(arr []byte) bool {
	length := len(arr)

	if length < 2 {
		return false // Array is too short to verify the checksum
	}

	lastTwoBytes := arr[length-2:]
	calculatedChecksum := calculateChecksum(arr[0 : length-2])

	return calculatedChecksum == lastTwoBytes[0]
}
