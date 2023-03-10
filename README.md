# GoPomodoro

GoPomodoro is a small app used to improve time management with the technique called [pomodoro](https://en.wikipedia.org/wiki/Pomodoro_Technique).
The application is written in Go with use of [Fyne v2](https://github.com/fyne-io/fyne) gui library.

# Goal

The aim was to create a small-windowed app which does not bother you while working but subtly prompts you to keep your eyes on the goal.

## Download & usage

### Windows

Download .exe from [releases page](github.com/wwicked/gopomodoro/releases).

### Linux

Requires libasound2-dev:
`apt install libasound2-dev`

## Features & caveats

Main window
<br>
![Main window](/assets/images/main.png)

Readable timer
<br>
![Timer](/assets/images/timer.png)

Breaks
<br>
![Breaks](/assets/images/break.png)

Pausing
<br>
![Pausing](/assets/images/pause.png)

Settings
<br>
![Settings](/assets/images/settings.png)

### Todo

-   Add an option to change background color for both work and break times.
-   Add an option to change short and long breaks frequency.
-   ~~Add an option to reset/skip the timer~~
