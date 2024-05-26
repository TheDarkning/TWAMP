#ifndef TWAMP_H
#define TWAMP_H

#include <iostream>
#include <chrono>
#include <thread>
#include <vector>
#include <cstring>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <unistd.h>
#include <algorithm>
#include <numeric>

struct TWAMP_Packet {
    uint32_t seqNumber;
    uint32_t timestampSec;
    uint32_t timestampUsec;
};

uint64_t current_time_us();

#endif // TWAMP_H

