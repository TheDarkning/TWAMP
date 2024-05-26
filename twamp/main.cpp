#include "twamp.h"
#include "utils.h"

#include <fstream>
#include <cstdio>


using namespace std;

int main(int argc, char *argv[]) {
    if (argc < 2) {
        print_usage(argv[0]);
        return 1;
    }

    string target = argv[1];
    size_t colon_pos = target.find(':');
    if (colon_pos == string::npos) {
        cerr << "Invalid target format. Use <ip:port>." << endl;
        return 1;
    }

    string ip = target.substr(0, colon_pos);
    int port = stoi(target.substr(colon_pos + 1));

    int packet_count = 10;
    int interval_ms = 1000;
    int payload_size = 12;
    string output_file = "";

    parse_arguments(argc, argv, packet_count, interval_ms, payload_size, output_file);

    string temp_file = output_file + ".tmp";

    ofstream out;
    if (!output_file.empty()) {
        out.open(temp_file);
        if (!out.is_open()) {
            cerr << "Failed to open output file: " << temp_file << endl;
            return 1;
        }
    }
    ostream& output = output_file.empty() ? cout : out;

    int sock = socket(AF_INET, SOCK_DGRAM, 0);
    if (sock < 0) {
        perror("socket");
        return 1;
    }

    struct sockaddr_in server_addr;
    memset(&server_addr, 0, sizeof(server_addr));
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(port);
    inet_pton(AF_INET, ip.c_str(), &server_addr.sin_addr);

    vector<double> latencies;

    for (int i = 0; i < packet_count; ++i) {
        TWAMP_Packet packet;
        uint64_t send_time_us = current_time_us();
        packet.seqNumber = htonl(i + 1);
        packet.timestampSec = htonl(send_time_us / 1000000);
        packet.timestampUsec = htonl(send_time_us % 1000000);

        vector<uint8_t> buf(payload_size);
        memcpy(buf.data(), &packet, sizeof(packet));

        if (sendto(sock, buf.data(), buf.size(), 0, (struct sockaddr*)&server_addr, sizeof(server_addr)) < 0) {
            perror("sendto");
            close(sock);
            return 1;
        }

        struct sockaddr_in from;
        socklen_t fromlen = sizeof(from);
        int n = recvfrom(sock, buf.data(), buf.size(), 0, (struct sockaddr*)&from, &fromlen);
        uint64_t recv_time_us = current_time_us();
        if (n < 0) {
            cerr << "Packet " << i + 1 << ": Request timed out" << endl;
        } else {
            uint32_t recv_sec, recv_usec;
            memcpy(&recv_sec, buf.data() + 4, 4);
            memcpy(&recv_usec, buf.data() + 8, 4);
            recv_sec = ntohl(recv_sec);
            recv_usec = ntohl(recv_usec);

            uint64_t sent_time_us = static_cast<uint64_t>(recv_sec) * 1000000 + recv_usec;
            double round_trip_time_ms = (recv_time_us - send_time_us) / 1000.0;

            latencies.push_back(round_trip_time_ms);
            output << "ID " << i + 1 << ": " << round_trip_time_ms << endl;
        }

        this_thread::sleep_for(chrono::milliseconds(interval_ms));
    }

    close(sock);

    if (!latencies.empty()) {
        double min_latency = *min_element(latencies.begin(), latencies.end());
        double max_latency = *max_element(latencies.begin(), latencies.end());
        double avg_latency = accumulate(latencies.begin(), latencies.end(), 0.0) / latencies.size();
        double packet_loss = ((packet_count - latencies.size()) / (double)packet_count) * 100.0;

        output << "Packet Loss: " << packet_loss << "%" << endl;
        output << "Min Latency: " << min_latency << " ms" << endl;
        output << "Max Latency: " << max_latency << " ms" << endl;
        output << "Avg Latency: " << avg_latency << " ms" << endl;
    } else {
        output << "No latency data available." << endl;
    }
    if (out.is_open()) {
        out.close();
        if (rename(temp_file.c_str(), output_file.c_str()) != 0) {
            cerr << "Error renaming temporary file to output file." << endl;
            return 1;
        }
    }
    return 0;
}
