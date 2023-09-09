export QEMU_PATH := $(HOME)/qemu
export QEMU_IMGS_PATH := $(HOME)/qemu/imgs

show:
		ls $$QEMU_PATH
		
init:
		mkdir $$QEMU_PATH

start-debian:
		qemu-system-x86_64 -nographic -m 2G -hda $$QEMU_IMGS_PATH/test.img