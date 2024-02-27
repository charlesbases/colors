package colors

// BlackSprint .
func BlackSprint(a ...interface{}) string {
	return New(ForegroundBlack).Sprint(a...)
}

// BlackSprintf .
func BlackSprintf(format string, a ...interface{}) string {
	return New(ForegroundBlack).Sprintf(format, a...)
}

// RedSprint .
func RedSprint(a ...interface{}) string {
	return New(ForegroundRed).Sprint(a...)
}

// RedSprintf .
func RedSprintf(format string, a ...interface{}) string {
	return New(ForegroundRed).Sprintf(format, a...)
}

// GreenSprint .
func GreenSprint(a ...interface{}) string {
	return New(ForegroundGreen).Sprint(a...)
}

// GreenSprintf .
func GreenSprintf(format string, a ...interface{}) string {
	return New(ForegroundGreen).Sprintf(format, a...)
}

// YellowSprint .
func YellowSprint(a ...interface{}) string {
	return New(ForegroundYellow).Sprint(a...)
}

// YellowSprintf .
func YellowSprintf(format string, a ...interface{}) string {
	return New(ForegroundYellow).Sprintf(format, a...)
}

// BlueSprint .
func BlueSprint(a ...interface{}) string {
	return New(ForegroundBlue).Sprint(a...)
}

// BlueSprintf .
func BlueSprintf(format string, a ...interface{}) string {
	return New(ForegroundBlue).Sprintf(format, a...)
}

// PurpleSprint .
func PurpleSprint(a ...interface{}) string {
	return New(ForegroundPurple).Sprint(a...)
}

// PurpleSprintf .
func PurpleSprintf(format string, a ...interface{}) string {
	return New(ForegroundPurple).Sprintf(format, a...)
}

// CyanSprint .
func CyanSprint(a ...interface{}) string {
	return New(ForegroundCyan).Sprint(a...)
}

// CyanSprintf .
func CyanSprintf(format string, a ...interface{}) string {
	return New(ForegroundCyan).Sprintf(format, a...)
}

// WhiteSprint .
func WhiteSprint(a ...interface{}) string {
	return New(ForegroundWhite).Sprint(a...)
}

// WhiteSprintf .
func WhiteSprintf(format string, a ...interface{}) string {
	return New(ForegroundWhite).Sprintf(format, a...)
}
