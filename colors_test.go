package colors

import (
	"fmt"
	"testing"
	"time"
)

func TestColors(t *testing.T) {
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

	fmt.Println(MagentaSprint("Magenta_Sprint"))
	fmt.Println(MagentaSprintf("Magenta_Sprintf: %s", time.Now().Format("2006-01-02 15:04:05.000")))

	fmt.Println(CyanSprint("Cyan_Sprint"))
	fmt.Println(CyanSprintf("Cyan_Sprintf: %s", time.Now().Format("2006-01-02 15:04:05.000")))

	fmt.Println(WhiteSprint("White_Sprint"))
	fmt.Println(WhiteSprintf("White_Sprintf: %s", time.Now().Format("2006-01-02 15:04:05.000")))

}
