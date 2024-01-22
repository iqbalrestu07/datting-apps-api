package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/iqbalrestu07/datting-apps-api/common"
	"github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// UploadHandler represents the handler for file uploads.
type UploadHandler struct {
	uploadUsecase domain.UploadUsecase
}

// NewUploadHandler will initialize the users/ resources endpoint
func NewUploadHandler(uu domain.UploadUsecase) *UploadHandler {
	return &UploadHandler{uploadUsecase: uu}
}

// Upload is the Echo handler function for file uploads.
func (h *UploadHandler) Upload(c echo.Context) error {

	// Input Tipe File
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": fmt.Sprintf("err: %s", err.Error())})
	}

	// Validate and sanitize file name
	validatedFileName, err := validateAndSanitizeFileName(file.Filename)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": fmt.Sprintf("invalid file name: %s", err.Error())})
	}

	// Use an absolute path
	absolutePath, err := filepath.Abs(filepath.Join("public", validatedFileName))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": fmt.Sprintf("failed to get absolute path: %s", err.Error())})
	}
	src, err := file.Open()
	if err != nil {
		return common.APIResponse(c, "failed to read file", err, absolutePath)
	}
	defer src.Close()

	// #nosec G304
	fl, err := os.Create(absolutePath)
	if err != nil {
		return common.APIResponse(c, "failed to read file", err, absolutePath)
	}
	defer fl.Close()

	if _, err = io.Copy(fl, src); err != nil {
		return common.APIResponse(c, "failed to store file", err, absolutePath)
	}
	userID, err := uuid.FromString(c.Get("user_id").(string))
	err = h.uploadUsecase.Create(c.Request().Context(), &domain.Photo{
		URL:    absolutePath,
		UserID: userID,
	})
	if err != nil {
		return common.APIResponse(c, "failed to store file", err, absolutePath)
	}

	// Response
	return common.APIResponse(c, "success", err, absolutePath)
}

func validateAndSanitizeFileName(fileName string) (string, error) {
	// Separate the extension
	ext := filepath.Ext(fileName)
	name := fileName[:len(fileName)-len(ext)]

	allowedPattern := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	var validatedName string
	for _, char := range name {
		if allowedPattern.MatchString(string(char)) {
			validatedName += string(char)
		}
	}

	if validatedName == "" {
		return "", fmt.Errorf("invalid characters in file name")
	}

	validatedFileName := validatedName + ext

	return validatedFileName, nil
}
