package commands

import (
	"fmt"
	"os/exec"
	"TWAMP/internal/models"
)

func StartServer(config models.PacketConfig) error {
	cmd := exec.Command("./twamp/twamp_test",
		fmt.Sprintf("%s:%d", config.IP, config.Port),
		fmt.Sprintf("--count=%d", config.Count),
		fmt.Sprintf("--interval=%d", config.Interval),
		fmt.Sprintf("--payload=%d", config.Payload),
		"--output=out.txt")
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

