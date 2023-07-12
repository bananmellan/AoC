#include <stdio.h>
#include <string.h>
#include "helper.h"

int main(void) {
	char** lines = get_input();

	int prev = atoi(lines[0]);
	int incs = 0;
	for (int i = 1; lines[i]; i++) {
		int curr = atoi(lines[i]);
		if (curr > prev)
			incs++;

		prev = curr;
	}

	printf("Part 1: %d\n", incs);

	incs = 0;
	prev = atoi(lines[0]) + atoi(lines[1]) + atoi(lines[2]);
	for (int i = 3; lines[i]; i++) {
		int curr = prev - atoi(lines[i - 3]) + atoi(lines[i]);
		if (curr > prev) incs++;
		prev = curr;
	}

	printf("Part 2: %d\n", incs);
}
