#!/bin/bash

aarch64-linux-gnu-as -o ARM64.o ARM64.s

if [ $? -ne 0 ]; then
    echo "Error al ensamblar."
    exit 1
fi

aarch64-linux-gnu-ld -o ARM64.elf ARM64.o

if [ $? -ne 0 ]; then
    echo "Error al enlazar."
    exit 1
fi

qemu-aarch64 ./ARM64.elf