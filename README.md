# LEM-IN

> Project: LEM-IN (ANT COLONY SIMULATION)

* This project simulates the movement of ants in a graph where nodes represent locations and edges represent paths between these locations. The objective is to find paths from a start node to an end node and simulate the ants' movement along these paths.

### Features
* Reads input from a file to configure the graph, coordinates, and paths.
* Validates the input data format.
* Constructs a bidirectional graph.
* Finds all possible paths between the start and end nodes using Breadth-First Search (BFS).
* Filters and selects paths to avoid overlapping nodes.
* Simulates the movement of ants along the selected paths.
* Measures the execution time of the program.
* Prerequisites
#### Go 1.16 or later
Usage
Clone the repository:

```go
git clone https://github.com/Emirtariksahin/Lem-in.git
```
```go 
cd Lem-in/service 
```
### Prepare the input file:

Ensure you have an input file formatted as follows and place it in the testler directory:

```go
import 
[Number of ants]
##start
[start node] [x-coordinate] [y-coordinate]
##end
[end node] [x-coordinate] [y-coordinate]
[node name] [x-coordinate] [y-coordinate]
[node name] [x-coordinate] [y-coordinate]
...
[node1]-[node2]
[node2]-[node3]
...
```
### Run the program:
```go
go run . input.txt
```
* Replace input.txt with the name of your input file.

### Input File Format
The input file should be structured in the following way:

* First line: The number of ants.
* Second line: ##start keyword.
* Third line: Start node with its coordinates.
* Fourth line: ##end keyword.
* Fifth line: End node with its coordinates.
* Subsequent lines: Other nodes with their coordinates.
* Following lines: Connections between nodes in the format [node1]-[node2].
>Example Input File
```
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1

```
>Output
```
Round 1: L1-2
Round 2: L1-3 L2-2
Round 3: L1-1 L2-3 L3-2
Round 4: L2-1 L3-3 L4-2
Round 5: L3-1 L4-3
Round 6: L4-1
```
### The program outputs the following information:

>>Number of ants.
* Coordinates of all nodes.
* Start and end coordinates.
* Connections between nodes.
* Constructed graph details.
* Found paths from the start node to the end node.
* Filtered paths for the ants to avoid overlapping.
* Simulation of ants moving along the paths.
>>>Functions
* readInputFile(fileName string): Reads the input file and returns the lines as a slice of strings.
* parseCoordinates(sentences []string): Parses coordinates from the input sentences and returns a map of coordinates.
* baglantilar(sentences []string): Extracts connections (edges) from the input sentences and returns a slice of strings.
* parseStartEndCoordinates(sentences []string): Parses the start and end coordinates from the input sentences.
* createGraph(coordinates map[string][2]int, connections []string): Creates a bidirectional graph from coordinates and connections.
* FindNodeByName(name string): Finds a node in the graph by its name.
* convertPathsToString(paths [][]Node): Converts paths represented by nodes to string format.
* convertToNodePaths(paths [][]string, graph Graph): Converts string paths back to node paths.
* printNodePaths(paths [][]Node): Prints the paths in terms of node names.
* FilterRoad(paths [][]string, antCount int): Filters paths to avoid node overlaps.
* SimulateAnts(graph *Graph, ants int, start, end *Node, allPaths [][]Node, finalNodePaths []Node): Simulates the movement of ants along the paths.
>>>>Notes
* Ensure the input file is correctly formatted, especially the coordinates and connections.
* The program checks for invalid data formats and outputs error messages if found.
The simulation considers the maximum number of non-overlapping paths and assigns them to the ants.