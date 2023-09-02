package thousands

import "fmt"

func ExampleThousands() {
	fmt.Printf("%s\n%#[1]s\n%1.0[1]v\n%1.1[1]s\n%-1.2[1]s\n%#[1]s\n%#1.1[1]s\n", Int(1000000000))
	// Output:
	// 1,000,000,000
	// 1,000,000,000
	// 1 000 000 000
	// 1,000,000k
	// 1.000M
	// 953Mi
	// 976,562ki
}

func ExampleThousandsNegatives() {
	fmt.Printf("%s\n%#[1]s\n%1.0[1]v\n%1.1[1]v\n%-1.2[1]s\n%#[1]s\n%#1.1[1]s\n", Int(-1<<30))
	// Output:
	// -1,073M
	// -1,024Mi
	// -1 073 741 824
	// -1 073 741k
	// -1.073M
	// -1,024Mi
	// -1,048,576ki
}

