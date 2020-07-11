# VeryGoodWare

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

In the second place, I also did this other script which proved to be way more complicated to write than the first one. Although it's probably due to the time it took me to work out the arithmetics in it and bit-wise operations.

I decided to take on this one because it was published by a Cylance researcher in its time, and I find them to be a very high level blog, so I went to write the thing which they had a hard time dealing with.

This one has two scripts:
- hideLSB.go
In this script we'll read the WAV file and the random string in the text file. Afterwards, we will iterate over the bytes in the WAV and over the bits in the bytes of the WAV. For each LSB in each byte, we'll AND it with 254. After the ANDing, we will OR it with each bit of the dummy sentence in the script.
After writing this I officially declare myself a brain-dead octopus.
- zylo.go
I named this file after the person without whom I would have died figuring out the maths and the one person who very agressively pressured me into working out the arithmetics. This one file however it's way simpler than the first one. In this case, we only have to grab each LSB of the WAV file and append it to the final txt file.

I'd love to have a working WAV file, because for now the sentence kinda messes with the header bits of the original WAV file and, even though it still is detected as WAV, it's also corrupted and you cannot play the music.
This was hard.
Enjoy.
