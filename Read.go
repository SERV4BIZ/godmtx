package godmtx

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/gfp/uuid"
)

// Read is decode data from bytes
func Read(buff []byte) (string, error) {
	pathtempdir := os.TempDir()
	pathname, errUUID := uuid.NewV4()
	if errUUID != nil {
		return "", errUUID
	}
	pathfile := fmt.Sprint(pathtempdir, "/", pathname, ".tmp")
	_, errWrite := files.WriteFile(pathfile, buff)
	if errWrite != nil {
		return "", errWrite
	}
	defer files.DeleteFile(pathfile)

	cmd := exec.Command("dmtxread", pathfile)
	stdout, errOutput := cmd.Output()
	if errOutput != nil {
		return "", errOutput
	}
	return string(stdout), nil
}
