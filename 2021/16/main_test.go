package main

import (
	"os"
	"testing"
)

func ReadFile(path string) string {
	f, _ := os.ReadFile(path)
	return string(f)
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "small example",
			input: "D2FE28",
			want:  6,
		},
		{
			name:  "medium example",
			input: "38006F45291200",
			want:  9,
		},
		{
			name:  "medium example 2",
			input: "EE00D40C823060",
			want:  14,
		},
		{
			name:  "main example 1",
			input: "8A004A801A8002F478",
			want:  16,
		},
		{
			name:  "main example 2",
			input: "620080001611562C8802118E34",
			want:  12,
		},
		{
			name:  "main example 3",
			input: "C0015000016115A2E0802F182340",
			want:  23,
		},
		{
			name:  "main example 4",
			input: "A0016C880162017C3686B18A3D4780",
			want:  31,
		},
		{
			name:  "main",
			input: ReadFile("input.txt"),
			want:  860,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.input); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "C200B40A82",
			input: "C200B40A82",
			want:  3,
		},
		{
			name:  "04005AC33890",
			input: "04005AC33890",
			want:  54,
		},
		{
			name:  "880086C3E88112",
			input: "880086C3E88112",
			want:  7,
		},
		{
			name:  "CE00C43D881120",
			input: "CE00C43D881120",
			want:  9,
		},
		{
			name:  "D8005AC2A8F0",
			input: "D8005AC2A8F0",
			want:  1,
		},
		{
			name:  "F600BC2D8F",
			input: "F600BC2D8F",
			want:  0,
		},
		{
			name:  "9C005AC2F8F0",
			input: "9C005AC2F8F0",
			want:  0,
		},
		{
			name:  "9C0141080250320F1802104A08",
			input: "9C0141080250320F1802104A08",
			want:  1,
		},
		{
			name:  "main",
			input: ReadFile("input.txt"),
			want:  470949537659,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.input); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadBits(t *testing.T) {
	type fields struct {
		buffer []byte
		offset int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "version",
			fields: fields{
				buffer: []byte{0b11010010, 0b11111110, 0b00101000},
			},
			args: args{
				n: 3,
			},
			want: 6,
		},
		{
			name: "packet type",
			fields: fields{
				buffer: []byte{0b11010010, 0b11111110, 0b00101000},
				offset: 3,
			},
			args: args{
				n: 3,
			},
			want: 4,
		},
		{
			name: "first literal",
			fields: fields{
				buffer: []byte{0b11010010, 0b11111110, 0b00101000},
				offset: 6,
			},
			args: args{
				n: 5,
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				buffer: tt.fields.buffer,
				offset: tt.fields.offset,
			}
			if got, _ := r.ReadBits(tt.args.n); got != tt.want {
				t.Errorf("ReadBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
