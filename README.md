# symmer
#### A CLI tool to make centralized configs a bit easier. 
___

## Description

symmer is my simplier take on NixOS's Home Manager system. symmer is an easier
way to manage your dotfiles and keep them all in one place using symlinks. This
in turn makes handling configurations for different apps way simplier while
using the power of git to keep everything in one spot and version managed. 



## Install options

### Releases
1. Download executable for your OS in the Releases page.
2. Move symmer to a sourced location or to a location in your system path if your on Windows.
3. Use ```symmer``` in your terminal to get started!

### Build
```go
git clone https://github.com/Tyy47/symmer.git
cd symmer  
go build
```

### Go Install
```go
git clone https://github.com/Tyy47/symmer.git
cd symmer  
go install
```

### Portable
1. Download the executable via the Releases page.
2. Place the symmer executable inside of your configurations folder.
3. Run symmer to generate a config and fill it out with your configuration paths.


## Usage
When running symmer for the first time, a JSON config file will be created for you to start filling out your pathways and apps. Below is an example of how to write a symmer config.
```json
{
    "Niri": {
        "cfg": "/home/$USER/Symmed_Configs/Niri/config.kdl", // Config stored in git repo
        "des": "/home/$USER/.config/niri/config.kdl" // Where config usually is normally
    },
    "Neovim": { 
        "cfg": "/home/tyler/Symmed_Configs/Neovim/*", // Supports wildcards!
        "des": "/home/tyler/.config/neovim/"
    }
}
```

Above is an example of how to setup a niri config using symmer. To create a link first you'll put the name as the identifier ( This can be anything but it's best to put the name of the application to make it visible to you in the terminal if it worked or not ). 

Next you'll put the location of where you stored your config, this will be placed inside of a git repo so it can be backed up and rewound if there is any issues.

Lastly, the symlink section is to where you want this config to be, for example, Niri's config loads in .config/niri/config.kdl so I assigned the symlink to be at that location so Niri can load it.

Finally, once all applications and pathways are filled out that you require, re-run symmer and it'll get to work symlinking all of your dotfiles!

## License
symmer is under the MIT license.
