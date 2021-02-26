package model

type (
	UserTable struct {
		UserID		string	`json:"id" gorm:"column:id"`
		NamaUser	string	`json:"nama_user" gorm:"column:nama_user"`
		TL			string	`json:"tanggal_lahir" gorm:"column:tanggal_lahir"`
		NoKTP		string	`json:"no_ktp" gorm:"column:no_ktp"`
		Pekerjaan	string	`json:"pekerjaan" gorm:"column:pekerjaan"`
		PT			string	`json:"pendidikan_terakhir" gorm:"column:pendidikan_terakhir"`
	}

	UserData struct {
		NamaUser	string	`json:"nama_user" gorm:"column:nama_user"`
		TL			string	`json:"tanggal_lahir" gorm:"column:tanggal_lahir"`
		NoKTP		string	`json:"no_ktp" gorm:"column:no_ktp"`
		Pekerjaan	string	`json:"pekerjaan" gorm:"column:pekerjaan"`
		PT			string	`json:"pendidikan_terakhir" gorm:"column:pendidikan_terakhir"`
	}
)
