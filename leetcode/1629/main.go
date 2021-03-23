package main

func main() {

}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	returnChar := keysPressed[0]
	max := releaseTimes[0]

	for i := 1; i < len(releaseTimes); i++ {

		temp := releaseTimes[i] - releaseTimes[i-1]
		if temp == max {
			if keysPressed[i] > returnChar {
				returnChar = keysPressed[i]
			}
		} else if temp >= max {
			max = temp
			returnChar = keysPressed[i]
		}
	}

	return byte(returnChar)
}
