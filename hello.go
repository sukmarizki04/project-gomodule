package projectmodule

func Hello(name string) string {
	return "Sukma" + name
}

/*menggagalkan unit test menggunakan panic merupakan hal yang bagus
Golang sendiri menggunakan cara sendiri untuk menggalkan unit test menggunakan  testing.T
Terdapat Function fail().Failnow() dan fatal() jika kitaingin menggagalkan unit test


#FAIL() DAN FAILNOW()

Sama-sama menggalkan unit test namun ada perbedaan di function ini
jika kita menggunakan Fail() akan menggalkan function dan akan terus berjalan mengeksekusi code

sedangkan fileNow() dia akan berhenti ketika code menemukan function failnow dan tidak melanjutkan code dibawahnya


pembahasan selanjutnya menggunakan t.Error dan t.fatal
ERROR Lebih seperti melakukan log (print) error,namun setelah melakukan log error diaakan secara otomatis memanggil function fail,sehingga mengakibatkan unit test dianggap gagal
namun karna hanya memanggil fail(),artinya eksekusi unit test	akan tetap berjalan sampai selesai

Fatal() mirip dengan error,hanya saja ,setelah melakukan log error dia akan memangil failnow ,sehingga mengakibatkan eksekusi unit test berhenti


Membahas Tentang Assertion

Melakukan pengecekan secara manual menggunakan if else sangatlah menyebalkan
Apa lagi result data yang harus dicek itu banyak
Oleh karena itu,sangat disarankan menggunakan assertion untuk melakukan pengecekan
Sayangnya ,Golang tidak menyediakan package untuk assertion ,sehingga kita butuh menambahkan library untuk melakukan assertion

salah satu library assertion yang paling populer di Golang adalah Testify
kita bisa menggunakan libbrary ini untuk melakukan assertion terhadap result data di unit test
kita bisa menambahkannya di gomodue kita:

Membahasa tentang skip test
*/
