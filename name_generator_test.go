package go90s

import (
	"strconv"
	"strings"
	"testing"
)

// GetRandomName should return random names
// Collisions should be rare
func TestGetRandomName(t *testing.T) {

	values := []string{
		GetRandomName(),
		GetRandomName(),
		GetRandomName(),
		GetRandomName(),
		GetRandomName(),
	}

	allUnique := true

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] == values[j] {
				allUnique = false
				break
			}
		}
	}

	if !allUnique {
		t.Errorf("GetRandomName is not generating random names, got: %v", strings.Join(values, ","))
	}

}

// GetName should return the same output for each input
// Very rare collisions are acceptable
func TestGetName(t *testing.T) {

	t.Run("GetName 'a' and GetName 'b' should yield different outputs", func(t *testing.T) {
		na := GetName("a")
		nb := GetName("b")

		if na == nb {
			t.Errorf("GetName 'a' and 'b' are returning the same value: %v", na)
		}
	})

	t.Run("GetName 'a' and GetName 'a' should yield the same outputs", func(t *testing.T) {
		na1 := GetName("a")
		na2 := GetName("a")

		if na1 != na2 {
			t.Errorf("GetName 'a' and 'a' are returning different values: %v != %v", na1, na2)
		}
	})

	t.Run("GetName should not allow collisions occur", func(t *testing.T) {
		const l = 1000
		t.Logf("Testing GetName(1-%d) for any collisions", l)
		names := [l]string{}
		for n := 0; n < l; n++ {
			names[n] = GetName(strconv.Itoa(n))
		}

		allUnique := true
		var collisiona int
		var collisionb int
		for i := 0; i < l; i++ {
			for j := i + 1; j < l; j++ {
				if names[i] == names[j] {
					allUnique = false
					collisiona = i
					collisionb = j
					break
				}
			}
		}

		if !allUnique {
			t.Errorf("GetName(1-1000) found collisions: idx %d and %d were the same", collisiona, collisionb)
		}
	})

}
