# Alfred My Mind🎗 [![Workflows](https://img.shields.io/badge/More%20Workflows-🎩-purple.svg)](https://github.com/learn-anything/alfred-workflows) [![Thanks](https://img.shields.io/badge/Say%20Thanks-💗-ff69b4.svg)](https://www.patreon.com/nikitavoloboev)
> [Alfred](https://www.alfredapp.com/) workflow to search through my notes and bookmarks on the web

<img src="http://i.imgur.com/4wvJNy6.png" width="500" alt="img">

This workflow lets you search through entirety of [my personal wiki](https://wiki.nikitavoloboev.xyz), [all the articles](https://wiki.nikitavoloboev.xyz/sharing/my-articles.html) I wrote, [GitHub repositories](https://wiki.nikitavoloboev.xyz/sharing/my-github.html) I made as well as [GitHub Gists](https://gist.github.com/nikitavoloboev) I share and a lot more.

You can read why I made the workflow [here](https://medium.com/@NikitaVoloboev/opening-up-my-mind-%EF%B8%8F-575c8ece8a24).

As this workflow is focused and optimised for fast access to all the knowledge and references I have indexed. It is advised you read through some parts of my wiki [here](https://wiki.nikitavoloboev.xyz) first.

My goals with sharing both this workflow and my [wiki knowledge base](https://wiki.nikitavoloboev.xyz) is to extend my idea of [knowledge bootstrapping](https://medium.com/@NikitaVoloboev/knowledge-bootstrapping-36c97e0dee19#.udmp9eotg) and tapping into a person's expertise in the most transparent way possible.

## Install
Download the workflow from [GitHub releases](../../releases/latest).

### Small Demo

<img src="media/demo.gif" width="500" alt="img">

## Contributing
[Suggestions](https://github.com/nikitavoloboev/alfred-my-mind/issues) and pull requests are highly encouraged!

This workflow heavily leverages curated lists from [Learn Anything](https://learn-anything.xyz/) that can be found [here](https://github.com/learn-anything/learn-anything/wiki/Curated-Lists). You can contribute directly to those lists so everyone benefits.

## Developing
If you want to add features and things to the workflow. It is best to use [this Alfred CLI tool](https://godoc.org/github.com/jason0x43/go-alfred/alfred) which you can install by running:

`go install github.com/jason0x43/go-alfred/alfred`

You can then run `alfred link` inside this repo you cloned. This will make a symbolic link of the [`workflow`](workflow) directory.

You can then make your changes to the code and after run `alfred build` to build the go binary to `workflow` directory. Which you can then use from inside Alfred [script filters](https://www.alfredapp.com/help/workflows/inputs/script-filter/).

## Credits
The workflow is built using [Deanishe](https://github.com/deanishe)'s amazing [AwGo](https://github.com/deanishe/awgo) library.

## Thank you 💜
You can support what I do on [Patreon](https://www.patreon.com/nikitavoloboev) or look into [other projects](https://nikitavoloboev.xyz/projects) I shared.

## License
MIT © [Nikita Voloboev](https://www.nikitavoloboev.xyz)
