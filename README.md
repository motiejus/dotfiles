dotfiles
========

This is my poor-man's configuration management.

Split into 3 parts:

Dotfiles in $HOME
-----------------

    $ git clone --recursive git@github.com:motiejus/dotfiles.git .dotfiles
    $ cd .dotfiles
    $ stow bash ctags tmux vim

Dotfiles in /
-------------

See `root/`.

Secrets & non-symlinks
----------------------

See `nonlinks`.
