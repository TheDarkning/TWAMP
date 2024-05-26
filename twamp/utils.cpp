#include "utils.h"

void print_usage(const char* progname) {
    std::cerr << "Usage: " << progname << " <ip:port> [--count=<count>] [--interval=<interval_ms>] [--payload=<payload_size>] [--output=<filename>]" << std::endl;
}

void parse_arguments(int argc, char *argv[], int &packet_count, int &interval_ms, int &payload_size, std::string &output_file) {
    for (int i = 2; i < argc; ++i) {
        std::string arg = argv[i];
        if (arg.find("--count=") == 0) {
            packet_count = std::stoi(arg.substr(8));
        } else if (arg.find("--interval=") == 0) {
            interval_ms = std::stoi(arg.substr(11));
        } else if (arg.find("--payload=") == 0) {
            payload_size = std::stoi(arg.substr(10));
        } else if (arg.find("--output=") == 0) {
            output_file = arg.substr(9);
        }
    }
}

