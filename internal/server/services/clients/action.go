package clients

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

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

func (s *service) CreateClientConfig(
	serverConfig *config.Configs,
	clientID string,
) error {
	cfg := models.ClientConfig{
		ClientID:       clientID,
		ServerName:     serverConfig.Env.Jwt.Issuer,
		ServerHost:     serverConfig.Server.Core.Host,
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

func (s *service) BuildClientExe(clientID string) error {
	dirPath := filepath.Join("clients", clientID)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	outputPath := filepath.Join(dirPath, "main")

	cmd := exec.Command(
		"go",
		"build",
		"-tags=embed",
		"-ldflags=-X main.useEmbed=true",
		"-o", outputPath,
		"./cmd/client/main.go",
	)
	cmd.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=amd64") // ya da kendi platformuna göre ayarla

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("build failed: %w\nOutput: %s", err, string(output))
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
