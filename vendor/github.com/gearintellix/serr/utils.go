package serr

// StandardFormat function
func StandardFormat() string {
	return "In %s[%s:%d] %s.%s"
}

// StandardColorFormat function
func StandardColorFormat() string {
	frmt := ""
	frmt += applyForeColor("In", colorDarkGray) + " "
	frmt += applyForeColor("%s", colorLightYellow)
	frmt += applyForeColor("[", colorDarkGray)
	frmt += applyForeColor("%s:%d", colorMagenta)
	frmt += applyForeColor("]", colorDarkGray)
	frmt += " %s.%s"
	return frmt
}
