package main

import (
	"errors"
	"fmt"
)

const (
	tax = 10
	app = 2000
)

func main() {
	totalHarga, err := HitungHargaTotal(15000, 10000, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Total Harga:", totalHarga)

	err = PembayaranBarang(600000, "credit", true)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Pembayaran berhasil!")
}

func HitungHargaTotal(hargaItem, ongkir float64, qty int) (float64, error) {
	if hargaItem <= 0 {
		return 0, errors.New("harga barang tidak boleh nol")
	}

	if qty <= 0 {
		return 0, errors.New("jumlah barang tidak boleh nol")
	}

	// 15000 * 2 = 30000
	hargaAkhirItem := hargaItem * float64(qty)

	if ongkir <= 0 {
		return 0, errors.New("harga ongkir tidak boleh nol")
	}

	// 30000 + 10000
	hargaSetelahOngkir := hargaAkhirItem + ongkir
	// 40000 + 3000
	pajak := hargaAkhirItem * tax / 100
	total := hargaSetelahOngkir + pajak + app

	return total, nil
}

func PembayaranBarang(hargaTotal float64, metodePembayaran string, dicicil bool) error {
	if hargaTotal <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	validMetodes := map[string]bool{"cod": true, "transfer": true, "debit": true, "credit": true, "gerai": true}
	if !validMetodes[metodePembayaran] {
		return errors.New("metode tidak dikenali")
	}

	if metodePembayaran == "credit" && !dicicil {
		return errors.New("credit harus dicicil")
	}

	if dicicil && metodePembayaran != "credit" {
		return errors.New("metode pembayaran harus credit untuk dicicil")
	}

	if dicicil && hargaTotal < 500000 {
		return errors.New("harga total harus >= 500.000 untuk pembayaran dicicil")
	}

	return nil
}
