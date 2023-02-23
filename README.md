dotfiles
========

This is my poor-man's configuration management.

Split into 3 parts:

Dotfiles in $HOME
-----------------

    $ git config protocol.file.allow always
    $ git clone --recursive git@github.com:motiejus/dotfiles.git .dotfiles
    $ cd .dotfiles
    $ stow bash ctags tmux vim

Updating submodules
-------------------

On a clean tree:

    $ cp .gitmodules.remote .gitmodules
    $ git diff  # make sure no submodules are missed
    $ git submodule foreach git pull
    $ git checkout -- .gitmodules
    $ git commit -am "update submodules"
    $ git subtrac update
    $ git push origin master master.trac

... and update this README when some of the steps turn out to be
wrong/misplaced.
