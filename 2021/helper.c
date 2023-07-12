char** get_input(void) {
	FILE *in = stdin;
	while (fgetc(in) != EOF);
	int size = ftell(in);
	char *content = malloc(size);

	rewind(in);
	for (int i = 0; i < size; i++)
		content[i] = fgetc(in);

	int line_count = 1;

	for (int i = 0; i < strlen(content); i++)
		if (content[i] == '\n')
			line_count++;

	char **lines = malloc(sizeof(char*) * line_count);

	int lineno = 0;
	int rowlen = 0;

	for (int i = 0; i <= strlen(content); i++) {
		if (content[i] == '\n' || content[i] == '\0') {
			lines[lineno] = malloc(sizeof(char) * rowlen);
			strncpy(lines[lineno], content + i - rowlen, rowlen);
			lineno++, rowlen = 0, i++;
		}

		rowlen++;
	}

	free(content);

	return lines;
}
