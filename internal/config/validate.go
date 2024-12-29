package config

type ValidationData struct {
	SelectedProfile string
}

func (c Config) Validate(validate ValidationData) error {
	if _, err := c.GetProfileByName(validate.SelectedProfile); err != nil {
		return err
	}

	return nil
}
