package cli

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func yaml2args[T any](file string) (args *T, err error) {
	// 读取YAML文件
	config, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read yaml file")
	}

	// 解析YAML文件
	args = new(T)
	err = yaml.Unmarshal(config, args)

	if err != nil {
		return nil, err
	}
	return args, nil
}