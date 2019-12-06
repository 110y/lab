#!/usr/bin/python2

# https://qiita.com/sg-matsumoto/items/8194320db32d4d8f7a16

from bcc import BPF

bpf_text = """
int trace_sys_clone(struct pt_regs *ctx) {
  bpf_trace_printk("sys_clone\\n");
  return 0;
}
"""

b = BPF(text=bpf_text)
b.attach_kprobe(event="__x64_sys_clone", fn_name="trace_sys_clone")

b.trace_print()
