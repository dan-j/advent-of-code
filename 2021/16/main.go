package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

const ByteLength int = 8

type Reader struct {
	buffer []byte
	offset int
	limits []int
}

func NewReader(input string) *Reader {
	bytes, err := hex.DecodeString(input)
	if err != nil {
		panic(fmt.Errorf("DecodeString(): err = %e", err))
	}

	return &Reader{
		buffer: bytes,
	}
}

func (r *Reader) ReadPacket() (p Packet, err error) {
	// first read may error, but others should be ok
	p.Version, err = r.ReadBits(3)
	if err != nil {
		return p, err
	}

	packetType, _ := r.ReadBits(3)
	p.Type = PacketType(packetType)

	if p.Type == LiteralPacketType {
		p.Literal = r.ReadLiteral()
	} else {
		if lengthTypeID, _ := r.ReadBits(1); lengthTypeID == 0 {
			// dynamic number of sub packets up to limit-bits in length
			limit, _ := r.ReadBits(15)
			r.PushLimit(limit)
			defer r.PopLimit()
			for {
				subPacket, err := r.ReadPacket()
				if err != nil {
					break
				}
				p.SubPackets = append(p.SubPackets, subPacket)
			}
		} else {
			// fixed number of sub packets
			numPackets, _ := r.ReadBits(11)
			p.SubPackets = make([]Packet, numPackets)
			for i := range p.SubPackets {
				p.SubPackets[i], _ = r.ReadPacket()
			}
		}
	}

	return p, nil
}

func (r *Reader) ReadLiteral() (value int) {
	for {
		// each block is 5 bits
		x, err := r.ReadBits(5)
		if err != nil {
			return value
		}

		// insert the 4 "value" bits into value
		value = (value << 4) | x&0b1111

		// if the 1st bit is not set, stop
		if x < 0b10000 {
			break
		}
	}

	return value
}

// ReadBits is the core functionality of the parser. It can read n many bits from the buffer and parse into an int,
// taking care of reading across bytes in the slice.
func (r *Reader) ReadBits(n int) (value int, err error) {
	idx, offset := r.offset/ByteLength, r.offset%ByteLength

	if n+r.offset > len(r.buffer)*8 {
		return 0, io.EOF
	}

	var bitsRead int

	// each iteration will either read n or up to the end of the current byte.
	for bitsRead < n {
		// remaining is how many bits left in the current byte
		remaining := ByteLength - offset
		// toRead is either the remaining number of bits to finish the read, or the number of bits remaining in the
		// current byte
		toRead := int(math.Min(float64(n-bitsRead), float64(remaining)))

		// limits is a stack, we check the last element and stop reading if the limit has been reached.
		if len(r.limits) > 0 {
			limit := r.limits[len(r.limits)-1]
			if toRead+bitsRead+r.offset > limit {
				// must set offset to be the end of the limit, since additional characters between packets are discarded
				r.offset = limit
				return value, io.EOF
			}
		}

		// shift the current byte to remove next bits which aren't part of this read.
		b := r.buffer[idx] >> (ByteLength - offset - toRead)

		// we may have previous bits in b, create a mask to remove them
		mask := int(math.Pow(2, float64(toRead)) - 1)

		// shift value by the number of bits we're reading, then OR them against the masked b
		value = value<<toRead | (int(b) & mask)
		bitsRead += toRead

		offset += toRead
		// if we've read upto the byte boundary, increment idx and reset offset
		if offset == 8 {
			offset = 0
			idx++
		}
	}

	// recalculate the reader's total offset
	r.offset = ByteLength*idx + offset

	return value, nil
}

func (r *Reader) PushLimit(limit int) {
	r.limits = append(r.limits, r.offset+limit)
}

func (r *Reader) PopLimit() {
	r.limits = r.limits[0 : len(r.limits)-1]
}

type PacketType int

const (
	SumPacketType PacketType = iota
	ProductPacketType
	MinPacketType
	MaxPacketType
	LiteralPacketType
	GreaterThanPacketType
	LessThanPacketType
	EqualPacketType
)

type Packet struct {
	Version    int
	Type       PacketType
	Literal    int
	SubPackets []Packet
}

func (p *Packet) Execute() int {
	switch p.Type {
	case SumPacketType:
		var sum int
		for _, packet := range p.SubPackets {
			sum += packet.Execute()
		}
		return sum
	case ProductPacketType:
		product := 1
		for _, packet := range p.SubPackets {
			product *= packet.Execute()
		}
		return product
	case MinPacketType:
		min := p.SubPackets[0].Execute()
		for _, packet := range p.SubPackets[1:] {
			next := packet.Execute()
			if next < min {
				min = next
			}
		}
		return min
	case MaxPacketType:
		max := p.SubPackets[0].Execute()
		for _, packet := range p.SubPackets[1:] {
			next := packet.Execute()
			if next > max {
				max = next
			}
		}
		return max
	case LiteralPacketType:
		return p.Literal
	case GreaterThanPacketType:
		if p.SubPackets[0].Execute() > p.SubPackets[1].Execute() {
			return 1
		}
		return 0
	case LessThanPacketType:
		if p.SubPackets[0].Execute() < p.SubPackets[1].Execute() {
			return 1
		}
		return 0
	case EqualPacketType:
		if p.SubPackets[0].Execute() == p.SubPackets[1].Execute() {
			return 1
		}
		return 0
	}

	panic("Not a valid type to Execute()")
}

func (p *Packet) SumVersions() int {
	sum := p.Version

	if p.SubPackets != nil {
		for _, o := range p.SubPackets {
			sum += o.SumVersions()
		}
	}

	return sum
}

func Part1(input string) int {
	r := NewReader(input)
	p, _ := r.ReadPacket()

	return p.SumVersions()
}

func Part2(input string) int {
	r := NewReader(input)
	p, _ := r.ReadPacket()

	return p.Execute()
}

func main() {
	start := time.Now()
	input, _ := os.ReadFile("input.txt")
	fmt.Printf("Part1(): %d, (%s)\n", Part1(string(input)), time.Since(start))
	start2 := time.Now()
	fmt.Printf("Part2(): %d, (%s)\n", Part2(string(input)), time.Since(start2))
	fmt.Printf("Total time: %s\n", time.Since(start))
}
