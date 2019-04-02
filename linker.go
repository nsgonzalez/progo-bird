package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/nsgonzalez/progo-bird/models"
)

func plGenConfigs(goal *models.Goal) []string {
	configs := []string{}
	configs = append(configs, "assert(params("+
		strconv.Itoa(CHAR_RUNSPEED)+", "+
		strconv.FormatFloat(TIME_AGENT_FACTOR, 'f', -1, 64)+")).")
	configs = append(configs, "assert(limitSup("+strconv.Itoa(SCENARIO_H-PLATFORM_TB_MARGIN)+")).")
	configs = append(configs, "assert(limitInf("+strconv.Itoa((SCENARIO_H*(-1))+PLATFORM_TB_MARGIN)+")).")
	configs = append(configs, "assert(charSize("+strconv.Itoa(CHAR_W)+", "+strconv.Itoa(CHAR_H)+")).")
	configs = append(configs, "assert(goal("+strconv.Itoa(int(goal.Pos.X))+", "+strconv.Itoa(int(goal.Pos.Y))+", "+strconv.Itoa(int(goal.Radius))+")).")
	return configs
}

func plGenPlatforms(platforms *[]models.Platform) []string {
	plPlatforms := []string{}

	for i, platform := range *platforms {
		// do not load base and top platforms
		if i >= 2 {
			var minY, maxY int
			minY = int(platform.Rect.Min.Y) - PLATFORM_MARGIN_V
			maxY = int(platform.Rect.Max.Y)
			ptype := "top"
			if platform.Type == models.PLATFORM_BOTTOM {
				minY = int(platform.Rect.Min.Y)
				maxY = int(platform.Rect.Max.Y) + PLATFORM_MARGIN_V
				ptype = "bottom"
			}

			plPlatforms = append(plPlatforms, "assert(platform("+
				strconv.Itoa(int(platform.Rect.Min.X)-PLATFORM_MARGIN_H)+", "+
				strconv.Itoa(minY)+", "+
				strconv.Itoa(int(platform.Rect.Max.X)+PLATFORM_MARGIN_H)+", "+
				strconv.Itoa(maxY)+", "+
				ptype+")).")
		}
	}

	return plPlatforms
}

func plGetGoal() string {
	return "solve(0, 0, " + fmt.Sprintf("%f", (TIME_AGENT_FACTOR*START_AGENT_FACTOR)) + ", Action, Actions), last(Actions, none), !."
}

func linkerExec() *exec.Cmd {
	cmd := exec.Command("./linker.py")
	cmd.Stdout = os.Stdout
	cmd.Start()
	time.Sleep(3 * time.Second)

	return cmd
}

func linkerQuery(configs []string, platforms []string, query string) string {
	conn, _ := net.Dial("tcp", "127.0.0.1:9999")
	defer conn.Close()

	fmt.Println(linkerReadResponse(&conn))

	for _, config := range configs {
		fmt.Fprintf(conn, config+"\n")
		fmt.Println(linkerReadResponse(&conn))
	}

	for _, platform := range platforms {
		fmt.Fprintf(conn, platform+"\n")
		fmt.Println(linkerReadResponse(&conn))
	}

	fmt.Fprintf(conn, query+"\n")
	response := linkerReadResponse(&conn)

	conn.Close()

	return response
}

func linkerReadResponse(conn *net.Conn) string {
	message, _ := bufio.NewReader(*conn).ReadString('\n')
	return message
}
