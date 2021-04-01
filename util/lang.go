package util

type Lang interface {
	Play() string
	Settings() string
	DarkTheme() string
	Language() string
	Save() string
	Cancel() string
	Turn() string
	SaveName() string
	Restart() string
	Exit() string
	ExitMsg() string
	Congrats() string
	Win() string
	Block() string
	BlockMsg() string
}

type English struct{}

type Russian struct{}

type Ukrainian struct{}

func (_ English) Play() string {
	return "Play"
}

func (_ English) Settings() string {
	return "Settings"
}

func (_ English) DarkTheme() string {
	return "Dark Theme"
}

func (_ English) Language() string {
	return "Language"
}

func (_ English) Save() string {
	return "Save"
}

func (_ English) Cancel() string {
	return "Cancel"
}

func (_ English) Turn() string {
	return "turn"
}

func (_ English) SaveName() string {
	return "Save name"
}

func (_ English) Restart() string {
	return "Restart"
}

func (_ English) Exit() string {
	return "Exit"
}

func (_ English) ExitMsg() string {
	return "Are you sure you want to exit?"
}

func (_ English) Congrats() string {
	return "Congratulations"
}

func (_ English) Win() string {
	return "win"
}

func (_ English) Block() string {
	return "Block!"
}

func (_ English) BlockMsg() string {
	return "No one wins, restart?"
}

func (_ Russian) Play() string {
	panic("implement me")
}

func (_ Russian) Settings() string {
	panic("implement me")
}

func (_ Russian) DarkTheme() string {
	panic("implement me")
}

func (_ Russian) Language() string {
	panic("implement me")
}

func (_ Russian) Save() string {
	panic("implement me")
}

func (_ Russian) Cancel() string {
	panic("implement me")
}

func (_ Russian) Turn() string {
	panic("implement me")
}

func (_ Russian) SaveName() string {
	panic("implement me")
}

func (_ Russian) Restart() string {
	panic("implement me")
}

func (_ Russian) Exit() string {
	panic("implement me")
}

func (_ Russian) ExitMsg() string {
	panic("implement me")
}

func (_ Russian) Congrats() string {
	panic("implement me")
}

func (_ Russian) Win() string {
	panic("implement me")
}

func (_ Russian) Block() string {
	panic("implement me")
}

func (_ Russian) BlockMsg() string {
	panic("implement me")
}

func (_ Ukrainian) Play() string {
	panic("implement me")
}

func (_ Ukrainian) Settings() string {
	panic("implement me")
}

func (_ Ukrainian) DarkTheme() string {
	panic("implement me")
}

func (_ Ukrainian) Language() string {
	panic("implement me")
}

func (_ Ukrainian) Save() string {
	panic("implement me")
}

func (_ Ukrainian) Cancel() string {
	panic("implement me")
}

func (_ Ukrainian) Turn() string {
	panic("implement me")
}

func (_ Ukrainian) SaveName() string {
	panic("implement me")
}

func (_ Ukrainian) Restart() string {
	panic("implement me")
}

func (_ Ukrainian) Exit() string {
	panic("implement me")
}

func (_ Ukrainian) ExitMsg() string {
	panic("implement me")
}

func (_ Ukrainian) Congrats() string {
	panic("implement me")
}

func (_ Ukrainian) Win() string {
	panic("implement me")
}

func (_ Ukrainian) Block() string {
	panic("implement me")
}

func (_ Ukrainian) BlockMsg() string {
	panic("implement me")
}
