package config

import "fmt"

type ValidationData struct {
	SelectedProfile string
}

func (c Config) Validate(validate ValidationData) error {
	if _, err := c.GetProfileByName(validate.SelectedProfile); err != nil {
		return err
	}

	if len(c.Profiles) == 0 {
		return fmt.Errorf("At least one profile must be present in the config")
	}

	return nil
}
