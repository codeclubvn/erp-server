package service

import (
	"context"
	"erp/cmd/lib"
	"erp/config"
	"erp/domain"
	"erp/handler/dto/request"
	"erp/repository"
	"erp/utils"
	"erp/utils/api_errors"
	"erp/utils/constants"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type (
	FileService interface {
		SaveFile(ctx context.Context, req request.UploadFileRequest) (*domain.File, error)
		Update(ctx context.Context, req request.UpdateFileRequest) (*domain.File, error)
		Delete(ctx context.Context, req request.DeleteFileRequest) error
		GetOne(ctx context.Context, id string) (*domain.File, error)
		Download(ctx context.Context, id string) (*domain.File, error)
		UploadImage(ctx context.Context, file *multipart.FileHeader) (*uploader.UploadResult, error)
	}
	fileService struct {
		fileRepository repository.FileRepository
		config         *config.Config
		cloudinary     lib.CloudinaryRepository
	}
)

func NewFileService(itemRepo repository.FileRepository, config *config.Config, cloudinary lib.CloudinaryRepository) FileService {
	return &fileService{
		fileRepository: itemRepo,
		config:         config,
		cloudinary:     cloudinary,
	}
}

func createFolder(fileId string, config *config.Config) string {
	firstChar := fileId[0:1]
	secondChar := fileId[1:2]
	uploadPath := config.Server.UploadPath + "/" + firstChar + "/" + secondChar + "/"

	// create folder if not exists
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadPath, 0755); err != nil {
			panic(err)
		}
	}
	return uploadPath
}

func saveToFolder(file *multipart.FileHeader, uploadPath, id, extensionName string) error {
	// Source
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(uploadPath + id + "." + extensionName)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func getExtensionNameFromFilename(fileName string) string {
	fileNameArr := strings.Split(fileName, ".")
	extensionName := ""
	if len(fileNameArr) > constants.NumberFileNameSplit {
		extensionName = fileNameArr[1]
	}
	return extensionName
}
func getFilename(fileName string) string {
	return strings.Split(fileName, ".")[0]
}

func (s *fileService) SaveFile(ctx context.Context, req request.UploadFileRequest) (*domain.File, error) {
	var err error
	fileId := uuid.NewV4().String()
	extensionName := getExtensionNameFromFilename(req.File.Filename)

	uploadPath := createFolder(fileId, s.config)
	if err = saveToFolder(req.File, uploadPath, fileId, extensionName); err != nil {
		return nil, errors.WithStack(err)
	}

	file := &domain.File{
		BaseModel: domain.BaseModel{
			ID: uuid.FromStringOrNil(fileId),
		},
		Path:          uploadPath,
		Size:          req.File.Size,
		ExtensionName: extensionName,
		FileName:      req.File.Filename,
		UserId:        req.UserId,
	}
	if err = s.fileRepository.Create(ctx, file); err != nil {
		return nil, err
	}

	return file, err
}

func (s *fileService) Update(ctx context.Context, req request.UpdateFileRequest) (*domain.File, error) {
	// get one file
	file, err := s.fileRepository.GetOneById(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// check file is belong to user
	if file.UserId != req.UserId {
		return nil, errors.New(api_errors.ErrUnauthorizedAccess)
	}

	if req.File != nil {
		// check filePath is not in ./domain/assets
		if !strings.Contains(file.Path, s.config.Server.UploadPath) {
			file.Path = createFolder(file.ID.String(), s.config)
		} else {
			// delete old file
			_ = os.Remove(file.Path + file.ID.String() + "." + file.ExtensionName)
		}

		file.ExtensionName = getExtensionNameFromFilename(req.File.Filename)
		file.Size = req.File.Size
		file.UpdaterID = req.UserId

		// create folder if not exists
		if _, err := os.Stat(file.Path); os.IsNotExist(err) {
			if err := os.MkdirAll(file.Path, 0755); err != nil {
				panic(err)
			}
		}
		if err := saveToFolder(req.File, file.Path, file.ID.String(), file.ExtensionName); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	if err = utils.Copy(file, req); err != nil {
		return nil, err
	}
	if req.FileName != "" {
		fileName := getFilename(req.FileName)
		file.FileName = fileName + "." + file.ExtensionName
	}

	if err = s.fileRepository.Update(ctx, file); err != nil {
		return nil, err
	}
	return file, err
}

func (s *fileService) Delete(ctx context.Context, req request.DeleteFileRequest) error {
	// get one file
	file, err := s.fileRepository.GetOneById(ctx, req.ID)
	if err != nil {
		return err
	}

	// check file is belong to user
	if file.UserId != req.UserId {
		return errors.New(api_errors.ErrUnauthorizedAccess)
	}

	return s.fileRepository.DeleteById(ctx, req.ID)
}

func (s *fileService) GetOne(ctx context.Context, id string) (*domain.File, error) {
	return s.fileRepository.GetOneById(ctx, id)
}

func (s *fileService) Download(ctx context.Context, id string) (*domain.File, error) {
	return s.fileRepository.GetOneById(ctx, id)
}

func (s *fileService) UploadImage(ctx context.Context, file *multipart.FileHeader) (*uploader.UploadResult, error) {
	res, err := s.cloudinary.UploadFileCloud(ctx, file)
	if err != nil {
		return nil, err
	}

	return res, err
}
