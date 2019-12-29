package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc2019/util"
)

type Planet struct {
	Name         string
	OrbitsPlanet string
}

var planets map[string]Planet

func main() {
	var fp = flag.String("filepath", "", "path to file")
	flag.Parse()

	if *fp == "" {
		panic("filepath is required")
	}

	fileContent := util.ReadFile(*fp)
	planets := map[string]*Planet{}
	for _, row := range fileContent {
		orbitString := strings.Split(row, ")")

		firstPlanet := orbitString[0]
		secondPlanet := orbitString[1]

		// Initialization boilerplate
		first, ok := planets[firstPlanet]
		if !ok {
			first = &Planet{Name: firstPlanet}
			planets[firstPlanet] = first
		}
		second, ok := planets[secondPlanet]
		if !ok {
			second = &Planet{Name: secondPlanet}
			planets[secondPlanet] = second
		}

		// Mark second planet as orbiting the first
		second.OrbitsPlanet = firstPlanet
	}

	// Visit all planets from you to the center
	planet := planets["YOU"]
	visited := map[string]bool{}
	visitOrder := map[string]int{}
	orbits := 1
	exists := false
	for planet.Name != "COM" {
		planet, exists = planets[planet.OrbitsPlanet]
		if !exists {
			panic("planet didnt exist LOL")
		}
		visited[planet.OrbitsPlanet] = true
		visitOrder[planet.OrbitsPlanet] = orbits

		orbits++
	}

	// Visit planets until a planet is recognized
	orbits = 0
	commonVisit := false
	planet = planets["SAN"]
	exists = false
	for !commonVisit {
		planet, exists = planets[planet.OrbitsPlanet]
		if !exists {
			panic("planet didnt exist LOL")
		}

		commonVisit = visited[planet.OrbitsPlanet]
		orbits++
	}

	fmt.Println(planet.Name)
	fmt.Println(orbits)
	fmt.Println(visitOrder[planet.OrbitsPlanet])
}

// Count orbits
func countOrbits() {
	totalOrbits := 0
	for _, planet := range planets {
		cur := planet
		orbits := 0
		var exists bool
		for cur.Name != "COM" {
			cur, exists = planets[cur.OrbitsPlanet]
			if !exists {
				panic("planet didnt exist LOL")
			}
			orbits++
		}
		totalOrbits += orbits
	}

	// Print total orbits
	fmt.Println(totalOrbits)
}
