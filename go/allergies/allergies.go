package allergies

/*
eggs          (1)
peanuts       (2)
shellfish     (4)
strawberries  (8)
tomatoes      (16)
chocolate     (32)
pollen        (64)
cats          (128)

Tests
{[]string{}, 0},
{[]string{"eggs"}, 1},
{[]string{"peanuts"}, 2},
{[]string{"strawberries"}, 8},
{[]string{"eggs", "peanuts"}, 3},
{[]string{"eggs", "shellfish"}, 5},
{[]string{"strawberries", "tomatoes", "chocolate", "pollen", "cats"}, 248},
{[]string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}, 255},
{[]string{"eggs", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}, 509},

*/

func Allergies(score int) (name []string) {
	return
}

/*
Tests
{false, 0, "peanuts"},
{false, 0, "cats"},
{false, 0, "strawberries"},
{true, 1, "eggs"},
{true, 5, "eggs"},
*/

func AllergicTo(score int, name string) (fact bool) {
	return
}
