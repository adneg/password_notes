# password_notes
This is a over network muntial ssl clinet to collected password.
You need to generate certificates and change the AES key in code before build.

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

first password: a

login admin

to work need run https://github.com/adneg/rserwer
