#!/bin/bash
set -exuo pipefail

EMACSVOLUMEN="${EMACSVOLUME:-/tmp/emacs}"

mkdir -p "$EMACSVOLUMEN/doomd"
mkdir -p ~/.config/

rm -rf ~/.emacs.d/

if test ! -d ~/.doom.d; then
    rm -rf ~/.doom.d/
    ln -s "$EMACSVOLUMEN/doomd" ~/.doom.d
fi

if test ! -f "$EMACSVOLUMEN/marker"; then
    rm -rf ~/.config/emacs
    rm -rf "$EMACSVOLUMEN/doomf"
    mkdir -p "$EMACSVOLUMEN/doomf" 
    ln -s "$EMACSVOLUMEN/doomf/" ~/.config/emacs
    git clone --depth 1 https://github.com/doomemacs/doomemacs ~/.config/emacs
    ~/.config/emacs/bin/doom install --force --verbose
    rm -rf ~/.doom.d/*
    cp -Rv ~/default_doom_conf/* ~/.doom.d/
    touch "$EMACSVOLUMEN/marker"
else
    if test ! -d ~/.config/emacs; then
        rm -rf ~/.config/emacs
        ln -s "$EMACSVOLUMEN/doomf/" ~/.config/emacs
    fi
fi

~/.config/emacs/bin/doom sync
emacs