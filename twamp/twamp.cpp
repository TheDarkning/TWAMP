#include "twamp.h"

uint64_t current_time_us() 
{
    using namespace std::chrono;
    return duration_cast<microseconds>(system_clock::now().time_since_epoch()).count();
}

