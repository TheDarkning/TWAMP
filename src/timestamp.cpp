#include "twamp.h"
#include <chrono>
#include <cstdint>
#include <winsock2.h>
#include <stdexcept>

namespace twamp {

    class Timestamp {
    public:
        static TWAMPTimestamp get_current_timestamp() {
            using namespace std::chrono;
            auto now = system_clock::now();
            auto duration_since_epoch = now.time_since_epoch();
            auto seconds = duration_cast<seconds>(duration_since_epoch);
            auto microseconds = duration_cast<microseconds>(duration_since_epoch - seconds);

            TWAMPTimestamp ts;
            // Convert Unix time to NTP time
            ts.integer = htonl(static_cast<uint32_t>(seconds.count() + 2208988800UL));
            ts.fractional = htonl(static_cast<uint32_t>((static_cast<double>(microseconds.count()) / 1e6) * (1ULL << 32)));
            return ts;
        }

        static uint64_t get_usec(const TWAMPTimestamp& ts) {
            struct timeval tv;
            tv.tv_sec = ntohl(ts.integer) - 2208988800UL;
            tv.tv_usec = static_cast<uint32_t>((static_cast<double>(ntohl(ts.fractional)) * 1e6) / (1ULL << 32));
            return tv.tv_sec * 1000000ULL + tv.tv_usec;
        }
    };

} // namespace twamp