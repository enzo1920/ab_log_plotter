package configer


//структура конфига 
type Configuration struct {
	Db_user string `json:"db_user"`
	Db_passwd string `json:"db_passwd"`
	Db_name string `json:"db_name"`
	Log_file_name string  `json:"log_file_name"`
}
