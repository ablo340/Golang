#ifndef FUNCTIONS_H
#define FUNCTIONS_H

#define STUDENT_NAME_MAX 32 

enum MENU_OPTION
{
	MENU_QUIT = 1,
	MENU_SHOW_ALL = 2,
	MENU_SHOW_UNPAID = 3,
	MENU_SEARCH = 4,
	MENU_EDIT = 5,
	MENU_SAVE = 6,
	MENU_LOAD = 7
};

typedef struct student
{
	char fn[STUDENT_NAME_MAX];
	char ln[STUDENT_NAME_MAX];
	int yob;
	int paid;
} student_t;

int init(student_t *array, size_t n);

int do_menu();

void show_student(student_t s);

void show_details(student_t *array, size_t n);

void show_unpaid(student_t *array, size_t n);

void search(student_t *array, size_t n);

void edit(student_t *array, size_t n);

void save(student_t *array, size_t n);

void load(student_t *array, size_t n);

void process_choice(int choice, student_t *array, size_t n);

#endif