package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/iovisor/gobpf/bcc"
)

const source string = `
int trace_sys_clone(struct pt_regs *ctx) {
  bpf_trace_printk("Hello Go BPF!\\n");
  return 0;
}
`

func main() {
	m := bcc.NewModule(source, []string{})
	defer m.Close()

	cloneKprobe, err := m.LoadKprobe("trace_sys_clone")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load kprobe__sys_clone: %s\n", err)
		os.Exit(1)
	}

	syscallName := bcc.GetSyscallFnName("clone")

	err = m.AttachKprobe(syscallName, cloneKprobe, -1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to attach kprobe__sys_fclone: %s\n", err)
		os.Exit(1)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	<-sig
}
