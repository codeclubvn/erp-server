package lib

import (
	"context"
	"erp/cmd/infrastructure"
	"erp/config"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go.uber.org/zap"
	"mime/multipart"
)

type CloudinaryRepository interface {
	UploadFileCloud(ctx context.Context, file *multipart.FileHeader) (resp *uploader.UploadResult, err error)
	GetAssetInfo(ctx context.Context)
	TransformImage(ctx context.Context)
}

type cloudinaryRepository struct {
	cld    *infrastructure.Cloudinary
	config *config.Config
	logger *zap.Logger
}

func NewCloudinaryRepository(cld *infrastructure.Cloudinary, logger *zap.Logger, config *config.Config) CloudinaryRepository {
	return &cloudinaryRepository{
		cld:    cld,
		config: config,
		logger: logger,
	}
}

func (r *cloudinaryRepository) UploadFileCloud(ctx context.Context, file *multipart.FileHeader) (resp *uploader.UploadResult, err error) {

	// SaveFile the cloudinary.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err = r.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:       r.config.Cloudinary.PublicId,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	if err != nil {
		fmt.Println("error")
	}

	return resp, err
}

func (r *cloudinaryRepository) GetAssetInfo(ctx context.Context) {
	// Get and use details of the cloudinary
	// ==============================
	resp, err := r.cld.Admin.Asset(ctx, admin.AssetParams{PublicID: r.config.Cloudinary.PublicId})
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("****3. Get and use details of the cloudinary****\nDetailed response:\n", resp, "\n")

	// Assign tags to the uploaded cloudinary based on its width. Save the response to the update in the variable 'update_resp'.
	if resp.Width > 900 {
		updateResp, err := r.cld.Admin.UpdateAsset(ctx, admin.UpdateAssetParams{
			PublicID: r.config.Cloudinary.PublicId,
			Tags:     []string{"large"}})
		if err != nil {
			fmt.Println("error")
		} else {
			// Log the new tag to the console.
			fmt.Println("New tag: ", updateResp.Tags, "\n")
		}
	} else {
		updateResp, err := r.cld.Admin.UpdateAsset(ctx, admin.UpdateAssetParams{
			PublicID: r.config.Cloudinary.PublicId,
			Tags:     []string{"small"}})
		if err != nil {
			fmt.Println("error")
		} else {
			// Log the new tag to the console.
			fmt.Println("New tag: ", updateResp.Tags, "\n")
		}
	}

}

func (r *cloudinaryRepository) TransformImage(ctx context.Context) {
	// Instantiate an object for the asset with public ID "my_image"
	qs_img, err := r.cld.Image("quickstart_butterfly")
	if err != nil {
		fmt.Println("error")
	}

	// Add the transformation
	qs_img.Transformation = "r_max/e_sepia"

	// Generate and log the delivery URL
	new_url, err := qs_img.String()
	if err != nil {
		fmt.Println("error")
	} else {
		print("****4. Transform the image****\nTransfrmation URL: ", new_url, "\n")
	}
}
