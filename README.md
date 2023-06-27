# SubCipher

I have created this simple project as part of my personal learning journey through the land of Rust programming.

Other than serving as a good excuse for me to try things out, this command line application serves the purpose of providing a fairly simple [substitution cipher](https://en.wikipedia.org/wiki/Substitution_cipher)--thus its name, SubCipher--that I used when I was a young kid.

There are way too many different types of substitution ciphers to name here, but the one I used was called *Zenit Polar*, which substituted single letters from any given word by swapping every instance of a *z* for a *p*, *e* for an *o*, *n* for *l*, *i* for *a*, and lastly *t* for a *r*, as well as their counterparts in the oposite direction.

![image][zenit polar]

Therefore, the words **Hello world** would be *encrypted* to **Honne wetnd**. I told you, it was a very simple cipher ;)

## Run it

  $ cargo run

  Message to encrypt:

  I don't know what to do
  
  Encrypted message: A del'r klew whir re de

[zenit polar]: subcipher.png