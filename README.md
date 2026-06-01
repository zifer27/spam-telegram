CARA JALANKAN TOOLS SPAM TELEGRAM! 

PERSIAPAN AWAL:

1. Pastikan komputer terhubung ke internet

2. Buka PowerShell (tekan Windows + R, ketik powershell, enter)

3. Install Go dengan perintah:
   winget install GoLang.Go

4. Tunggu proses install selesai

5. Buat folder kerja:
   cd C:\
   mkdir spam-telegram
   cd spam-telegram

6. Buat file program:
   notepad main.go

7. Copy paste kode program ke notepad, lalu save dan tutup

8. Install library yang dibutuhkan:
   go mod init spam-telegram
   go get github.com/gotd/td/telegram
   go get github.com/gotd/td/telegram/auth
   go get github.com/gotd/td/tg

9. Tunggu semua library selesai terinstall

MENJALANKAN TOOLS:

10. Di PowerShell yang sama, ketik:
    go run main.go

11. Program akan menampilkan tampilan awal

12. Masukkan nomor target dengan format +62...
    Contoh: +628123456789

13. Tekan Enter

14. Program akan mengirim OTP setiap 30 detik

MEMBERHENTIKAN TOOLS:

15. Ketik .stop lalu Enter

16. Program akan berhenti dan menampilkan total OTP terkirim

MENJALANKAN LAGI NANTI:

- Buka folder C:\otp-bot
- Klik kanan -> Open in Terminal
- Ketik: go run main.go

ATAU

- Buka PowerShell
- Ketik: cd C:\spam-telegram
- Ketik: go run main.go

PENTING:

- Nomor target harus pakai format internasional +
- Jangan gunakan untuk spam
- Gunakan sewajarnya
