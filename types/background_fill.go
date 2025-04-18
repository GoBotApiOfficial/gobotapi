// Code AutoGenerated; DO NOT EDIT.

package types

// BackgroundFill This object describes the way a background is filled based on the selected colors
// Currently, it can be one of
//   - BackgroundFillSolid
//   - BackgroundFillGradient
//   - BackgroundFillFreeformGradient
type BackgroundFill struct {
	BottomColor   int    `json:"bottom_color"`
	Color         int    `json:"color"`
	Colors        []int  `json:"colors"`
	RotationAngle int    `json:"rotation_angle"`
	TopColor      int    `json:"top_color"`
	Type          string `json:"type"`
}

func (x BackgroundFill) Kind() int {
	switch x.Type {
	case "solid":
		return TypeBackgroundFillSolid
	case "gradient":
		return TypeBackgroundFillGradient
	case "freeform_gradient":
		return TypeBackgroundFillFreeformGradient
	default:
		return -1
	}
}
