root directory
--------------

    $ git clone --recursive git@github.com:motiejus/dotfiles.git .dotfiles
    $ cd .dotfiles/root/
    $ sudo stow --ignore='\.sw[op]' -v -t / $(hostname)
