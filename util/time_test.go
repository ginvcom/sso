package util

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeToString(t *testing.T) {
	now := time.Now().Local()
	fmt.Println(TimeToString(now, "YYYY"))
	fmt.Println(TimeToString(now, "YY-MM-DD"))
	fmt.Println(TimeToString(now, "YYYY-M-DD"))
	fmt.Println(TimeToString(now, "YYYY-MM-DD HH:mm"))
	fmt.Println(TimeToString(now, "YYYY-MM-DD hh:mm:ss a Z"))
	fmt.Println(TimeToString(now, "YYYY-MM-DD h:mm:ss A ZZ"))
	fmt.Println("-------------------------")
	time2 := now.Add(time.Hour * -12)
	fmt.Println(TimeToString(time2, "YYYY"))
	fmt.Println(TimeToString(time2, "YYYY-MM-DD"))
	fmt.Println(TimeToString(time2, "YYYY-MM-DD"))
	fmt.Println(TimeToString(time2, "YYYY-MM-DD HH:mm"))
	fmt.Println(TimeToString(time2, "YYYY-MM-DD hh:mm:ss a"))
	fmt.Println(TimeToString(time2, "YYYY-MM-DD h:mm:ss A"))
}
