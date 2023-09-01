#include <stdio.h>
#include <string.h>
#include "helper.h"

void rating(char** lines, int line_count, int cols, char t1, char t2) {
	int len = line_count;
	for (int i = 0; i <= cols; i++) {
		int lbal = 0;
		int nlen = 0;
		char com;

		for (int j = 0; j < len; j++)
			lbal += lines[j][i] == '1' ? 1 : -1;

		com = (lbal == 0) ? t1 : (lbal > 0 ? t1 : t2);
		for (int j = 0; j < len; j++)
			if (lines[j][i] == com)
				lines[nlen++] = lines[j];

		len = nlen;

		if (len == 1) break;
	}
}

int main(void) {
	char** lines = get_input();
	int cols = strlen(lines[0]);
	int bal[cols];
	memset(bal, 0, sizeof bal);

	int line_count = 0;
	for (; lines[line_count]; line_count++)
		for (int j = 0; lines[line_count][j]; j++)
			bal[j] += lines[line_count][j] == '1' ? 1 : -1;

	int gamma, epsilon;

	for (int i = 0; i < cols; i++) {
		gamma   |= (bal[i] >= 0) << (cols - i - 1);
		epsilon |= (bal[i] <  0) << (cols - i - 1);
	}

	printf("Part 1: %d\n", gamma*epsilon);

	char* olines[line_count]; char* clines[line_count];
	for (int i = 0; i < line_count; i++) {
		olines[i] = lines[i]; clines[i] = lines[i];
	}

	rating(olines, line_count, cols, '1', '0');
	rating(clines, line_count, cols, '0', '1');
	int oxy = strtol(olines[0], &olines[1], 2);
	int co2 = strtol(clines[0], &clines[1], 2);
	printf("Part 2: %d\n", oxy * co2);
}
