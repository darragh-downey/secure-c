#include <stdio.h>
#include <string.h>
#include <stdlib.h>

// Buffer overflow example
void bufferOverflow()
{
    char buffer[10];
    strcpy(buffer, "This string is too long for the buffer");
    printf("Buffer content: %s\n", buffer);
}

// Format string vulnerability example
void formatStringVulnerability(char *userInput)
{
    char buffer[100];
    snprintf(buffer, sizeof(buffer), userInput);
    printf(buffer);
}

// Integer overflow example
void integerOverflow()
{
    unsigned int max = 4294967295;
    unsigned int result = max + 1;
    printf("Integer overflow result: %u\n", result);
}

// Use of gets() function (unsafe)
void useOfGets()
{
    char buffer[50];
    gets(buffer); // gets() is unsafe and should not be used
    printf("You entered: %s\n", buffer);
}

// Use of system() function without validation
void useOfSystem(char *userInput)
{
    char command[100];
    snprintf(command, sizeof(command), "echo %s", userInput);
    system(command); // Potential command injection
}

int main()
{
    char userInput[100];

    // Testing buffer overflow
    bufferOverflow();

    // Testing format string vulnerability
    printf("Enter a format string: ");
    scanf("%99s", userInput);
    formatStringVulnerability(userInput);

    // Testing integer overflow
    integerOverflow();

    // Testing unsafe use of gets()
    printf("Enter a string: ");
    useOfGets();

    // Testing use of system() function
    printf("Enter a command: ");
    scanf("%99s", userInput);
    useOfSystem(userInput);

    return 0;
}
