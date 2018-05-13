# password_notes

build script:
	#!/bin/bash
	place=`pwd`
	program=`basename $place`
	program=`echo $program".sh"`
	date
	qtdeploy -docker build desktop #&&\
	cp $place/pliki/* deploy/linux/
	date
	cd deploy/linux && ./$program

first password a login admin
