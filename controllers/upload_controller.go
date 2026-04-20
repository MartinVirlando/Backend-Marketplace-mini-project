package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"backend/utils"

	"github.com/labstack/echo/v4"
)

type UploadController struct{}

func NewUploadController() *UploadController {
	return &UploadController{}
}

func (ctrl *UploadController) UploadImage(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "No image provided")
	}

	// Validasi ekstensi
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	if !allowedExt[ext] {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid file type")
	}

	// Validasi ukuran (max 2MB)
	if file.Size > 2*1024*1024 {
		return utils.ErrorResponse(c, http.StatusBadRequest, "File too large (max 2MB)")
	}

	// Buat folder uploads/ jika belum ada
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create upload directory")
	}

	// Generate nama file unik
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dst := filepath.Join("uploads", filename)

	// Buka & simpan file
	src, err := file.Open()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to open file")
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to save file")
	}
	defer out.Close()

	if _, err = out.ReadFrom(src); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to write file")
	}

	return utils.SuccessResponse(c, http.StatusOK, "Image uploaded", map[string]string{
		"path": dst,
	})
}
