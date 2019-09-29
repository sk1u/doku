package doku

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestRemove(t *testing.T) {
    assert := assert.New(t)

    d := value("123456789")
    d = d.remove("3")
    assert.Equal(d, value("12456789"))
}

func TestFind(t *testing.T) {
    assert := assert.New(t)

    var i int
    var found bool

    var example = []index{"0", "1", "7", "2", "3", "4"}
    var findTests = []struct {
        elementToFind index
        expectedIndex int
        expectedFound bool
    }{
        {"7", 2, true},
        {"14", -1, false},
    }

    for _, ft := range findTests {
        i, found = find(example, ft.elementToFind)
        assert.Equal(i, ft.expectedIndex)
        assert.Equal(found, ft.expectedFound)
    }
}

func TestCross(t *testing.T) {
    assert := assert.New(t)

    const (
        letters = "AB"
        numbers = "12"
    )

    var result = []index{"A1", "A2", "B1", "B2"}

    crossed := cross(letters, numbers)

    assert.Equal(len(crossed), 4)
    for i, v := range result {
        assert.Equal(crossed[i], v)
    }
}

const (
    puzzle       = "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"
    easypuzzle   = "003020600900305001001806400008102900700000008006708200002609500800203009005010300"
    solvedpuzzle = "417369825632158947958724316825437169791586432346912758289643571573291684164875293"
)

func TestNewSudoku(t *testing.T) {
    assert := assert.New(t)
    s := NewSudoku(easypuzzle)

    // Testing s.squares
    assert.Equal(len(s.squares), 81)
    // fmt.Println("squares: ", s.squares)

    // Testing s.units
    c2units := [][]index{
        {"A1", "A2", "A3", "B1", "B2", "B3", "C1", "C2", "C3"},
        {"A2", "B2", "C2", "D2", "E2", "F2", "G2", "H2", "I2"},
        {"C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9"},
    }
    for _, u := range c2units {
        assert.Contains(s.units["C2"], u)
    }
    // fmt.Println(s.units["A2"])

    // Testing s.peers
    c2peers := []index{
        "A1", "A2", "A3", "B1",
        "B2", "B3", "C1", "C3",
        "D2", "E2", "F2", "G2",
        "H2", "I2", "C4", "C5",
        "C6", "C7", "C8", "C9",
    }
    for _, p := range c2peers {
        assert.Contains(s.peers["C2"], p)
    }
    assert.Equal(len(s.peers["C2"]), 20)
    // fmt.Println(s.peers["A2"])

    // Testing s.grid
    assert.Equal(s.grid["A1"], value("4"))
    assert.Equal(s.grid["A2"], value("123456789"))
    assert.Equal(s.grid["A7"], value("8"))
    assert.Equal(s.grid["I9"], value("123456789"))

    // Testing s.Display()
    assert.Equal(len(s.Display()), 264)

    fmt.Println("Initial Puzzle:")
    fmt.Println(s.Display())

    // Testing s.issolved()
    assert.False(s.issolved())
    assert.True(NewSudoku(solvedpuzzle).issolved())

    fmt.Println("Solved Puzzle:")
    fmt.Println(NewSudoku(solvedpuzzle).Display())
}

func TestAssign(t *testing.T) {
    assert := assert.New(t)

    assert.False(true)
}

func TestEliminate(t *testing.T) {
    assert := assert.New(t)

    assert.False(true)
}

func TestRemoveFromPeers(t *testing.T) {
    assert := assert.New(t)

    assert.False(true)
}

func TestSinglePossibility(t *testing.T) {
    assert := assert.New(t)

    assert.False(true)
}

// TODO: Broken
func TestConstraintPropagation(t *testing.T) {
    assert := assert.New(t)
    s := NewSudoku(easypuzzle)
    fmt.Println(s.Display())
    err := s.Solve()
    fmt.Println(err)
    fmt.Println(s.Display())

    assert.True(s.issolved())
    // fmt.Println(s.grid)

}
