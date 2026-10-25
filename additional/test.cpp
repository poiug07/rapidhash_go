#include <iostream>
#include <fstream>
#include <cstring>
#include <cstdint>
#include <string>
#include "rapidhash.h"

int main() {
    std::ofstream out("file.csv", std::ios::trunc);

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
        out << "\"" << input << "\"," << std::hex << hash << std::endl;
    }

    // "a" repeated N times, around chunk boundaries.
    for (int len : {15, 16, 17, 47, 48, 49, 111, 112, 113, 223, 224, 225}) {
        std::string s(len, 'a');
        uint64_t hash = rapidhash(s.data(), s.size());
        out << "\"" << s << "\"," << std::hex << hash << std::endl;
    }

    return 0;
}
