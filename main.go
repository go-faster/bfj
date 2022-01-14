package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-faster/jx"

	"golang.org/x/tools/benchmark/parse"
)

func write(b *parse.Benchmark) {
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	e.Obj(func(e *jx.Encoder) {
		e.FieldStart("name_full")
		e.Str(b.Name)

		name := strings.TrimPrefix(b.Name, "Benchmark")
		var cpus int
		{
			idx := strings.LastIndex(name, "-")
			if idx > 0 {
				v, err := strconv.Atoi(name[idx+1:])
				if err == nil {
					cpus = v
					name = name[:idx]
				}
			}
		}
		e.FieldStart("name")
		e.Str(name)
		if cpus > 0 {
			e.FieldStart("cpus")
			e.Int(cpus)
		}

		e.FieldStart("count")
		e.Int(b.N)

		if (b.Measured & parse.NsPerOp) != 0 {
			e.FieldStart("duration_ns")
			e.Float64(b.NsPerOp)
		}
		if (b.Measured & parse.MBPerS) != 0 {
			e.FieldStart("mb_per_sec")
			e.Float64(b.MBPerS)
		}
		if (b.Measured & parse.AllocedBytesPerOp) != 0 {
			e.FieldStart("alloc_bytes")
			e.UInt64(b.AllocedBytesPerOp)
		}
		if (b.Measured & parse.AllocsPerOp) != 0 {
			e.FieldStart("alloc_count")
			e.UInt64(b.AllocsPerOp)
		}
	})

	fmt.Println(e)
}

func run() error {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		b, err := parse.ParseLine(s.Text())
		if err != nil {
			continue
		}
		write(b)
	}
	return s.Err()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
	}
}
