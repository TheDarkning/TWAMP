#define CPPHTTPLIB_OPENSSL_SUPPORT
#include "httplib.h"
#include "json/json.h"
#include "inja/inja.hpp"
#include "xxhash.h"
#include "xxh3.h"

#include <iostream>
#include <sstream>
#include <iomanip>

void load_json(const char* fname, Json::Value& value){
    std::ifstream ifs(fname);
    if (!ifs.is_open()){
        throw std::exception("Error: failed to open json file!");
    }

    Json::CharReaderBuilder builder;
    builder["collectComments"] = true;

    std::string errs;
    if (!parseFromStream(builder, ifs, &value, &errs)){
        throw std::exception("Error: Failed to parse json!");
    }
}

int main(){
    std :: cout << "Start" << std :: endl;
    //Loading the config file
    Json::Value conf;
    load_json("./config.jsonc", conf);

    // Create server using paths from config
    httplib::SSLServer srv(//"./certificate.pem", "./private_key.pem"
        conf.get("ssl_cert", "crt.pem").asCString(),
        conf.get("sll_key", "key.pem").asCString()
    );
    // std::cout << conf.get("ssl_cert", "crt.pem").asCString() << ' ' << conf.get("sll_key", "key.pem").asCString() << std::endl;
    
    
    // Register www-data as content/static directory
    srv.set_mount_point("/", "./webpage");

    // Main index
    srv.Get("/", [&](const httplib::Request& req, httplib::Response& res) {
         res.set_redirect("/index.html");
    });

    srv.Get("/submit", [&](const httplib::Request& req, httplib::Response& res) {
        std::string target_ip = req.get_param_value("target_ip");
        std::cout << "target_ip: " << target_ip << std::endl;

        std::string test_duration = req.get_param_value("test_duration");
        std::cout << "test_duration: " << test_duration << std::endl;

        std::string packet_size = req.get_param_value("packet_size");
        std::cout << "packet_size: " << packet_size << std::endl;
    });


    // Stop the server when the user access /stop
    srv.Get("/stop", [&](const httplib::Request& req, httplib::Response& res) {
         srv.stop();
         res.set_redirect("/");
    });

    // Listen server to port
    //httplib::Server srv;
    if (!srv.is_valid()) {
        printf("server has an error...\n");
        return -1;
    }
    //srv.Get("/", [&](const httplib::Request& /*req*/, httplib::Response& res) {
     //   res.set_content("Hello World", "text/html");
      //  });

    srv.listen( ///"localhost", 443
        conf.get("server_ip", "0.0.0.0").asCString(),
        conf.get("server_port", 443).asInt()
    );

    return 0;
}