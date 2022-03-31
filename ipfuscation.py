# ipfuscation.py
# reads a list of IP addresses and translates them to shellcode
# purpose is to hide malicious code as IP addresses, which could be interpreted as C&C data
# idea from article: https://www.bleepingcomputer.com/news/security/hive-ransomware-uses-new-ipfuscation-trick-to-hide-payload/

from capstone import *

def disassemble_me(hex_val):
    CODE = bytes(hex_val)
    md = Cs(CS_ARCH_X86, CS_MODE_64)

    for i in md.disasm(CODE, 0x1000):
        print("0x%x:\t%s\t%s" %(i.address, i.mnemonic, i.op_str))

with open("inputIP.txt") as f:
    lines = f.readlines()
    hex_code = []
    for line in lines:
        line = line.split(".")[::-1]
        for num in line:
            hex_code.append(int(num))
    disassemble_me(hex_code)
