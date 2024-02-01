package colors

import (
	"fmt"
	"testing"
	"time"
)

// print .
func print() () {
	fmt.Println(BlackSprint("Black_Sprint"))
	fmt.Println(BlackSprintf("Black_Sprintf: %s", time.Now().Format("2006-01-02 15:04:05.000")))

	fmt.Println(RedSprint("Red_Sprint"))
	fmt.Println(RedSprintf("Red_Sprintf: %s", time.Now().Format("2006-01-02 15:04:05.000")))

	fmt.Println(GreenSprint("Green_Sprint"))
	fmt.Println(GreenSprintf("Green_Sprintf: %s", time.Now().Format("2006-01-02 15:04:05.000")))

	fmt.Println(YellowSprint("Yellow_Sprint"))
	fmt.Println(YellowSprintf("Yellow_Sprintf: %s", time.Now().Format("2006-01-02 15:04:05.000")))

	fmt.Println(BlueSprint("Blue_Sprint"))
	fmt.Println(BlueSprintf("Blue_Sprintf: %s", time.Now().Format("2006-01-02 15:04:05.000")))

	fmt.Println(PurpleSprint("Purple_Sprint"))
	fmt.Println(PurpleSprintf("Purple_Sprintf: %s", time.Now().Format("2006-01-02 15:04:05.000")))

	fmt.Println(CyanSprint("Cyan_Sprint"))
	fmt.Println(CyanSprintf("Cyan_Sprintf: %s", time.Now().Format("2006-01-02 15:04:05.000")))

	fmt.Println(WhiteSprint("White_Sprint"))
	fmt.Println(WhiteSprintf("White_Sprintf: %s", time.Now().Format("2006-01-02 15:04:05.000")))
}

func TestColors(t *testing.T) {
	print()
}

//    1879            649608 ns/op            2276 B/op         64 allocs/op
func BenchmarkColors(b *testing.B) {
	var bench = func(f func()) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			f()
		}
		b.StopTimer()
	}

	bench(print)
}
