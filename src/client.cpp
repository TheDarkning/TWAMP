#include <iostream>
#include <string>
#include <winsock2.h>
#include <ws2tcpip.h>
#include "twamp.h"

class TwampClient {
public:
    TwampClient(const std::string& server_ip, uint16_t server_port)
        : server_ip_(server_ip), server_port_(server_port) {
        WSADATA wsaData;
        if (WSAStartup(MAKEWORD(2, 2), &wsaData) != 0) {
            throw std::runtime_error("WSAStartup failed.");
        }
        connect_to_server();
    }

    ~TwampClient() {
        closesocket(sockfd_);
        WSACleanup();
    }

    void connect_to_server() {
        struct sockaddr_in server_addr;
        memset(&server_addr, 0, sizeof(server_addr));
        server_addr.sin_family = AF_INET;
        server_addr.sin_port = htons(server_port_);

        if (inet_pton(AF_INET, server_ip_.c_str(), &server_addr.sin_addr) <= 0) {
            throw std::runtime_error("Invalid address/ Address not supported.");
        }

        sockfd_ = socket(AF_INET, SOCK_STREAM, 0);
        if (sockfd_ == INVALID_SOCKET) {
            throw std::runtime_error("Socket creation failed.");
        }

        if (connect(sockfd_, (struct sockaddr*)&server_addr, sizeof(server_addr)) < 0) {
            throw std::runtime_error("Connection Failed.");
        }

        std::cout << "Connected to server." << std::endl;
    }

private:
    SOCKET sockfd_;
    std::string server_ip_;
    uint16_t server_port_;
};

int main(int argc, char* argv[]) {
    if (argc != 3) {
        std::cerr << "Usage: " << argv[0] << " <IP> <port>" << std::endl;
        return 1;
    }

    std::string server_ip = argv[1];
    uint16_t server_port = static_cast<uint16_t>(std::stoi(argv[2]));

    try {
        TwampClient client(server_ip, server_port);
        // Additional client operations
    }
    catch (const std::exception& e) {
        std::cerr << "Exception: " << e.what() << std::endl;
        return 1;
    }

    return 0;
}