package main

// Deklarasi Library
import ( 
 "bufio"
 "fmt"
 "os"
 "strings"
 "time"
 "strconv"
 "gorm.io/driver/postgres"
 "gorm.io/gorm"
 "github.com/satori/go.uuid"
)

//Deklarasi Struct (penyimpanan sementara)
type Product struct {
 gorm.Model
 ID uuid.UUID 
 Code string 
 Name string 
 Price int 

 CreatedAt time.Time
 UpdatedAt time.Time
 DeletedAt *time.Time
}

//Deklarasi Fungsi untuk membaca data inputan pengguna
func userInputData() string {
var input string
reader := bufio.NewReader(os.Stdin)
input, _ = reader.ReadString('\n')
input = strings.Replace(input, "\n", "", -1)
return input
}

//Program utama
func main() {
    //Deklarasi Variabel
     var isRepeat bool
     var input string
     var product Product 
     var baca int

     isRepeat = true;
 
     //koneksi ke Database
     dsn := "user=postgres password=123 dbname=latihan port=5432"
     db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
     if err != nil {
        panic("failed to connect database")
     }

     //perulangan menu utama
     for isRepeat {    
      fmt.Println("1. Pencarian Data")
      fmt.Println("2. Tambah Data")
      fmt.Println("3. Ubah Data")
      fmt.Println("4. Hapus Data")
      fmt.Println("0. Keluar Dari Applikasi")

         //input = userInputData()
         fmt.Scanln(&baca)
         if baca == 0 {
            isRepeat = false
            fmt.Println("Terimakasih, Selamat tinggal")
         } else if baca == 1 {
             fmt.Println("Masukkan Kode Produk")
             input = userInputData()
             result := db.Where("code = ?", input).First(&product)
             if result.Error != nil {
                 fmt.Println("Data Tidak Ditemukan")
             } else {
                 fmt.Println("______________________")
                 fmt.Println("     Informasi Data   ")
                 fmt.Println("\nKode Barang :  \n", product.Code)
                 fmt.Println("\nNama Barang :  \n", product.Name)
                 fmt.Println("\nHarga Barang : \n", product.Price)
                 fmt.Println("______________________")

                }
         } else if baca == 2 {
             product.ID = uuid.Must(uuid.NewV4())
             fmt.Println("Masukkan Kode Produk")
             input = userInputData()
             product.Code = input


             fmt.Println("Masukkan Nama Produk")
             input = userInputData()
             product.Name = input


             fmt.Println("Masukkan Harga Produk")
             input = userInputData()
             product.Price,_ = strconv.Atoi(input)
             db.Create(&product)

             fmt.Println("Data Berhasil DiInput")
           
         } else if baca == 3 {
             fmt.Println("Masukkan Kode Produk yang akan diupdate :")
             input = userInputData()
             product.Code = input
             result := db.Where("code = ?", input).First(&product)
             if result.Error != nil {
                fmt.Println("Data Tidak Ditemukan")
            } else {
                fmt.Println("______________________")
                fmt.Println("     Informasi Data   ")
                fmt.Println("\nKode Barang :  \n", product.Code)
                fmt.Println("\nNama Barang :  \n", product.Name)
                fmt.Println("\nMasukkan Nama Baru Produk : \n")
                input = userInputData()
                product.Name = input
                fmt.Println("\nHarga Barang :  \n", product.Price)
                fmt.Println("Masukkan Harga Baru Produk : \n")
                input = userInputData()
                product.Price, _ = strconv.Atoi(input)
                db.Updates(&product)
             fmt.Println("Data Berhasil Diupdate")



                }
            } else if baca == 4 {
                fmt.Println("Masukkan Kode Produk yang akan dihapus :")
                input = userInputData()
                product.Code = input
                result := db.Where("code = ?", input).First(&product)
                if result.Error != nil {
                   fmt.Println("Data Tidak Ditemukan")
               } else {
                   db.Delete(&product)
                fmt.Println("Data Berhasil Dihapus")
               }
            }
    }
}    
    