#include <stdio.h>
#include <string.h>

void safeFunction()
{
    char buffer[50];
    strncpy(buffer, "This is a safe string", sizeof(buffer) - 1);
    buffer[sizeof(buffer) - 1] = '\0'; // Ensure null termination
    printf("Buffer content: %s\n", buffer);
}

int main()
{
    safeFunction();
    return 0;
}
