package health

import "community50/config"

func Ready(c config.Config) bool { return c.Secure() }
