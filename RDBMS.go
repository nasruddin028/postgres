package main

// Penggunaan Library
import ( 
 "bufio"
 "fmt"
// "os"
 "strings"
 "time"
 "gorm.io/driver/postgres"
 "gorm.io/gorm"
 "github.com/gofrs/uuid"
)

type User struct {
 gorm.Model
 ID uuid.UUID 
 Name string 
 Email string 
 Password string
 CreatedAt time.Time
 UpdatedAt time.Time
 DeletedAt *time.Time
}

type Category struct {
	gorm.Model
	ID uuid.UUID 
	Name string 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
   }

type Article struct {
	gorm.Model
	ID uuid.UUID        "gorm:(not null;column:id)"
	Title string        "gorm:not null;column:title"
	Content string 		"gorm:not null;column:content"
	UserID uuid.UUID  		"gorm:not null;column:user_id"
	CategoryID uuid.UUID   "gorm:not null;column:category_id"
	CreatedAt time.Time "gorm:not null;column:created_at"
	UpdatedAt time.Time	"gorm:not null;column:updated_at"
	DeletedAt *time.Time "gorm:not null;column:deleted_at"
   }

func userInputData() string {
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	return input
	}

	func main() {
		 var isRepeat, menuRepeat bool
		 var input string
		 var user User
		 var category Category
		 var article Article
		 var menu int
		 var count int64
		 var ulang int

		 menuRepeat = true
		 isRepeat = true

		 dsn := "user=postgres password=123 dbname=paper port=5432"
		 db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
		 if err != nil {
			panic("failed to connect userbase")
		 }

		 for menuRepeat {			 
			fmt.Println("ǁ▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬ǁ")
			fmt.Println("ǁ           APLIKASI MONITORING ARTIKEL            ǁ")
			fmt.Println("ǁ▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬ǁ")
			fmt.Println("ǁ  1. Tabel User                                   ǁ")
			fmt.Println("ǁ__________________________________________________ǁ")             
			fmt.Println("ǁ                                                  ǁ")
			fmt.Println("ǁ  2. Tabel Artikel                                ǁ")
			fmt.Println("ǁ__________________________________________________ǁ")
			fmt.Println("ǁ                                                  ǁ")
			fmt.Println("ǁ  3. Tabel Kategori                               ǁ")
			fmt.Println("ǁ__________________________________________________ǁ")
			fmt.Println("ǁ                                                  ǁ")
			fmt.Println("ǁ  4. Tampilkan Jumlah Artikel Setiap Penulis      ǁ")
			fmt.Println("ǁ__________________________________________________ǁ")
			fmt.Println("ǁ                                                  ǁ")
			fmt.Println("ǁ  5. Tampilakn Jumlah Artikel Setiap Kategori     ǁ")
			fmt.Println("ǁ__________________________________________________ǁ")
			fmt.Println("ǁ                                                  ǁ")
			fmt.Println("ǁ  0. Keluar dari aplikasi                         ǁ")
			fmt.Println("ǁ__________________________________________________ǁ")
			fmt.Println(" ")
			fmt.Print("      Silahkan pilih menu :     ")
			fmt.Scanln(&menu)		
			if menu == 0 {
				menuRepeat = false
				fmt.Println("Terimakasih, Selamat tinggal")
// FUNGSI 1
//___________________________________________________________________________________________________				
        	} else if menu == 1 {
				isRepeat = true
				for isRepeat { 
					fmt.Println("")
					fmt.Println("")
					fmt.Println("ANDA AKAN MENGEDIT DATA TABEL USER")   
					fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
					fmt.Println("1. Pencarian Data")
					fmt.Println("2. Tambah Data")
					fmt.Println("3. Ubah Data")
					fmt.Println("4. Hapus Data")
					fmt.Println("0. Kembali ke menu utama")
					fmt.Println("")
					fmt.Print("   Input Pilihan : ")
					fmt.Scanln(&ulang)
					fmt.Println("")					
					if ulang == 0 {
						isRepeat = false
						fmt.Println("Terimakasih, kembali ke menu utama")
// FUNGSI 1.1
//___________________________________________________________________________________________________			
					} else if ulang == 1 {
					 fmt.Print("Masukkan Nama Penulis : ")
					 input = userInputData()
					 result := db.Where("name = ?", input).First(&user)
						 if result.Error != nil {
						 fmt.Println("Data Tidak Ditemukan")
						 } else {
						 fmt.Println("______________________")
						 fmt.Println("     Informasi Data   ")
						 fmt.Print("\nNama penulis :  ", user.Name,"\n")
						 fmt.Print("\nEmail penulis :  ", user.Email,"\n")						 
						 fmt.Println("______________________")		
						}
// FUNGSI 1.2
//___________________________________________________________________________________________________			
					 } else if ulang == 2 {
						user.ID = uuid.Must(uuid.NewV4())
						 fmt.Print("Masukkan Nama Penulis : ")
						 input = userInputData()
						 user.Name = input	
						 fmt.Print("Masukkan email : ")
						 input = userInputData()
						 user.Email = input		
						fmt.Print("Masukkan Password : ")
						input = userInputData()
						user.Password = input					  
 					 db.Create(&user)
					 fmt.Println("")
					 fmt.Println("Data Berhasil DiInput")				 
// FUNGSI 1.3
//___________________________________________________________________________________________________			
					 } else if ulang == 3 {
						fmt.Print("Masukkan nama yang akan diupdate : ")
						input = userInputData()
						user.Name = input
						result := db.Where("name = ?", input).First(&user)
						if result.Error != nil {
						   fmt.Println("Data Tidak Ditemukan")		
					   } else {
						   fmt.Println("______________________")
						   fmt.Println("     Informasi Data   ")
						   fmt.Print("\n Nama Penulis :  ", user.Name, "\n")
						   fmt.Print("\n Email Penulis :  ", user.Email,"\n")						   
						   fmt.Println("______________________")		
						   fmt.Print("\nEntri Nama Baru : ")
						   input = userInputData()
						   user.Name = input
						   fmt.Print("Entri Email Baru : ")
						   input = userInputData()
						   user.Email = input
						   fmt.Print("Entri Password Baru : ")
						   input = userInputData()
						   user.Password = input		
						   db.Updates(&user)
						fmt.Println("Data Berhasil Diupdate")		   
						   }
// FUNGSI 1.4
//___________________________________________________________________________________________________	
					   } else if ulang == 4 {
						   fmt.Print("Masukkan Nama Penulis yang akan dihapus :")
						   input = userInputData()
						   user.Name = input
						   result := db.Where("name = ?", input).First(&user)
						   if result.Error != nil {
							  fmt.Println("Data Tidak Ditemukan")
						  } else {
							  db.Delete(&user)
						   fmt.Println("Data Berhasil Dihapus")
						  }
					   }		
				 }
// FUNGSI 2
//___________________________________________________________________________________________________		 
		 } else if menu == 2 {
			isRepeat = true
			for isRepeat { 
				fmt.Println("")
				fmt.Println("")
				fmt.Println("ANDA AKAN MENGOLAH DATA TABEL ARTIKEL")   
				fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
				fmt.Println("1. Pencarian Data")
				fmt.Println("2. Tambah Data")
				fmt.Println("3. Ubah Data")
				fmt.Println("4. Hapus Data")
				fmt.Println("0. Kembali ke menu utama")
				fmt.Println("")
				fmt.Print("   Input Pilihan : ")
				fmt.Scanln(&ulang)
				if ulang == 0 {
					isRepeat = false
					fmt.Println("Terimakasih, kembali ke menu utama")
// FUNGSI 2.1
//___________________________________________________________________________________________________	
				} else if ulang == 1 {
				 fmt.Print("Masukkan Nama Artikel : ")
				 input = userInputData()
					 result := db.Where("title = ?", input).First(&article)
						 if result.Error != nil {
						 fmt.Println("Data Tidak Ditemukan")
						 } else {
						 fmt.Println("___________________________________________________")
						 fmt.Println("     Informasi Artikel   ")
						 fmt.Println("\n Title Artikel :  ", article.Title, "\n")
						 fmt.Println("\n Content Artikel :  ", article.Content, "\n")
						 fmt.Println("\n User id :  ", article.UserID, "\n")
						 fmt.Println("\n Kategori id :  ", article.CategoryID, "\n")
						 fmt.Println("___________________________________________________")
						}
// FUNGSI 2.2
//___________________________________________________________________________________________________		
					 } else if ulang == 2 {
						article.ID = uuid.Must(uuid.NewV4())
						 fmt.Print("Judul Artikel : ")
						 input = userInputData()
						 article.Title = input		
						 fmt.Print("Konten Artikel : ")
						 input = userInputData()
						 article.Content = input		
						 fmt.Print("Masukkan Nama Penulis : ")
						 input = userInputData()
						 user.Name = input				
						 db.Model(&User{}).Where("name = ?", input).First(&user)
						 fmt.Print("User ID Penulis : ")
						 fmt.Println(user.ID)
						 article.UserID = user.ID
						 fmt.Print("Masukkan Jenis kategori : ")
						 input = userInputData()
						 category.Name = input

						 result := db.Model(&Category{}).Where("name = ?", input).First(&category)
						 if result.Error != nil {
							 fmt.Println()
							 fmt.Println("Data Tidak Ditemukan")
							 fmt.Println()		 
						} else {				
							fmt.Print("Kategori Id : ") 
							fmt.Println(category.ID)
						 article.CategoryID = category.ID						  
						 db.Create(&article)
					 fmt.Println(" ")
					 fmt.Println("Data Berhasil Di Input")
						}
					 } else if ulang == 3 {
// FUNGSI 2.3
//___________________________________________________________________________________________________	
						fmt.Println("")
						fmt.Print("Masukkan title yang akan diupdate : ")
						input = userInputData()
						article.Title = input
						result := db.Where("title = ?", input).First(&article)
						if result.Error != nil {
						   fmt.Println("Data Tidak Ditemukan")		
					   	   } else {
						   fmt.Println("_____________________________________")
						   fmt.Println("     Informasi Data   ")
						   fmt.Print("\nTitle Artikel :  ", article.Title, "\n")
						   fmt.Print("\nContent Artikel :  ", article.Content, "\n")
						   fmt.Print("\nUser Id :  ", article.UserID, "\n")
						   fmt.Print("\nKategori Id :  ", article.CategoryID, "\n")
						   fmt.Println("_____________________________________")		
						   fmt.Print("Entri Title Baru : ")
						   input = userInputData()
						   article.Title = input
						   fmt.Print("Entri Content Baru : ")
						   input = userInputData()
						   article.Content = input
						   fmt.Print("Masukkan Nama Penulis : ")
						   input = userInputData()
						   user.Name = input				  
						   db.Model(&User{}).Where("name = ?", input).First(&user)
						   article.UserID = user.ID
						   fmt.Print("Masukkan Jenis kategori : ")
						   input = userInputData()
						   category.Name = input  
						   result := db.Model(&Category{}).Where("name = ?", input).First(&category)
						   if result.Error != nil {
							   fmt.Println()
							   fmt.Println("Data Tidak Ditemukan")
							   fmt.Println()		   
						  } else {
						   article.CategoryID = category.ID
						  }
						   db.Updates(&article)
						   fmt.Println("Data Berhasil Diupdate")
					   		}	
// FUNGSI 2.4
//___________________________________________________________________________________________________		   						   
					   } else if ulang == 4 {
						   fmt.Println("")
						   fmt.Print("Masukkan Judul Artikel yang akan dihapus :")
						   input = userInputData()
						   article.Title = input
						   result := db.Where("title = ?", input).First(&article)
						   if result.Error != nil {
							  fmt.Println("Data Tidak Ditemukan")
						  } else {
							  db.Delete(&article)
						   fmt.Println("Data Berhasil Dihapus")
						  }
					   }
				}
// FUNGSI 3			
//___________________________________________________________________________________________________			
			} else if menu == 3 {
				isRepeat = true
				for isRepeat { 
					fmt.Println("")
					fmt.Println("")
					fmt.Println("ANDA AKAN MENGOLAH DATA TABEL KATEGORI")   
					fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
					fmt.Println("1. Pencarian Data")
					fmt.Println("2. Tambah Data")
					fmt.Println("3. Ubah Data")
					fmt.Println("4. Hapus Data")
					fmt.Println("0. Kembali ke menu utama")
					fmt.Println("")
					fmt.Print("   Input Pilihan : ")
					fmt.Scanln(&ulang)
					if ulang == 0 {
						isRepeat = false
						fmt.Println("Terimakasih, kembali ke menu utama")
// FUNGSI 3.1			
//___________________________________________________________________________________________________		
					} else if ulang == 1 {
					 fmt.Print("Masukkan Nama Kategori : ")
					 input = userInputData()
						 result := db.Where("name = ?", input).First(&category)
							 if result.Error != nil {
							 fmt.Println("Data Tidak Ditemukan")
							 } else {
							 fmt.Println("______________________")
							 fmt.Println("     Informasi Kategori   ")
							 fmt.Println("\n Nama category :  ", category.Name, "\n")
							 fmt.Println("______________________")			
							}
// FUNGSI 3.2			
//___________________________________________________________________________________________________			
						 } else if ulang == 2 {
							category.ID = uuid.Must(uuid.NewV4())
							 fmt.Print("Nama Kategori : ")
							 input = userInputData()
							 category.Name = input
			 				 db.Create(&category)	
						 fmt.Println("Data Berhasil DiInput")
// FUNGSI 3.2			
//___________________________________________________________________________________________________						 
			
						 } else if ulang == 3 {
							fmt.Print("Masukkan nama yang akan diupdate : ")
							input = userInputData()
							category.Name = input
							result := db.Model(&Category{}).Where("name = ?", input).First(&category)
							if result.Error != nil {
								fmt.Println()
								fmt.Println("Data Tidak Ditemukan")
							    fmt.Println()		
						   } else {
							   fmt.Println("__________________________")
							   fmt.Println("     Informasi Kategori   ")
							   fmt.Print("\nNama Kategori :  ", category.Name,"\n")	
							   fmt.Println("__________________________")
			
							   fmt.Print("\nEntri Kategori Baru : ")
							   input = userInputData()
							   category.Name = input
							   
							   db.Updates(&category)
							fmt.Println("Data Berhasil Diupdate")
			   
							   }
// FUNGSI 3.4			
//___________________________________________________________________________________________________							   
						   } else if ulang == 4 {
							   fmt.Print("Masukkan Kategori yang akan dihapus : ")
							   input = userInputData()
							   category.Name = input
							   result := db.Where("name = ?", input).First(&category)
							   if result.Error != nil {
								  fmt.Println("Data Tidak Ditemukan")
							  } else {
								  db.Delete(&category)
							   fmt.Println("Data Berhasil Dihapus")
							  }
						   }
					}
// FUNGSI 4
//___________________________________________________________________________________________________			
			} else if menu == 4 {
				fmt.Println("")
				fmt.Println("")
				fmt.Println(" Fungsi 4 : Menghitung Jumlah Artikel Setiap Penulis")
				fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
				fmt.Print("Masukkan Nama Penulis : ")
							input = userInputData()
							user.Name = input
				
				db.Model(&User{}).Where("name = ?", input).First(&user)
				fmt.Print("User ID Penulis : ")
				fmt.Println(user.ID)
				fmt.Print("Jumlah Tulisan : ")
				db.Model(&Article{}).Where("user_id = ?", user.ID).Count(&count)
				fmt.Println(count)
				fmt.Println("")
				fmt.Println("")
// FUNGSI 5
//___________________________________________________________________________________________________			
			} else if menu == 5 { 				   
				fmt.Println("")
				fmt.Println("")
				fmt.Println(" Fungsi 5 : Menghitung Jumlah Artikel Setiap Kategori ")
				fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
				fmt.Print("Masukkan Jenis kategori : ")
							input = userInputData()
							category.Name = input

							result := db.Model(&Category{}).Where("name = ?", input).First(&category)
							if result.Error != nil {
								fmt.Println()
								fmt.Println("Data Tidak Ditemukan")
							    fmt.Println()		
						   } else {
							    fmt.Print("Kategori Id : ") 
							    fmt.Println(category.ID)
							    fmt.Print("Jumlah Artikel : ")
								db.Model(&Article{}).Where("category_id = ?", category.ID).Count(&count)
								fmt.Println(count)
								fmt.Println("")
								fmt.Println("")
						   }
//---------------------------------------------------------------------------

			} 

	}
	
}