# VeryGoodWare
A collection of scripts that I didn't recreate from malware

I created this file with the goal in mind to learn a bit more about Go and recreate some of the malware behaviours I found to be interesting in some of the samples.
This repo contains the following

## NotLokiBot

These scripts were base on a LokiBot piece in which the authors, i order to hide the files from the AV, they hid information within image files. The type of files the used were PNGs, which are the same file type I used to hide information, partly because of the lossy compression algorithm of JPEGs and partly to recreate what I saw.
I wrote 3 different scripts to do what I wanted to do:

- Stego.go
This one is just a simple script made with the goal to open a new calc.exe instance every time is executed. I used this to pretend my "evil" executable was launched.
- Encrypter.go
Encrypter will, as the name suggests, encrypt the file. Instead of using its own algorithm just like the LokiBot sample, I will just use AES to encrypt the executable. Afterwards, this one also appends the encrypted file byte by byte at the end of the PNG.
The resulting file will be way heavier than the normal PNG file.
- Loader.go
This script is the one that should be sent an executed by the end user in their own system. To find the beginning of the encrypted file, I first of all implemented a check that verifies that the PNG file has certain bytes, if it has, then the image that we read contains the payload.
Afterwards, it will look for the IEND bytes and read everything right after them.
My final goal also consisted on execute the extracted executable in memory, but that was something I wasn't able to do in Golang. Maybe there are methods out the I was unable to find, and if so please feel free to fork it or submit a PR.

I also attached the two PNG files, the one without anything (rose.png) and the one with the encrypted data (smile.png). You'll notice the size differences.

## NotOceanLotus

