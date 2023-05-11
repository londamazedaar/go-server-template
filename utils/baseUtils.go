package utils

import (
	"vas/errorCodes"
	"vas/logger"

	"github.com/gofiber/fiber/v2"
)

type IError struct {
	Message      string `json:"message"`
	Code         int    `json:"code"`
	DebugMessage string `json:"appMessage"`
}

type AppError struct {
	DebugMessage string                      `json:"appMessage"`
	ErrorCode    errorCodes.ErrorCodeMessage `json:"errorCode"`
}

func ErrorResponse(c *fiber.Ctx, appError AppError, reqBody interface{}) error {
	errResp := IError{
		Message:      appError.ErrorCode.Message,
		Code:         appError.ErrorCode.Code,
		DebugMessage: appError.DebugMessage,
	}
	logger.LogError(c, "RESPONSE: "+appError.ErrorCode.Message, logger.LogReqResp{
		RequestBody:  reqBody,
		ResponseData: errResp,
	})
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success":   false,
		"requestId": c.Locals("requestId"),
		"data":      nil,
		"errors":    errResp,
	})
}

func SuccessResponse(c *fiber.Ctx, respData interface{}, reqBody interface{}) error {
	logger.LogInfo(c, "RESPONSE", logger.LogReqResp{
		RequestBody:  reqBody,
		ResponseData: respData,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":   true,
		"requestId": c.Locals("requestId"),
		"data":      respData,
		"errors":    nil,
	})
}
