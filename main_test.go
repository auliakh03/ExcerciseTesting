package main

import (
	"testing"
)

func TestHitungHargaTotal(t *testing.T) {
	type args struct {
		hargaItem float64
		ongkir    float64
		qty       int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "failed in harga item",
			args: args{
				hargaItem: 0,
				ongkir:    1,
				qty:       1,
			},
			want:    0,
			wantErr: true,
		},

		{
			name: "failed in qty",
			args: args{
				hargaItem: 1,
				ongkir:    1,
				qty:       0,
			},
			want:    0,
			wantErr: true,
		},

		{
			name: "failed in ongkir",
			args: args{
				hargaItem: 1,
				ongkir:    0,
				qty:       1,
			},
			want:    0,
			wantErr: true,
		},

		{
			name: "sukses",
			args: args{
				hargaItem: 15000,
				ongkir:    10000,
				qty:       2,
			},
			want:    45000,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HitungHargaTotal(tt.args.hargaItem, tt.args.ongkir, tt.args.qty)
			if (err != nil) != tt.wantErr {
				t.Errorf("HitungHargaTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HitungHargaTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPembayaranBarang(t *testing.T) {
	tests := []struct {
		name             string
		hargaTotal       float64
		metodePembayaran string
		dicicil          bool
		wantErr          string
	}{
		{
			name:             "failed harga tidak boleh nol",
			hargaTotal:       0,
			metodePembayaran: "cod",
			dicicil:          false,
			wantErr:          "harga tidak bisa nol",
		},
		{
			name:             "failed metode pembayaran tidak dikenali",
			hargaTotal:       10000,
			metodePembayaran: "emoney",
			dicicil:          false,
			wantErr:          "metode tidak dikenali",
		},
		{
			name:             "failed metode pembayaran credit harus dicicil",
			hargaTotal:       200000,
			metodePembayaran: "credit",
			dicicil:          false,
			wantErr:          "credit harus dicicil",
		},
		{
			name:             "success",
			hargaTotal:       600000,
			metodePembayaran: "credit",
			dicicil:          true,
			wantErr:          "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PembayaranBarang(tt.hargaTotal, tt.metodePembayaran, tt.dicicil)
			if (err != nil && err.Error() != tt.wantErr) || (err == nil && tt.wantErr != "") {
				t.Errorf("PembayaranBarang() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Menambahkan pesan cetak metode pembayaran
			if err == nil {
				t.Logf("Pembayaran berhasil dengan metode: %s", tt.metodePembayaran)
			}
		})
	}
}
