// Copyright 2018 The gooid Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package camera

import (
	"github.com/gooid/gooid/internal/ndk"
)

const (
	FACING_FRONT = 1 // 前摄
	FACING_BACK  = 0 // 后摄

	FLASH_MODE_AUTO    = app.CAMERA_FLASH_MODE_AUTO
	FLASH_MODE_OFF     = app.CAMERA_FLASH_MODE_OFF
	FLASH_MODE_ON      = app.CAMERA_FLASH_MODE_ON
	FLASH_MODE_RED_EYE = app.CAMERA_FLASH_MODE_RED_EYE
	FLASH_MODE_TORCH   = app.CAMERA_FLASH_MODE_TORCH
	FLASH_MODES_NUM    = app.CAMERA_FLASH_MODES_NUM

	FOCUS_MODE_AUTO               = app.CAMERA_FOCUS_MODE_AUTO
	FOCUS_MODE_CONTINUOUS_VIDEO   = app.CAMERA_FOCUS_MODE_CONTINUOUS_VIDEO
	FOCUS_MODE_EDOF               = app.CAMERA_FOCUS_MODE_EDOF
	FOCUS_MODE_FIXED              = app.CAMERA_FOCUS_MODE_FIXED
	FOCUS_MODE_INFINITY           = app.CAMERA_FOCUS_MODE_INFINITY
	FOCUS_MODE_MACRO              = app.CAMERA_FOCUS_MODE_MACRO
	FOCUS_MODE_CONTINUOUS_PICTURE = app.CAMERA_FOCUS_MODE_CONTINUOUS_PICTURE
	FOCUS_MODES_NUM               = app.CAMERA_FOCUS_MODES_NUM

	WHITE_BALANCE_AUTO             = app.CAMERA_WHITE_BALANCE_AUTO
	WHITE_BALANCE_CLOUDY_DAYLIGHT  = app.CAMERA_WHITE_BALANCE_CLOUDY_DAYLIGHT
	WHITE_BALANCE_DAYLIGHT         = app.CAMERA_WHITE_BALANCE_DAYLIGHT
	WHITE_BALANCE_FLUORESCENT      = app.CAMERA_WHITE_BALANCE_FLUORESCENT
	WHITE_BALANCE_INCANDESCENT     = app.CAMERA_WHITE_BALANCE_INCANDESCENT
	WHITE_BALANCE_SHADE            = app.CAMERA_WHITE_BALANCE_SHADE
	WHITE_BALANCE_TWILIGHT         = app.CAMERA_WHITE_BALANCE_TWILIGHT
	WHITE_BALANCE_WARM_FLUORESCENT = app.CAMERA_WHITE_BALANCE_WARM_FLUORESCENT
	WHITE_BALANCE_MODES_NUM        = app.CAMERA_WHITE_BALANCE_MODES_NUM

	ANTIBANDING_50HZ      = app.CAMERA_ANTIBANDING_50HZ
	ANTIBANDING_60HZ      = app.CAMERA_ANTIBANDING_60HZ
	ANTIBANDING_AUTO      = app.CAMERA_ANTIBANDING_AUTO
	ANTIBANDING_OFF       = app.CAMERA_ANTIBANDING_OFF
	ANTIBANDING_MODES_NUM = app.CAMERA_ANTIBANDING_MODES_NUM

	FOCUS_DISTANCE_NEAR_INDEX    = app.CAMERA_FOCUS_DISTANCE_NEAR_INDEX
	FOCUS_DISTANCE_OPTIMAL_INDEX = app.CAMERA_FOCUS_DISTANCE_OPTIMAL_INDEX
	FOCUS_DISTANCE_FAR_INDEX     = app.CAMERA_FOCUS_DISTANCE_FAR_INDEX
)

type Camera = app.Camera
type CameraCallback = app.CameraCallback

func Connect(cameraId int, cb CameraCallback) Camera {
	return app.CameraConnect(cameraId, cb)
}
