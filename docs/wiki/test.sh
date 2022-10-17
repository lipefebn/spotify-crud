#!/bin/bash
# sirene.sh
# Bom para chamar a atenção dos colegas de trabalho :)

echo -ne "\033[11;2000]"   # Cada bipe dura quase um segundo
while :                   # Loop eterno
do
	echo -ne "\033[100;100]\a" ; sleep 0.5  # Tom alto (agudo)
	echo -ne "\033[10;400]\a" ; sleep 0.5  # Tom baixo (grave)
	echo -ne "\033[10;400]\a" ; sleep 0.5  # Tom baixo (grave)
	echo -ne "\033[10;400]\a" ; sleep 0.5  # Tom baixo (grave)
	echo -ne "\033[10;400]\a" ; sleep 1  # Tom baixo (grave)
	echo -ne "\033[10;400]\a" ; sleep 1  # Tom baixo (grave)
done