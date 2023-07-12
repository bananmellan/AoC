#include "helper.h"
#include <stdio.h>
#include <string.h>

int main(void) {
	char **lines = get_input();

	int depth = 0, hpos = 0;

	for (int i = 0; lines[i]; i++) {
		char dir[8];
		int len;

		sscanf(lines[i], "%s %d", dir, &len);

		switch (dir[0]) {
		case 'u':
			depth -= len;
			break;
		case 'd':
			depth += len;
			break;
		default:
			hpos += len;
			break;
		}
	}

	printf("Part 1: %d\n", depth * hpos);

	int aim = 0; depth = 0, hpos = 0;

	for (int i = 0; lines[i]; i++) {
		char dir[8];
		int len;

		sscanf(lines[i], "%s %d", dir, &len);

		switch (dir[0]) {
		case 'u':
			aim -= len;
			break;
		case 'd':
			aim += len;
			break;
		default:
			hpos += len;
			depth += aim * len;
			break;
		}
	}

	printf("Part 2: %d\n", depth * hpos);
}
