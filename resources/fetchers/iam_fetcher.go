package fetchers

import (
	"context"
)

const IAMType = "aws-iam"

type IAMFetcher struct {
	iamProvider *IAMProvider
	cfg         IAMFetcherConfig
}

type IAMFetcherConfig struct {
	BaseFetcherConfig
	RoleName string `config:"roleName"`
}

type IAMResource struct {
	Data interface{}
}

func NewIAMFetcher(awsCfg AwsFetcherConfig, cfg IAMFetcherConfig) (Fetcher, error) {
	iam := NewIAMProvider(awsCfg.Config)

	return &IAMFetcher{
		cfg:         cfg,
		iamProvider: iam,
	}, nil
}

func (f IAMFetcher) Fetch(ctx context.Context) ([]FetchedResource, error) {
	results := make([]FetchedResource, 0)

	result, err := f.iamProvider.GetIAMRolePermissions(ctx, f.cfg.RoleName)
	results = append(results, IAMResource{result})

	return results, err
}

func (f IAMFetcher) Stop() {
}

//TODO: Add resource id logic to all AWS resources
func (r IAMResource) GetID() (string, error) {
	return "", nil
}

func (r IAMResource) GetData() interface{} {
	return r.Data
}
