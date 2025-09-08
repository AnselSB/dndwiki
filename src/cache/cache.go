package cache

import (
	"fmt"
	"os"
)

// the following return objects from the cache, throw an error if the file in the cache doesn't exist based on the provided path
func GetObject(path string) ([]byte, error) {
	// check to see if the path exists
	filePath := fmt.Sprintf("public/cache/%v.json", path)
	if _, err := os.Stat(filePath); err != nil {
		// could be a schrodinger case here but still treat it like it doesnt exist
		return nil, err
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// creates a new file within the cache and writes the contents of the object into the file
func SetObject(filePath string, contents []byte) error {

	fullPath := fmt.Sprintf("public/cache/%v.json", filePath)
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	writeErr := os.WriteFile(fullPath, contents, 0666)
	if writeErr != nil {
		os.Remove(fullPath)
		return err
	}
	return nil
}

func SetDirectory(dirName string) string {
	// first attempt to create the path based on the passed directory name
	fullPath := fmt.Sprintf("public/cache/%v", dirName)
	err := os.Mkdir(fullPath, 0755)
	if err != nil {
		return "Error occurred making new cache. Check directory name inputted"
	}
	// make a .gitkeep file so the directory shows in git
	file, err := os.Create(fullPath + "/.gitkeep")
	if err != nil {
		return "unable to keep directory in git"
	}
	defer file.Close()
	return "Successfully made new cache entry!"
}
