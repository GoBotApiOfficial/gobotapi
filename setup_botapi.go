package gobotapi

func (ctx *BasicClient) SetupBotAPI(hostName string, port int, https bool) {
	ctx.BotApiConfig = Config{
		HostName: hostName,
		Port:     port,
		HTTPS:    https,
	}
}
