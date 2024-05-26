#ifndef UTILS_H
#define UTILS_H

#include <iostream>
#include <string>

void print_usage(const char* progname);
void parse_arguments(int argc, char *argv[], int &packet_count, int &interval_ms, int &payload_size, std::string &output_file);

#endif // UTILS_H

