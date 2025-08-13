package clients

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/serviceerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/config"
	"github.com/OpsOMI/S.L.A.M/internal/server/config/core/models"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/clients"
	"gopkg.in/yaml.v3"
)

func (s *service) GetByID(
	ctx context.Context,
	id string,
) (*clients.Client, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return nil, err
	}

	domainModel, err := s.repositories.Clients().GetByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return domainModel, nil
}

func (s *service) GetByClientKey(
	ctx context.Context,
	clientKey string,
) (*clients.Client, error) {
	key, err := s.utils.Parse().ParseRequiredUUID(clientKey)
	if err != nil {
		return nil, err
	}

	domainModel, err := s.repositories.Clients().GetByClientKey(ctx, key)
	if err != nil {
		return nil, err
	}

	return domainModel, nil
}

func (s *service) GetByUserID(
	ctx context.Context,
	userID string,
) (*clients.Clients, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(userID)
	if err != nil {
		return nil, err
	}

	return s.repositories.Clients().GetByUserID(ctx, uid)
}

func (s *service) RevokeByID(
	ctx context.Context,
	id string,
) error {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return err
	}

	return s.repositories.Clients().RevokeByID(ctx, uid)
}

func (s *service) DeleteByID(
	ctx context.Context,
	id string,
) error {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return err
	}

	return s.repositories.Clients().DeleteByID(ctx, uid)
}

func (s *service) IsExistByID(
	ctx context.Context,
	id string,
) (*bool, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return nil, err
	}

	return s.repositories.Clients().IsExistByID(ctx, uid)
}

func (s *service) IsExistByClientKey(
	ctx context.Context,
	clientKey string,
) (*bool, error) {
	key, err := s.utils.Parse().ParseRequiredUUID(clientKey)
	if err != nil {
		return nil, err
	}

	return s.repositories.Clients().IsExistByClientKey(ctx, key)
}

func (s *service) IsRevoked(
	ctx context.Context,
	id string,
) (*bool, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return nil, err
	}

	return s.repositories.Clients().IsRevoked(ctx, uid)
}

func (s *service) CreateClient(
	serverConfig *config.Configs,
	clientKey, nickname string,
) error {
	if err := s.CreateClientConfig(serverConfig, clientKey); err != nil {
		return serviceerrors.BadRequest(clients.ErrConfigCreateFailed)
	}
	if err := s.CopyServerCert(serverConfig.Server.Core.TSLCertPath); err != nil {
		return serviceerrors.BadRequest(clients.ErrTslCertCopyFailed)
	}
	if err := s.BuildClientExe(nickname); err != nil {
		return serviceerrors.BadRequest(clients.ErrBuildClientFailed)
	}
	if err := s.DeleteEmbeddedFiles(); err != nil {
		return serviceerrors.BadRequest(clients.ErrDeleteEmbededFilesFailed)
	}

	return nil
}

func (s *service) CreateClientConfig(
	serverConfig *config.Configs,
	clientID string,
) error {
	var host string
	if serverConfig.Server.Core.ExternalHost != "" {
		host = serverConfig.Server.Core.ExternalHost
	} else {
		host = serverConfig.Server.Core.Host
	}

	cfg := models.ClientConfig{
		ClientKey:      clientID,
		TSLServerName:  serverConfig.Env.TSL.ServerName,
		ServerHost:     host,
		ServerPort:     serverConfig.Server.Core.Port,
		TimeoutSeconds: 1,
		ReconnectRetry: 3,
	}

	dir := s.getClientDirPath()
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal client config: %w", err)
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filePath := filepath.Join(dir, "embeded.yaml")
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write client.yaml: %w", err)
	}

	return nil
}

func (s *service) CopyServerCert(src string) error {
	destDir := s.getClientDirPath()
	dest := filepath.Join(destDir, "embeded.crt")

	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source cert file: %w", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create destination cert file: %w", err)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return fmt.Errorf("failed to copy cert file: %w", err)
	}

	return nil
}

func (s *service) DeleteEmbeddedFiles() error {
	destDir := s.getClientDirPath()

	filesToDelete := []string{
		filepath.Join(destDir, "embeded.crt"),
		filepath.Join(destDir, "embeded.yaml"),
	}

	for _, filePath := range filesToDelete {
		err := os.Remove(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return fmt.Errorf("failed to delete file %s: %w", filePath, err)
		}
	}

	return nil
}

func (s *service) BuildClientExe(nickname string) error {
	platforms := []struct {
		GOOS   string
		GOARCH string
	}{
		{"darwin", "amd64"},
		{"linux", "amd64"},
		{"windows", "amd64"},
	}

	for _, p := range platforms {
		dirPath := filepath.Join("clients", nickname, fmt.Sprintf("%s_%s", p.GOOS, p.GOARCH))

		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
		}

		outputName := "main"
		if p.GOOS == "windows" {
			outputName = "main.exe"
		}

		outputPath := filepath.Join(dirPath, outputName)

		cmd := exec.Command(
			"go",
			"build",
			"-tags=embed",
			"-ldflags=-X main.useEmbed=true",
			"-o", outputPath,
			"./cmd/client/main.go",
		)

		cmd.Env = append(os.Environ(),
			"GOOS="+p.GOOS,
			"GOARCH="+p.GOARCH,
		)

		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("build failed for %s/%s: %w\nOutput: %s", p.GOOS, p.GOARCH, err, string(output))
		}

		fmt.Printf("Built client for %s/%s: %s\n", p.GOOS, p.GOARCH, outputPath)
	}

	return nil
}

func (s *service) getClientDirPath() string {
	root, err := filepath.Abs(".")
	if err != nil {

	}

	configPath := filepath.Join(root, "internal", "client", "config")

	return configPath
}
