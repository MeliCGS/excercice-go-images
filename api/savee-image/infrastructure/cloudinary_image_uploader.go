package image_infrastructure

import (
	"context"
	"io"
	"log"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func NewCloudinaryImageUploader() *CloudinaryImageUploader {
	instance, err := cloudinary.NewFromParams("dx9ihxd2q", "244977292583613", "w604923bi0sl4opBm2QkCrzKGJg")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &CloudinaryImageUploader{
		cloudinaryInstance: instance,
	}
}

type CloudinaryImageUploader struct {
	cloudinaryInstance *cloudinary.Cloudinary
}

func (c *CloudinaryImageUploader) UploadImage(image io.Reader) string {
	context := context.Background()

	uploadResult, err := c.cloudinaryInstance.Upload.Upload(
		context,
		image,
		uploader.UploadParams{PublicID: "logo"})

	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}

	return uploadResult.SecureURL
}
