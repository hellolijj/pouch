package opts

import (
	"fmt"
	"strings"
)

// ParseDiskQuota parses diskquota configurations of container.
func ParseDiskQuota(quotas []string) (map[string]string, error) {
	var quotaMaps = make(map[string]string)

	if quotas == nil {
		return nil, fmt.Errorf("invalid format for disk quota: %s", quotas)
	}

	for _, quota := range quotas {
		if strings.TrimSpace(quota) == "" {
			return nil, fmt.Errorf("invalid format for disk quota: %s", quota)
		}

		parts := strings.Split(quota, "=")
		switch len(parts) {
		case 1:
			quotaMaps["/"] = strings.TrimSpace(parts[0])
		case 2:
			quotaMaps[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		default:
			return nil, fmt.Errorf("invalid format for disk quota: %s", quota)
		}
	}

	return quotaMaps, nil
}

// ValidateDiskQuota verifies diskquota configurations of container.
func ValidateDiskQuota(quotaMaps map[string]string) error {
	// TODO
	return nil
}
