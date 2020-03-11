# ProGo Bird
ProGo Bird is a simple Agent written in Prolog that plans a game for a custom and basic version of Flappy Bird where the bird has to reach a goal. It's pretty basic and it does not calculate what action it should do online, instead it does all the calculation before the game start (calling swipl through python), then store the actions to execute in a slice and pops them until the game is finished.

The game is based on a modified version of **Platform** under the Pixel's examples (https://github.com/faiface/pixel-examples/tree/master/platformer).

## Requisites
  - debian packages
    - git
    - golang
    - swi-prolog
    - python3-pip
    - libgl1-mesa-dev
    - libxcursor-dev
    - libxrandr-dev
    - libxinerama-dev
    - libxi-dev
  - pip packages
    - pyswip (refactor branch)
  - OpenGL >= 3.0

###### Note: it only was tested under Debian 9.4 x86, Ubuntu 18.04.1 x64 and Deepin 15.9.2 x64

##### Install dependencies

```sh
$ sudo apt install -y git golang swi-prolog python3-pip libgl1-mesa-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev
$ wget -O pyswip.zip https://github.com/yuce/pyswip/archive/refactor.zip
$ unzip pyswip.zip
$ cd pyswip-refactor
$ python3 setup.py install --user
```

## How it works
When the game sarts the bird is at (0, 0) and we need it to arrive to a goal located at (X, Y) with a radius R. The speed of the loop in which the game runs is 60fps, so in order to make it work it is multiplied by a factor, which is passed to the agent program so it calculates which moves it should make according to that factor starting from the point (0, 0). Then, the result of the multiplication give us the fps at which the agent performs an action. For example, if the factor would be 0.5, then the result would be 30 and the agent would calculate which actions it should perform every 30fps (0.5 seconds approximately).

I could't find or get successful results with any go library implementing a prolog interpreter, so before the game's loop starts a few calls to a python script implementing pyswip (swipl library for python) is made through a tcp socket and reading back the actions it should perform to play a successful game.

## Run
Clone the repository, install the go dependencies and run the program.

```sh
$ git clone https://github.com/nsgonzalez/progo-bird.git
$ cd progo-bird
$ go get ./...
$ go run *.go
```
![](game.gif)

#### Successful games
The games that proved to be successful under this weird implementation are the following, and all of them require to modify  **parameters.go** as follows. 

##### Game 1 - 30fps
```golang
TIME_AGENT_FACTOR = 0.5
START_AGENT_FACTOR = 4
PLATFORM_MARGIN_H  = 18
PLATFORM_MARGIN_V  = 0
PLATFORM_TB_MARGIN = 0
```
##### Game 2 - 24fps
```golang
TIME_AGENT_FACTOR = 0.40
START_AGENT_FACTOR = 4
PLATFORM_MARGIN_H  = 10
PLATFORM_MARGIN_V  = 0
PLATFORM_TB_MARGIN = 0
```
##### Game 3 - 18fps
```golang
TIME_AGENT_FACTOR = 0.30
START_AGENT_FACTOR = 5
PLATFORM_MARGIN_H  = 10
PLATFORM_MARGIN_V  = 0
PLATFORM_TB_MARGIN = 0
```
##### Game 4 - 15fps
```golang
TIME_AGENT_FACTOR = 0.25
START_AGENT_FACTOR = 4
PLATFORM_MARGIN_H  = 10
PLATFORM_MARGIN_V  = 10
PLATFORM_TB_MARGIN = 10
```
##### Game 5 - 12fps
```golang
TIME_AGENT_FACTOR = 0.20
START_AGENT_FACTOR = 7
PLATFORM_MARGIN_H  = 0
PLATFORM_MARGIN_V  = 10
PLATFORM_TB_MARGIN = 0
```

Just a little explanation about the variables (constants actually):
- **TIME_AGENT_FACTOR** is the factor that indicates how time is discretized.
- **START_AGENT_FACTOR** is multiplied with **TIME_AGENT_FACTOR** to simulate the fall when game starts.
- **PLATFORM_MARGIN_H** AND **PLATFORM_MARGIN_V** are horizontal and vertical margins for the platforms, so the agent doesn't calculate the moves too close to them. 
- **PLATFORM_TB_MARGIN** same as above but for the top and bottom platforms.
