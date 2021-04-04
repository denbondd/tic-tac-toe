package util

type Lang interface {
	GetLangNum() int
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

func (_ English) GetLangNum() int {
	return 0
}

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
	return "'s turn"
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

func (_ Russian) GetLangNum() int {
	return 1
}

func (_ Russian) Play() string {
	return "Играть"
}

func (_ Russian) Settings() string {
	return "Настройки"
}

func (_ Russian) DarkTheme() string {
	return "Темная тема"
}

func (_ Russian) Language() string {
	return "Язык"
}

func (_ Russian) Save() string {
	return "Сохранить"
}

func (_ Russian) Cancel() string {
	return "Отмена"
}

func (_ Russian) Turn() string {
	return " очередь"
}

func (_ Russian) SaveName() string {
	return "Сохранить имя"
}

func (_ Russian) Restart() string {
	return "Заново"
}

func (_ Russian) Exit() string {
	return "Выход"
}

func (_ Russian) ExitMsg() string {
	return "Вы уверены что хотите выйти?"
}

func (_ Russian) Congrats() string {
	return "Поздравляем"
}

func (_ Russian) Win() string {
	return "выиграл"
}

func (_ Russian) Block() string {
	return "Ничья"
}

func (_ Russian) BlockMsg() string {
	return "Никто не выиграл, начать заново?"
}

func (_ Ukrainian) GetLangNum() int {
	return 2
}

func (_ Ukrainian) Play() string {
	return "Грати"
}

func (_ Ukrainian) Settings() string {
	return "Настройки"
}

func (_ Ukrainian) DarkTheme() string {
	return "Темна тема"
}

func (_ Ukrainian) Language() string {
	return "Мова"
}

func (_ Ukrainian) Save() string {
	return "Зберегти"
}

func (_ Ukrainian) Cancel() string {
	return "Скасувати"
}

func (_ Ukrainian) Turn() string {
	return " черга"
}

func (_ Ukrainian) SaveName() string {
	return "Зберегти ім'я"
}

func (_ Ukrainian) Restart() string {
	return "Переграти"
}

func (_ Ukrainian) Exit() string {
	return "Вийти"
}

func (_ Ukrainian) ExitMsg() string {
	return "Ви впевнені що хочете вийти?"
}

func (_ Ukrainian) Congrats() string {
	return "Вітання"
}

func (_ Ukrainian) Win() string {
	return "переміг"
}

func (_ Ukrainian) Block() string {
	return "Нічия"
}

func (_ Ukrainian) BlockMsg() string {
	return "Нікто не переміг, переграти?"
}
