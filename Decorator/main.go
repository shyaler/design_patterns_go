package main

import (
	writer "main/Writer"
	"time"
)

func main() {
	WriteFoldeWithProperties()

	WriteFileWithMetadata()

	WriteFolderWithMetadataAndProperties()
}

func WriteFoldeWithProperties() {
	folderWriter := &writer.FolderWriter{FolderPath: "https://folderPath.com"}
	folderWithProperties := writer.NewPropertiesWriter(folderWriter, time.Now(), time.Now())

	folderWithProperties.Write("folder with properties")
}

func WriteFileWithMetadata() {
	fileWriter := &writer.FileWriter{FilePath: "https://filePath.com"}

	metadata := make(map[string]string, 3)
	metadata["Description"] = "File description"
	metadata["Author"] = "shyaler"
	metadata["AdditionalInfo"] = "Additional information about the file"

	filesWithMetadata := writer.NewMetadataWriter(fileWriter, metadata)

	filesWithMetadata.Write("file with metadata")
}

func WriteFolderWithMetadataAndProperties() {
	folderWriter := &writer.FolderWriter{FolderPath: "https://folderPath.com"}
	folderWithProperties := writer.NewPropertiesWriter(folderWriter, time.Now(), time.Now())
	metadata := make(map[string]string, 1)
	metadata["Description"] = "Folder description"

	folderWithMetadata := writer.NewMetadataWriter(folderWithProperties, metadata)
	folderWithMetadata.Write("folder with metadata and properties")
}
