root directory
--------------

    $ git clone --recursive git@github.com:motiejus/dotfiles.git .dotfiles
    $ cd .dotfiles/secrets/
    $ git crypt unlock <(gpg -d <(sed -n '/BEGIN PGP/,/END PGP/ p' README.md))
    $ make -C $(hostname) install

This is the symmetric encryption key for `git-crypt`:

```
-----BEGIN PGP MESSAGE-----

jA0ECQMCWhpsUYzqWGDq0sAAAYTR9moDry/FpymMu2OhEERTLumZdCsytGtZFBCJ
US3IGVBavQago+jSw92hDYupruCL7oNZs50wobS80e5a6Tw+Pw+t1LUpmmwxLXnX
S5cDBvVox7NMUAx7v0SzNhoIPc4S4lfP+zS1CQSpsVPRBujQQoZBYdGXBosyO2lq
ebv4WpgncbBKQhT2RqZV4+DR4NYqmZqz1A4TKZx5b1ViaZKjnQwwkL2TtcSPGsU4
cJk5Si6XZ5ItOmFBDjHPCnKt
=ADT3
-----END PGP MESSAGE-----
```
