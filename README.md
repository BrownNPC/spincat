# Spin Cat


### An annoying little desktop pet

 Spin cat will chase your mouse while making an annoying sound.


![spincat](./spin.gif)

> Perfect for installing on your school computer. (not recommended for legal reasons)

## [Download](https://github.com/BrownNPC/spincat/releases/latest)



## Configure

A file *`spincat-config.json`* is **created in the same folder as the program.** You can edit it in any text editor.

> **NOTE:** Changes apply without restarting the program.

`Size` **(default: 80)** *Size of the window in pixels.*

`Speed` **(default: 4.0)** *How fast the cat travels to the cursor.*

`SpinSpeed` **(default: 0.75)** *How fast the cat spins while moving.*

`MousePassthrough` **(default: true)** *If the window should pass mouse clicks to the program behind it.*

`WindowDecorations` **(default: false)** *If the title bar and close & minimize buttons should show.*

`Quiet` **(default: false)** *If sound should be disabled.*

> **TIP**: You can control volume from your system tray


## How to exit the program?

> *Good question...*


## Building

### Windows:
> No C compiler needed.
```
go build .
```

### Linux/MacOS

> [⚠️ C compiler & Dependencies must be installed.](https://ebitengine.org/en/documents/install.html#Installing_dependencies)

```
go build .
```
