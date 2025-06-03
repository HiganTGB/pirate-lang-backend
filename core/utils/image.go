package utils

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"pirate-lang-go/core/config"
	"pirate-lang-go/core/errors"
	"pirate-lang-go/core/logger"
)

func ResizeImage(file io.Reader) ([]byte, error) {

	// Đọc và giải mã ảnh
	img, format, err := image.Decode(file)
	if err != nil {
		logger.Error("AccountService:ResizeImage:Failed to decode image", "error", err)
		return nil, err
	}

	resizedImg := resize.Resize(config.Get().AvatarSize.Width, config.Get().AvatarSize.Height, img, resize.Bilinear)

	buf := new(bytes.Buffer)

	switch format {
	case "jpeg":

		err = jpeg.Encode(buf, resizedImg, &jpeg.Options{Quality: 75})
	case "png":

		err = png.Encode(buf, resizedImg)
	default:

		logger.Warn("AccountService:ResizeImage:Unsupported image format for encoding, returning error", "format", format)
		return nil, errors.NewAppError(errors.ErrInvalidInput, "AccountService:ResizeImage:Unsupported image format for encoding (only JPEG and PNG are supported)", nil)
	}

	if err != nil {
		logger.Error("AccountService:ResizeImage:Failed to encode image", "error", err)
		return nil, err
	}

	return buf.Bytes(), nil
}
