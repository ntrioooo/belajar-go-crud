package resp

import "github.com/gin-gonic/gin"

type envelope map[string]any

func OK(c *gin.Context, data any)          { c.JSON(200, envelope{"data": data}) }
func Created(c *gin.Context, data any)     { c.JSON(201, envelope{"data": data}) }
func NoContent(c *gin.Context)             { c.Status(204) }
func BadRequest(c *gin.Context, msg any)   { c.JSON(400, envelope{"error": msg}) }
func Unauthorized(c *gin.Context, msg any) { c.JSON(401, envelope{"error": msg}) }
func NotFound(c *gin.Context, msg any)     { c.JSON(404, envelope{"error": msg}) }
func Internal(c *gin.Context, msg any)     { c.JSON(500, envelope{"error": msg}) }
