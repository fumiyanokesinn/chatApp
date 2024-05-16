package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GetEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file:\n%s", err))
	}
}

func GetTestEnv() {
	// 現在の作業ディレクトリを取得
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Error getting current working directory:\n%s", err))
	}

	// プロジェクトのルートディレクトリを取得
	rootPath := findProjectRoot(wd)
	if rootPath == "" {
		panic("Error: could not find project root directory")
	}

	// .env.testファイルの絶対パスを取得
	envPath := filepath.Join(rootPath, ".env.test")

	// .env.testファイルをロード
	err = godotenv.Load(envPath)
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file at %s:\n%s", envPath, err))
	}
}

// プロジェクトのルートディレクトリを見つける関数
func findProjectRoot(dir string) string {
	for {
		if _, err := os.Stat(filepath.Join(dir, ".env.test")); err == nil {
			return dir
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}
	return ""
}
