#include <iostream>
#include <cstring>
#include <cstdint>
#include "rapidhash.h"

int main() {
    const char* testInputs[] = {
        "test",
        "hello",
        "world",
        "",  // Empty string
        "123456789",
        "!@#$%^&*()",
        "    ",  // Four spaces
        "\t",     // Tab character
        " \n ",   // Space, newline, space
        "a very long string that exceeds normal lengths to check for performance and correctness of the hash function",
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
        "こんにちは",  // Japanese for "Hello"
        "😊🚀🌟",     // Emojis
        "abcabcabcabcabc",
        "123123123123123",
        "A",                     // Single character
        "AB",                    // Two characters
        "ABCDEFGHIJKLMNOPQRSTUVWXYZ",  // All uppercase letters
        "Mix3d!@#Ch4racters$%^&*()"
    };

    for (const char* input : testInputs) {
        uint64_t hash = rapidhash(input, strlen(input));
        std::cout << "\"" << input << "\"," << std::hex << hash << std::endl;
    }

    return 0;
}