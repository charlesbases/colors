package colors

// BlackSprint .
func BlackSprint(a ...interface{}) string {
	return New(F_Black).Sprint(a...)
}

// BlackSprintf .
func BlackSprintf(format string, a ...interface{}) string {
	return New(F_Black).Sprintf(format, a...)
}

// RedSprint .
func RedSprint(a ...interface{}) string {
	return New(F_Red).Sprint(a...)
}

// RedSprintf .
func RedSprintf(format string, a ...interface{}) string {
	return New(F_Red).Sprintf(format, a...)
}

// GreenSprint .
func GreenSprint(a ...interface{}) string {
	return New(F_Green).Sprint(a...)
}

// GreenSprintf .
func GreenSprintf(format string, a ...interface{}) string {
	return New(F_Green).Sprintf(format, a...)
}

// YellowSprint .
func YellowSprint(a ...interface{}) string {
	return New(F_Yellow).Sprint(a...)
}

// YellowSprintf .
func YellowSprintf(format string, a ...interface{}) string {
	return New(F_Yellow).Sprintf(format, a...)
}

// BlueSprint .
func BlueSprint(a ...interface{}) string {
	return New(F_Blue).Sprint(a...)
}

// BlueSprintf .
func BlueSprintf(format string, a ...interface{}) string {
	return New(F_Blue).Sprintf(format, a...)
}

// PurpleSprint .
func PurpleSprint(a ...interface{}) string {
	return New(F_Purple).Sprint(a...)
}

// PurpleSprintf .
func PurpleSprintf(format string, a ...interface{}) string {
	return New(F_Purple).Sprintf(format, a...)
}

// CyanSprint .
func CyanSprint(a ...interface{}) string {
	return New(F_Cyan).Sprint(a...)
}

// CyanSprintf .
func CyanSprintf(format string, a ...interface{}) string {
	return New(F_Cyan).Sprintf(format, a...)
}

// WhiteSprint .
func WhiteSprint(a ...interface{}) string {
	return New(F_White).Sprint(a...)
}

// WhiteSprintf .
func WhiteSprintf(format string, a ...interface{}) string {
	return New(F_White).Sprintf(format, a...)
}
