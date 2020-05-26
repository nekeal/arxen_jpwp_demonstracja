package main

import "testing"

func Test_PrimesCounter(t *testing.T) {
	type args struct {
		start uint64
		limit uint64
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{name: "test error", args: struct {
			start uint64
			limit uint64
		}{start: 100, limit: 99}, want: 0, wantErr: true},
		{name: "test numberOfPrimes", args: struct {
			start uint64
			limit uint64
		}{start: uint64(10000), limit: uint64(200000)}, want: 16755, wantErr: false},
		{name: "test for big numbers", args: struct {
			start uint64
			limit uint64
		}{start: uint64(100), limit: uint64(1000)}, want: 143, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrimesCounter(tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrimesCounter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PrimesCounter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PrimesCounterHandler(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	type args struct {
		start uint64
		limit uint64
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{name: "test error", args: struct {
			start uint64
			limit uint64
		}{start: 100, limit: 99}, want: 0, wantErr: true},
		{name: "test numberOfPrimes", args: struct {
			start uint64
			limit uint64
		}{start: uint64(10000), limit: uint64(200000)}, want: 16755, wantErr: false},
		{name: "test for big numbers", args: struct {
			start uint64
			limit uint64
		}{start: uint64(100), limit: uint64(1000)}, want: 143, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrimesCounterHandler(tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrimesCounterHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PrimesCounterHandler() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_PrimesCounter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if _, err := PrimesCounter(10000,500000); err != nil {
			panic(err)
		}
	}
}


func Benchmark_PrimesCounterHandler(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if _, err := PrimesCounterHandler(10000,500000); err != nil {
			panic(err)
		}
	}
}