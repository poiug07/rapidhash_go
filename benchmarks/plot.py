#!/usr/bin/env python3
"""Parse `go test -bench` output and render a throughput comparison chart.

Usage:
    go test -bench=. | python3 plot.py [output.png]

Reads bench output on stdin, writes a grouped bar chart (default
benchmarks/benchmark.png). With no matplotlib installed it still prints a
summary table to stdout.
"""

import re
import sys

try:
    import matplotlib.pyplot as plt
except ImportError:
    plt = None


def parse(lines):
    # e.g. BenchmarkRapidhash/input_size_2048-12  7802228  154.7 ns/op  7641.79 MB/s
    pat = re.compile(
        r"Benchmark(?P<name>\w+)/input_size_(?P<size>\d+)"
        r"(?:-\d+)?\s+\d+\s+(?P<ns>[\d.]+)\s+ns/op(?:\s+(?P<mbs>[\d.]+)\s+MB/s)?"
    )
    results = {}
    for line in lines:
        m = pat.search(line)
        if m:
            results.setdefault(m["name"], {})[int(m["size"])] = (
                float(m["ns"]),
                float(m["mbs"]) if m["mbs"] else None,
            )
    return results


def main():
    lines = sys.stdin.read().splitlines()
    results = parse(lines)
    if not results:
        sys.exit("no benchmark lines found on stdin")

    sizes = sorted(next(iter(results.values())).keys())
    names = list(results.keys())

    print(f"{'size':>8} | " + " | ".join(f"{n:>22}" for n in names))
    for s in sizes:
        row = []
        for n in names:
            ns, mbs = results[n].get(s, (float("nan"), None))
            row.append(f"{ns:8.2f} ns/op" if mbs is None else f"{ns:8.2f} ns {mbs:10.2f} MB/s")
        print(f"{s:>8} | " + " | ".join(f"{r:>22}" for r in row))

    if plt is None:
        print("\nmatplotlib not installed; skipping chart", file=sys.stderr)
        return

    out = sys.argv[1] if len(sys.argv) > 1 else "benchmark.png"
    x = range(len(sizes))
    width = 0.8 / len(names)
    fig, ax = plt.subplots(figsize=(10, 5))
    for i, n in enumerate(names):
        vals = [results[n][s][0] for s in sizes]
        ax.bar([xi + i * width for xi in x], vals, width, label=n)
    ax.set_xticks([xi + width * (len(names) - 1) / 2 for xi in x])
    ax.set_xticklabels([str(s) for s in sizes])
    ax.set_xlabel("input size (bytes)")
    ax.set_ylabel("ns/op")
    ax.set_yscale("log")
    ax.set_title("rapidhash vs xxhash (lower is better)")
    ax.legend()
    fig.tight_layout()
    fig.savefig(out, dpi=150)
    print(f"\nwrote {out}", file=sys.stderr)


if __name__ == "__main__":
    main()
